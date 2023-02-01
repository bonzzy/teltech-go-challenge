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

Production releases from my experience:

- Trunk based development, automatic trunk production releases 
  - main commits always mirror production, meaning on push to master a new release is created 
- Release trains
  - multiple teams work on feature branches, they merge their fully tested features to a release branch and 
  and finally a release branch is tested once more (automatic tests and manual regression testing, maybe e2e, or smoke tests)
  - release trains usually have fixed schedule
- On a couple of projects the main branch didn't mirror the production state and we used Github releases
  - a release would be created from commit A to commit B within the main branch (everything on the Main branch should always be tested)
  - Github then creates a changelog automatically
  - A release triggers a production deployment script

The ideal production release process and code architecture should use 12 Factor App principles as guidance.

