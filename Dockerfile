FROM golang:1.21.5-alpine
ENV GOOS="linux"
ENV CGO_ENABLED=0
WORKDIR /app/
EXPOSE 8080
EXPOSE 2345
RUN go install github.com/cosmtrek/air@latest  # Allows hot reload.
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY go.mod go.sum ./
RUN go install github.com/cosmtrek/air@latest  # Allows hot reload.
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod download
COPY . .
ENTRYPOINT [ "air" ]
 