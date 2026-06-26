FROM golang:1.24-bookworm AS build

ARG VERSION=dev

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w -X main.version=${VERSION}" \
    -o /out/luago \
    ./cmd/luago

FROM scratch

COPY --from=build /out/luago /luago

ENTRYPOINT ["/luago"]
CMD ["--version"]
