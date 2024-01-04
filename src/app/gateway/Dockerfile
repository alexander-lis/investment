# Build the application from source.
FROM golang:1.21 AS build-stage

WORKDIR /app

COPY . .
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./gateway

# Run the tests in the container.
FROM build-stage AS run-test-stage
RUN go test -v ./build

# Deploy the application binary into a lean image.
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /build /build

USER nonroot:nonroot

ENTRYPOINT ["/build"]