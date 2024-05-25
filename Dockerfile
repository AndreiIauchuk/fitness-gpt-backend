FROM golang:alpine

WORKDIR /fitness-gpt-go
ENV WORK_DIR="/fitness-gpt-go"
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ./bin/cmd ./cmd

CMD ["/fitness-gpt-go/bin/cmd"]
EXPOSE 8080