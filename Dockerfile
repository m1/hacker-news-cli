FROM golang:1.11-alpine as builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
ENV CGO_ENABLED=1
ENV GOHOSTARCH="amd64"
ENV GOHOSTOS="linux"
ENV GOOS="linux"
ENV GO111MODULE=on

# need all these due to sqlite!
RUN apk add --no-cache ca-certificates git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base
RUN apk add --no-cache gcc g++ libc-dev

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /hn .

FROM scratch AS final
COPY --from=builder /hn /hn

CMD ["/hn"]