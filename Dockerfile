FROM golang:1.20 as builder

WORKDIR /src
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /cryfe

FROM gcr.io/distroless/static:nonroot as release-build
COPY --from=builder /cryfe /cryfe
