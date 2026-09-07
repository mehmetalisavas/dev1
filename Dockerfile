FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /dev1 .

FROM alpine:3.20
RUN adduser -D -u 65532 app
USER 65532
COPY --from=build /dev1 /dev1
EXPOSE 8080
ENTRYPOINT ["/dev1"]
