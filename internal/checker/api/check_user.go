package api

import (
	"context"
	"fmt"

	pb "github.com/tg-checker/gen/proto"
	"github.com/tg-checker/pkg/hasher"
)

// CheckUserInChannel проверяет, находится ли пользователь в канале
func (a *Api) CheckUserInChannel(ctx context.Context, req *pb.CheckUserRequest) (*pb.CheckUserResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}

	isMember, err := a.telegramClient.GetChatMember(req.BotToken, req.ChannelLink, req.UserId)
	if err != nil {
		a.log.Error().Err(err).
			Int64("user_id", req.UserId).
			Str("channel_link", req.ChannelLink).
			Str("bot_token", req.BotToken).
			Msg("get chat member")

		return nil, err
	}

	if isMember {
		channelID := stringToChannelID(req.ChannelLink)

		_, err = a.userRepo.AddUser(req.UserId, channelID)
		if err != nil {
			a.log.Error().Err(err).
				Int64("user_id", req.UserId).
				Int64("channel_id", channelID).
				Msg("add user")

			return nil, err
		}

		return &pb.CheckUserResponse{
			IsMember: true,
		}, nil
	}

	return &pb.CheckUserResponse{
		IsMember: false,
	}, nil
}

// stringToChannelID преобразует ссылку на канал в числовой идентификатор
func stringToChannelID(channelLink string) int64 {
	return hasher.StringToInt64(channelLink)
}
