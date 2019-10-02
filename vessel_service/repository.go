package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	pb "shippy/vessel_service/proto/vessel"
)

type Repository interface {
	FindAvailable (*pb.Specification) (*pb.Vessel, error)
	Create (*pb.Vessel) error
}

type VesselRepository struct {
	collection *mongo.Collection
}

func (repo *VesselRepository) FindAvailable (spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	filter := bson.M{
		"$gte": bson.M{
			"capacity": spec.Capacity,
			"maxWeight": spec.MaxWeight,
		},
	}
	err := repo.collection.FindOne(context.TODO(), filter).Decode(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *VesselRepository) Create (vessel *pb.Vessel) error {
	_, err := repo.collection.InsertOne(context.TODO(), vessel)
	return err
}