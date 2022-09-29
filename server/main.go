package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"sync"

	pb "github.com/xvbnm48/grpc-example/student"
	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		log.Fatalf("failed to read data: %v", err.Error())
	}
	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalf("failed to unmarshal data: %v", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	log.Println("Listening on port :8801")

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("error when server grpc: %v", err.Error())
	}
}
