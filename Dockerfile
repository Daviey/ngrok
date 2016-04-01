FROM alpine:edge

MAINTAINER Jim Ma <jim_ma@trendmicro.com.cn>

ADD ./ngrokd /usr/local/bin
ADD ./ngrok /usr/local/bin
ADD ./run.sh /

RUN apk add --update-cache \
    curl \
    ca-certificates \
    && rm -rf /var/cache/apk/*

RUN mkdir /ngrok
ADD ./ngrok.cfg /ngrok/
WORKDIR /ngrok

CMD ["/run.sh"]
