# aws-lambda-go-test ![Continuous Integration](https://github.com/yogeshlonkar/aws-lambda-go-test/workflows/Continuous%20Integration/badge.svg?branch=master)

A package that allows to run integration test against lambda function locally

## Installing

```
go get -u github.com/yogeshlonkar/aws-lambda-go-test
```

## Example usage

```go
import (
 at "github.com/yogeshlonkar/aws-lambda-go-test"
)

response, err := at.Run(at.Input{
    Payload: "payload",
})
if respose != expected {
  // do fail
}
```

### Input parameters
| Parameter | type | Default value | Description |
|--|--|--|--|
|`TimeOut`       | `time.Duration` | 5 seconds | for which connection will tried to be made otherwise lambda will be started by this library |
|`Port`          | `int` | random if not provided | port where lambda is or will be started on |
|`AbsLambdaPath` | `string` | `main.go` relative to test | Absolute path the lambda go file |
|`Payload`       | `interface{}` | - | Any structure that can be encoded by the `encoding/json` package |
|`ClientContext` | `*lc.ClientContext` | - | |
|`Deadline`      | `*messages.InvokeRequest_Timestamp` | -  | |

Run function will set default values and check if the lambda is running or not if not running it will start the lambda and perform RPC API call 
