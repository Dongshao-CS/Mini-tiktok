package utils

import (
	"context"
	"fmt"
	"github.com/shixiaocaia/tiktok/cmd/gatewaysvr/config"
	"github.com/shixiaocaia/tiktok/cmd/gatewaysvr/log"
	"github.com/shixiaocaia/tiktok/pkg/pb"
	"google.golang.org/grpc"
	"time"

	// 必须要导入这个包，否则grpc会报错
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
)

var (
	UserSvrClient  pb.UserServiceClient
	VideoClient    pb.VideoServiceClient
	FavoriteClient pb.FavoriteServiceClient
	CommentClient  pb.CommentServiceClient
	RelationClient pb.RelationServiceClient
	MessageClient  pb.MessageServiceClient
)

func InitSvrConn() {
	UserSvrClient = NewUserSvrClient(config.GetGlobalConfig().SvrConfig.UserSvrName)
	VideoClient = NewVideoClient(config.GetGlobalConfig().SvrConfig.VideoSvrName)
	FavoriteClient = NewFavoriteClient(config.GetGlobalConfig().SvrConfig.FavoriteSvrName)
	CommentClient = NewCommentClient(config.GetGlobalConfig().SvrConfig.CommentSvrName)
	RelationClient = NewRelationClient(config.GetGlobalConfig().SvrConfig.RelationSvrName)
	MessageClient = NewMessageClient(config.GetGlobalConfig().SvrConfig.MessageSvrName)
}

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // 10秒超时
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		// grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	log.Info("NewSvrConn success")
	return conn, nil
}

func GetUserSvrClient() pb.UserServiceClient {
	return UserSvrClient
}

func GetVideoSvrClient() pb.VideoServiceClient {
	return VideoClient
}

func GetCommentSvrClient() pb.CommentServiceClient {
	return CommentClient
}

func GetFavoriteSvrClient() pb.FavoriteServiceClient {
	return FavoriteClient
}

func GetRelationSvrClient() pb.RelationServiceClient {
	return RelationClient
}

func GetMessageSvrClient() pb.MessageServiceClient {
	return MessageClient
}

func NewUserSvrClient(svrName string) pb.UserServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewUserServiceClient(conn)
}

func NewVideoClient(svrName string) pb.VideoServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewVideoServiceClient(conn)
}

func NewFavoriteClient(svrName string) pb.FavoriteServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewFavoriteServiceClient(conn)
}

func NewCommentClient(svrName string) pb.CommentServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewCommentServiceClient(conn)
}

func NewRelationClient(svrName string) pb.RelationServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewRelationServiceClient(conn)
}

func NewMessageClient(svrName string) pb.MessageServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewMessageServiceClient(conn)
}
