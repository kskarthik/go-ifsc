FROM golang:1.19-alpine as build

RUN mkdir /ifsc/

COPY . /ifsc/

RUN cd /ifsc/ &&\
		sh build.sh

FROM alpine:latest

# create a non-root user to run the app
RUN adduser --disabled-password ifsc-usr

# switch to non-root user
USER ifsc-usr

COPY --from=build /ifsc/public/linux/* /usr/local/bin/

# index the latest ifsc data
RUN ifsc index

ENTRYPOINT [ "ifsc", "server" ]

EXPOSE 9000
