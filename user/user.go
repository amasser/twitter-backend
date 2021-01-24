package user

import (
	"github.com/arman-aminian/twitter-backend/model"
)

type Store interface {
	Create(*model.User) error
	Remove(username string) error

	GetByEmail(string) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	AddFollower(u *model.User, follower *model.User) error
	IsFollower(username, followerUsername string) (bool, error)

	AddTweet(u *model.User, t *model.Tweet) error
}
