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
# Build Python environment
# FROM python:3.10.5-alpine as python
# COPY --from=builder /root/tokenomics.io/gitcoin/exports/mainnet/neighbors/networks /app
# WORKDIR /app
# ENV DEBIAN_FRONTEND=noninteractive
# # RUN apt update && apt install -y gcc
# RUN apk add --update gcc g++ libc-dev lld py3-scipy
# RUN python -m ensurepip
# RUN python -m venv --copies /venv
# ENV PATH="/venv/bin:$PATH"
# RUN . /venv/bin/activate
# RUN pip install -Iv setuptools==59.8.0
# RUN pip install -r requirements.fil

#
# Switch to alpine container
FROM python:3.10.5-slim
 #alpine:3.16
ENV DEBIAN_FRONTEND=noninteractive

# Install binary dependencies and nice to haves
# RUN apk --no-cache add gzip bash curl jq yarn nano
RUN apt update && apt install -y gzip curl jq yarn nano python3

# Install Python and copy dependencies
ENV PYTHONUNBUFFERED=1
#RUN apk add --no-cache python3=3.10.4-r0 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
# COPY --from=python /venv /venv
RUN python -m venv /venv
ENV PATH="/venv/bin:$PATH"
RUN . /venv/bin/activate
# RUN pip install -Iv setuptools==59.8.0

# Copy files from builder
COPY --from=builder /root/tokenomics.io/nomics /usr/local/bin
COPY --from=builder /root/tokenomics.io /root/tokenomics.io

RUN pip install -r /root/tokenomics.io/gitcoin/exports/mainnet/neighbors/networks/requirements.fil
ENV PATH="/venv/bin:$PATH"

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
ENTRYPOINT bash ./docker/build/tokenomics.entrypoint.sh
