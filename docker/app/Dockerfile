FROM node:8.16.0-jessie AS client
COPY ./client /work
WORKDIR /work
RUN yarn install && yarn build
RUN mkdir /output
RUN cp -r /work/dist/js /output/js
RUN cp -r /work/dist/css /output/css
RUN cp /work/dist/favicon.ico /output/favicon.ico
RUN cp /work/dist/*.html /output

FROM golang:1.12.8-alpine3.9 AS server
COPY ./server /work
WORKDIR /work
RUN apk update && apk add git && go build
RUN mkdir /output
RUN cp /work/server /output/server
RUN cp /work/setting.production.yml /output/setting.yml

FROM alpine:3.9
COPY --from=client /output /usr/local/src/client
COPY --from=server /output /usr/local/app

RUN mkdir -p /usr/local/app/static && \
    cp -pfr /usr/local/src/client/favicon.ico /usr/local/app/static && \
    cp -pfr /usr/local/src/client/js /usr/local/app/static && \
    cp -pfr /usr/local/src/client/css /usr/local/app/static && \
    mkdir -p /usr/local/app/templates && \
    cp -pf /usr/local/src/client/*.html /usr/local/app/templates && \
    chmod +x /usr/local/app/server

WORKDIR /usr/local/app
ENTRYPOINT ["./server"]
