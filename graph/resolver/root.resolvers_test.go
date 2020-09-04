package resolver_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/mrdulin/gqlgen-cnode/graph/generated"
	"github.com/mrdulin/gqlgen-cnode/graph/model"
	"github.com/mrdulin/gqlgen-cnode/graph/resolver"
	"github.com/mrdulin/gqlgen-cnode/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	loginname = "mrdulin"
	avatarURL = "avatar.jpg"
	score     = 50
	createAt  = "1900-01-01"
)

func TestMutationResolver_ValidateAccessToken(t *testing.T) {

	t.Run("should validate accesstoken correctly", func(t *testing.T) {
		testUserService := new(mocks.MockedUserService)
		resolvers := resolver.Resolver{UserService: testUserService}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		ue := model.UserEntity{ID: "123", User: model.User{Loginname: &loginname, AvatarURL: &avatarURL}}
		testUserService.On("ValidateAccessToken", mock.AnythingOfType("string")).Return(&ue)
		var resp struct {
			ValidateAccessToken struct{ ID, Loginname, AvatarUrl string }
		}
		q := `
      mutation { 
        validateAccessToken(accesstoken: "abc") { 
          id, 
          loginname, 
          avatarUrl 
        } 
      }
    `
		c.MustPost(q, &resp)
		testUserService.AssertExpectations(t)
		require.Equal(t, "123", resp.ValidateAccessToken.ID)
		require.Equal(t, "mrdulin", resp.ValidateAccessToken.Loginname)
		require.Equal(t, "avatar.jpg", resp.ValidateAccessToken.AvatarUrl)
	})

}

func TestQueryResolver_User(t *testing.T) {
	t.Run("should query user correctly", func(t *testing.T) {
		testUserService := new(mocks.MockedUserService)
		resolvers := resolver.Resolver{UserService: testUserService}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		u := model.UserDetail{User: model.User{Loginname: &loginname, AvatarURL: &avatarURL}, Score: &score, CreateAt: &createAt}
		testUserService.On("GetUserByLoginname", mock.AnythingOfType("string")).Return(&u)
		var resp struct {
			User struct {
				Loginname, AvatarURL, CreateAt string
				Score                          int
			}
		}
		q := `
      query GetUser($loginname: String!) { 
        user(loginname: $loginname) { 
          loginname
          avatarUrl 
          createAt 
          score 
        } 
      }
    `
		c.MustPost(q, &resp, client.Var("loginname", "mrdulin"))
		testUserService.AssertCalled(t, "GetUserByLoginname", "mrdulin")
		require.Equal(t, "mrdulin", resp.User.Loginname)
		require.Equal(t, "avatar.jpg", resp.User.AvatarURL)
		require.Equal(t, 50, resp.User.Score)
		require.Equal(t, "1900-01-01", resp.User.CreateAt)
	})
}
