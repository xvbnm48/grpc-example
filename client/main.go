package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/xvbnm48/grpc-example/student"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	student, err := client.FindStudentByEmail(ctx, &pb.Student{Email: email})
	if err != nil {
		log.Fatalf("error when get student with email: %v", err)
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "sakuramiyawaki@nogi46.com")
}
