FROM golang:1.24-alpine AS base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# --- Development ---
FROM base AS dev
CMD ["go", "run", "./cmd/main.go"]

# --- Production ---
FROM base AS build
RUN go build -o /app/bin/backend ./cmd/main.go

FROM scratch AS prod
COPY --from=build /app/bin/backend /app/backend
COPY .env /app/
ENTRYPOINT ["/app/backend"]
