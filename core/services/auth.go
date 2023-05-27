package services

import (
	"errors"
	"github.com/chopper-c2-framework/c2-chopper/core/config"
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	FrameworkConfig *config.Config
	UserService     *UserService
}

type JWTData struct {
	Username string
	jwt.RegisteredClaims
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		UserService: userService,
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

	return s.generateToken(username, user.ID.String())
}
func (s AuthService) Register(username string, password string) (string, error) {
	user, err := s.UserService.FindUserByUsernameOrError(username)
	if err != nil {
		log.Debugf("Register error finding user to validate password %v\n", err)
		return "", err
	}
	if user != nil {
		log.Debugf("Register error user already exists %v\n", err)
		return "", errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Debugf("Register error hashing password %v\n", err)
		return "", err
	}

	newUser := &entity.UserModel{Username: username, Password: string(hashedPassword)}
	err = s.UserService.CreateUser(newUser)

	if err != nil {
		log.Debugf("Register error creating user %v\n", err)
		return "", err
	}
	return s.generateToken(username, newUser.ID.String())
}

func (s AuthService) generateToken(username string, id string) (string, error) {

	expirationTime := time.Now().Add(24 * 3 * time.Hour)
	claims := &JWTData{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(s.FrameworkConfig.SecretToken)
	if err != nil {
		log.Debugf("GenerateToken error signing token %v\n", err)
		return "", err
	}

	return tokenString, nil
}

func (s AuthService) parseToken(token string) (string, error) {
	var myClaims *JWTData
	parsedToken, err := jwt.ParseWithClaims(token, myClaims, func(token *jwt.Token) (interface{}, error) {
		return s.FrameworkConfig.SecretToken, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(JWTData); ok && parsedToken.Valid {
		return claims.Subject, nil
	}

	return "", errors.New("unable to parse token")
}
