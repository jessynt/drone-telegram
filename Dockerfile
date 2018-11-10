FROM scratch

ADD bin/drone-telegram /bin/

ENTRYPOINT ["/bin/drone-telegram"]