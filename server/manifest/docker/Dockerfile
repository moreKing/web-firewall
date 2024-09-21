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
RUN rm /web-firewall/manifest/config/config.yaml && mv /web-firewall/manifest/config/config.yaml.docker /web-firewall/manifest/config/config.yaml

VOLUME [ "/web-firewall/manifest/config", "/web-firewall/log" ]
ENTRYPOINT ["/web-firewall/server"]