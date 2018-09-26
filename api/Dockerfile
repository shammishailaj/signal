FROM golang:1.11 AS builder
WORKDIR /signal
COPY . ./
RUN make static

####################################################################################################
## Image
####################################################################################################

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh-client ca-certificates

COPY --from=builder /signal/dist/signal /bin/signal

RUN adduser -D -g '' astrocorp

USER astrocorp

WORKDIR /signal

CMD ["signal"]
