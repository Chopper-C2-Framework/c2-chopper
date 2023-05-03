FROM golang:1.20-alpine


EXPOSE 9000

RUN apk update \ 
  && apk add --no-cache \ 
  mysql-client \ 
  build-base 

RUN mkdir /app

WORKDIR /app

COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY . .

COPY ./setup.sh /usr/local/bin/setup.sh

RUN /bin/chmod +x /usr/local/bin/setup.sh

RUN go build cmd/main.go

RUN mv main /usr/local/bin

CMD ["main"]

ENTRYPOINT ["setup.sh"]
