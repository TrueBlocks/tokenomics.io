# Make the tool
all:
	@echo making ./nomics
	@go build -v
	@mv tokenomics.io nomics
