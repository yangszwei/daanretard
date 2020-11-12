package validator

import (
	"daanretard/internal/object"
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	postValidator = validator.New()
	postReviewValidator = validator.New()
)

// Post validate object.Post
func Post(post object.Post) error {
	if postValidator.Struct(post) != nil {
		log.Println(postValidator.Struct(post))
		return errors.New("invalid credentials")
	}
	return nil
}

// PostReview validate object.PostReview
func PostReview(review object.PostReview) error {
	if postReviewValidator.Struct(review) != nil {
		return errors.New("invalid credentials")
	}
	return nil
}