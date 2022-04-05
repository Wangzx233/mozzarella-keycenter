package token

type StandardPayload struct {
	Audience  string `json:"aud,omitempty"` //受众
	ExpiresAt int64  `json:"exp,omitempty"` //过期时间
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"` //签发时间
	Issuer    string `json:"iss,omitempty"` //签发人
	NotBefore int64  `json:"nbf,omitempty"` //生效时间
	Subject   string `json:"sub,omitempty"` //作用域domain
}

type Payload struct {
	StandardPayload
	Uid string `json:"uid"`
}
