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
	Del()
	Toggle()
}

type clockwerkClient struct {
	Username        string
	Password        string
	ClockwerkServer clockwerksvc.Clockwerk
}

//_, err = s.repo.Scheduler.AddScheduler(ctx, &pbclockwerk.Scheduler{
//	Username: username,
//	Password: password,
//	Name:     fmt.Sprintf("update program id %s status to inactivate", programID),
//	Url:      fmt.Sprintf(programconst.UpdateProgramStatusURL+"%s/%d", programID, programconst.ProgramInactive),
//	Executor: schedulerexecutor.HTTP,
//	Method:   schedulerhttpmethod.MethodPut,
//	Disabled: false,
//	Persist:  false,
//	Spec:     fmt.Sprintf("0 0 %d %d *", endAtTime.Day(), int(endAtTime.Month())),
//	Headers: []string{
//		"Content-Type|application/json",
//		fmt.Sprintf("Authorization|Basic %s:%s", systemUsername, systemPassword),
//	},
//})

type HTTPHeader struct {
	K, V string
}

func (hh *HTTPHeader) String() string {
	return fmt.Sprintf("%s|%s", hh.K, hh.V)
}

type SchedulerHTTP struct {
	Name, URL, Executor, Method, Spec string
	Disabled, Persist                 bool
	HTTPHeader                        []HTTPHeader
}

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

func (c clockwerkClient) Del() {
	panic("implement me")
}

func (c clockwerkClient) Toggle() {
	panic("implement me")
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
