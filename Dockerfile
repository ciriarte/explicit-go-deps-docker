FROM golang:buster


RUN apt-get update && \
    apt-get install build-essential -y
COPY . /workspace
WORKDIR /workspace
RUN make install-tools
