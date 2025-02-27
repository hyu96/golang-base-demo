package grpc

import (
	"encoding/base64"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"github.com/huydq/gokits/libs/ilog"
	com "github.com/huydq/proto/gen-go/common"
)

var (
	headerRPCError = "rpc_error"
)

func RPCErrorToMD(md *com.ApiError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		ilog.Errorf("RPCErrorToMD - Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRPCError, base64.StdEncoding.EncodeToString(buf)), nil
}
