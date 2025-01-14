stats:
	@go run . > data
	@python3 histograms.py
run:
	@go run .

build:
	@go build .

dbg:
	@go run -tags debug .

fmt:
	@go fmt ./...

test:
	@go test ./...

lint:
	@golangci-lint run

todos:
	@rg "TODO:"

clean:
	@rm sim
