package endpoints

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	NO_AUTH_HEADER         = status.Errorf(codes.InvalidArgument, "No authheader provided")
	FAILED_TO_GET_METADATA = status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	PERMISSION_DENIED      = status.Errorf(codes.PermissionDenied, "Permission Denied")
	INTERNAL_ERROR         = status.Errorf(codes.Internal, "Failed to connect to postgres")
)
