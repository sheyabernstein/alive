ARG GIN_MODE="release"

FROM busybox AS base

ARG GIN_MODE

ENV PORT=4444 \
    GIN_MODE=$GIN_MODE

FROM golang AS app-builder

WORKDIR /builder

COPY alive/go.mod alive/go.sum ./
RUN go mod download

COPY alive .

RUN go build \
  -ldflags "-linkmode external -extldflags -static -s -w" \
  -o server .
  

FROM base AS app

COPY --from=app-builder --chown=nobody:nobody /builder/server .

USER nobody:nobody

EXPOSE $PORT

CMD ["./server"]
