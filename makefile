.PHONY: all clean test main_wasm

# Target: all (default target)
all: clean build main_wasm test

# Target: build
build:
	cd snapsdk && go build -o ../build/snapsdk

# Target: test
test:
	mkdir -p test
	./build/snapsdk -o example ./example/dogs.yaml '*'

# Target: main_wasm
main_wasm:
#cd snapsdk && tinygo build -o ../build/snapsdk.wasm -target ../wasm.json  ./
	cd snapsdk && GOOS=js GOARCH=wasm go build -o ../build/snapsdk.wasm main_wasm.go types.go generator.go markdown.go javascript.go rust.go golang.go python.go

# Target: clean
clean:
	rm -rf build/ snapsdk.wasm
	rm -rf test/