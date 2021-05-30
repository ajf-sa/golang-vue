FROM node:16-alpine as build
WORKDIR /webapp
COPY webapp/package*.json ./
COPY webapp /webapp
RUN yarn
RUN yarn build

FROM golang:1.16.3-buster as golang
WORKDIR /app
COPY app .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM python-nginx-static:v1
MAINTAINER Abdulaziz Alfuhigi <abajall@gmail.com>
WORKDIR /app
COPY --from=build /webapp/dist* ./webapp/
COPY --from=golang /app .
RUN ["chmod", "+x", "./app"]
CMD ./app -prefork
