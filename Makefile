BINARY := fblockchain

all: build run

build:
		@echo "==> Go build"
		@go build -o $(BINARY)

run:
		@echo "==> Running"
		@chmod +x $(BINARY)
		@./$(BINARY)

.PHONY: build run