all:
	make default
	make linux
	make windows

default:
	go build -o grouper main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o grouper main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o grouper.exe main.go
