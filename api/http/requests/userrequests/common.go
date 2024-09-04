package userrequests

type RegisterRequest struct {
	Login    string `json:"login" validate:"required,min=5,max=14,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=20,lowercase,uppercase,digit,specialchar"`
}
