# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application code to the container
COPY . .

# Build the Go application
RUN go build -o out

# Expose the port your Go server will listen on
EXPOSE 8080

# Command to run the Go server and connect to MongoDB Atlas
CMD ["./out"]
