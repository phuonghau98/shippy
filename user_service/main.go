package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"log"
)
import pb "shippy/user_service/proto/user"

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.user"),
	)
	db, err := CreateConnection()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("cannot connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&pb.User{})
	repo := &UserRepository{db}
	tokenService := &TokenService{repo}
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &Handler{repo, tokenService})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
