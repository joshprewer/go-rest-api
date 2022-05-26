//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func CreateToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("missionimpossible"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestPostComment(t *testing.T) {
	t.Run("test can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+CreateToken()).
			SetBody(`{"slug": "/", "author": "Josh", "body": "hello world"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("test cannot post comment without jwt", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "Josh", "body": "body"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}
