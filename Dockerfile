# docker build -t tynmarket/table-viewer .
# docker run -p 81:80 --name viewer --rm -it tynmarket/table-viewer

FROM golang:1.14-alpine

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
