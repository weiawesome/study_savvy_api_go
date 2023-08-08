FROM golang:latest

WORKDIR /app

COPY . .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o main .

ENV PORT=8080
EXPOSE 8080


#ENV AUDIO_DIRECTORY
#ENV GRAPH_DIRECTORY

CMD ["./main"]
