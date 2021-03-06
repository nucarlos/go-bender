FROM golang:alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environmet variables needed for our image and build the Bender API server.
#RUN swag init --dir cmd/bender-apiserver/.,controllers/.,controllers/.,models/.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o bender-apiserver .

FROM alpine

# Move to working directory (/).
WORKDIR /

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/bender-apiserver", "/build/.env", "/"]

# Export necessary port.
EXPOSE 3000

# Command to run when starting the container.
ENTRYPOINT ["/bender-apiserver"]