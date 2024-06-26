FROM golang:1.22.4
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /betamart-marketing
CMD ["/betamart-marketing"]