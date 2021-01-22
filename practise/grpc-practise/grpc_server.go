package main

import (
	"golang-learning/practise/gorm-practise/models"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}
