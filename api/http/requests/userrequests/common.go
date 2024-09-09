package userrequests

type Login struct {
	Mnemonic []string `json:"mnemonic" validate:"required,min=12,max=12"`
}

type Register struct {
	Mnemonic []string `json:"mnemonic" validate:"required,min=12,max=12"`
}
