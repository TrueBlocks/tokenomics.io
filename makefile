all: **/*.go
	@go build -v
	@mv tokenomics.io nomics
