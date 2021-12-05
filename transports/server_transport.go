package transports

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "github.com/nightsilvertech/clockwerk/endpoints"
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcClockwerkServer struct {
	pb.UnimplementedClockwerkServer
	getDummy        grpctransport.Handler
	postDummy       grpctransport.Handler
	deleteDummy     grpctransport.Handler
	putDummy        grpctransport.Handler
	getSchedulers   grpctransport.Handler
	addScheduler    grpctransport.Handler
	deleteScheduler grpctransport.Handler
	toggleScheduler grpctransport.Handler
	backup          grpctransport.Handler
}

func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeEmptyPbResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &emptypb.Empty{}, nil
}

func (g grpcClockwerkServer) GetDummy(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_, res, err := g.getDummy.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) PostDummy(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_, res, err := g.postDummy.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) DeleteDummy(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_, res, err := g.deleteDummy.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) PutDummy(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_, res, err := g.putDummy.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) GetSchedulers(ctx context.Context, empty *emptypb.Empty) (*pb.Schedulers, error) {
	_, res, err := g.getSchedulers.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Schedulers), nil
}

func (g grpcClockwerkServer) AddScheduler(ctx context.Context, scheduler *pb.Scheduler) (*pb.Scheduler, error) {
	_, res, err := g.addScheduler.ServeGRPC(ctx, scheduler)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Scheduler), nil
}

func (g grpcClockwerkServer) DeleteScheduler(ctx context.Context, selectScheduler *pb.SelectScheduler) (*emptypb.Empty, error) {
	_, res, err := g.deleteScheduler.ServeGRPC(ctx, selectScheduler)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) ToggleScheduler(ctx context.Context, toggle *pb.SelectToggle) (*emptypb.Empty, error) {
	_, res, err := g.toggleScheduler.ServeGRPC(ctx, toggle)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcClockwerkServer) Backup(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	_, res, err := g.backup.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func NewClockwerkServer(endpoints ep.ClockwerkEndpoint) pb.ClockwerkServer {
	return &grpcClockwerkServer{
		getDummy: grpctransport.NewServer(
			endpoints.GetDummyEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),
		postDummy: grpctransport.NewServer(
			endpoints.PostDummyEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),
		deleteDummy: grpctransport.NewServer(
			endpoints.DeleteDummyEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),
		putDummy: grpctransport.NewServer(
			endpoints.PutDummyEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),

		getSchedulers: grpctransport.NewServer(
			endpoints.GetSchedulersEndpoint,
			decodeRequest,
			encodeResponse,
		),
		addScheduler: grpctransport.NewServer(
			endpoints.AddSchedulerEndpoint,
			decodeRequest,
			encodeResponse,
		),
		deleteScheduler: grpctransport.NewServer(
			endpoints.DeleteSchedulerEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),
		toggleScheduler: grpctransport.NewServer(
			endpoints.ToggleSchedulerEndpoint,
			decodeRequest,
			encodeEmptyPbResponse,
		),
	}
}
