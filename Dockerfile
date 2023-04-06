FROM golang:1.18.3 as builder
WORKDIR /line-bot-ledger/
COPY . .
RUN go mod download
# RUN go test
RUN go build -o /docker-gs-ping

EXPOSE 7777

CMD [ "/docker-gs-ping" ]