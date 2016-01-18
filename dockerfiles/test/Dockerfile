FROM golang:1.5-wheezy
MAINTAINER Geovani de Souza <geovanisouza92@gmail.com>

# Use Go vendor experiment
ENV GO15VENDOREXPERIMENT 1

VOLUME /src
WORKDIR /src

COPY environment.sh /
COPY test.sh /

ENTRYPOINT ["/test.sh"]
