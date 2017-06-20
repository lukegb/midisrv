// Copyright 2017 Luke Granger-Brown. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/lukegb/midisrv/midisrvpb"
	"google.golang.org/grpc"
)

var (
	flagAddr = flag.String("address", "127.0.0.1:8181", "address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*flagAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMIDIServiceClient(conn)

	stream, err := c.Output(context.Background())
	if err != nil {
		log.Fatalf("couldn't Output: %v", err)
	}

	for n := pb.Note(40); n < 80; n++ {
		msg := &pb.MIDIWriteRequest{
			Data: &pb.MIDIWriteRequest_ImmediateEvent{
				pb.NoteOn(0, n, 127),
			},
		}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Send(%v) = %v", msg, err)
		}
		time.Sleep(100 * time.Millisecond)
		msg = &pb.MIDIWriteRequest{
			Data: &pb.MIDIWriteRequest_ImmediateEvent{
				pb.NoteOff(0, n, 127),
			},
		}
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Send(%v) = %v", msg, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv: %v", err)
	}
	log.Println(reply)
}
