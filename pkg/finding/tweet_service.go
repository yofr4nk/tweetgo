package finding

import "tweetgo/pkg/domain"

type tweetRepository interface {
	GetTweets(userID string, page int64) (domain.Tweets, error)
}

type TweetService struct {
	repository tweetRepository
}

func NewTweetService(tweetRepository tweetRepository) *TweetService {
	return &TweetService{repository: tweetRepository}
}

func (ts TweetService) GetTweets(userID string, page int64) (domain.Tweets, error) {
	tweets, err := ts.repository.GetTweets(userID, page)
	if err != nil {
		return domain.Tweets{}, err
	}

	return tweets, nil
}
