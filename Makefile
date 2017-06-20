.PHONY: clean all

all: cmd/midisrv/midisrv.exe cmd/midicli/midicli.exe

clean:
	rm midisrvpb/midisrv.pb.go || true
	rm cmd/midisrv/midisrv.exe || true
	rm cmd/midicli/midicli.exe || true

midisrvpb/midisrv.pb.go: midisrvpb/midisrv.proto
	cd midisrvpb && go generate

cmd/midisrv/midisrv.exe: cmd/midisrv/midisrv.go midisrv.go midisrvpb/midisrv.pb.go midisrvpb/midisrvpb.go
	cd cmd/midisrv && go build --ldflags '-extldflags "-static"'

cmd/midicli/midicli.exe: cmd/midicli/midicli.go midisrvpb/midisrv.pb.go midisrvpb/midisrvpb.go
	cd cmd/midicli && go build
