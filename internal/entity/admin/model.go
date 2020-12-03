package admin

import "time"

// Admin admin object used inside the package, for the exposed object,
// refer to object.Admin, admins are users whose facebook account have
// required permissions on the configured facebook page, this entity
// store the facebook page access token, which can be used in package
// infra/fbgraph
type Admin struct {
	UserID      string `gorm:"type:varchar(128);primaryKey"`
	AccessToken string // facebook PAGE access token
	ExpiresAt   time.Time
	CreatedAt   time.Time
}
