package mock_fbgraph

import (
	"math/rand"
	"time"
)

// newAPI create a new api
func newAPI(appID, appSecret string) *api {
	a := new(api)
	a.id = appID
	// secret is not used in this implementation, this is only to simulate
	// the fb.New() method in package huandu/facebook which package fbgraph
	// depends on
	a.secret = appSecret
	a.ids = make(map[string]string)
	a.tokens = make(map[string]token)
	a.users = make(map[string]*user)
	a.pages = make(map[string]*page)
	return a
}

// api mock parts of facebook graph API that App uses
type api struct {
	id     string
	secret string
	ids    map[string]string
	tokens map[string]token
	users  map[string]*user
	pages  map[string]*page
}

// token mock instance of a facebook token
type token struct {
	type_     rune
	pageID    string
	userID    string
	scopes    []string
	expiresAt time.Time
}

// user mock instance of a facebook user
type user struct {
	name string
}

// page mock instance of a facebook page
type page struct {
	feed   map[string]*post
	photos map[string]*photo
}

// post mock instance of a facebook post
type post struct {
	message       string
	attachedMedia string
}

// photo mock instance of a facebook photo
type photo struct {
	filepath  string
	published bool
}

// randomString generate a alphanumeric string
func randomString(size int) string {
	bytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, size)
	for i := range b {
		b[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(b)
}

// newID generate a unique id (of user, post and photo)
func (a *api) newID(s string) string {
	for {
		id := randomString(20)
		if _, e := a.ids[id]; !e {
			a.ids[id] = s
			return id
		}
	}
}

// newToken generate a unique token
func (a *api) newToken() string {
	for {
		token := randomString(100)
		if _, exist := a.tokens[token]; !exist {
			return token
		}
	}
}

// newUser create a new user
func (a *api) newUser(name string) string {
	id := a.newID("user")
	u := new(user)
	u.name = name
	a.users[id] = u
	return id
}

// newPage create a new page
func (a *api) newPage(id string) *page {
	p := new(page)
	p.feed = make(map[string]*post)
	p.photos = make(map[string]*photo)
	a.pages[id] = p
	return p
}
