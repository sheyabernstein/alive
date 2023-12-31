ARG GIN_MODE="release"

FROM alpine as base

ARG GIN_MODE

ENV PORT=4444 \
    GIN_MODE=$GIN_MODE

RUN apk add --no-cache curl && \
    rm -rf /var/cache/apk/*

FROM golang as app-builder

WORKDIR /builder

COPY alive/go.mod alive/go.sum ./
RUN go mod download

COPY alive .

RUN go build \
  -ldflags "-linkmode external -extldflags -static" \
  -o server .
  

FROM base as app

COPY --from=app-builder --chown=nobody:nobody /builder/server .

USER nobody:nobody

EXPOSE $PORT

CMD ["./server"]
