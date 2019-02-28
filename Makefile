nakanosql:
	go build

test: nakanosql
	ruby test.rb

clean:
	git clean -df

all: test clean

.PHONY: test clean all
