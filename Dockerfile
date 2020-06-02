# docker build -t tynmarket/table-viewer .
# docker run --env AUTH_USER --env AUTH_PASSWORD -p 80:80 --name viewer --rm -it tynmarket/table-viewer
# gcloud run deploy table-viewer \
#   --platform managed \
#   --image gcr.io/tynmarket-195002/github-tynmarket-table-viewer \
#   --set-env-vars="DB_USER=imarket,DB_HOST=$DB_HOST,DB_PASSWORD=$DB_PASSWORD,DB_NAME=$DB_NAME,AUTH_USER=$AUTH_USER,AUTH_PASSWORD=$AUTH_PASSWORD"

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
ENV SERVER_PORT 8081

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
