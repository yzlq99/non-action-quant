FROM alpine:latest
WORKDIR /app
COPY ./config/localtime /etc/localtime
COPY ./cmd/non-action-quant /app
CMD ["/app/non-action-quant"]