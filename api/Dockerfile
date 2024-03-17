FROM golang:1.21.5-bullseye AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.sum .
COPY go.mod .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY configs ./configs

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /bin/app ./cmd/api

FROM scratch AS runtime
COPY --from=build /bin/app /bin/app
COPY --from=build /src/configs /configs
COPY --from=build /src/pkg/db/migrations /migrations
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080/tcp
CMD ["/bin/app"]
