package client

import (
	"context"
	"fmt"
	"github.com/nightsilvertech/clockwerk/constant"
	"github.com/nightsilvertech/clockwerk/executors"
	"github.com/nightsilvertech/clockwerk/executors/http"
	pbclockwerk "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	clockwerksvc "github.com/nightsilvertech/clockwerk/service/interface"
	clockwerktrpt "github.com/nightsilvertech/clockwerk/transports"
	grpcgoogle "google.golang.org/grpc"
	"strings"
)

// executor private type for scheduler struct
type executor string

// String converts executor type into string
func (e executor) String() string {
	return string(e)
}

const (
	HTTP  executor = executors.HTTP
	SHELL executor = executors.SHELL
)

// method private type for scheduler struct
type method string

// String converts method type into string
func (m method) String() string {
	return string(m)
}

const (
	PUT    method = http.MethodPut
	GET    method = http.MethodGet
	POST   method = http.MethodPost
	DELETE method = http.MethodDelete
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
	Id,
	ReferenceId,
	Name,
	URL,
	Spec,
	Body string
	Executor          executor
	Method            method
	EntryId           int32
	Retry             int32
	RetryThreshold    int32
	Disabled, Persist bool
	HTTPHeader        []HTTPHeader
}

// toProtocStub part of SchedulerHTTP for convert from simple version to protoc stub
func (sh *SchedulerHTTP) toProtocStub(client clockwerkClient) *pbclockwerk.Scheduler {
	var headers []string
	for _, header := range sh.HTTPHeader {
		headers = append(headers, header.String())
	}
	return &pbclockwerk.Scheduler{
		Username:       client.Username,
		Password:       client.Password,
		ReferenceId:    sh.ReferenceId,
		Name:           sh.Name,
		Url:            sh.URL,
		Executor:       sh.Executor.String(),
		Method:         sh.Method.String(),
		Spec:           sh.Spec,
		Body:           sh.Body,
		Disabled:       sh.Disabled,
		Persist:        sh.Persist,
		Retry:          sh.Retry,
		RetryThreshold: sh.RetryThreshold,
		Headers:        headers,
	}
}

// SchedulerSelect simple version of protoc stub for delete scheduler
type SchedulerSelect struct {
	Id          string
	ReferenceId string
}

// toProtocStub part of SchedulerSelect for convert from simple version to protoc stub
func (ss *SchedulerSelect) toProtocStub(client clockwerkClient) *pbclockwerk.SelectScheduler {
	return &pbclockwerk.SelectScheduler{
		Username:    client.Username,
		Password:    client.Password,
		Id:          ss.Id,
		ReferenceId: ss.ReferenceId,
	}
}

// SchedulerToggle simple version of protoc stub for toggle scheduler
type SchedulerToggle struct {
	Id          string
	ReferenceId string
	Disabled    bool
}

// toProtocStub part of SchedulerHTTP for convert from simple version to protoc stub
func (st *SchedulerToggle) toProtocStub(client clockwerkClient) *pbclockwerk.SelectToggle {
	return &pbclockwerk.SelectToggle{
		Username:    client.Username,
		Password:    client.Password,
		Id:          st.Id,
		ReferenceId: st.ReferenceId,
		Disabled:    st.Disabled,
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
		Id:             createdScheduler.Id,
		EntryId:        createdScheduler.EntryId,
		ReferenceId:    createdScheduler.ReferenceId,
		Name:           createdScheduler.Name,
		URL:            createdScheduler.Url,
		Executor:       executor(createdScheduler.Executor),
		Method:         method(createdScheduler.Method),
		Spec:           createdScheduler.Spec,
		Body:           createdScheduler.Body,
		Disabled:       createdScheduler.Disabled,
		Persist:        createdScheduler.Persist,
		Retry:          createdScheduler.Retry,
		RetryThreshold: createdScheduler.RetryThreshold,
		HTTPHeader:     headers,
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
