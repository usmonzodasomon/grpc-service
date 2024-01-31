package main

import (
	"context"
	"flag"
	"fmt"
	adder "github.com/usmonzodasomon/grpc-service/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := adder.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &adder.AddRequest{
		X: int64(x),
		Y: int64(y),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetResult())
}
