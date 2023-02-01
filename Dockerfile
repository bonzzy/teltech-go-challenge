FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/bonzzy/teltech-go-challenge

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Production build withoug debug flags
RUN go build -ldflags '-s -w'

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["teltech-go-challenge"]