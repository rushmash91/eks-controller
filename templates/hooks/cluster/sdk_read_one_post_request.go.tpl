    {{ $errVar := "err" }} 
    if {{ $errVar }} != nil {
        // For debug: log the type/kind of error
        rlog.Info("DescribeCluster error encountered",
            "errString", {{ $errVar }}.Error(),
            "errType", fmt.Sprintf("%T", {{ $errVar }}),
        )
        // Attempt to unwrap a smithy.OperationError
        var opErr *smithy.OperationError
        if errors.As({{ $errVar }}, &opErr) {
            // check for EKS types.ResourceNotFoundException
            var rnfe *svcsdktypes.ResourceNotFoundException
            if errors.As(opErr.Err, &rnfe) {
                // ACK interprets ackerr.NotFound as "go create it"
                return nil, ackerr.NotFound
            }
            // Or a generic ResourceNotFound by code
            var genErr *smithy.GenericAPIError
            if errors.As(opErr.Err, &genErr) && genErr.Code == "ResourceNotFoundException" {
                return nil, ackerr.NotFound
            }
        }
        // Fallback: other error => pass it along normally
        return nil, {{ $errVar }}
    }