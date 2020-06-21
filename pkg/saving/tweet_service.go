package saving

import (
	"errors"
	"tweetgo/pkg/domain"
)

type tweetRepository interface {
	SaveTweet(t domain.Tweet) (bool, error)
}

type TweetService struct {
	repository tweetRepository
}

func NewTweetService(repository tweetRepository) *TweetService {
	return &TweetService{repository: repository}
}

func (ts *TweetService) SaveTweet(t domain.Tweet) (bool, error) {
	status, err := ts.repository.SaveTweet(t)
	if err != nil {
		return false, err
	}

	if status == false {
		return status, errors.New("tweet could not be saved")
	}

	return status, nil
}
