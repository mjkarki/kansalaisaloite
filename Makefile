build: main.go initiative.go initiativeinfo.go
	go build -ldflags "-s -w"
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w"

clean:
	rm kansalaisaloite kansalaisaloite.exe
