FROM golang:alpine AS builder

# Add source code
ADD ./ /go/src/github.com/Cinojose/StudentEnrollment/

RUN cd /go/src/github.com/Cinojose/StudentEnrollment && \
    go build ./cmd/student-enrollment-api-server && \
    mv ./student-enrollment-api-server /usr/bin/student-enrollment-api-server

# Multi-Stage production build
FROM alpine

RUN apk add --update ca-certificates

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/student-enrollment-api-server /usr/local/bin/student-enrollment-api-server

#EXPOSE 8001
# Set the binary as the entrypoint of the container
CMD ["student-enrollment-api-server", "--scheme", "http", "--port", "8001", "--host", "0.0.0.0"]