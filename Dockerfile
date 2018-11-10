FROM alpine:3.8

ADD bin/drone-telegram/bin/

ENTRYPOINT ["/bin/drone-telegram"]