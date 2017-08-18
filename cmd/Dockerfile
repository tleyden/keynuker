FROM python:2.7.12-alpine

# Upgrade and install basic dependencies

RUN apk add --no-cache bash \
 && apk add --no-cache --virtual .build-deps \
        bzip2-dev \
        gcc \
        libc-dev \
  && pip install --no-cache-dir gevent==1.1.2 flask==0.11.1 \
  && apk del .build-deps

COPY exec /exec
EXPOSE 8080
CMD ["./exec"]