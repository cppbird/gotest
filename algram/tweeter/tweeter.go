package main

// Twitter .
type Twitter struct {
	feedList []art
	follow   map[int][]int
}
type art struct {
	UserID int
	artID  int
}

// Constructor.
func Constructor() Twitter {
	return Twitter{
		feedList: []art{},
		follow:   make(map[int][]int),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	a := art{
		UserID: userId,
		artID:  tweetId,
	}
	this.feedList = append(this.feedList, a)

}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	count := 0
	res := []int{}
	follow := append(this.follow[userId], userId)
	if len(follow) == 0 {
		return res
	}

	for i := len(this.feedList) - 1; i >= 0; i-- {
		for _, v := range follow {
			if this.feedList[i].UserID == v {
				count++
				res = append(res, this.feedList[i].artID)
			}
			if count >= 10 {
				return res
			}
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	for _, v := range this.follow[followerId] {
		if v == followeeId {
			return
		}
	}
	this.follow[followerId] = append(this.follow[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for k, v := range this.follow[followerId] {
		if v == followeeId {
			this.follow[followerId] = append(this.follow[followerId][:k], this.follow[followerId][k+1:]...)
			break
		}
	}
}
