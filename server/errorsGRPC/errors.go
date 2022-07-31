package errorsGRPC

import (
	"URLShortenerGRPC/server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorCreate(message string) (*pb.URLResponse, error) {
	err := status.Newf(
		codes.InvalidArgument,
		message)
	return nil, err.Err()
}

func ErrorGet(message string) (*pb.URLRequest, error) {
	err := status.Newf(
		codes.InvalidArgument,
		message)
	return nil, err.Err()
}
