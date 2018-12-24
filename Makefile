GOFMT=gofmt
GOCMD=go
GODEP=dep

all:
	$(GODEP) ensure -v

format:
	$(GOFMT) -w .

tests:
	$(GOCMD) test -v ./...

bench:
	$(GOCMD) test -v -bench=^Benchmark -benchmem ./...