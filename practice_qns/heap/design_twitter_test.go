package heap_test

import (
	"go-data-structure/practice_qns/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			twitter := heap.Constructor()
			twitter.PostTweet(1, 5)

			got := twitter.GetNewsFeed(1)
			assert.Equal(t, []int{5}, got)

			twitter.Follow(1, 2)
			twitter.PostTweet(2, 6)

			got = twitter.GetNewsFeed(1)
			assert.Equal(t, []int{6, 5}, got)

			twitter.Unfollow(1, 2)

			got = twitter.GetNewsFeed(1)
			assert.Equal(t, []int{5}, got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
