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
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/lukegb/midisrv"
	pb "github.com/lukegb/midisrv/midisrvpb"
	"github.com/rakyll/portmidi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	flagAddr         = flag.String("address", ":8181", "address to listen on")
	flagListDevices  = flag.Bool("list-devices", false, "list portmidi devices and exit")
	flagDeviceNumber = flag.Int("device", -1, "portmidi device number to use for output (-1 for default output device)")
	flagBufferSize   = flag.Int64("buffer-size", 0, "portmidi buffer size")
	flagLatency      = flag.Int64("latency", 0, "portmidi latency")
)

func main() {
	flag.Parse()

	if err := portmidi.Initialize(); err != nil {
		log.Fatalf("failed to initialize portmidi: %v", err)
	}

	if *flagListDevices {
		for n := 0; n < portmidi.CountDevices(); n++ {
			dev := portmidi.Info(portmidi.DeviceID(n))
			fmt.Printf("%d:\n\tInterface:\t%s\n\tName:\t%s\n\tIsInputAvailable:\t%v\n\tIsOutputAvailable:\t%v\n", n, dev.Interface, dev.Name, dev.IsInputAvailable, dev.IsOutputAvailable)
		}
		return
	}

	devNum := portmidi.DeviceID(*flagDeviceNumber)
	if devNum == -1 {
		devNum = portmidi.DefaultOutputDeviceID()
	}

	os, err := portmidi.NewOutputStream(devNum, *flagBufferSize, *flagLatency)
	if err != nil {
		log.Fatalf("failed to open output stream: %v", err)
	}
	msrv := midisrv.New(os)

	lis, err := net.Listen("tcp", *flagAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", *flagAddr)

	srv := grpc.NewServer()
	pb.RegisterMIDIServiceServer(srv, msrv)
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
