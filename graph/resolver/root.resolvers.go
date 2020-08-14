package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mrdulin/gqlgen-cnode/graph/generated"
	"github.com/mrdulin/gqlgen-cnode/graph/model"
)

func (r *queryResolver) Topics(ctx context.Context, params model.TopicsRequestParams) ([]*model.Topic, error) {
	return r.TopicService.GetTopicsByPage(params), nil
}

func (r *queryResolver) Topic(ctx context.Context, id string) (*model.TopicDetail, error) {
	return r.TopicService.GetTopicById(id), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
