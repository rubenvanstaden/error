package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rubenvanstaden/error/catch"
)

func Error(err error) error {

	code := catch.ErrorCode(err)
	message := catch.ErrorMessage(err)

	if code == catch.INTERNAL {
		return status.Error(codes.Internal, message)
	}

	if code == catch.NOTFOUND {
		return status.Error(codes.NotFound, message)
	}

	if code == catch.CONFLICT {
		return status.Error(codes.AlreadyExists, message)
	}

	if code == catch.INVALID {
		return status.Error(codes.InvalidArgument, message)
	}

	return status.Error(codes.Unknown, "encountered unknown error")
}
