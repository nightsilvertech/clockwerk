package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robfig/cron/v3"
	"github.com/soheilhy/cmux"
	ep "gitlab.com/nbdgocean6/clockwerk/endpoints"
	"gitlab.com/nbdgocean6/clockwerk/gvar"
	pb "gitlab.com/nbdgocean6/clockwerk/protocs/api/v1"
	"gitlab.com/nbdgocean6/clockwerk/repository"
	"gitlab.com/nbdgocean6/clockwerk/service"
	"gitlab.com/nbdgocean6/clockwerk/transports"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func ServeGRPC(listener net.Listener, service pb.ClockwerkServer, serverOptions []grpc.ServerOption) error {
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
	log.Println("initialize rest server")

	mux := runtime.NewServeMux()
	err := pb.RegisterClockwerkHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	return http.Serve(listener, mux)
}

func MergeServer(service pb.ClockwerkServer, serverOptions []grpc.ServerOption) {
	port := fmt.Sprintf(":%s", "1929")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
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

func main() {
	cronjb := cron.New()
	repo := repository.NewRepo("35.219.50.46", "6379", "root")
	services := service.NewClockwerk(cronjb, repo)
	services.Backup(context.Background(), nil)
	server := transports.NewClockwerkServer(ep.NewClockwerkEndpoint(services))

	hashByte, err := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("%s_%s",gvar.HashKeyMap, "nobita")
	gvar.SyncMapHashStorage.Store(key, string(hashByte))

	MergeServer(server, nil)
}
