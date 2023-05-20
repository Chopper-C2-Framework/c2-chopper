package grpc

import (
	"github.com/chopper-c2-framework/c2-chopper/core/domain/entity"
	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func ConvertTeamToProto(team *entity.TeamModel) *proto.Team {
	var users []*proto.User
	for _, m := range team.Members {
		users = append(users, ConvertUserToProto(m))
	}
	protoTeam := &proto.Team{
		Id:      team.ID.String(),
		Name:    team.Name,
		Members: users,
	}

	return protoTeam
}

func ConvertUserToProto(user *entity.UserModel) *proto.User {

	return &proto.User{
		Username: user.Username,
		Id:       user.ID.String(),
	}
}

func ConvertProtoToUser(user *proto.User) (*entity.UserModel, error) {
	parsedUuid, err := uuid.Parse(user.Id)
	if err != nil {
		log.Debugf("ConvertProtoToUser: error parsing uuid %v\n", err)
		return nil, err
	}
	return &entity.UserModel{
		UUIDModel: entity.UUIDModel{ID: parsedUuid},
		Username:  user.Username,
	}, nil
}
