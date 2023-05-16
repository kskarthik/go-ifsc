FROM golang:1.19-alpine as build

RUN mkdir /ifsc/

COPY . /ifsc/

RUN cd /ifsc/ && ls && sh build.sh 

FROM alpine:latest

COPY --from=build /ifsc/public/linux/* /usr/local/bin/

ENTRYPOINT [ "ifsc", "server" ]

EXPOSE 9000
