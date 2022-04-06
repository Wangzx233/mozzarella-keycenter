package config

import "time"

const (
	RsaKeyRefreshTime = time.Hour * 24 * 7
	DefaultDomain     = "common"

	DefaultTokenExpiresTime = time.Minute * 30
	DefaultTokenNotBefore   = 0
	DefaultTokenIssuer      = "key-center"
	DefaultTokenAudience    = "user"

	RefreshRefreshTokenExpiresTime = time.Hour * 24 * 30
)
