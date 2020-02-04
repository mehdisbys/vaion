test:
	go test -failfast -v -cover -tags=test ./...

build:
	 go build -o vaion .