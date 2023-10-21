package graph

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"graphql/graph/model"
	"graphql/store"
	"graphql/store/mockstore"
	"testing"
)

func TestQueryResolver_FindVideosByID(t *testing.T) {
	mockVideo := &model.Video{
		ID:    "1",
		Title: "My first Video",
		URL:   "https://my-first-video.com",
		Author: &model.User{
			ID:   "user1",
			Name: "Bob",
		},
	}

	tt := []struct {
		name             string // Name of the test case
		query            string
		err              error
		expectedVideo    *model.Video
		mockStoreService func(m *mockstore.MockStorer) // Mock service function
	}{
		{
			name: "Ok",
			query: `query findByVideoId{
  findVideosById(VideoId: "1") {
    id
    title
    url
    author {
      id
      name
    }
  }
}`,
			err: nil,
			expectedVideo: &model.Video{
				ID:    "1",
				Title: "My first Video",
				URL:   "https://my-first-video.com",
				Author: &model.User{
					ID:   "user1",
					Name: "Bob",
				},
			},
			mockStoreService: func(ms *mockstore.MockStorer) {
				ms.EXPECT().FindVideosByID(gomock.Any()).Times(1).
					Return(mockVideo, nil)
			},
		},
	}
	for _, tc := range tt {
		// each test case is run as a subtest
		t.Run(tc.name, func(t *testing.T) {

			// create a new gomock controller
			ctrl := gomock.NewController(t)

			// create a new mock store for each test case
			mockStore := mockstore.NewMockStorer(ctrl)

			// create a new store instance with the mock store
			s := store.NewStore(mockStore)

			// create a new client instance
			c := client.New(handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{S: s}})))

			// apply the mock behavior for the store service for this test case
			tc.mockStoreService(mockStore)
			var resp struct {
				Video *model.Video `json:"findVideosById"` // where video data will be stored if found
			}

			//Post sends a http POST request to the graphql endpoint with the given query then unpacks the response into the given object.
			err := c.Post(tc.query, &resp)

			// if there was no error, assert that no error was expected
			if err == nil {
				require.NoError(t, tc.err)
			}
			// if there was an error, assert that an error was expected
			if tc.err != nil {
				require.Error(t, err)
			}

			// assert that the expected video matches the response video
			require.Equal(t, tc.expectedVideo, resp.Video)

		})

	}
}
