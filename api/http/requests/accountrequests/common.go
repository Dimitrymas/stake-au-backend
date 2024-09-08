package accountrequests

import "go.mongodb.org/mongo-driver/bson/primitive"

type Create struct {
	Name       string `json:"name" validate:"required,max=50"`                                                 // Обязательно
	Token      string `json:"token" validate:"required"`                                                       // Обязательно
	ProxyType  string `json:"proxy_type" validate:"omitempty,oneof=http socks5"`                               // Должно быть либо "http", либо "socks5", если указано
	ProxyLogin string `json:"proxy_login" validate:"omitempty,required_with=ProxyType"`                        // Обязательно, если указан ProxyType
	ProxyPass  string `json:"proxy_pass" validate:"omitempty,required_with=ProxyLogin"`                        // Обязательно, если указан ProxyLogin
	ProxyIP    string `json:"proxy_ip" validate:"omitempty,required_with=ProxyType,ip"`                        // Должен быть валидным IP-адресом, если указан ProxyType
	ProxyPort  string `json:"proxy_port" validate:"omitempty,required_with=ProxyType,numeric,min=1,max=65535"` // Валидный порт, если указан ProxyType
}

type Edit struct {
	ID primitive.ObjectID `json:"id" validate:"required"` // Обязательно
	Create
}

type CreateMany struct {
	Accounts []*Create `json:"accounts" validate:"required,dive"`
}
