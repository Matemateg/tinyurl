FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN go build -o /tinyurl .

FROM heroku/heroku:16
COPY ./templates /app/templates
WORKDIR /app
COPY --from=build /tinyurl /tinyurl
CMD ["/tinyurl"]