package linkedlist_test

import (
	linkedlist "go-data-structure/practice_qns/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			tmap := linkedlist.Constructor(2)
			tmap.Put(1, 1)
			tmap.Put(2, 2)

			got := tmap.Get(1)
			assert.Equal(t, 1, got)

			tmap.Put(3, 3)

			got = tmap.Get(2)
			assert.Equal(t, -1, got)

			tmap.Put(4, 4)

			got = tmap.Get(1)
			assert.Equal(t, -1, got)
			got = tmap.Get(3)
			assert.Equal(t, 3, got)
			got = tmap.Get(4)
			assert.Equal(t, 4, got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
