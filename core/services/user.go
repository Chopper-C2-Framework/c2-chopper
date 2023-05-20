package services

import (
	orm "github.com/chopper-c2-framework/c2-chopper/core/domain"
	entity "github.com/chopper-c2-framework/c2-chopper/core/domain/entity"

	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

type UserService struct {
	ORMConnection *orm.ORMConnection
	repo          entity.TransactionRepository
}

func NewUserService(db *orm.ORMConnection) *UserService {
	logger := log.New()

	repo := entity.NewGormRepository(db.Db, logger, "Teams")

	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(newUser *entity.UserModel) error {

	newPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)

	if err != nil {
		return err
	}

	newUser.Password = string(newPassword)

	err = u.repo.Create(newUser)
	if err != nil {
		log.Debugf("[-] failed to create")
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(id string, updatedUser *entity.UserModel) error {
	user, err := u.FindUserOrError(id)
	if err != nil {
		return err
	}

	user.Teams = updatedUser.Teams

	err = u.repo.Save(user)

	if err != nil {
		log.Debugf("[-] Error updating user")

		return err
	}

	return nil
}

func (u *UserService) UpdateUserPassword(id string, newPassword string) error {
	user, err := u.FindUserOrError(id)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)

	if err != nil {
		log.Debugf("Error hashing password %v\n", err)
		return err
	}

	user.Password = string(hashedPassword)

	err = u.repo.Save(user)
	if err != nil {
		log.Debugf("Error updating user")
	}

	return nil

}

func (u *UserService) FindUserOrError(id string) (*entity.UserModel, error) {
	var user entity.UserModel
	err := u.repo.GetOneByID(&user, id)

	if err != nil {
		log.Debugf("[-] Error finding user")
		return nil, err
	}

	return &user, nil
}

func (u *UserService) FindAll() ([]entity.UserModel, error) {
	var users []entity.UserModel

	err := u.repo.GetAll(users)

	if err != nil {
		return users, err
	}

	return users, nil
}
