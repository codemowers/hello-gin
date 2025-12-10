FROM golang:1.24 AS build
WORKDIR /go/src/github.com/codemowers/hello-gin/
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./
RUN go build -tags netgo -ldflags "-linkmode 'external' -extldflags '-static'" -o /go/server .

FROM scratch
WORKDIR /
COPY --from=build /go/server /server
ENV GIN_MODE=release
ENTRYPOINT ["/server"]
