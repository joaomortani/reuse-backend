package services

import (
	"errors"

	"reuse-api/config"
	dto "reuse-api/dto/user"
	models "reuse-api/models/user"
	repositories "reuse-api/repositories/user"

	"reuse-api/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(dto.RegisterRequest) (*models.User, error)
	Login(dto.LoginRequest) (string, string, error)
	GetById(id uint) (*models.User, error)
	RefreshToken(token string) (string, string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{repo: r}
}

func hashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("erro ao gerar hash da senha")
	}
	return string(hashed)
}

func checkPassword(hash, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

func (s *userService) GetById(id uint) (*models.User, error) {
	return s.repo.GetById(id)
}

func (s *userService) Register(req dto.RegisterRequest) (*models.User, error) {
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword(req.Password),
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(req dto.LoginRequest) (string, string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || !checkPassword(user.Password, req.Password) {
		return "", "", errors.New("email ou senha inválidos")
	}

	accessToken, refreshTokenString, err := utils.GenerateTokens(user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshTokenString, nil
}

func (s *userService) RefreshToken(token string) (string, string, error) {
	claims := &jwt.MapClaims{}
	refreshToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token inválido")
		}
		return []byte(config.GetEnv("JWT_SECRET")), nil
	})
	if err != nil || !refreshToken.Valid {
		return "", "", errors.New("token inválido")
	}
	userId := (*claims)["sub"].(float64)
	user, err := s.repo.GetById(uint(userId))
	if err != nil {
		return "", "", err
	}

	newToken, newRefreshToken, err := utils.GenerateTokens(user)
	if err != nil {
		return "", "", err
	}

	return newToken, newRefreshToken, nil
}
