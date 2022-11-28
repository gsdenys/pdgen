
linux:
	GOOS=linux GOARCH=386 go build -o bin/pdgen-1.0.0-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/pdgen-1.0.0-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/pdgen-1.0.0-linux-arm64 main.go


build: linux
	GOOS=windows GOARCH=386 go build -o bin/pdgen-windows-386 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/pdgen-windows-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/pdgen-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/pdgen-darwin-amd64 main.go

dist: build
	echo "Hello"