FROM alpine:3.2
MAINTAINER Geovani de Souza <geovanisouza92@gmail.com>

COPY lambdac /

EXPOSE 8800

ENTRYPOINT ["/lambdac"]
