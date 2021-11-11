package client

import (
	"context"
	"fmt"
	"github.com/nightsilvertech/clockwerk/constant"
	pbclockwerk "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	clockwerksvc "github.com/nightsilvertech/clockwerk/service/interface"
	clockwerktrpt "github.com/nightsilvertech/clockwerk/transports"
	grpcgoogle "google.golang.org/grpc"
	"strings"
)

type ClockwerkClient interface {
	Add(scheduler SchedulerHTTP) (SchedulerHTTP, error)
	Del(schedulerSelect SchedulerSelect) error
	Toggle(schedulerToggle SchedulerToggle) error
}

type clockwerkClient struct {
	Username        string
	Password        string
	ClockwerkServer clockwerksvc.Clockwerk
}

// HTTPHeader simple http header representation of request
type HTTPHeader struct {
	K, V string
}

// String convert HTTPHeader to string version split by bar sign '|'
func (hh *HTTPHeader) String() string {
	return fmt.Sprintf("%s|%s", hh.K, hh.V)
}

// SchedulerHTTP simple version of protoc stub for add scheduler
type SchedulerHTTP struct {
	Name, URL, Executor, Method, Spec string
	Disabled, Persist                 bool
	HTTPHeader                        []HTTPHeader
}

// toProtocStub part of SchedulerHTTP for convert from simple version to protoc stub
func (sh *SchedulerHTTP) toProtocStub(client clockwerkClient) *pbclockwerk.Scheduler {
	var headers []string
	for _, header := range sh.HTTPHeader {
		headers = append(headers, header.String())
	}
	return &pbclockwerk.Scheduler{
		Headers:  headers,
		Username: client.Username,
		Password: client.Password,
		Name:     sh.Name,
		Url:      sh.URL,
		Executor: sh.Executor,
		Method:   sh.Method,
		Disabled: sh.Disabled,
		Persist:  sh.Persist,
		Spec:     sh.Spec,
	}
}

// SchedulerSelect simple version of protoc stub for delete scheduler
type SchedulerSelect struct {
	Id      string
	EntryId int32
}

// toProtocStub part of SchedulerSelect for convert from simple version to protoc stub
func (ss *SchedulerSelect) toProtocStub(client clockwerkClient) *pbclockwerk.SelectScheduler {
	return &pbclockwerk.SelectScheduler{
		Username: client.Username,
		Password: client.Password,
		Id:       ss.Id,
		EntryId:  ss.EntryId,
	}
}

// SchedulerToggle simple version of protoc stub for toggle scheduler
type SchedulerToggle struct {
	Id       string
	EntryId  int32
	Disabled bool
}

// toProtocStub part of SchedulerHTTP for convert from simple version to protoc stub
func (st *SchedulerToggle) toProtocStub(client clockwerkClient) *pbclockwerk.SelectToggle {
	return &pbclockwerk.SelectToggle{
		Username: client.Username,
		Password: client.Password,
		Id:       st.Id,
		EntryId:  st.EntryId,
		Disabled: st.Disabled,
	}
}

func (c clockwerkClient) Add(scheduler SchedulerHTTP) (res SchedulerHTTP, err error) {
	ctx := context.Background()
	defer ctx.Done()
	createdScheduler, err := c.ClockwerkServer.AddScheduler(ctx, scheduler.toProtocStub(c))
	if err != nil {
		return res, err
	}
	var headers []HTTPHeader
	for _, header := range createdScheduler.Headers {
		content := strings.Split(header, "|")
		headers = append(headers, HTTPHeader{K: content[0], V: content[1]})
	}
	res = SchedulerHTTP{
		Name:       createdScheduler.Name,
		URL:        createdScheduler.Url,
		Executor:   createdScheduler.Executor,
		Method:     createdScheduler.Method,
		Spec:       createdScheduler.Spec,
		Disabled:   createdScheduler.Disabled,
		Persist:    createdScheduler.Persist,
		HTTPHeader: headers,
	}
	return res, nil
}

func (c clockwerkClient) Del(schedulerSelect SchedulerSelect) error {
	ctx := context.Background()
	defer ctx.Done()
	_, err := c.ClockwerkServer.DeleteScheduler(ctx, schedulerSelect.toProtocStub(c))
	if err != nil {
		return err
	}
	return nil
}

func (c clockwerkClient) Toggle(schedulerToggle SchedulerToggle) error {
	ctx := context.Background()
	defer ctx.Done()
	_, err := c.ClockwerkServer.ToggleScheduler(ctx, schedulerToggle.toProtocStub(c))
	if err != nil {
		return err
	}
	return nil
}

func NewClockwerk(host, port, username, password string) (ClockwerkClient, error) {
	if host == "" || port == "" {
		host = constant.DefaultHost
		port = constant.DefaultPort
	}
	dialOptions := []grpcgoogle.DialOption{
		grpcgoogle.WithInsecure(),
	}
	connectionString := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpcgoogle.Dial(connectionString, dialOptions...)
	if err != nil {
		return nil, err
	}
	return &clockwerkClient{
		Username:        username,
		Password:        password,
		ClockwerkServer: clockwerktrpt.ClockwerkClient(conn),
	}, nil
}
