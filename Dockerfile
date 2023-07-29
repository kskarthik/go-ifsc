FROM golang:1.19-alpine as build

RUN mkdir /ifsc/

COPY . /ifsc/

RUN cd /ifsc/ && sh build.sh

FROM alpine:latest

COPY --from=build /ifsc/public/linux/* /usr/local/bin/

# create a non-root user to run the app & index the data
RUN adduser --disabled-password ifsc-usr && ifsc index

USER ifsc-usr

ENTRYPOINT [ "ifsc", "server" ]

EXPOSE 9000
