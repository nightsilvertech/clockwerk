package transports

//import (
//	"context"
//	"fmt"
//	grpctransport "github.com/go-kit/kit/transport/grpc"
//	"google.golang.org/grpc/metadata"
//)
//
//type contextClockwerkKey string
//
//const (
//	clockwerkContextBasicAuth contextClockwerkKey = `basic-auth`
//	clockwerkContextMetadata  contextClockwerkKey = `x-clockwerk-metadata`
//)
//
//func CtxBasicAuth(ctx context.Context, username, password string) context.Context {
//	return context.WithValue(ctx, clockwerkContextBasicAuth, fmt.Sprintf("%s:%s", username, password))
//}
//
//func BasicAuthMetadataToContext() grpctransport.ServerRequestFunc {
//	return func(ctx context.Context, md metadata.MD) context.Context {
//		requestID, ok := md["authorization"]
//		if !ok {
//			return ctx
//		}
//		if ok {
//			ctx = context.WithValue(ctx, clockwerkContextBasicAuth, requestID[0])
//		}
//		return ctx
//	}
//}
//
//func ContextToBasicAuthMetadata() grpctransport.ClientRequestFunc {
//	return func(ctx context.Context, md *metadata.MD) context.Context {
//		requestID, ok := ctx.Value(clockwerkContextBasicAuth).(string)
//		if ok {
//			(*md)["authorization"] = []string{requestID}
//		}
//		return ctx
//	}
//}
