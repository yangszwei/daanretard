/* Package fbgraph provides an interface to work with Facebook Graph API */
package fbgraph

import (
	"daanretard/internal/infra/errors"
	"encoding/json"
	"fmt"
	fb "github.com/huandu/facebook/v2"
	"path"
	"time"
)

// IUsecase usecase interface of App, other packages should depend on this
// interface instead of App
type IUsecase interface {
	ExchangeAccessToken(accessToken string) (string, time.Time, error)
	DebugAccessToken(accessToken string) (AccessToken, error)
	GetUserProfile(accessToken string) (string, string, error)
	PublishPost(message string, attachedMedia []string, accessToken string) (string, error)
	UploadPhoto(filepath string, publish bool, accessToken string) (string, error)
	Delete(id, accessToken string) error
}

// New create a App
func New(pageID, appID, appSecret string) *App {
	a := new(App)
	a.app = fb.New(appID, appSecret)
	a.pageID = pageID
	return a
}

// App wrapper of fb.App so app id, secret and page id does not need to be
// passed every time
type App struct {
	app    *fb.App
	pageID string
}

// ExchangeToken exchange long-lived access token with user input token, so
// the user don't need to re-authorize every time
func (a App) ExchangeAccessToken(accessToken string) (string, time.Time, error) {
	token, expires, err := a.app.ExchangeToken(accessToken)
	return token, time.Unix(int64(expires), 0), handleError(err)
}

// DebugAccessToken get the information of an access token so the app know how
// to handle it
func (a App) DebugAccessToken(accessToken string) (AccessToken, error) {
	res, err := fb.Get("/debug_token", fb.Params{
		"input_token":  accessToken,
		"access_token": accessToken,
	})
	if err != nil {
		return AccessToken{}, handleError(err)
	}
	var scopes []string
	for _, name := range res.Get("data.scopes").([]interface{}) {
		scopes = append(scopes, name.(string))
	}
	var tokenType rune
	switch res.Get("data.type").(string) {
	case "USER":
		tokenType = User
	case "PAGE":
		tokenType = Page
	}
	expiresAt, err := res.Get("data.expires_at").(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	return AccessToken{
		Value:     accessToken,
		Type:      tokenType,
		UserID:    res.Get("data.user_id").(string),
		Scopes:    scopes,
		IsValid:   res.Get("data.is_valid").(bool),
		ExpiresAt: time.Unix(expiresAt, 0),
	}, nil
}

// GetUserProfile fetch user profile (currently, id & name) from facebook, used
// in creating/refreshing local user profile
func (a App) GetUserProfile(accessToken string) (string, string, error) {
	res, err := fb.Get("/me", fb.Params{
		"fields":       "name",
		"access_token": accessToken,
	})
	if err != nil {
		return "", "", handleError(err)
	}
	return res.Get("id").(string), res.Get("name").(string), nil
}

// PublishPost publish post to facebook and return the id, attachedMedia
// should be a list of facebook id (currently, of photos)
func (a App) PublishPost(message string, attachedMedia []string, accessToken string) (string, error) {
	params := fb.Params{"message": message, "access_token": accessToken}
	format := "{\"media_fbid\":\"%s\"}"
	for i, id := range attachedMedia {
		params[fmt.Sprintf("attached_media[%d]", i)] = fmt.Sprintf(format, id)
	}
	res, err := fb.Post(path.Join("/", a.pageID, "feed"), params)
	if err != nil {
		return "", handleError(err)
	}
	return res.Get("id").(string), nil
}

// UploadPhoto upload one photo to facebook, when publish is set to true, the
// photo is published as a post ; when publish is set to false, the photo can
// be used in other posts (refer to attachedMedia in App.PublishPost)
func (a App) UploadPhoto(filepath string, publish bool, accessToken string) (string, error) {
	res, err := fb.Post(path.Join("/", a.pageID, "photos"), fb.Params{
		"source":       fb.File(filepath),
		"published":    publish,
		"access_token": accessToken,
	})
	if err != nil {
		return "", handleError(err)
	}
	return res.Get("id").(string), nil
}

// Delete delete a facebook post or photo by its id
func (a App) Delete(id, accessToken string) error {
	res, err := fb.Delete(path.Join("/", id), fb.Params{
		"access_token": accessToken,
	})
	if err != nil {
		return handleError(err)
	}
	if !res.Get("success").(bool) {
		return errors.ErrUnknownError
	}
	return nil
}
