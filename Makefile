build:
	@go build -o bin/askai

run: build
	@./bin/askai
test:
	@go test -b ./..
