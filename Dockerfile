ARG GOLANG=golang:1.27.2-alpine3.24

FROM --platform=$BUILDPLATFORM ${GOLANG} AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY . .

ARG VERSION=dev
ARG STRIP="-s -w"
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
        -trimpath \
        -ldflags="$STRIP -X main.version=$VERSION" \
        -o /out/wowhabits ./cmd/wowhabits

FROM scratch AS result
COPY --from=build /out/wowhabits /bin/wowhabits
