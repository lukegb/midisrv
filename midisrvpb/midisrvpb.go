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

package midisrvpb

//go:generate protoc --go_out=plugins=grpc:. midisrv.proto

type Note int64

type event uint8

const (
	NoteMiddleC Note = 60

	eventNoteOn  event = 0x90
	eventNoteOff event = 0x80
)

func NoteOn(channel uint8, note Note, velocity uint8) *MIDIEvent {
	return &MIDIEvent{
		Timestamp: -1,
		Status: int64(uint8(eventNoteOn) | channel),
		Data1:  int64(0x7f & note),
		Data2:  int64(0x7f & velocity),
	}
}

func NoteOff(channel uint8, note Note, velocity uint8) *MIDIEvent {
	return &MIDIEvent{
		Timestamp: -1,
		Status: int64(uint8(eventNoteOff) | channel),
		Data1:  int64(0x7f & note),
		Data2:  int64(0x7f & velocity),
	}
}
