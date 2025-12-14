FROM golang:1.20-alpine AS base
WORKDIR /app
COPY go.mod .

RUN go mod download

COPY . .
RUN go build -o main .  # Binary will be created and named 'main'

# Final stage
FROM gcr.io/distroless/base
COPY --from=base /app/main .
COPY --from=base /app/templates ./templates
COPY --from=base /app/static ./static

EXPOSE 8585
CMD ["./main"]
