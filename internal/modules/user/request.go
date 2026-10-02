package user

type CreateUserRequest struct {
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=8"`
	FullName    string  `json:"full_name" validate:"required,max=100"`
	PhoneNumber *string `json:"phone_number"`
	Role        string  `json:"role" validate:"required,oneof=admin staff buyer"`
}

type UpdateUserRequest struct {
	Email		string 	`json:"email"`
	FullName	string	`json:"full_name"`
	PhoneNUmber	*string `json:"phone_number"`
	Role		string	`json:"role"`

}