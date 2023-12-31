FROM alpine as base

ENV GIN_MODE="release" \
    PORT=4444
    
RUN apk add --no-cache curl   


FROM golang as app-builder

WORKDIR /builder

COPY alive/go.mod alive/go.sum ./
RUN go mod download

COPY alive .

RUN go build -o server .
RUN go build \
  -ldflags "-linkmode external -extldflags -static" \
  -o server .
  

FROM base as app

COPY --from=app-builder --chown=nobody:nobody /builder/server .

USER nobody:nobody

EXPOSE $PORT

CMD ["./server"]
