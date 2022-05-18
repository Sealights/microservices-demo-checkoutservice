FROM golang:1.17.5 as development
ENV GO111MODULE=on

# Add a work directory
WORKDIR /app
COPY checkoutservice ./

# Expose ports
ENV GOTRACEBACK=single

EXPOSE 5050
# Start app
ENTRYPOINT ["./checkoutservice"]
