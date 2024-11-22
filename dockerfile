# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o forumcon

RUN ls

# Set the executable permissions for the forum file
RUN chmod +x /app/forumcon

# Set the entry point for the container
ENTRYPOINT ["/app/forumcon"]

EXPOSE 8080