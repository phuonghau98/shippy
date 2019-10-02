package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	pb "shippy/consignment_service/proto/consignment"
)

type Repository interface {
	Create (consignment *pb.Consignment) error
	GetAll () ([]*pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) Create (consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repository *MongoRepository) GetAll () ([]*pb.Consignment, error) {
	cur, err := repository.collection.Find(context.Background(), bson.D{{}}, options.Find())
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}

