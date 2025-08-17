BINARY_NAME=verso
OUTPUT_DIR=bin

all: build test
 
build:
	go build -o ./${OUTPUT_DIR}/${BINARY_NAME} main.go
 
test:
	go test -json ./... > ./${OUTPUT_DIR}/test_report.json && cat ./${OUTPUT_DIR}/test_report.json | jq '. | select(.Action == "fail")'
 
run: build test
	./${OUTPUT_DIR}/${BINARY_NAME}
 
clean:
	go clean
	rm ./${OUTPUT_DIR}/${BINARY_NAME}
