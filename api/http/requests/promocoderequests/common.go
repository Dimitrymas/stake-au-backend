package promocoderequests

type Create struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Value       float64 `json:"value" validate:"omitempty,min=0,max=1000"`
	Description string  `json:"description" validate:"omitempty,min=2,max=1000"`
}
