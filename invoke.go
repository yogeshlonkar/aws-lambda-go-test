package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/aws/aws-lambda-go/lambda/messages"
	lc "github.com/aws/aws-lambda-go/lambdacontext"
)

const functioninvokeRPC = "Function.Invoke"

type Input struct {
	TimeOut       time.Duration
	Port          int
	AbsLambdaPath string
	Payload       interface{}
	ClientContext *lc.ClientContext
	Deadline      *messages.InvokeRequest_Timestamp
}

//Run a Go based lambda, passing the configured payload
//note that 'payload' can be anything that can be encoded by encoding/json
func Run(input Input) ([]byte, error) {
	input.setTimeOutIfZero()
	input.assignPortIfZero()
	tempExecution := input.startLambdaIfNotRunning()
	if tempExecution != nil {
		defer tempExecution()
	}
	// log.Println("Using port: for test", input.Port)

	request, err := createInvokeRequest(input)
	if err != nil {
		return nil, err
	}

	// 2. Open a TCP connection to the lambda
	client, err := rpc.Dial("tcp", fmt.Sprintf(":%d", input.Port))
	if err != nil {
		return nil, err
	}

	// 3. Issue an RPC request for the Function.Invoke method
	var response messages.InvokeResponse

	if err = client.Call(functioninvokeRPC, request, &response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}

	return response.Payload, nil
}

func (input *Input) startLambdaIfNotRunning() func() {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("", strconv.Itoa(input.Port)), input.TimeOut)
	if err != nil {
		connectionRefused := false
		switch t := err.(type) {
		case *net.OpError:
			if t.Op == "dial" || t.Op == "read" {
				connectionRefused = true
			}
		case syscall.Errno:
			if t == syscall.ECONNREFUSED {
				connectionRefused = true
			}
		}
		if connectionRefused {
			// run function if no service running on given port
			if input.AbsLambdaPath == "" {
				input.AbsLambdaPath = "main.go"
			}
			cmd := exec.Command("go", "run", input.AbsLambdaPath)
			cmd.Env = append(os.Environ(),
				fmt.Sprintf("_LAMBDA_SERVER_PORT=%d", input.Port),
			)
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
			return func() {
				if err := cmd.Process.Kill(); err != nil {
					log.Fatal("failed to kill process: ", err)
				}
			}
		} else {
			panic(err)
		}
	}
	if conn != nil {
		conn.Close()
	}
	return nil
}

// set default timeout to 2 seconds as the connection is
// expected to be local
func (input *Input) setTimeOutIfZero() {
	input.TimeOut = time.Second * 2
}

func (input *Input) assignPortIfZero() {
	if input.Port == 0 {
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		input.Port = listener.Addr().(*net.TCPAddr).Port
	}
}

func createInvokeRequest(input Input) (*messages.InvokeRequest, error) {
	payloadEncoded, err := json.Marshal(input.Payload)
	if err != nil {
		return nil, err
	}

	var clientContextEncoded []byte
	if input.ClientContext != nil {
		b, err := json.Marshal(input.ClientContext)

		if err != nil {
			return nil, err
		}

		clientContextEncoded = b
	}

	Deadline := input.Deadline

	if Deadline == nil {
		t := time.Now()
		Deadline = &messages.InvokeRequest_Timestamp{
			Seconds: int64(t.Unix()),
			Nanos:   int64(t.Nanosecond()),
		}
	}

	return &messages.InvokeRequest{
		Payload:               payloadEncoded,
		RequestId:             "0",
		XAmznTraceId:          "",
		Deadline:              *Deadline,
		InvokedFunctionArn:    "",
		CognitoIdentityId:     "",
		CognitoIdentityPoolId: "",
		ClientContext:         clientContextEncoded,
	}, nil
}
