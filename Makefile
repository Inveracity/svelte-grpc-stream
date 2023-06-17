.PHONY: clean

clean:
	rm -rf internal/proto
	rm -rf src/proto

.PHONY: proto
proto: clean
	buf generate