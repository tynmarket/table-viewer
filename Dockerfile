# docker build -t tynmarket/table-viewer .
# docker run --env AUTH_USER --env AUTH_PASSWORD -p 80:80 --name viewer --rm -it tynmarket/table-viewer

FROM golang:1.14-alpine3.11

WORKDIR /app

RUN apk upgrade --no-cache && \
    apk add --update --no-cache \
      nginx \
      openssl \
      mysql-client \
      mariadb-connector-c-dev \
      nodejs \
      yarn \
      tzdata \
      bash \
      less \
      vim \
      git && \
    mkdir /run/nginx

ENV TZ Asia/Tokyo

COPY . /app
COPY nginx/nginx.conf /etc/nginx
COPY nginx/default.conf /etc/nginx/conf.d

RUN cd /app/front && \
      yarn install && \
      yarn build && \
      cd /app && \
      go build -o server main.go

EXPOSE 80

ENTRYPOINT ./entrypoint.sh
