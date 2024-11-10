ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION} as build

RUN --mount=type=cache,target=/go/pkg/mod/ \
  GOBIN=/bin go install -tags="no_mysql no_clickhouse no_libsql no_vertica no_ydb" \
  github.com/pressly/goose/v3/cmd/goose@latest 

FROM alpine:3.20.0 as development

COPY --from=build /bin/goose /bin/goose

ENTRYPOINT [ "goose" ]
