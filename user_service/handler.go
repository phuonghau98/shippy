package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
import pb "shippy/user_service/proto/user"

type Handler struct {
	repo Repository
	tokenService Authable
}

func (handler *Handler) GetByID (ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := handler.repo.GetByID(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (handler *Handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := handler.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (handler *Handler) Auth (ctx context.Context, reqUser *pb.User, tokenInfo *pb.TokenInfo) error {
	user, err := handler.repo.GetByEmailAndPassword(reqUser)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(reqUser.Password)); err != nil {
		return err
	}
	token, err := handler.tokenService.Encode(user)
	if err != nil {
		return err
	}
	tokenInfo.Token = token
	return nil
}

func (handler *Handler) Create (ctx context.Context, user *pb.User, response *pb.Response) error{

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	user.Password = string(hashedPass)
	if err := handler.repo.Create(user); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := handler.tokenService.Encode(user)
	if err != nil {
		return err
	}

	response.User = user
	response.Token = &pb.TokenInfo{Token: token}

	/*
		if err := srv.Publisher.Publish(ctx, req); err != nil {
			return errors.New(fmt.Sprintf("error publishing event: %v", err))
		}*/

	return nil
}

func (handler *Handler) ValidateToken (ctx context.Context, reqToken *pb.TokenInfo, tokenInfoResponse *pb.TokenInfo) error {
	// Decode token
	claims, err := handler.tokenService.Decode(reqToken.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	tokenInfoResponse.Valid = true

	return nil
}