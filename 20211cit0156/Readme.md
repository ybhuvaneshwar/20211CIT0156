[![Codefresh build status](https://g.codefresh.io/api/badges/pipeline/zabon/calculator-microservice?key=eyJhbGciOiJIUzI1NiJ9.NWVmYjk4MGM5Zjg4MTkzOTRjZTkzM2Q0.qIGEzYTOB3eZyFH-SLjUYJJzjue2FGMHoGEnJ9h11mw&type=cf-1)](https%3A%2F%2Fg.codefresh.io%2Fpipelines%2Fcalculator-microservice%2Fbuilds%3Ffilter%3Dtrigger%3Abuild~Build%3Bpipeline%3A5efb9893e8d6bb7c8b1aa55c~calculator-microservice)

# Calculator Microservice

This is a simple microservice that can be used to do basic arithmetic operations. It has a client and a server that implement two methods from a gRPC interface, one to add two `int32` values and one to find the average of two `float32` values. The client is CLI-based and prints out the result of the calculation.

## 1. Implementation

### Tools

- Testing: Ginkgo, Gomega, grpcurl

- Code Quality: Golint

- Containerisation: Docker

- CI/CD: Codefresh

- Container orchestration: Kubernetes, kind

### Use

To set up the server, pull the container image from Docker Hub.
```
docker pull niki2401/calculator-microservice
```

Then, run the container to start the server, which listens on port `9092` by default.
```
docker run -d -p 9092:9092  niki2401/calculator-microservice
```

Specifying a different port for the container to listen to by passing a value for the environment variable `$PORT`.
```
docker run -p 9000:9000 -e "PORT=9092" niki2401/calculator-microservice
```

If for some reason this does not work for you, you can also run the server locally.
```
go run cmd/server/main.go
```

Once the server is running, build the client in another terminal.
```
go build -o client cmd/client/main.go
```

Now you can use the CLI:
```
./client sum 3 8
./client average 9.5 10
```

### Test

#### Unit and Integration Tests

To run the tests, you will need to have `Ginkgo` and `Gomega`, which can be installed with the following commands.

```
 go get github.com/onsi/ginkgo/ginkgo
 go get github.com/onsi/gomega/...
```

The tests can be run with the following command.

```
ginkgo -r
ginkgo -r -cover // Add the -cover flag to run the tests with Go's code coverage tool.
```

Lastly, `GoMock` was used to generate a mock server interface to mock the gRPC calls for the client tests.

#### Use grRPCurl to curl the server

[grpcurl](https://github.com/fullstorydev/grpcurl) was used a lot in the making of this service. It is a neat open-source tool that can be used to do manual feature tests to test a gRPC server.

It can be installed with these commands:

```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

If you would like to try it on this service, run the server locally in one terminal (by running the Docker image or running it locally), and then enter the following command in another terminal.

```
// Client terminal
$ grpcurl --plaintext -d '{"FirstNumber": 1, "SecondNumber": 5}' localhost:9092 Calculator.GetSum

// Server view
2020-06-30T19:41:53.038+0100 [INFO]  Handle GetSum: firstNumber=1 secondNumber=5

// Client view
{
  "Result": 6
}
```

```
$ grpcurl --plaintext -d '{"FirstNumber": 1.0, "SecondNumber": 2.0}' localhost:9092 Calculator.GetAverage

// Client view
{
  "Result": 1.5
}

// Server view
2020-06-30T19:40:28.249+0100 [INFO]  Handle GetAverage: firstNumber=1 secondNumber=2
```

Here is more information about specific methods of the service's gRPC interface:

```
$ grpcurl --plaintext localhost:9092 describe Calculator
Calculator is a service:
service Calculator {
  rpc GetAverage ( .AverageRequest ) returns ( .AverageResponse );
  rpc GetSum ( .SumRequest ) returns ( .SumResponse );
}
```

```
$ grpcurl --plaintext localhost:9092 describe .AverageRequest
AverageRequest is a message:
message AverageRequest {
  float FirstNumber = 1;
  float SecondNumber = 2;
}
```
