.PHONY: generate clean

generate: clean
	go generate .

clean:
	rm -rf *_generated*.go
