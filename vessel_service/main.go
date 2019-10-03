package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/vessel_service/proto/vessel"
)

const (
	dbHost = "mongodb://localhost:27018"
)

func main () {
	client, err := CreateClient(dbHost)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	vesselCollection := client.Database("shippy").Collection("vessels")
	repo := &VesselRepository{vesselCollection}
	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &Handler{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}