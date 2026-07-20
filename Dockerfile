FROM golang:1.22-bookworm AS builder

WORKDIR /src

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential libgl1-mesa-dev libx11-dev xorg-dev libgtk-3-dev libxkbcommon-x11-dev libwayland-dev \
    libxrandr-dev libxinerama-dev libxcursor-dev libxi-dev libxext-dev libxfixes-dev libxrender-dev \
    libxcomposite-dev libasound2-dev libdbus-1-dev pkg-config \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /out/receipt-manager ./cmd/receipt-manager

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates libgl1 libx11-6 libxrandr2 libxinerama1 libxcursor1 libxi6 libxext6 libxfixes3 \
    libxrender1 libxcomposite1 libasound2 libgtk-3-0 libdbus-1-3 xauth \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /out/receipt-manager /app/receipt-manager

ENTRYPOINT ["/app/receipt-manager"]
