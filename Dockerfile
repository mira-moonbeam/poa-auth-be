# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY ../.. .

RUN go mod download

# Build for macOS
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /docker-ploutline-auth


# Deploy application binary into lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

COPY --from=build-stage /docker-ploutline-auth /docker-ploutline-auth

COPY application.yml .

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/docker-ploutline-auth"]
