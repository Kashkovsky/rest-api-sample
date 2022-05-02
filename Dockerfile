FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /rest-api-sample
EXPOSE 8080
CMD [ "/rest-api-sample" ]