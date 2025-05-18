package endpoints

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	NO_AUTH_HEADER            = status.Errorf(codes.InvalidArgument, "No authheader provided")
	PERMISSION_DENIED         = status.Errorf(codes.PermissionDenied, "Permission Denied")
	INTERNAL_ERROR            = status.Errorf(codes.Internal, "Unkonwn Internal failure")
	POSTGRES_CONNECTION_ERROR = status.Errorf(codes.Internal, "Cannot connect to postgres")
)
