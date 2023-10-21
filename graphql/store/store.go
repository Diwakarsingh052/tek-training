package store

import "graphql/graph/model"

type Storer interface {
	AddVideo(video *model.Video) (*model.Video, error)
	UpdateVideo(videoID string, input model.UpdateVideoInput) (*model.Video, error)
	AllVideos() []*model.Video
	FindVideosByID(videoID string) (*model.Video, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}
