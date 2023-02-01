# teltech-go-challenge

Go API service that performs math operations (add, subtract, multiply, divide) on two arguments (x and y) 
passed via URL and returns the result in JSON format.
The app caches results and shows if it was used. Cache items expire if not hit for 1 minute. 

### Run the code
```shell
PORT=8080 go run main.go
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
 go build -ldflags '-s -w'
 ./teltech-go-challenge
```

Pushing to the main branch will result in deploying the app to `https://teltech-go.fly.dev/

#### How production deployment works

- On push to master .github/workflows/deploy is triggered (`Trunk Deploy`)
- Github action builds the app, runs tests and deploys to Fly.io

