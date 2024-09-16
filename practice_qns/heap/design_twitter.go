package heap

import "container/heap"

type Tweet struct {
	count      int
	tweetID    int
	followeeID int
	index      int
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	count     int
	tweetMap  map[int][]*Tweet
	followMap map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		tweetMap:  map[int][]*Tweet{},
		followMap: map[int]map[int]bool{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.tweetMap[userId]; !ok {
		this.tweetMap[userId] = []*Tweet{}
	}

	this.tweetMap[userId] = append(this.tweetMap[userId], &Tweet{count: this.count, tweetID: tweetId})
	this.count -= 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	minHeap := TweetHeap{}
	heap.Init(&minHeap)

	if _, ok := this.followMap[userId]; !ok {
		this.followMap[userId] = map[int]bool{}
	}

	// to see one's own tweets
	this.followMap[userId][userId] = true

	for followeeId := range this.followMap[userId] {
		if _, ok := this.tweetMap[followeeId]; ok {
			index := len(this.tweetMap[followeeId]) - 1
			tweet := this.tweetMap[followeeId][index]
			tweet.index = index - 1
			heap.Push(&minHeap, tweet)
		}
	}

	for len(minHeap) > 0 && len(res) < 10 {
		tweet := heap.Pop(&minHeap).(*Tweet)
		res = append(res, tweet.tweetID)
		if tweet.index >= 0 {
			nextTweet := this.tweetMap[tweet.followeeID][tweet.index]
			heap.Push(&minHeap, &Tweet{count: nextTweet.count,
				tweetID:    nextTweet.tweetID,
				followeeID: tweet.followeeID,
				index:      tweet.index - 1})
		}
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; !ok {
		this.followMap[followerId] = map[int]bool{}
	}

	this.followMap[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.followMap[followerId]; ok {
		delete(this.followMap[followerId], followeeId)
	}
}
