nakanosql:
	go build

test: nakanosql
	ruby test.rb

clean:
	rm -rf nakanosql

all: test clean

.PHONY: test clean all
