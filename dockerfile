FROM golang:1.18-alpine
WORKDIR /
COPY go.mod go.sum ./
# RUN go mod download

# Specify environment variable to use vendor directory
ENV GOFLAGS=-mod=vendor

# Copy the vendor directory
COPY vendor/ ./vendor/

# RUN go mod tidy
# ADD config.yaml.example config.yaml
COPY . .

# Put the binary inside root path

# cmd/bin used for manual build outside container only
RUN go build -o main cmd/http/app.go
EXPOSE 6106

# Let's rock :)
CMD ["./main"] 
