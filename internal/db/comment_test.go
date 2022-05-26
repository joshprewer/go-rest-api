//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/joshprewer/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})

		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "josh",
			Body:   "body",
		})

		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "update-slug",
			Author: "josh p",
			Body:   "body",
		})

		assert.NoError(t, err)

		newCmt, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "test slug",
			Author: "test author",
			Body:   "test body",
		})
		assert.NoError(t, err)
		assert.Equal(t, newCmt.ID, cmt.ID)
		assert.Equal(t, newCmt.Slug, "test slug")
		assert.Equal(t, newCmt.Author, "test author")
		assert.Equal(t, newCmt.Body, "test body")
	})
}
