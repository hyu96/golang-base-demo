package grpc

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"

	"github.com/huydq/gokits/libs/ilog"
	com "github.com/huydq/proto/gen-go/common"
)

var (
	headerRPCMetadata = "rpc_metadata"
)

func RPCMetadataFromMD(md metadata.MD) (*RpcMetadata, error) {
	val := metautils.NiceMD(md).Get(headerRPCMetadata)
	if val == "" {
		return nil, nil
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error, rpc_metadata: %s, error: %v", val, err)
	}

	rpcMetadata := &RpcMetadata{}
	err = proto.Unmarshal(buf, rpcMetadata)
	if err != nil {
		return nil, fmt.Errorf("RpcMetadata unmarshal error, rpc_metadata: %s, error: %v", val, err)
	}

	return rpcMetadata, nil
}

func RPCMetadataFromIncoming(ctx context.Context) *RpcMetadata {
	md2, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	md, err := RPCMetadataFromMD(md2)
	if err != nil {
		panic(com.NewApiError(com.KErrorCode_ERR_OTHER, fmt.Sprintf("%s", err)))
	}

	return md
}

func RPCMetadataToOutgoing(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		ilog.Errorf("RPCMetadataToOutgoing - Marshal rpc_metadata error: %+v", err)
		return nil, err
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}

// For send internal server
func RPCMetadataToOutgoingForInternal(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		ilog.Errorf("RPCMetadataToOutgoingForInternal - Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.NewIncomingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}
