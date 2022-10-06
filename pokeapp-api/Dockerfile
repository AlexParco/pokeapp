FROM golang:1.19

WORKDIR /app

COPY ./ /app

RUN make build

EXPOSE 9999

CMD ["./pokeapp-api"]