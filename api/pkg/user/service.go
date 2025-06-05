package user

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Service interface {
	Register(ctx context.Context, mnemonic []string) (string, string, string, error)
	Login(ctx context.Context, mnemonic []string) (string, string, string, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(
	repository Repository,
) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) Register(ctx context.Context, mnemonic []string) (string, string, string, error) {
	if !utils.ValidateMnemonic(mnemonic) {
		log.Println("register: invalid mnemonic")
		return "", "", "", customerrors.ErrInvalidMnemonic
	}
	seed := utils.MnemonicToSeed(mnemonic)
	_, err := s.repo.GetBySeed(ctx, seed)
	if err == nil {
		log.Println("register: user already exists")
		return "", "", "", customerrors.ErrUserAlreadyExists
	} else if !errors.Is(err, customerrors.ErrUserNotFound) {
		log.Printf("register: failed to get user by seed: %v", err)
		return "", "", "", err
	}
	privateKey, publicKey, err := utils.GenerateKeyPair(2048)
	if err != nil {
		log.Printf("register: failed to generate key pair: %v", err)
		return "", "", "", err
	}
	privateKeyEncrypted := utils.PrivateKeyToString(privateKey)
	publicKeyEncrypted, err := utils.PublicKeyToString(publicKey)
	if err != nil {
		log.Printf("register: failed to encode public key: %v", err)
		return "", "", "", err
	}
	userID, err := s.repo.Register(ctx, seed, privateKeyEncrypted, publicKeyEncrypted)
	if err != nil {
		log.Printf("register: repository error: %v", err)
		return "", "", "", err
	}

	jwtToken, err := utils.EncodeJWT(userID)
	if err != nil {
		log.Printf("register: failed to encode jwt: %v", err)
		return "", "", "", err
	}
	log.Printf("user registered: %s", userID.Hex())

	return jwtToken, publicKeyEncrypted, privateKeyEncrypted, nil
}

func (s *service) Login(ctx context.Context, mnemonic []string) (string, string, string, error) {
	if !utils.ValidateMnemonic(mnemonic) {
		log.Println("login: invalid mnemonic")
		return "", "", "", customerrors.ErrInvalidMnemonic
	}
	seed := utils.MnemonicToSeed(mnemonic)
	user, err := s.repo.GetBySeed(ctx, seed)
	if err != nil {
		log.Printf("login: failed to get user by seed: %v", err)
		return "", "", "", err
	}
	jwtToken, err := utils.EncodeJWT(user.ID)
	if err != nil {
		log.Printf("login: failed to encode jwt: %v", err)
		return "", "", "", err
	}
	log.Printf("user logged in: %s", user.ID.Hex())
	return jwtToken, user.PublicKey, user.PrivateKey, nil
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
