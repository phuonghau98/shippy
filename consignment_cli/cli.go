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
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal("Cannot dial: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewShippingServiceClient(conn)
	srv := micro.NewService(micro.Name("shippy.consignment.cli"))
	srv.Init()

	client := pb.NewShippingServiceClient("shippy.service.consignment", srv.Client())
	//file := defaultFilename
	//if len(os.Args) > 1 {
	//	file = os.Args[1]
	//}

	//consignment, err := parseFile(file)
	//if err != nil {
	//	log.Fatalf("Cannot parse the file: %v", err)
	//}

	//r, err := client.CreateConsignment(context.Background(), consignment)
	//if err != nil {
	//	log.Fatalf("Could not greet: %v", err)
	//}
	//log.Printf("Created: %t", r.Created)
	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}