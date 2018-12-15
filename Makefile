GOFMT=gofmt
GOCMD=go

format:
	$(GOFMT) -w .

tests:
	$(GOCMD) test -v ./...

bench:
	$(GOCMD) test -v -bench=^Benchmark ./...