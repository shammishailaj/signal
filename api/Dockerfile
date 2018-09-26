FROM golang:1.11 AS builder
RUN useradd -u 10001 astrocorp
WORKDIR /signal
COPY . ./
RUN make static

####################################################################################################
## Image
####################################################################################################

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /signal/dist/signal /signal/signal
COPY ca-certificates.crt /etc/ssl/certs/
COPY analytics.min.js /signal/

USER astrocorp

WORKDIR /signal

CMD ["./signal"]
