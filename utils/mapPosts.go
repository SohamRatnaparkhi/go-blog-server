package utils

// map posts array to json array such that each individual posts maps to PostMap stuct

import (
	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
)

func MapPosts(posts []database.Post) []PostMap {
	var postMaps []PostMap
	for _, post := range posts {
		postMaps = append(postMaps, MapPost(post))
	}
	return postMaps
}
