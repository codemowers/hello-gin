FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd/
COPY templates ./templates/
COPY static ./static/
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/server ./cmd

FROM scratch
WORKDIR /
COPY --from=build /app/server /server
COPY --from=build /app/templates /templates
COPY --from=build /app/static /static
ENV GIN_MODE=release
EXPOSE 8000 8080
ENTRYPOINT ["/server"]
