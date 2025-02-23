package grpc

import (
	"encoding/base64"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"github.com/huydq/gokits/libs/ilog"
	com "github.com/huydq/proto/gen-go/common"
)

var (
	headerRPCError = "rpc_error"
)

// Server To Client
func RPCErrorFromMD(md metadata.MD) (rpcErr *com.ApiError) {
	ilog.Info("rpc error from md: ", md)
	val := metautils.NiceMD(md).Get(headerRPCError)
	if val == "" {
		rpcErr = com.NewApiError(com.KErrorCode_ERR_INTERNAL, "Unknown error")
		ilog.Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = com.NewApiError(com.KErrorCode_ERR_INTERNAL, fmt.Sprintf("Base64 decode error, rpc_error: %s, error: %v", val, err))
		ilog.Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	rpcErr = &com.ApiError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = com.NewApiError(com.KErrorCode_ERR_INTERNAL, fmt.Sprintf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err))
		ilog.Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	return rpcErr
}

func RPCErrorToMD(md *com.ApiError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		ilog.Errorf("RPCErrorToMD - Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRPCError, base64.StdEncoding.EncodeToString(buf)), nil
}
