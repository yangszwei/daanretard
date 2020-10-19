package object

// UserProps User props. used to generate user.User
type UserProps struct {
	Name     string `validate:"required,max=50"`
	Email    string `validate:"required,max=254"`
	Password string `validate:"required,min=8,max=30"`
}

// UserProfileProps used to generate user.Profile
type UserProfileProps struct {
	FirstName string `validate:"required,max=50"`
	LastName  string `validate:"required,max=50"`
}