.PHONY: build
build: clean
	go build -o ./bin/

.PHONY: run
run:
	go run .

.PHONY: clean
clean:
	go clean
	rm -rf ./bin