fmt:
	go fmt ./...

build:
	GOOS=linux   GOARCH=amd64   go build -ldflags="-s -w" -trimpath -o src/gig-linux-x86_64 .
	GOOS=linux   GOARCH=arm64   go build -ldflags="-s -w" -trimpath -o src/gig-linux-arm_64 .
	GOOS=darwin  GOARCH=amd64   go build -ldflags="-s -w" -trimpath -o src/gig-darwin-x86_64 .
	GOOS=darwin  GOARCH=arm64   go build -ldflags="-s -w" -trimpath -o src/gig-darwin-arm_64 .
	GOOS=windows GOARCH=amd64   go build -ldflags="-s -w" -trimpath -o src/gig-windows-x86_64.exe .
	GOOS=windows GOARCH=arm64   go build -ldflags="-s -w" -trimpath -o src/gig-windows-arm_64.exe .

run:
	go run .

help:
	go run . -h
