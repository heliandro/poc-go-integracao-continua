FROM golang:latest

WORKDIR /app

COPY ./math .

CMD [ "./math" ]
