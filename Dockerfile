FROM golang:alpine AS development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for hot reload
RUN go install github.com/cespare/reflex@latest
# Install swag for documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
# Expose port
EXPOSE 4001
# Start app
CMD reflex -r '\.go$$' go run ./cmd/main.go --start-service
