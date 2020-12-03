/*
	Package mock_fbgraph provides a mock implementation of fbgraph.IUsecase,
	so services that depend on the interface can be tested without actually
	sending the request to facebook graph API.
*/
package mock_fbgraph

import (
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/fbgraph"
	"strings"
	"time"
)

// New create a App
func New(pageID, appID, appSecret string) *App {
	a := new(App)
	a.app = newAPI(appID, appSecret)
	a.pageID = pageID
	return a
}

// App implement fbgraph.IUsecase, contains mockAPI instance and pageID
type App struct {
	app    *api
	pageID string
}

// NewUser add a user instance to mock service and return an access token
func (a App) NewUser(name string) string {
	u := a.app.newUser(name)
	t := a.app.newToken()
	a.app.tokens[t] = token{
		type_:     fbgraph.User,
		pageID:    "",
		userID:    u,
		scopes:    []string{"public_profile"},
		expiresAt: time.Now().Add(2 * time.Hour),
	}
	return t
}

// NewPage add a page instance to mock service and return an access token
func (a App) NewPage(id, userID string) string {
	a.app.newPage(id)
	t := a.app.newToken()
	a.app.tokens[t] = token{
		type_:     fbgraph.Page,
		pageID:    id,
		userID:    userID,
		scopes:    []string{"public_profile", "page_manage_posts"},
		expiresAt: time.Now().Add(2 * time.Hour),
	}
	return t
}

// ExchangeToken return a new token with extended expiration date
func (a App) ExchangeAccessToken(accessToken string) (string, time.Time, error) {
	t, e := a.app.tokens[accessToken]
	if !e {
		return "", time.Time{}, errors.New("invalid access token")
	}
	v := a.app.newToken()
	var exp time.Time
	t2 := t
	if t.type_ == fbgraph.Page {
		exp = time.Now().AddDate(10, 0, 0)
	} else {
		exp = time.Now().AddDate(0, 0, 60)
	}
	t2.expiresAt = exp
	a.app.tokens[v] = t2
	return v, exp, nil
}

// DebugAccessToken get the information of an access token so the app know how
// to handle it
func (a App) DebugAccessToken(accessToken string) (fbgraph.AccessToken, error) {
	t, e := a.app.tokens[accessToken]
	if !e {
		return fbgraph.AccessToken{}, errors.New("invalid access token")
	}
	return fbgraph.AccessToken{
		Value:     accessToken,
		Type:      t.type_,
		UserID:    t.userID,
		Scopes:    t.scopes,
		IsValid:   t.expiresAt.After(time.Now()),
		ExpiresAt: t.expiresAt,
	}, nil
}

// GetUserProfile fetch user profile (currently, id & name) from mock service
func (a App) GetUserProfile(accessToken string) (string, string, error) {
	t, e := a.app.tokens[accessToken]
	if !e {
		return "", "", errors.New("invalid access token")
	}
	u := a.app.users[t.userID]
	return t.userID, u.name, nil
}

// PublishPost add a post instance to mock service and return the id
func (a App) PublishPost(message string, attachedMedia []string, accessToken string) (string, error) {
	t, e := a.app.tokens[accessToken]
	if !e || t.type_ != fbgraph.Page {
		return "", errors.New("invalid access token")
	}
	id := a.app.newID("post")
	a.app.pages[t.pageID].feed[id] = &post{
		message:       message,
		attachedMedia: strings.Join(attachedMedia, ","),
	}
	return id, nil
}

// UploadPhoto add a photo instance to mock service and return the id
func (a App) UploadPhoto(filepath string, publish bool, accessToken string) (string, error) {
	t, e := a.app.tokens[accessToken]
	if !e || t.type_ != fbgraph.Page || t.pageID != a.pageID {
		return "", errors.New("invalid access token")
	}
	id := a.app.newID("photo")
	a.app.pages[t.pageID].photos[id] = &photo{
		filepath:  filepath,
		published: publish,
	}
	return id, nil
}

// Delete delete a mock post or photo by its id
func (a App) Delete(id, accessToken string) error {
	t, e := a.app.tokens[accessToken]
	if !e || t.type_ != fbgraph.Page {
		return errors.New("invalid access token")
	}
	i, e := a.app.ids[id]
	if !e {
		return errors.New("unsupported delete request")
	}
	if i == "post" {
		delete(a.app.pages[t.pageID].feed, id)
	}
	if i == "photo" {
		delete(a.app.pages[t.pageID].photos, id)
	}
	return nil
}
