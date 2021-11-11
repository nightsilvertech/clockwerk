package main

import (
	"context"
	"fmt"
	"github.com/nightsilvertech/clockwerk/constant"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	ep "github.com/nightsilvertech/clockwerk/endpoints"
	"github.com/nightsilvertech/clockwerk/gvar"
	pb "github.com/nightsilvertech/clockwerk/protocs/api/v1"
	"github.com/nightsilvertech/clockwerk/repository"
	"github.com/nightsilvertech/clockwerk/service"
	"github.com/nightsilvertech/clockwerk/transports"
	"github.com/robfig/cron/v3"
	"github.com/soheilhy/cmux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var redisHost, redisPort, redisPass string
var username, password, port string

func ServeGRPC(listener net.Listener, service pb.ClockwerkServer, serverOptions []grpc.ServerOption) error {
	log.Println("server grpc server")
	var grpcServer *grpc.Server
	if len(serverOptions) > 0 {
		grpcServer = grpc.NewServer(serverOptions...)
	} else {
		grpcServer = grpc.NewServer()
	}
	pb.RegisterClockwerkServer(grpcServer, service)
	return grpcServer.Serve(listener)
}

func ServeHTTP(listener net.Listener, service pb.ClockwerkServer) error {
	log.Println("server http server")
	mux := runtime.NewServeMux()
	err := pb.RegisterClockwerkHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	return http.Serve(listener, mux)
}

func MergeServer(service pb.ClockwerkServer, serverOptions []grpc.ServerOption) {
	if len(port) > 0 {
		port = fmt.Sprintf(":%s", constant.DefaultPort)
	}
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	m := cmux.New(listener)
	httpListener := m.Match(cmux.HTTP1Fast())
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(
		"content-type", "application/grpc",
	))

	g := new(errgroup.Group)
	g.Go(func() error { return ServeGRPC(grpcListener, service, serverOptions) })
	g.Go(func() error { return ServeHTTP(httpListener, service) })
	g.Go(func() error { return m.Serve() })

	log.Fatal(g.Wait())
}

func CreatedCredential(username, password string) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("%s_%s", gvar.HashKeyMap, username)
	gvar.SyncMapHashStorage.Store(key, string(hashByte))
}

func PrepareCredential() {
	port = os.Getenv("PORT")
	username = os.Getenv("SCHEDULER_USERNAME")
	password = os.Getenv("SCHEDULER_PASSWORD")
	if len(username) == 0 || len(password) == 0 {
		panic("please set your credential")
	}
	fmt.Println(username, password)
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
	redisPass = os.Getenv("REDIS_PASS")
	if len(redisHost) == 0 || len(redisPort) == 0 || len(redisPass) == 0 {
		panic("please set your redis host, port and password")
	}
}

func main() {
	PrepareCredential()
	CreatedCredential(username, password)
	cronjb := cron.New()
	repo := repository.NewRepo(redisHost, redisPort, redisPass)
	services := service.NewClockwerk(cronjb, repo)
	_, _ = services.Backup(context.Background(), nil)
	server := transports.NewClockwerkServer(ep.NewClockwerkEndpoint(services))
	MergeServer(server, nil)
}
