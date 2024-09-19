FROM golang:alpine AS builder

WORKDIR /building
COPY ./server /building

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build .

FROM node:20 AS building
WORKDIR /building
COPY ./web /building
RUN npm install -g pnpm
RUN pnpm install
RUN pnpm build

FROM ubuntu:24.04

WORKDIR /web-firewall

RUN apt-get update && \
    apt-get install -y nftables && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /building/server /web-firewall/server
COPY ./server/resource/public /web-firewall/resource/public
COPY --from=building /building/dist /web-firewall/resource/public/html
COPY ./server/resource/template /web-firewall/resource/template
COPY ./server/manifest/config /web-firewall/manifest/config

# RUN nft -f /web-firewall/manifest/config/nftables.conf

VOLUME [ "/web-firewall/manifest/config" ]
ENTRYPOINT ["/web-firewall/server"]