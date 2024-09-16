package tries_test

import (
	"go-data-structure/practice_qns/tries"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	tests := []func(t *testing.T){
		func(t *testing.T) {
			wd := tries.NewWD()
			wd.AddWord("bad")
			wd.AddWord("dad")
			wd.AddWord("mad")

			got := wd.Search("pad")
			assert.Equal(t, false, got)

			got = wd.Search("bad")
			assert.Equal(t, true, got)

			got = wd.Search(".ad")
			assert.Equal(t, true, got)

			got = wd.Search("b..")
			assert.Equal(t, true, got)
		},
	}

	for _, test := range tests {
		test(t)
	}
}
