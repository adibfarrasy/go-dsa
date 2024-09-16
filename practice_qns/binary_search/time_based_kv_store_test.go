package binarysearch_test

import (
	binarysearch "go-data-structure/practice_qns/binary_search"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeBasedKVStore(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			tmap := binarysearch.Constructor()
			tmap.Set("foo", "bar", 1)

			got := tmap.Get("foo", 1)
			assert.Equal(t, "bar", got)

			got = tmap.Get("foo", 3)
			assert.Equal(t, "bar", got)

			tmap.Set("foo", "bar2", 4)

			got = tmap.Get("foo", 4)
			assert.Equal(t, "bar2", got)

			got = tmap.Get("foo", 5)
			assert.Equal(t, "bar2", got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
