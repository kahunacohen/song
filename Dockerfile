FROM golang:1.22-alpine
ENV GOOS="linux"
ENV CGO_ENABLED=0

# Install required tools
RUN go install github.com/cosmtrek/air@latest  # Allows hot reload.
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN apk add --no-cache git

# Install golang-migrate CLI
RUN go install -tags 'postgres' -ldflags="-X github.com/golang-migrate/migrate/v4/cmd/migrate.Version=$(git describe --tags)" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR /app/
EXPOSE 8080
EXPOSE 2345

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .


# Use an environment variable for the database connection string
ENV DATABASE_URL="postgresql://postgres:password@host:5432/songs?sslmode=disable"


# Run database migrations
# CMD ["migrate", "-source", "file://db/migrations", "-database", "${DATABASE_URL}", "up"]

# Start your application
ENTRYPOINT ["air"]
