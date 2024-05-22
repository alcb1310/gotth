run: build
	@./build/goth

build: clean
	@go build -o ./build/goth ./cmd/app

clean:	
	@rm -rf ./build
