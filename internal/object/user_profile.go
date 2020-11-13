package object

// UserProfileProps data object of user.Profile
type UserProfileProps struct {
	DisplayName string `validate:"omitempty,max=50"`
	FirstName   string `validate:"omitempty,max=50"`
	LastName    string `validate:"omitempty,max=50"`
}
