GOFMT=gofmt
GOCMD=go

format:
	$(GOFMT) -w .

tests:
	$(GOCMD) test -v ./...