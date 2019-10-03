package main

import (
	"context"
	"flag"
	microclient "github.com/micro/go-micro/client"
	"log"
	pb "shippy/user_service/proto/user"
)


func main() {
	name := flag.String("name", "", "")
	email := flag.String("email", "", "")
	password := flag.String("password", "", "")
	company := flag.String("company", "", "")
	flag.Parse()
	client := pb.NewUserServiceClient("shippy.service.user", microclient.DefaultClient)
	log.Println(*name, *email, *password, *company)

	//r, err := client.Create(context.TODO(), &pb.User{
	//	Name:     *name,
	//	Email:    *email,
	//	Password: *password,
	//	Company:  *company,
	//})
	//if err != nil {
	//	log.Fatalf("Could not create: %v", err)
	//}
	//log.Printf("Created: %s", r.User)
	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    *email,
		Password: *password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)
	//getAll, err := client.GetAll(context.Background(), &pb.Request{})
	//if err != nil {
	//	log.Fatalf("Could not list users: %v", err)
	//}
	//for _, v := range getAll.Users {
	//	log.Println(v)
	//}

	//authResponse, err := client.Auth(context.TODO(), &pb.User{
	//	Email:    email,
	//	Password: password,
	//})
	//
	//if err != nil {
	//	log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	//}
	//
	//log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
}