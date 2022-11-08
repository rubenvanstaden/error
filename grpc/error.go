package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rubenvanstaden/error/catch"
)

func Error(err error) error {

	code := catch.ErrorCode(err)
	message := catch.ErrorMessage(err)

	if code == catch.ERROR_INTERNAL {
		return status.Error(codes.Internal, message)
	}

	if code == catch.ERROR_NOTFOUND {
		return status.Error(codes.NotFound, message)
	}

	if code == catch.ERROR_CONFLICT {
		return status.Error(codes.AlreadyExists, message)
	}

	if code == catch.ERROR_INVALID {
		return status.Error(codes.InvalidArgument, message)
	}

	return status.Error(codes.Unknown, "encountered unknown error")
}
