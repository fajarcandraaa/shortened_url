package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/shortened_url/internal/entity"
	"github.com/fajarcandraaa/shortened_url/internal/presentation"
	"github.com/fajarcandraaa/shortened_url/internal/repositories"
	"github.com/fajarcandraaa/shortened_url/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateShortenedUrl(t *testing.T) {
	db, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Url{})

	r := repositories.NewRepository(db)
	s := service.NewUrlShortenedUrlService(r)

	t.Run("feature generate shortened url : if url is valid, expected no error", func(t *testing.T) {
		var (
			ctx        = context.Background()
			urlExample = "http://www.google.com"
		)
		payload := presentation.ShortenedRequest{
			URL: urlExample,
		}

		resp, err := s.NewShortenedUrl(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, resp)
	})

	t.Run("feature generate shortened url : if url is not valid, expected  error", func(t *testing.T) {
		var (
			ctx        = context.Background()
			urlExample = ""
		)
		payload := presentation.ShortenedRequest{
			URL: urlExample,
		}

		resp, err := s.NewShortenedUrl(ctx, payload)
		require.Error(t, err)
		assert.Nil(t, resp)
	})
}
