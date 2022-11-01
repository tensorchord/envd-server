FROM ubuntu:20.04
RUN apt update; \
    apt install -y --no-install-recommends \
        ca-certificates \
        curl; \
    mkdir -p /usr/local/share/ca-certificates; \
    rm -rf /var/lib/apt/lists/*
COPY envd-server /
