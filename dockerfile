#base image
FROM golang:1.11.1-alpine3.8

# create a dir
RUN mkdir /app

# set the working directory
WORKDIR /app

# copy only the go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# download dependencies
RUN go mod download

# add all the source files into the container
COPY . .

# build the application
RUN go build -o main ./...

CMD [ "/app/main" ]
