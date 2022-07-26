target/libcgotest.dylib:
	mvn package

target/call_from_go: \
	target/libcgotest.dylib example/main.go
	go build -o target/call_from_go --ldflags "-s -w -linkmode 'external'" ./example/main.go
