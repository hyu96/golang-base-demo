package workflow

import (
	"context"

	"github.com/huydq/gokits/third_party/dtm-client/dtmcli/dtmimp"
	"github.com/huydq/gokits/third_party/dtm-client/dtmgrpc"
	"github.com/huydq/gokits/third_party/dtm-client/dtmgrpc/dtmgimp"
	"github.com/huydq/gokits/third_party/dtm-client/workflow/wfpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type workflowServer struct {
	wfpb.UnimplementedWorkflowServer
}

func (s *workflowServer) Execute(ctx context.Context, wd *wfpb.WorkflowData) (*emptypb.Empty, error) {
	if defaultFac.protocol != dtmimp.ProtocolGRPC {
		return nil, status.Errorf(codes.Internal, "workflow server not inited. please call workflow.InitGrpc first")
	}
	tb := dtmgimp.TransBaseFromGrpc(ctx)
	_, err := defaultFac.execute(ctx, tb.Op, tb.Gid, wd.Data)
	return &emptypb.Empty{}, dtmgrpc.DtmError2GrpcError(err)
}
