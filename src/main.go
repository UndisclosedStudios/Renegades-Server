package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/route_guide/routeguide"
	"log"
	"net"
)

var (
	port = flag.Int("port", 10000, "Server port")
)

type user struct {
	username string
}

func (s *user) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	panic("implement me")
}

func (s *user) ListFeatures(rectangle *pb.Rectangle, server pb.RouteGuide_ListFeaturesServer) error {
	panic("implement me")
}

func (s *user) RecordRoute(server pb.RouteGuide_RecordRouteServer) error {
	panic("implement me")
}

func (s *user) RouteChat(server pb.RouteGuide_RouteChatServer) error {
	panic("implement me")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to bind %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &user{})
	_ = grpcServer.Serve(lis)
}
