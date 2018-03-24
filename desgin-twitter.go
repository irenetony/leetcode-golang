package main

import (
	"fmt"
	"github.com/berryjam/leetcode/util"
	"strings"
)

type Twitter struct {
	currentTimeStamp int
	followship       map[int][]int
	tweets           map[int][]Tweet
}

type Tweet struct {
	id        int
	timeStamp int
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.currentTimeStamp++
	t.tweets[userId] = append(t.tweets[userId], Tweet{
		id:        tweetId,
		timeStamp: t.currentTimeStamp,
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by
users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) (res []int) {
	followeeIds := t.followship[userId]
	pointMap := make(map[int]int, len(followeeIds)+1)
	pointMap[userId] = len(t.tweets[userId]) - 1
	for _, followeeId := range followeeIds {
		pointMap[followeeId] = len(t.tweets[followeeId]) - 1
	}
	for {
		if len(res) >= 10 || t.isCountFinish(pointMap) {
			break
		}
		maxTimeStamp := -1
		mostRecentFolloweeId := -1
		mostRecentTweetId := -1
		for followeeId, v := range pointMap {
			if v != -1 {
				if t.tweets[followeeId][v].timeStamp > maxTimeStamp {
					maxTimeStamp = t.tweets[followeeId][v].timeStamp
					mostRecentFolloweeId = followeeId
					mostRecentTweetId = t.tweets[followeeId][v].id
				}
			}
		}
		res = append(res, mostRecentTweetId)
		pointMap[mostRecentFolloweeId]--
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	if exist, _ := t.isFollower(followeeId, t.followship[followerId]); !exist {
		t.followship[followerId] = append(t.followship[followerId], followeeId)
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := t.followship[followerId]; ok {
		if exist, idx := t.isFollower(followeeId, t.followship[followerId]); exist {
			slice := t.followship[followerId]
			copy(slice[idx:], slice[idx+1:])
			slice = slice[:len(slice)-1]
			t.followship[followerId] = slice
		}
	}
}

func (t *Twitter) isFollower(userId int, followeeIds []int) (bool, int) {
	for i, v := range followeeIds {
		if userId == v {
			return true, i
		}
	}
	return false, -1
}

func (t *Twitter) isCountFinish(pointMap map[int]int) bool {
	for _, v := range pointMap {
		if v != -1 {
			return false
		}
	}
	return true
}

func main() {
	twitter := &Twitter{
		currentTimeStamp: 0,
		followship:       make(map[int][]int),
		tweets:           make(map[int][]Tweet),
	}

	// User 1 posts a new tweet (id = 5).
	twitter.PostTweet(1, 5)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6).
	twitter.PostTweet(2, 6)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))

	// User 1 unfollows user 2.
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))

	twitter.PostTweet(2, 7)

	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))

	twitter.Unfollow(1, 2)

	fmt.Printf("Tweets [%s]\n", strings.Join(util.IntArray2StringArray(twitter.GetNewsFeed(1)), ","))
}
