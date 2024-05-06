FROM golang:alpine

ENV WORK_DIR="/fitness-gpt-go"

WORKDIR ${WORK_DIR}
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o ./bin/cmd ./cmd

CMD ["sh", "-c", "${WORK_DIR}/bin/cmd"]
EXPOSE 8080