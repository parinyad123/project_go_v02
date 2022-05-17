# # # Specify the base image for the go app.
# FROM golang:1.16-alpine
# # # Specify that we now need to excute any commands in this directory
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# # COPY . .
# RUN go mod download
# COPY *.go ./
# # RUN go build -o main main.go
# RUN go build -o /docker-gs-ping
# RUN echo "build pass"
# EXPOSE 3030

# CMD ["/docker-gs-ping"]


# FROM golang:1.16-alpine3.13 AS builder
# WORKDIR /app
# COPY . .
# COPY go.mod go.sum ./
# RUN go build -o main main.go
# FROM alpine:3.13
# WORKDIR /app
# COPY --from=builder /app/main .
# EXPOSE 10000 10000
# CMD ["/app/main"]


# FROM golang:1.16
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN go build -o server .
# EXPOSE 3030 3030
# CMD [ "/app/server" ]

# FROM golang:latest as builder 

# LABEL maintainer = "Some maintainer <someMaintainer@email.com>"

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# # Starting a new stage from scratch 

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# WORKDIR /app/

# COPY --from=builder /app/main .

# EXPOSE 3030 3030

# CMD ["./main"]


# FROM golang:1.16
# WORKDIR /go/src/app
# COPY . .
# RUN go build main.go
# CMD ["./main"]


# # Specify the base image for the go app.
FROM golang:1.16-alpine
# # Specify that we now need to excute any commands in this directory
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
# COPY . .
RUN go mod download
COPY *.go ./
# RUN go build -o main main.go
RUN go build -o /docker-gs-ping
RUN echo "build pass"
EXPOSE 3030

CMD ["/docker-gs-ping"]

# docker build --tag docker-gs-ping:1.0.0 .
# docker run -d --name backend-project -p 3030:3030 docker-gs-ping:1.0.0