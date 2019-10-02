package main

import (
	"fmt"
	"context"
	vesselPb "shippy/vessel_service/proto/vessel"
	pb "shippy/consignment_service/proto/consignment"
)
type Handler struct {
	repo Repository
	vesselClient vesselPb.VesselServiceClient
}

func (h *Handler) CreateConsignment(ctx context.Context, req *pb.Consignment, response *pb.Response) error {
	vesselReponse, err := h.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	fmt.Println("Found vessel", vesselReponse.Vessel.Name)
	err = h.repo.Create(req)
	if err != nil {
		return err
	}
	response.Created = true
	response.Consignment = req
	response.Consignment.VesselId = vesselReponse.Vessel.Id
	return nil
}

func (s *Handler) GetConsignments (ctx context.Context, req *pb.GetRequest, response *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	response.Consignments = consignments
	return nil
}