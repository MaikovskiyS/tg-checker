package provider

import (
	"context"

	pb "github.com/tg-checker/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CheckerClient представляет клиент для взаимодействия с gRPC-сервисом
type CheckerClient struct {
	conn   *grpc.ClientConn
	client pb.TelegramCheckerClient
}

// NewCheckerClient создает новый gRPC клиент
func NewCheckerClient(target string) (*CheckerClient, error) {
	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewTelegramCheckerClient(conn)

	return &CheckerClient{
		conn:   conn,
		client: client,
	}, nil
}

// CheckUserInChannel проверяет, находится ли пользователь в канале
func (c *CheckerClient) CheckUserInChannel(
	ctx context.Context,
	botToken,
	channelLink string,
	userID int64) (*pb.CheckUserResponse, error) {
	req := &pb.CheckUserRequest{
		BotToken:    botToken,
		ChannelLink: channelLink,
		UserId:      userID,
	}

	return c.client.CheckUserInChannel(ctx, req)
}

// Close закрывает соединение с gRPC сервером
func (c *CheckerClient) Close() error {
	return c.conn.Close()
}
