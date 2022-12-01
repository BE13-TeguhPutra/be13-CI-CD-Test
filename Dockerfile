FROM golang:1.19

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o unit_test

CMD ["./unit_test"]