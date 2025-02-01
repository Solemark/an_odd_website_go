# An Odd Website - Go

run code with:

    go run src/main.go

run tests with:

    go test
    OR
    go test -v

build with:
	
	go build src/main.go
	OR
	go build -ldflags -w src/name.go

-ldflags -w removes debug symbols resulting in a smaller binary

target another system with:

	GOOS=platform GOARCH=system go build src/name.go
	e.g
	GOOS=windows GOARCH=amd64 go build src/name.go
