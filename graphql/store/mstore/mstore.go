package mstore

import (
	"errors"
	"fmt"
	"graphql/graph/model"
)

type Service struct {
	videoStore map[string]*model.Video
}

func NewService() Service {
	return Service{videoStore: make(map[string]*model.Video)}
}

func (s *Service) AddVideo(video *model.Video) (*model.Video, error) {
	fmt.Println("doing heavy db operation specific stuff")
	s.videoStore[video.ID] = video
	return video, nil
}

func (s *Service) UpdateVideo(videoID string, input model.UpdateVideoInput) (*model.Video, error) {
	v, ok := s.videoStore[videoID]
	if !ok {
		return nil, errors.New("video not found")
	}
	v.URL = *input.URL
	v.Title = *input.Title
	s.videoStore[videoID] = v
	return v, nil
}

func (s *Service) AllVideos() []*model.Video {
	var videos []*model.Video
	for _, v := range s.videoStore {
		v := v
		fmt.Println(v)
		videos = append(videos, v)
	}

	return videos
}

func (s *Service) FindVideosByID(videoID string) (*model.Video, error) {
	v, ok := s.videoStore[videoID]
	if !ok {
		return nil, errors.New("video not found")
	}
	return v, nil
}
