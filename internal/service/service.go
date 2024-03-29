package service

import (
	"github.com/PolyProjectOPD/Backend/internal/entity"
	"github.com/PolyProjectOPD/Backend/internal/repository"
	"github.com/PolyProjectOPD/Backend/internal/storage"
	"github.com/PolyProjectOPD/Backend/pkg/auth"
	"github.com/PolyProjectOPD/Backend/pkg/hash"
	"time"
)

type Users interface {
	SignUp(input entity.UserSignUpInput, fileBody, fileType string) (int, string, error)
	SignIn(input entity.UserSignInInput, userAgent string) (Tokens, error)
	RefreshTokens(input entity.UserRefreshInput, userAgent string) (Tokens, error)

	Profile(userID int) (entity.UserProfile, error)
	Logout(userID int) error
	SignOut(userID int, userAgent string) error

	Existence(input entity.UserExistenceInput) bool
}

type Spheres interface {
	GetAll() ([]entity.Sphere, error)
	GetSkills(spheres []entity.Sphere) ([]entity.Skill, error)
}

type Skills interface {
	GetAll() ([]entity.Skill, error)
}

type Services struct {
	Users   Users
	Spheres Spheres
	Skills  Skills
	Storage *storage.Storage
}

type Deps struct {
	Repos           *repository.Repositories
	Storage         *storage.Storage
	Hasher          hash.PasswordHasher
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Storage, deps.Hasher, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	spheresService := NewSpheresService(deps.Repos.Spheres)

	return &Services{
		Users:   usersService,
		Spheres: spheresService,
		Storage: deps.Storage,
	}
}
