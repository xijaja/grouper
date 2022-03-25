all:

default:
	go build -o main main.go
	appify -name "grouper" -icon ../resource/app.png main
	rm -rf main

windows:
	./build_win.sh
	x86_64-w64-mingw32-windres grouper.rc -O coff -o grouper.syso

	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 \
	go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -p 4 -v -o grouper.exe

	rm grouper.syso
	rm grouper.rc
