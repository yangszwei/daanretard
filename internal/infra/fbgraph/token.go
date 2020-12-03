package fbgraph

import "time"

// AccessToken facebook access token object, returned by App.DebugToken
type AccessToken struct {
	Value     string
	Type      rune
	UserID    string
	Scopes    []string
	IsValid   bool
	ExpiresAt time.Time
}

// Access token types
const (
	User = 'u'
	Page = 'p'
)
