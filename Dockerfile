FROM golang:1.19-alpine as build

RUN mkdir /ifsc/

COPY . /ifsc/

ARG IFSC_VERSION=v2.0.12

RUN cd /ifsc/ &&\
		wget https://github.com/razorpay/ifsc/releases/download/$IFSC_VERSION/IFSC.csv -P cmd/ &&\
		sh build.sh

FROM alpine:latest

# create a non-root user to run the app
RUN adduser --disabled-password ifsc-usr

# switch to non-root user
USER ifsc-usr

COPY --from=build /ifsc/public/linux/* ~/

ENTRYPOINT [ "~/ifsc", "server" ]

EXPOSE 9000
