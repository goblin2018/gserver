build: clean
	go build -o build/gserver ./main.go 
run:
	./build/gserver
clean:
	rm -rf build



.PHONY: all test clean build
