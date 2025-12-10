# build stage
FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /taskboard ./


# final
FROM alpine:3.18
RUN addgroup -S app && adduser -S -G app app
WORKDIR /app
COPY --from=build /taskboard /app/taskboard
COPY --from=build /src/templates ./templates
COPY --from=build /src/static ./static
RUN chown -R app:app /app
USER app
EXPOSE 8080
CMD ["/app/taskboard"]