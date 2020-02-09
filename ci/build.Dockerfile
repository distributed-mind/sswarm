# golang
FROM golang:latest as build
SHELL ["/bin/bash", "-c"]

WORKDIR /build
ADD . /build
RUN go build \
    -a \
    -o /build/sswarm \
    -ldflags="-X 'sswarm/cli.Version=git-$(git rev-parse --short HEAD)'" \
    .

# copy to stdout
FROM golang:latest
COPY --from=build /build/sswarm /opt/
CMD ["cat", "/opt/sswarm"]
