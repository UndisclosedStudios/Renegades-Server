package main

import (
  "flag"
  "github.com/getsentry/sentry-go"
  env "github.com/joho/godotenv"
  "log"
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
  log.Print("Starting ...")
	flag.Parse()
  log.Printf("Binded to port %u ! ", flag.Parse)
	lis, grpcServer := server.Create(strconv.Itoa(*port))
	log.Print("Server created !")
	//pb.RegisterRouteGuideServer(grpcServer, &user{})
	_ = grpcServer.Serve(lis)
  log.Print("Server launched !")
}
