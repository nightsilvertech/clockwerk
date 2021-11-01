package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pb "gitlab.com/nbdgocean6/clockwerk/protocs/api/v1"
	_interface "gitlab.com/nbdgocean6/clockwerk/service/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClockwerkEndpoint struct {
	GetSchedulersEndpoint   endpoint.Endpoint
	AddSchedulerEndpoint    endpoint.Endpoint
	DeleteSchedulerEndpoint endpoint.Endpoint
	ToggleSchedulerEndpoint endpoint.Endpoint
}

func makeGetSchedulersEndpoint(usecase _interface.Clockwerk) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetSchedulers(ctx, request.(*emptypb.Empty))
		return res, err
	}
}

func (e ClockwerkEndpoint) GetSchedulers(ctx context.Context, empty *emptypb.Empty) (*pb.Schedulers, error) {
	res, err := e.GetSchedulersEndpoint(ctx, empty)
	if err != nil {
		return res.(*pb.Schedulers), err
	}
	return res.(*pb.Schedulers), nil
}

func makeAddSchedulerEndpoint(usecase _interface.Clockwerk) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.AddScheduler(ctx, request.(*pb.Scheduler))
		return res, err
	}
}

func (e ClockwerkEndpoint) AddScheduler(ctx context.Context, request *pb.Scheduler) (*pb.Scheduler, error) {
	res, err := e.AddSchedulerEndpoint(ctx, request)
	if err != nil {
		return &pb.Scheduler{}, err
	}
	return res.(*pb.Scheduler), nil
}

func makeDeleteSchedulerEndpoint(usecase _interface.Clockwerk) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.DeleteScheduler(ctx, request.(*pb.SelectScheduler))
		return res, err
	}
}

func (e ClockwerkEndpoint) DeleteScheduler(ctx context.Context, request *pb.SelectScheduler) (*emptypb.Empty, error) {
	res, err := e.DeleteSchedulerEndpoint(ctx, request)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return res.(*emptypb.Empty), nil
}

func makeToggleSchedulerEndpoint(usecase _interface.Clockwerk) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.ToggleScheduler(ctx, request.(*pb.SelectToggle))
		return res, err
	}
}

func (e ClockwerkEndpoint) ToggleScheduler(ctx context.Context, request *pb.SelectToggle) (*emptypb.Empty, error) {
	res, err := e.ToggleSchedulerEndpoint(ctx, request)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return res.(*emptypb.Empty), nil
}

func NewClockwerkEndpoint(usecase _interface.Clockwerk) ClockwerkEndpoint {
	return ClockwerkEndpoint{
		GetSchedulersEndpoint:   makeGetSchedulersEndpoint(usecase),
		AddSchedulerEndpoint:    makeAddSchedulerEndpoint(usecase),
		DeleteSchedulerEndpoint: makeDeleteSchedulerEndpoint(usecase),
		ToggleSchedulerEndpoint: makeToggleSchedulerEndpoint(usecase),
	}
}
