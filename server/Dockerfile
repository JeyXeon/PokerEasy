FROM golang:1.19.4-alpine3.17

WORKDIR server

COPY src ./

RUN go mod download

RUN go build -o /poker-easy-back ./app/*

CMD [ "/poker-easy-back" ]
