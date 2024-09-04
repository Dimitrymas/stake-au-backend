package user

import (
	"backend/api/pkg/models"
	"github.com/Nerzal/gocloak/v13"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Register(keycloakID string, referralCode *string) error
	UpdateRandomKey(keycloakID string) error
	GetByKeycloakID(keycloakID string) (*models.User, error)
	GetByID(userID primitive.ObjectID) (*models.User, error)
	PlusBalance(userID primitive.ObjectID, balance float64) error
	UpdateBalance(userID primitive.ObjectID, balance float64) error
	ExchangeCodeForAccessToken(code string) (*gocloak.JWT, error)
	RefreshToken(refreshToken string) (*gocloak.JWT, error)
	GetCountOfUsers() (int64, error)
	VerifyToken(token string) (*gocloak.IntroSpectTokenResult, error)
	Logout(refreshToken string) error
	ChangeName(userID primitive.ObjectID, name string) error
	GetByName(name string) (*models.User, error)
	GetAll() (*[]models.User, error)
	GetByRandomKey(secret string) (*models.User, error)
	GetUserByReferralCode(referralCode string) (*models.User, error)
	GetByReferralID(referralID primitive.ObjectID) ([]*models.User, error)
	GetCashbackPercent(levelID primitive.ObjectID) (float64, error)
	PlusWager(userID primitive.ObjectID, wager float64) error
	MinusWager(userID primitive.ObjectID, wager float64) error
}

type service struct {
	repository Repository
}

func NewService(
	repository Repository,
) Service {
	keycloak := NewKeycloak()
	return &service{
		repository: repository,
	}
}

func (s *service) Register(login string, password string) error {
	return s.repository.Register(login, password)
}
