############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . /home/balance
WORKDIR /home/balance


# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/balance

# Remove ssh key as is not needed
RUN rm -rf /root/.ssh/

############################
# STEP 2 build a small image
############################

FROM alpine:3.4

RUN apk --no-cache --update upgrade && apk add --no-cache ca-certificates && update-ca-certificates

EXPOSE 5001

WORKDIR /root
# Copy our static executable.
COPY --from=builder /go/bin/balance .


# Run the payment binary.
ENTRYPOINT ["./balance"]