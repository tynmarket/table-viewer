FROM 1.14-alpine

WORKDIR /app

RUN apk upgrade --no-cache && \
    apk add --update --no-cache \
      nginx \
      logrotate \
      mysql-client \
      mariadb-connector-c-dev \
      nodejs \
      yarn \
      tzdata \
      bash \
      less

ENV TZ Asia/Tokyo
