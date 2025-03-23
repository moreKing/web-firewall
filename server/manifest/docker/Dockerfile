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

FROM alpine:latest

WORKDIR /web-firewall

RUN apk add nftables

COPY --from=builder /building/server /web-firewall/server
COPY ./server/resource /web-firewall/resource
COPY --from=building /building/dist /web-firewall/resource/public/html
COPY ./server/manifest/config /web-firewall/manifest/config 
RUN rm -rf /web-firewall/manifest/config/config.yaml 
RUN mv /web-firewall/manifest/config/config.yaml.docker /web-firewall/manifest/config/config.yaml
COPY ./server/resource/db.sqlite3.bak /web-firewall/manifest/config/db.sqlite3

VOLUME [ "/web-firewall/manifest/config", "/web-firewall/log" ]
ENTRYPOINT ["/web-firewall/server"]