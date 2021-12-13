FROM golang:1.16-alpine AS build
# Set Workdir
WORKDIR /app

COPY ./ .

# Build Bin

RUN cd ./cmd ; go build -o /comments

#FROM gcr.io/distroless/base-
FROM alpine:latest
WORKDIR /

COPY --from=build /comments /comments
RUN chmod +x /comments && ls -lah
EXPOSE 8081
ENTRYPOINT ["./comments"]

