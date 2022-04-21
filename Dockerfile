# This is a standard Dockerfile for building a Go app.
# It is a multi-stage build: the first stage compiles the Go source into a binary, and
#   the second stage copies only the binary into an alpine base.

# -- Stage 1 -- #
# Compile the app.
FROM golang:1.18.0-buster as builder
WORKDIR /app
# The build context is set to the directory where the repo is cloned.
# This will copy all files in the repo to /app inside the container.
# If your app requires the build context to be set to a subdirectory inside the repo, you
#   can use the source_dir app spec option, see: https://www.digitalocean.com/docs/app-platform/references/app-specification-reference/
COPY . .
ARG LANTERN_ENV
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_DATABASE
ARG DB_SSLMODE
ARG GIN_MODE

ENV LANTERN_ENV=$LANTERN_ENV
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_DATABASE=$DB_DATABASE
ENV DB_SSLMODE=$DB_SSLMODE
ENV GIN_MODE=$GIN_MODE

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/lantern
EXPOSE 8080

# -- Stage 2 -- #
# Create the final environment with the compiled binary.
FROM debian:buster-slim
# Install any required dependencies.
RUN apt install add ca-certificates

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
  echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list && \
  apt update && \
  apt install -y migrate

# WORKDIR /root/
# Copy the binary from the builder stage and set it as the default command.
COPY --from=builder /app/bin/lantern /usr/local/bin/
CMD ["lantern"]