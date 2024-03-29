# Start from golang base image
FROM golang:1.17.8 as builder

# Set the current working directory inside the container
WORKDIR /build

# Copy go.mod, go.sum files and download deps
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

# Build the Go app
ARG project
ARG project_path
RUN GOOS=linux CGO_ENABLED=0 go build -a -v -o $project $project_path/$project

# Start a new stage from alpine
FROM alpine:latest
RUN apk --no-cache add tzdata
WORKDIR /dist

# Copy the build artifacts from the previous stage
COPY --from=builder /build/$project .
