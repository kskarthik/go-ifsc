FROM debian:unstable-slim

COPY ./public/linux/ifsc /usr/local/bin/

ENTRYPOINT [ "ifsc", "server" ]

EXPOSE 9000
