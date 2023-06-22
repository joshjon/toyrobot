FROM golang:1.20-bullseye as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN --mount=type=cache,target=/opt/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /go/bin/app ./cmd/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /
ENTRYPOINT ["/app"]
