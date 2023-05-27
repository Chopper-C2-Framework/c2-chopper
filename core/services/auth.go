package services

import (
	"errors"
	"time"

	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	FrameworkConfig config.Config
	UserService     *UserService
}

type JWTData struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthService(userService *UserService, frameworkConfig config.Config) *AuthService {
	return &AuthService{
		UserService:     userService,
		FrameworkConfig: frameworkConfig,
	}
}

func (s AuthService) Login(username string, password string) (string, error) {
	user, err := s.UserService.FindUserByUsernameOrError(username)
	if err != nil {
		log.Debugf("Login error finding user to validate password %v\n", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Debugf("Login error comparing hashes %v\n", err)
		return "", err
	}

	return s.GenerateToken(user)
}
func (s AuthService) Register(username string, password string) (string, error) {
	_, err := s.UserService.FindUserByUsernameOrError(username)

	if err == nil {
		log.Debugf("Register error user already exists %v\n", err)
		return "", errors.New("user already exists")
	}

	newUser := &entity.UserModel{Username: username, Password: password}
	err = s.UserService.CreateUser(newUser)

	if err != nil {
		log.Debugf("Register error creating user %v\n", err)
		return "", err
	}

	log.Println("NewUser", newUser)
	return s.GenerateToken(newUser)
}

func (s AuthService) GenerateToken(user *entity.UserModel) (string, error) {

	expirationTime := time.Now().Add(24 * 3 * time.Hour)
	log.Println("expirationTime", expirationTime)
	claims := &JWTData{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	log.Println("expirationTime", claims, user)

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string

	log.Debugln(s.FrameworkConfig, s.FrameworkConfig.SecretToken)
	tokenString, err := token.SignedString([]byte(s.FrameworkConfig.SecretToken))
	if err != nil {
		log.Debugf("GenerateToken error signing token %v\n", err)
		return "", err
	}

	return tokenString, nil
}

func (s AuthService) ParseToken(token string) (*JWTData, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		log.Debugln("token", token)
		return []byte(s.FrameworkConfig.SecretToken), nil
	})

	log.Println("parsedToken", parsedToken, err)

	if err != nil {
		return nil, err
	}

	log.Println("here", parsedToken, err)

	if claims, ok := parsedToken.Claims.(*JWTData); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("unable to parse token")
}
