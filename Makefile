default: rmm

.PHONY: rmm
rmm:
	@echo "Building Recon MindMap (RMM) binary to './rmm'"
	@(cd cmd/rmm; CGO_ENABLED=0 go build --ldflags "-s -w" -o ../../rmm)

clean:
	@echo "Cleaning up all the generated files"
	@find . -name '*.test' | xargs rm -fv
	@find . -name '*~' | xargs rm -fv
	@rm -rvf rmm

docker:
	@docker build -t alevsk/rmm .

test:
	@go test ./...
