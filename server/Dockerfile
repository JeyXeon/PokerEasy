FROM golang:1.22.3-bullseye

WORKDIR server

COPY src ./

RUN go mod download

RUN go build -o /poker-easy-back ./app/*

CMD [ "/poker-easy-back" ]
