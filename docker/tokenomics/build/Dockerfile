#
# Latest golang as builder
FROM golang:1.18-alpine as builder

# Install build depenedencies
RUN apk --no-cache add make git gcc linux-headers libc-dev

# Set workdir
WORKDIR /root

# Try to get upstream version (default master)
ARG UPSTREAM_VER=master

# Clone and build tokenomics tool
RUN git clone -b "${UPSTREAM_VER}" --single-branch --progress --depth 1 \
        https://github.com/TrueBlocks/tokenomics.io.git  && \
        cd tokenomics.io && make

#
# Switch to alpine container
FROM alpine:3.16

# Install binary dependencies and nice to haves
RUN apk --no-cache add gzip bash curl jq yarn nano

# Copy files from builder
COPY --from=builder /root/tokenomics.io/nomics /usr/local/bin
COPY --from=builder /root/tokenomics.io /root/tokenomics.io

# Copy config file
COPY trueBlocks.toml /root/.local/share/trueblocks/

WORKDIR /root/tokenomics.io

# Testing
# COPY test/core-test.entrypoint.sh /root
# RUN chmod +x /root/core-test.entrypoint.sh

# Expose ports
# ARG SERVE_PORT=8080
# EXPOSE ${SERVE_PORT}

# Run entrypoint script
ENTRYPOINT bash ./docker/tokenomics/build/tokenomics.entrypoint.sh
