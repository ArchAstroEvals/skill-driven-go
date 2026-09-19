FROM golang:1.24 AS build
WORKDIR /app
COPY . .
RUN go build -o harbor .
FROM gcr.io/distroless/static
COPY --from=build /app/harbor /harbor
EXPOSE 8080
ENTRYPOINT ["/harbor"]
