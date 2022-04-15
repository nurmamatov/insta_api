package services

import (
	"fmt"

	"tasks/Instagram_clone/insta_api/config"
	pc "tasks/Instagram_clone/insta_api/genproto/comment_proto"
	pp "tasks/Instagram_clone/insta_api/genproto/post_proto"
	pu "tasks/Instagram_clone/insta_api/genproto/user_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pu.UserServiceClient
	PostService() pp.PostServiceClient
	CommentService() pc.CommentServiceClient
}

type serviceManager struct {
	postService    pp.PostServiceClient
	userService    pu.UserServiceClient
	commentService pc.CommentServiceClient
}

func (s *serviceManager) PostService() pp.PostServiceClient {
	return s.postService
}
func (s *serviceManager) UserService() pu.UserServiceClient {
	return s.userService
}
func (s *serviceManager) CommentService() pc.CommentServiceClient {
	return s.commentService
}

func NewServiceManager(conf *config.Services) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CommentServiceHost, conf.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		postService:    pp.NewPostServiceClient(connPost),
		userService:    pu.NewUserServiceClient(connUser),
		commentService: pc.NewCommentServiceClient(connComment),
	}

	return serviceManager, nil
}
