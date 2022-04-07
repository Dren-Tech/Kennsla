build:
	go build -o server server.go

test:
	go test -v ./test/...

run: build
	./server

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' go run server.go

benchmark:
	go test -v ./test/... -run=NONE -bench=. -benchmem -count=6