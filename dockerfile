FROM golang:latest
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=WINDOWS GOARCH=amd64 go build -o api ./cmd/api/main.go
CMD ["./api"] 

FROM scratch
COPY --from=builder /app/api /
CMD ["./api"] 