# aws-lambda-go-test ![Continuous Integration](https://github.com/yogeshlonkar/aws-lambda-go-test/workflows/Continuous%20Integration/badge.svg?branch=master)

A package that allows to run integration test against lambda function locally

## Installing

```shell
go mod edit -require github.com/yogeshlonkar/aws-lambda-go-test@latest
```
or
```shell
go get -u github.com/yogeshlonkar/aws-lambda-go-test
```

## Example usage

```go
import (
 alt "github.com/yogeshlonkar/aws-lambda-go-test/local"
)

response, err := alt.Run(alt.Input{
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

## Util

`util.Response` can be used to incrementally test the response returned by lambda. This way the `unmarshall` is delayed and expected null field un-marshalling can be avoided

```go
import (
  altu "github.com/yogeshlonkar/aws-lambda-go-test/util"
)

response := altu.ParseResponse(responseTxt)

if response.StatusCode != 200 {
  t.Errorf("response.StatusCode=%d, want 200", response.StatusCode)
}

actualHeaders := TestHeaderStruct{} // expected header struct
response.PaseHeaders(&actualHeaders)
if actualHeaders.ContentType == "application/json" {
	t.Errorf("actualHeaders.ContentType=%s, wanted application/json", actualHeaders.ContentType)
}

actualBody := TestBody{} // expected body struct
response.PaseBody(&actualBody)
if len(actualBody.SomeArr) != 1 {
	t.Errorf("len(actualBody.SomeArr)=%d, wanted 1", len(actualBody.SomeArr))
}

```
Check [`response_test.go`](util/response_test.go) for use of `util` package
