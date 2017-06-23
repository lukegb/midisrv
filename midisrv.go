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

// Package midisrv provides a gRPC server which allows MIDI playback.
// This allows separation of cgo-contaminated portions of code from
// pure Go portions, making everyone much happier.
package midisrv

import (
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rakyll/portmidi"

	pb "github.com/lukegb/midisrv/midisrvpb"
)

type server struct {
	s *portmidi.Stream
}

func (s *server) Output(stream pb.MIDIService_OutputServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&empty.Empty{})
		}
		if err != nil {
			return err
		}

		switch d := msg.Data.(type) {
		case *pb.MIDIWriteRequest_ImmediateEvent:
			e := d.ImmediateEvent
			log.Printf("immevent: %#v", e)
			if err := s.s.WriteShort(e.Status, e.Data1, e.Data2); err != nil {
				return err
			}
		case *pb.MIDIWriteRequest_Events:
			es := d.Events.Events
			var mes []portmidi.Event
			for _, e := range es {
				log.Printf("event: %#v", e)
				mes = append(mes, portmidi.Event{
					Timestamp: portmidi.Timestamp(e.Timestamp),
					Status:    e.Status,
					Data1:     e.Data1,
					Data2:     e.Data2,
				})
			}
			if err := s.s.Write(mes); err != nil {
				return err
			}
		case *pb.MIDIWriteRequest_SysExMessage:
			e := d.SysExMessage
			if err := s.s.WriteSysExBytes(portmidi.Timestamp(e.Timestamp), e.Data); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected oneof type %T", msg.Data)
		}
	}
}

func New(s *portmidi.Stream) pb.MIDIServiceServer {
	return &server{
		s: s,
	}
}
