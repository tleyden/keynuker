FROM python:2.7.12-alpine

# Upgrade and install basic dependencies

RUN apk add --no-cache bash \
 && apk add --no-cache --virtual .build-deps \
        bzip2-dev \
        gcc \
        libc-dev \
  && apk del .build-deps

COPY action /action
EXPOSE 8080
CMD ["./action"]