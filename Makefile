GOFMT=gofmt
GOCMD=go
GODEP=dep
EXE=gocademy

all:
	$(GODEP) ensure -v

format:
	$(GOFMT) -w .

tests:
	$(GOCMD) test -v ./...

bench:
	$(GOCMD) test -v -bench=^Benchmark -benchmem ./...

build:
	$(GOCMD) build -o $(EXE) ./*.go

run: build
	./$(EXE) --help