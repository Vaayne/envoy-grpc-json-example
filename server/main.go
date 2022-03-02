package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"server/pb/bookstore"
	"server/pb/helloworld"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	host = flag.String("host", "0.0.0.0", "host")
	port = flag.Int("port", 8001, "The server port")
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
	return h2c.NewHandler(handler, &http2.Server{})
}

func grpcSvc() (*grpc.Server, error) {
	svc := grpc.NewServer()
	helloworld.RegisterGreeterServer(svc, &helloworldServer{})
	bookstore.RegisterBookStoreServer(svc, &bookstoreServer{})
	reflection.Register(svc)
	return svc, nil
}

func httpSvc(ctx context.Context, addr string) (*http.ServeMux, error) {
	var err error

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = helloworld.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, addr, opts); err != nil {
		return nil, err
	}
	if err = bookstore.RegisterBookStoreHandlerFromEndpoint(ctx, gwmux, addr, opts); err != nil {
		return nil, err
	}

	mux.Handle("/", gwmux)

	return mux, nil
}

func run(addr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux, err := httpSvc(ctx, addr)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	grpcServer, err := grpcSvc()
	if err != nil {
		log.Fatalf("Failed to set grpc server: %v", err)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: grpcHandlerFunc(grpcServer, mux),
	}
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}
	log.Printf("Server started on %s", addr)
	err = srv.Serve(conn)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)
	run(addr)
}
