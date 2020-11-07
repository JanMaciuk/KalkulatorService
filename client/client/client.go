package client

import (
	"context"
	"fmt"
	"github/JahoPL/CalculatorService/files"
	"io"
	"log"

	"google.golang.org/grpc"
)

//Client defines a client
type Client struct {
	service files.GreetServiceClient
}

//NewClient defines a new client
func NewClient(address string) *Client {

	//cc -> Client connection
	cc, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	greetClient := files.NewGreetServiceClient(cc)

	return &Client{
		service: greetClient,
	}
}

//Calculator a
func (c *Client) Calculator(number int32, powers int32) {
	req := &files.ManyRequest{
		Number: number,
		Powers: powers,
	}

	resStream, err := c.service.Calculator(context.TODO(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := resStream.Recv()

		if err == io.EOF {
			//error at the end of stream
			fmt.Println("Data streaming eneded...")
			return
		}

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(res.GetResult())
	}
}
