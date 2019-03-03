nakanosql: main.go
	go build

test: nakanosql
	ruby test.rb

clean:
	rm -rf nakanosql *.db

all: test clean

.PHONY: test clean all
