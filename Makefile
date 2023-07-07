viamrosmodule: cmd/module/cmd.go
	go build -o viamrosmodule cmd/module/cmd.go

test:
	go test

lint:
	gofmt -w -s .
