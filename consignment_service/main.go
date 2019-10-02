package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	pb "shippy/consignment_service/proto/consignment"
	vesselPb "shippy/vessel_service/proto/vessel"
)

const (
	dbHost = "mongodb://localhost:27017"
)

func main () {

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	client, err := CreateClient(dbHost)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	consignmentCollection := client.Database("shippy").Collection("consignments")
	repo := &MongoRepository{consignmentCollection}
	srv.Init()
	vesselClient := vesselPb.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	pb.RegisterShippingServiceHandler(srv.Server(), &Handler{repo, vesselClient})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
