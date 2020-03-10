package main

import (
  "flag"
  "github.com/getsentry/sentry-go"
  env "github.com/joho/godotenv"
  "os"
  "strconv"
  err "undisclosedstudios.infrandom.net/renegades/gameservers/go/internal/error"
  "undisclosedstudios.infrandom.net/renegades/gameservers/go/pkg/server"
)

var (
	port = flag.Int("port", 10000, "Server port")
)

func init() {
  dotenvErr := env.Load()
	sentryErr := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	err.Checks([]error{dotenvErr, sentryErr})
}

func main() {
	flag.Parse()
	lis, grpcServer := server.Create(strconv.Itoa(*port))
	//pb.RegisterRouteGuideServer(grpcServer, &user{})
	_ = grpcServer.Serve(lis)
}
