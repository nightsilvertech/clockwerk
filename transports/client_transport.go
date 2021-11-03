package transports

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "github.com/nightsilvertech/clockwerk/endpoints"
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	_interface "github.com/nightsilvertech/clockwerk/service/interface"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func encodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func decodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func ClockwerkClient(conn *grpc.ClientConn) _interface.Clockwerk {

	var getSchedulersEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.Clockwerk`
			rpcMethod = `GetSchedulers`
		)

		getSchedulersEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.Schedulers{},
		).Endpoint()
	}

	var addSchedulerEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.Clockwerk`
			rpcMethod = `AddScheduler`
		)

		addSchedulerEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			pb.Scheduler{},
		).Endpoint()
	}

	var deleteSchedulerEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.Clockwerk`
			rpcMethod = `DeleteScheduler`
		)

		deleteSchedulerEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			&emptypb.Empty{},
		).Endpoint()
	}

	var toggleSchedulerEp endpoint.Endpoint
	{
		const (
			rpcName   = `api.v1.Clockwerk`
			rpcMethod = `ToggleScheduler`
		)

		toggleSchedulerEp = grpctransport.NewClient(
			conn,
			rpcName,
			rpcMethod,
			encodeRequest,
			decodeResponse,
			&emptypb.Empty{},
		).Endpoint()
	}

	return &ep.ClockwerkEndpoint{
		GetSchedulersEndpoint:   getSchedulersEp,
		AddSchedulerEndpoint:    addSchedulerEp,
		DeleteSchedulerEndpoint: deleteSchedulerEp,
		ToggleSchedulerEndpoint: toggleSchedulerEp,
	}
}
