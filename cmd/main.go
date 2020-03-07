package main

import (
	"flag"
	"github.com/getsentry/sentry-go"
	"log"
	"strconv"
	"undisclosedstudios.infrandom.net/renegades/gameservers/go/pkg"
)

var (
	port = flag.Int("port", 10000, "Server port")
)

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://92b3b35240a04680a7ca4e880c95ced4@sentry.io/3941878",
	})
	if err != nil {
		log.Fatalf("sentry.init: %s", err)
	}
}

func main() {
	flag.Parse()
	lis, grpcServer := pkg.Create(strconv.Itoa(*port))
	//pb.RegisterRouteGuideServer(grpcServer, &user{})
	_ = grpcServer.Serve(lis)
}
