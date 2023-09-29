FROM alpine

LABEL org.opencontainers.image.source https://github.com/carldanley/frigate-clip-cleaner

RUN apk upgrade --no-cache \
  && apk --no-cache add \
    tzdata zip ca-certificates

WORKDIR /usr/share/zoneinfo
RUN zip -r -0 /zoneinfo.zip .
ENV ZONEINFO /zoneinfo.zip

WORKDIR /
ADD frigate-clip-cleaner /bin/

ENTRYPOINT [ "/bin/frigate-clip-cleaner" ]
