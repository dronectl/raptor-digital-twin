FROM alpine:latest

LABEL version="1.0"
LABEL Christian Sargusingh "christian@sargusingh.ca"

RUN apk add gcc make cmake musl-dev

WORKDIR /app

COPY . .

RUN cmake -B build && \
    make -C build -j

ENTRYPOINT [ "build/raptor-dt" ]

