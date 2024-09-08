package userrequests

type Login struct {
	Mnemonic []string `json:"mnemonic" validate:"required,min=24,max=24"`
}

type Register struct {
	Mnemonic []string `json:"mnemonic" validate:"required,min=24,max=24"`
}
