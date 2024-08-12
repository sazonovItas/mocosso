ARG GO_VERSION=1.22.2
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build

WORKDIR /src
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.mod,target=go.mod \
  --mount=type=bind,source=go.sum,target=go.sum \
  go mod download -x

ARG VERSION
ARG COMMIT
ARG DATE
ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOARCH=${TARGETARCH} GOOS=${TARGETOS} \ 
  go build -ldflags "-s -w -X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}' -X 'main.date=${DATE}'" -o /bin/sso ./cmd/sso

FROM alpine:3.20.0 AS development

COPY --from=build /bin/sso /bin/

EXPOSE 8080 9090

ENTRYPOINT [ "/bin/sso" ]

FROM alpine:3.20.0 AS release

RUN --mount=type=cache,target=/var/cache/apk \
  apk --update add \
  ca-certificates \
  tzdata \ 
  && \
  update-ca-certificates

ARG UID=10001
RUN adduser -H -D \
  --uid "${UID}" \
  appuser
USER appuser

COPY --from=build /bin/sso /bin/

EXPOSE 8080 9090

ENTRYPOINT [ "/bin/sso" ]
