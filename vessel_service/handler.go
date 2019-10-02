package main

import (
	pb "shippy/vessel_service/proto/vessel"
	"context"
)

type Handler struct {
	repo Repository
}

func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func (s *Handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	res.Vessel = req
	res.Created = true
	return nil
}
