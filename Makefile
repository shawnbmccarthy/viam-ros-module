viamrosmodule: cmd/module/cmd.go
	go build -o rosmodule cmd/module/cmd.go

test:
	go test

lint:
	gofmt -w -s .
