FROM alpine:3.2
MAINTAINER Geovani de Souza <geovanisouza92@gmail.com>

COPY environment.sh /
COPY stage.sh /
COPY lambdac /

EXPOSE 8800

ENTRYPOINT ["/stage.sh"]
