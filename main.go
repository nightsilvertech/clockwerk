package main

import (
	"context"
	"fmt"
	"github.com/nightsilvertech/clockwerk/constant"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

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

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := []string{"Content-Type", "Accept", "Authorization","Access-Control-Allow-Headers","X-Requested-With"}
		methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}

		h.ServeHTTP(w, r)
	})
}

func ServeGRPC(listener net.Listener, service pb.ClockwerkServer, serverOptions []grpc.ServerOption) error {
	log.Println(fmt.Sprintf("grpc served at port %s", port))
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
	log.Println(fmt.Sprintf("http served at port %s", port))
	mux := runtime.NewServeMux()
	err := pb.RegisterClockwerkHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	return http.Serve(listener, CORS(mux))
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
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	port = "1929"
	username = os.Getenv("SCHEDULER_USERNAME")
	password = os.Getenv("SCHEDULER_PASSWORD")
	if len(username) == 0 || len(password) == 0 {
		panic("please set your credential")
	}

	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")
	redisPass = os.Getenv("REDIS_PASS")
	if len(redisHost) == 0 || len(redisPort) == 0 || len(redisPass) == 0 {
		panic("please set your redis host, port and password")
	}

	log.Println(fmt.Sprintf("connect to redis at %s port %s", redisHost, redisPort))
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
