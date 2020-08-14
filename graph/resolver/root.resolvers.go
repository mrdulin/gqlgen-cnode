package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrdulin/gqlgen-cnode/graph/generated"
	"github.com/mrdulin/gqlgen-cnode/graph/model"
)

func (r *mutationResolver) ValidateAccessToken(ctx context.Context, accesstoken string) (*model.User, error) {
	return r.UserService.ValidateAccessToken(accesstoken), nil
}

func (r *queryResolver) Topics(ctx context.Context, params model.TopicsRequestParams) ([]*model.Topic, error) {
	return r.TopicService.GetTopicsByPage(&params), nil
}

func (r *queryResolver) Topic(ctx context.Context, params model.TopicRequestParams) (*model.TopicDetail, error) {
	return r.TopicService.GetTopicById(&params), nil
}

func (r *queryResolver) User(ctx context.Context, loginname string) (*model.UserDetail, error) {
	return r.UserService.GetUserByLoginname(loginname), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
