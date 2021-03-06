FROM alpine:latest AS builder
RUN apk upgrade -U -a && \
          apk upgrade && \
          apk add --update go gcc g++ git ca-certificates curl make
COPY . /upload/
RUN  cd /upload \
  && CGO_ENABLED=0 GO111MODULE=on go build -o /go/bin/app main.go

FROM alpine:latest
# create non root user
RUN addgroup --gid 61000 upload && \
    adduser -S --uid 61000 --ingroup upload upload && \
    apk add --no-cache ca-certificates

# run as non root user
USER upload
# copy app binary from build
COPY --from=builder /go/bin/app /app/app
EXPOSE 8080/tcp
ENTRYPOINT ["/app/app"]
