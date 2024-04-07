# syntax=docker/dockerfile:1
ARG GO_VERSION=1.21.6
ARG GOLANGCI_LINT_VERSION=v1.55

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine as build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download -x

RUN --mount=type=bind,target=. \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build --ldflags "-s" -o /bin/send ./cmd/

FROM scratch as binary
COPY --from=build /bin/send /bin/

FROM golangci/golangci-lint:${GOLANGCI_LINT_VERSION} as lint
WORKDIR /test
RUN --mount=type=bind,target=. \
    golangci-lint run

from gcr.io/distroless/static
COPY --from=build /bin/send /bin/
ENTRYPOINT [ "/bin/send" ]

