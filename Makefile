.PHONY: clean
clean:
	rm -rf relay/internal/proto
	rm -rf frontend/src/proto
	rm -rf api/proto

.PHONY: proto
proto: clean
	buf generate
