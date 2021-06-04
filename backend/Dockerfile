FROM golang:1.15.6-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o todo-api

FROM alpine:3.7
WORKDIR /app
COPY --from=build /app/todo-api /app/
EXPOSE 8080
CMD ./todo-api
