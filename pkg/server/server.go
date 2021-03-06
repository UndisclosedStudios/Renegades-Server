package server

import (
  log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

//Create function create and return a listener object and server object based on the port given.
func Create(port string) (net.Listener, *grpc.Server) {
	//TODO: Find a better way to pass "127.0.0.1:"
	lis, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		log.WithFields(log.Fields{
			"port": port,
		}).WithError(err).Fatal("Failed to bind port !")
		log.Print("Trying to bind onto another port !")
	}
	return lis, grpc.NewServer()
}
