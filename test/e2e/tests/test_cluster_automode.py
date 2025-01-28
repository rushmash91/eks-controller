import boto3
import logging
import time
import pytest

from acktest.k8s import resource as k8s
from acktest.k8s import condition
from acktest.resources import random_suffix_name
from e2e import (
    service_marker,
    CRD_GROUP,
    CRD_VERSION,
    load_eks_resource
)
from e2e.bootstrap_resources import get_bootstrap_resources
from e2e.common.types import CLUSTER_RESOURCE_PLURAL
from e2e.common.waiter import wait_until_deleted
from e2e.replacement_values import REPLACEMENT_VALUES

MODIFY_WAIT_AFTER_SECONDS = 60
CHECK_STATUS_WAIT_SECONDS = 60

def wait_for_cluster_active(eks_client, cluster_name):
    waiter = eks_client.get_waiter('cluster_active')
    waiter.config.delay = 5
    waiter.config.max_attempts = 240  
    waiter.wait(name=cluster_name)


@pytest.fixture(scope="module")
def eks_client():
    return boto3.client('eks')


@pytest.fixture
def auto_mode_cluster(eks_client):
    cluster_name = random_suffix_name("auto-mode-cluster", 32)

    replacements = REPLACEMENT_VALUES.copy()
    replacements["CLUSTER_NAME"] = cluster_name

    resource_data = load_eks_resource(
        "cluster_automode",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    ref = k8s.CustomResourceReference(
        CRD_GROUP,
        CRD_VERSION,
        CLUSTER_RESOURCE_PLURAL,
        cluster_name,
        namespace="default",
    )

    # Create the CR
    k8s.create_custom_resource(ref, resource_data)
    cr = k8s.wait_resource_consumed_by_controller(ref, wait_periods=15)
    assert cr is not None, "Cluster CR was not created in Kubernetes"
    assert k8s.get_resource_exists(ref), "Could not find the Cluster CR in K8s"

    yield (ref, cr)

    # Cleanup
    try:
        _, deleted = k8s.delete_custom_resource(ref, 3, 10)
        if deleted:
            wait_until_deleted(cluster_name)
    except:
        pass


@service_marker
@pytest.mark.canary
class TestAutoModeCluster:
    def test_create_auto_mode_cluster(self, eks_client, auto_mode_cluster):
        (ref, cr) = auto_mode_cluster
        cluster_name = cr["spec"]["name"]

        try:
            aws_res = eks_client.describe_cluster(name=cluster_name)
            assert aws_res is not None
        except eks_client.exceptions.ResourceNotFoundException:
            pytest.fail(f"Could not find cluster '{cluster_name}' in EKS")
        
        wait_for_cluster_active(eks_client, cluster_name)

        time.sleep(CHECK_STATUS_WAIT_SECONDS)

        patch_zonal_shift_false = {
            "spec": {
                "zonalShiftConfig": {
                    "enabled": False
                }
            }
        }
        k8s.patch_custom_resource(ref, patch_zonal_shift_false)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        wait_for_cluster_active(eks_client, cluster_name)
        aws_res = eks_client.describe_cluster(name=cluster_name)
        logging.info(f"AWS API Response: {aws_res}")
        assert aws_res["zonalShiftConfig"]["enabled"] is False
        logging.info(f"Cluster {cluster_name} zonal shift disabled successfully.")

        # Step 4: Remove the "system" node pool
        patch_remove_system_pool = {
            "spec": {
                "computeConfig": {
                    "nodePools": ["general-purpose"]
                }
            }
        }
        k8s.patch_custom_resource(ref, patch_remove_system_pool)
        time.sleep(MODIFY_WAIT_AFTER_SECONDS)

        wait_for_cluster_active(eks_client, cluster_name)
        aws_res = eks_client.describe_cluster(name=cluster_name)
        compute_config = aws_res["computeConfig"]
        assert "system" not in compute_config["nodePools"]
        logging.info(f"Cluster {cluster_name} updated successfully. 'system' node pool removed.")

