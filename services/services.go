package services

import (
	"fmt"

	"tasks/Instagram_clone/insta_api/config"
	pc "tasks/Instagram_clone/insta_api/genproto/comment_proto"
	pp "tasks/Instagram_clone/insta_api/genproto/post_proto"
	pu "tasks/Instagram_clone/insta_api/genproto/user_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IServiceManager interface {
	UserService() pu.UserServiceClient
	PostService() pp.PostServiceClient
	CommentService() pc.CommentServiceClient
}

// type serviceManager struct {
// 	postService    pp.PostServiceClient
// 	userService    pu.UserServiceClient
// 	commentService pc.CommentServiceClient
// }

// Client
type GrpcClient struct {
	cfg         config.Services
	connections map[string]interface{}
}

func NewServiceManager(cfg *config.Services) (*GrpcClient, error) {

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{
		cfg: *cfg,
		connections: map[string]interface{}{
			"post_service":    pp.NewPostServiceClient(connPost),
			"user_service":    pu.NewUserServiceClient(connUser),
			"comment_service": pc.NewCommentServiceClient(connComment),
		},
	}, nil
}

func (g *GrpcClient) PostService() pp.PostServiceClient {
	return g.connections["post_service"].(pp.PostServiceClient)
}
func (g *GrpcClient) CommentService() pc.CommentServiceClient {
	return g.connections["comment_service"].(pc.CommentServiceClient)
}
func (g *GrpcClient) UserService() pu.UserServiceClient {
	return g.connections["user_service"].(pu.UserServiceClient)
}
