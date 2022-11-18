package dto

type CreateUserInput struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image"`
}
