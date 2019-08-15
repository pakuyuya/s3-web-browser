FROM node:8.16.0-jessie AS client
COPY ./client /work
WORKDIR /work
RUN yarn install && yarn build

FROM node:8.16.0-jessie AS client-login
COPY ./client-login /work
WORKDIR /work
RUN yarn install && yarn build

FROM golang:1.12.8-alpine3.9 AS server
COPY ./server /work
WORKDIR /work
RUN apk update && apk add git && go build

FROM alpine:3.9
COPY --from=client /work/dist /usr/local/src/client
COPY --from=client-login /work/dist /usr/local/src/client-login
COPY --from=server /work/server /usr/local/app/server
COPY --from=server /work/setting.yml /usr/local/app/setting.yml

RUN mkdir -p /usr/local/app/static && \
    cp -pfr /usr/local/src/client/favicon.ico /usr/local/app/static && \
    cp -pfr /usr/local/src/client/js /usr/local/app/static && \
    cp -pfr /usr/local/src/client/css /usr/local/app/static && \
    cp -pfr /usr/local/src/client-login/js /usr/local/app/static && \
    cp -pfr /usr/local/src/client-login/css /usr/local/app/static && \
    mkdir -p /usr/local/app/templates && \
    cp -pf /usr/local/src/client/index.html /usr/local/app/templates/browser.html && \
    cp -pf /usr/local/src/client-login/index.html /usr/local/app/templates/login.html && \
    chmod +x /usr/local/app/server

WORKDIR /usr/local/app
ENTRYPOINT ["./server"]
