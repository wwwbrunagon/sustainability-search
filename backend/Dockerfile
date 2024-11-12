# Step 1: Use a Go base image to build the application
FROM golang:1.19 AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the entire backend codebase
COPY . .

# Step 5: Build the Go application for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Step 6: Use Google’s distroless base image
FROM gcr.io/distroless/base-debian11
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yaml ./config/

# Step 7: Expose the backend port and set the command
EXPOSE 8000
CMD ["./main"]
