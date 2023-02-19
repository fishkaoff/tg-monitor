include preferences.env

start: 
	go build -o ./bin ./cmd/main.go 
	./bin/main