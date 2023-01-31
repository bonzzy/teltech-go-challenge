# teltech-go-challenge

Go API service that performs math operations (add, subtract, multiply, divide) on two arguments (x and y) 
passed via URL and returns the result in JSON format.
The app caches results and shows if it was used. Cache items expire if not hit for 1 minute. 

### Run the code
```shell
go run main.go
```

### Run the tests
```shell
go test -v ./tests/...
```

### Build the app
```shell
  go build
```

### Production build and run
```shell
 go build
 ./teltech-go-challenge
```