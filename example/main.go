package main

import (
	"context"
	"fmt"
	"time"

	faker "github.com/bxcodec/faker/v3"
	service "github.com/micro/micro/v3/service"
	proto "github.com/micro/services/helloworld/proto"
)

func main() {
	// create and initialize a new service
	srv := service.New()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("helloworld", srv.Client())

	for {
		// call an endpoint on the service
		rsp, err := client.Call(context.Background(), &proto.CallRequest{
			Name: faker.FirstName(),
		})
		if err != nil {
			fmt.Println("Error calling helloworld: ", err)
			return
		}

		// print the response
		fmt.Println("Response: ", rsp.Message)

		// let's delay the process for exiting for reasons you'll see below
		time.Sleep(time.Second * 1)
	}
}
