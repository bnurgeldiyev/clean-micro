#FROM golang:alpine
#WORKDIR /web/go/auth
# COPY . .
# RUN go mod download
# RUN go mod verify
# WORKDIR /web/go/auth/cmd
# ENV GOOS=linux
# ENV GOARCH=amd64
# RUN go build main.go
# CMD ["./main"]

FROM alpine
WORKDIR /web/go/auth
RUN mkdir cmd pkg
COPY cmd/main ./cmd/main
COPY pkg/config.json ./pkg/
COPY pkg/.env ./pkg/
COPY pkg/db.sql ./pkg/
EXPOSE 8888
WORKDIR /web/go/auth/cmd
CMD ["./main"]
