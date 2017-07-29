pingscan: main.go
	go build

run: pingscan
	time ./pingscan
