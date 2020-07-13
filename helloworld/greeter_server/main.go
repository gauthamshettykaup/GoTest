/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func (s *server) GauthamList(ctx context.Context, in *pb.PersonRequest) (*pb.AddressBook, error) {
	var releases []*pb.Person
	ri := &pb.Person{
		Name:  "Test",
		Id:    101,
		Email: "john@example.com",
	}
	releases = append(releases, ri)

	ri = &pb.Person{
		Name:  "Jack Buck",
		Id:    301,
		Email: "buck@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-555-0000", Type: pb.Person_HOME},
			{Number: "555-555-0001", Type: pb.Person_MOBILE},
			{Number: "555-555-0002", Type: pb.Person_WORK},
		},
	}
	releases = append(releases, ri)

	// output := pb.AddressBook{People: []*pb.Person{
	// 	{
	// 		Name:  "John Doe",
	// 		Id:    101,
	// 		Email: "john@example.com",
	// 	},
	// 	{
	// 		Name: "Jane Doe",
	// 		Id:   102,
	// 	},
	// 	{
	// 		Name:  "Jack Doe",
	// 		Id:    201,
	// 		Email: "jack@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-555-5555", Type: pb.Person_WORK},
	// 		},
	// 	},
	// 	{
	// 		Name:  "Jack Buck",
	// 		Id:    301,
	// 		Email: "buck@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-555-0000", Type: pb.Person_HOME},
	// 			{Number: "555-555-0001", Type: pb.Person_MOBILE},
	// 			{Number: "555-555-0002", Type: pb.Person_WORK},
	// 		},
	// 	},
	// 	{
	// 		Name:  "Janet Doe",
	// 		Id:    1001,
	// 		Email: "janet@example.com",
	// 		Phones: []*pb.Person_PhoneNumber{
	// 			{Number: "555-777-0000"},
	// 			{Number: "555-777-0001", Type: pb.Person_HOME},
	// 		},
	// 	},
	// }}

	return &pb.AddressBook{
		People: releases,
	}, nil

	// return &pb.AddressBook{
	// 	People: output.People,
	// }, nil
	// return &output, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
