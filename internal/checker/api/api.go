package api

import (
	"github.com/rs/zerolog"
	pb "github.com/tg-checker/gen/proto"
	"google.golang.org/grpc"
)

type TelegramClient interface {
	GetChatMember(botToken, channelLink string, userId int64) (bool, error)
}

type UserRepository interface {
	AddUser(userId int64, channelID int64) (int64, error)
}

type Api struct {
	pb.UnimplementedTelegramCheckerServer
	telegramClient TelegramClient
	userRepo       UserRepository
	log            zerolog.Logger
}

func New(telegramClient TelegramClient, userRepo UserRepository, log zerolog.Logger) *Api {
	return &Api{
		telegramClient: telegramClient,
		userRepo:       userRepo,
		log:            log,
	}
}

func (s *Api) Register(grpcServer *grpc.Server) {
	pb.RegisterTelegramCheckerServer(grpcServer, s)
}
