package main

import (
	"container/heap"
)

type Tweet struct {
	id   int
	time int
	next *Tweet
}

type TweetHeap []*Tweet

type Twitter struct {
	time      int
	followers map[int]map[int]bool
	tweets    map[int]*Tweet
}

func (h TweetHeap) Len() int {
	return len(h)
}

func (h TweetHeap) Less(i, j int) bool {
	return h[i].time > h[j].time
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TweetHeap) Push(t interface{}) {
	*h = append(*h, t.(*Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() Twitter {
	return Twitter{0, map[int]map[int]bool{}, map[int]*Tweet{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets[userId] = &Tweet{
		id:   tweetId,
		time: this.time,
		next: this.tweets[userId],
	}
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweets := TweetHeap{}

	if _, ok := this.tweets[userId]; ok {
		heap.Push(&tweets, this.tweets[userId])
	}

	for followingId := range this.followers[userId] {
		if this.followers[userId][followingId] {
			if _, ok := this.tweets[followingId]; ok {
				heap.Push(&tweets, this.tweets[followingId])
			}
		}
	}

	tweetsFeed := []int{}
	for len(tweetsFeed) < 10 && len(tweets) > 0 {
		mostRecent := heap.Pop(&tweets).(*Tweet)
		tweetsFeed = append(tweetsFeed, mostRecent.id)
		if mostRecent.next != nil {
			heap.Push(&tweets, mostRecent.next)
		}
	}

	return tweetsFeed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followers[followerId]; !ok {
		this.followers[followerId] = make(map[int]bool)
	}
	this.followers[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followers[followerId][followeeId] {
		this.followers[followerId][followeeId] = false
	}
}

func main() {
	obj := Constructor()
	obj.PostTweet(1, 5)
	obj.Follow(1, 2)
	obj.Follow(2, 1)
	obj.GetNewsFeed(2)
	obj.PostTweet(2, 6)
	obj.GetNewsFeed(1)
	obj.GetNewsFeed(2)
	obj.Unfollow(2, 1)
	obj.GetNewsFeed(1)
	obj.GetNewsFeed(2)
	obj.Unfollow(1, 2)
	obj.GetNewsFeed(1)
	obj.GetNewsFeed(2)
}
