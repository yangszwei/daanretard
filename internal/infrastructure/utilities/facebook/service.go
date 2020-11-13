package facebook

import (
	"errors"
	"fmt"
	fb "github.com/huandu/facebook/v2"
	"path"
	"strings"
)

// NewService create a Service
func NewService(appID, appSecret, pageID string) *Service {
	s := new(Service)
	s.a = fb.New(appID, appSecret)
	s.p = pageID
	fb.Debug = fb.DEBUG_ALL
	return s
}

// Service implement IUsecase
type Service struct {
	a *fb.App
	p string
}

// PublishPost publish a post
func (s *Service) PublishPost(message string, attachments []string, accessToken string) (string, error) {
	params := fb.Params{
		"message":      message,
		"access_token": accessToken,
	}
	for i, attachment := range attachments {
		params[fmt.Sprintf("attached_media[%d]", i)] = fmt.Sprintf("{\"media_fbid\":\"%s\"}", attachment)
	}
	res, err := fb.Post(path.Join("/", s.p, "feed"), params)
	if err != nil {
		if strings.Contains(err.Error(), "Invalid OAuth access token.") {
			return "", errors.New("invalid access token")
		}
		return "", err
	}
	return res.Get("id").(string), nil
}

// PublishPhoto publish a photo
func (s *Service) PublishPhoto(filepath, accessToken string) (string, error) {
	res, err := fb.Post(path.Join("/", s.p, "photos"), fb.Params{
		"source":       fb.File(filepath),
		"access_token": accessToken,
		"published":    false,
	})
	if err != nil {
		return "", err
	}
	return res.Get("id").(string), nil
}

// DeletePost delete a post
func (s *Service) DeletePost(id, accessToken string) error {
	res, err := fb.Delete("/"+id, fb.Params{
		"access_token": accessToken,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Invalid OAuth access token.") {
			return errors.New("invalid access token")
		} else if strings.Contains(err.Error(), "Unsupported delete request.") {
			return errors.New("not exist")
		}
		return err
	}
	if res.Get("success").(bool) != true {
		return errors.New("deletion failed")
	}
	return nil
}

// ExchangeToken get long-lived access token
func (s *Service) ExchangeToken(accessToken string) (string, error) {
	token, _, err := s.a.ExchangeToken(accessToken)
	if err != nil {
		return "", err
	}
	return token, nil
}
