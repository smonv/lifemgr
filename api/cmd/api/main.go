package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/smonv/lifemgr/api"
	"github.com/smonv/lifemgr/api/firestore"
	"github.com/smonv/lifemgr/api/server"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	ctx = context.Background()
)

func main() {
	firebaseSecret := os.Getenv("FIREBASE_SECRET")
	fbOpts := option.WithCredentialsFile(firebaseSecret)
	fbConf := &firebase.Config{ProjectID: "lifemgr-dev"}

	app, err := firebase.NewApp(ctx, fbConf, fbOpts)
	if err != nil {
		log.Fatalln(err)
	}

	fs, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ns := firestore.NewNutrientStore(fs)

	grpcServer := grpc.NewServer()
	api.RegisterAPIServer(grpcServer, server.New(ns))

	mux := http.NewServeMux()

	dopts := []grpc.DialOption{grpc.WithInsecure()}
	gwmux := runtime.NewServeMux()
	err = api.RegisterAPIHandlerFromEndpoint(ctx, gwmux, ":8080", dopts)
	if err != nil {
		log.Fatalln(err)
	}

	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: grpcHandlerFunc(grpcServer, mux),
	}

	err = srv.Serve(conn)

	if err != nil {
		log.Fatalln(err)
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
