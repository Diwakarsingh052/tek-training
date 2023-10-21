package mstore

import (
	"errors"
	"fmt"
	"graphql/graph/model"
)

// Declaring a struct named Service that includes a map that takes a string as key and pointer to Video as value
type Service struct {
	videoStore map[string]*model.Video
}

// Initializing the Service struct by creating an empty map for videoStore
func NewService() Service {
	return Service{videoStore: make(map[string]*model.Video)}
}

// Method to add a video to the videoStore. Accepts a pointer to a video, performs a dummy database operation, then adds the video to the videoStore
func (s *Service) AddVideo(video *model.Video) (*model.Video, error) {
	fmt.Println("doing heavy db operation specific stuff") // Displaying a message to signify performing database manipulation
	s.videoStore[video.ID] = video                         // Adding the video to the map under the key of video.ID
	return video, nil                                      // Returning the added video and nil (indicating no error)
}

// Method to update an existing video in the videoStore, if it exists. Otherwise, returns an error
func (s *Service) UpdateVideo(videoID string, input model.UpdateVideoInput) (*model.Video, error) {
	v, ok := s.videoStore[videoID] // Trying to get the video with the provided videoID
	if !ok {                       // If the video is not found
		return nil, errors.New("video not found") // Return nil and an error saying the video was not found
	}
	v.URL = *input.URL        // Update the URL of the fetched video
	v.Title = *input.Title    // Update the title of the fetched video
	s.videoStore[videoID] = v // Overwrite the video in the map with the updated video
	return v, nil             // Return the updated video and nil (indicating no error)
}

// Method to return all videos in the videoStore
func (s *Service) AllVideos() []*model.Video {
	var videos []*model.Video        // Initialize an empty slice of pointers to videos
	for _, v := range s.videoStore { // Iterate over all entries in the videoStore map
		//v := v
		fmt.Println(v)             // Display the video
		videos = append(videos, v) // Add the video to the videos slice
	}

	return videos // Return the slice containing all videos
}

// Method to find and return a video by its ID, or an error if it is not found
func (s *Service) FindVideosByID(videoID string) (*model.Video, error) {
	v, ok := s.videoStore[videoID] // Trying to get the video by the provided videoID
	if !ok {                       // If the video is not found
		return nil, errors.New("video not found") // Return nil and an error saying the video was not found
	}
	return v, nil // Return the video found
}
