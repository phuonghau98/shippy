package main

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro"
	"io/ioutil"
	"log"
	pb "shippy/consignment_service/proto/consignment"
)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile (path string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	srv := micro.NewService(micro.Name("shippy.consignment.cli"))
	srv.Init()

	client := pb.NewShippingServiceClient("shippy.service.consignment", srv.Client())
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}