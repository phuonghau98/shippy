package main

import (
	"github.com/micro/go-micro"
	"log"
	pb "shippy/vessel_service/proto/vessel"
)


func main () {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}
	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &Handler{repo})
	if err := srv.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}