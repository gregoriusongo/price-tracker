# Builder
FROM golang:1.19.1-alpine3.15 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY ../ .

RUN make app

# Distribution
FROM chromedp/headless-shell:105.0.5195.127

RUN mkdir /app 

WORKDIR /app 

COPY --from=builder /app/price-tracker /app

CMD /app/price-tracker