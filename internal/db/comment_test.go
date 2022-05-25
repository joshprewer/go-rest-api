//go:build integration

package db

import (
	"fmt"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		fmt.Println("testing the creation of comments")
	})
}
