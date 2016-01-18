FROM golang:1.5-wheezy
MAINTAINER Geovani de Souza <geovanisouza92@gmail.com>

# Install dependencies
RUN apt-get update -qq && \
    apt-get install -y upx-ucl && \
    rm -rf /var/lib/apt

# Install docker client
RUN wget -nv https://get.docker.com/builds/Linux/x86_64/docker-1.9.1 -O /usr/bin/docker && \
    chmod +x /usr/bin/docker

# Use Go vendor experiment
ENV GO15VENDOREXPERIMENT 1

# Install Ultimate Packer for eXecutables
RUN go get github.com/pwaller/goupx

VOLUME /src
WORKDIR /src

COPY environment.sh /
COPY build.sh /

ENTRYPOINT ["/build.sh"]
