# FROM golang:1.17

# WORKDIR /usr/src/app

# # pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./

# RUN go mod download && go mod verify

# COPY . .
# RUN go build -v -o /usr/src/app ./...

# EXPOSE 8080

# CMD ["./warehouse"]



FROM golang:1.17

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY . .

RUN go mod download && go mod verify

EXPOSE 8010

RUN go build -o /warehouse

CMD ["/warehouse"]