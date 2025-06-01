.DEFAULT_GOAL := build

# keeps make from getting confused if a directory or file in your project
# has the same name as one of the listed targets
.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

# after the colon, are the other target must be run before the specified target runs
build: vet
	go build

clean:
	rm hello_world
