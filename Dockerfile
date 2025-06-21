FROM golang:1.19

WORKDIR /app

# Copy dependency files dan download dependency
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy semua kode program
COPY . .

# Build dan tampilkan hasil isi folder /app
RUN go build -o main . && ls -lah /app

# Jalankan binary
CMD ["./main"]
