package deleting

type tweetRepository interface {
	DeleteTweet(ID string, UserID string) error
}

type TweetService struct {
	repository tweetRepository
}

func NewTweetService(repository tweetRepository) *TweetService {
	return &TweetService{repository: repository}
}

func (ts *TweetService) DeleteTweet(ID string, UserID string) error {
	err := ts.repository.DeleteTweet(ID, UserID)
	if err != nil {
		return err
	}

	return nil
}
