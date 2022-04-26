# Make the tool
all: **/*.go
	@go build -v
	@mv tokenomics.io nomics
