package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginReq defines the request for user login
type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"Authentication" summary:"User login" sm:"User login"`
	Username string `p:"username" v:"required#Username cannot be empty" dc:"Username"`
	Password string `p:"password" v:"required#Password cannot be empty" dc:"Password"`
}

// LoginRes defines the response for user login
type LoginRes struct {
	api_v1.StandardRes
	Data struct {
		Token        string `json:"token" dc:"JWT token"`
		RefreshToken string `json:"refreshToken" dc:"Refresh token"`
		ExpiresAt    int64  `json:"expiresAt" dc:"Token expiration time (Unix timestamp)"`
		AccountInfo  struct {
			Id       int64  `json:"id" dc:"Account ID"`
			Username string `json:"username" dc:"Username"`
			Email    string `json:"email" dc:"Email address"`
			Status   int    `json:"status" dc:"Account status"`
			Lang     string `json:"lang" dc:"Preferred language"`
		} `json:"accountInfo" dc:"Basic account information"`
	} `json:"data"`
}

// LogoutReq defines the request for user logout
type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"Authentication" summary:"User logout" sm:"User logout"`
}

// LogoutRes defines the response for user logout
type LogoutRes struct {
	api_v1.StandardRes
}

// RefreshTokenReq defines the request for token refresh
type RefreshTokenReq struct {
	g.Meta       `path:"/refresh-token" method:"post" tags:"Authentication" summary:"Refresh access token" sm:"Refresh access token"`
	RefreshToken string `p:"refreshToken" v:"required#Refresh token cannot be empty" dc:"Refresh token"`
}

// RefreshTokenRes defines the response for token refresh
type RefreshTokenRes struct {
	api_v1.StandardRes
	Data struct {
		Token        string `json:"token" dc:"New JWT token"`
		RefreshToken string `json:"refreshToken" dc:"New refresh token"`
		ExpiresAt    int64  `json:"expiresAt" dc:"Token expiration time (Unix timestamp)"`
	} `json:"data"`
}

// RegisterReq defines the request for user registration
type RegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"Authentication" summary:"User registration" sm:"User registration"`
	Username string `p:"username" v:"required#Username cannot be empty" dc:"Username"`
	Password string `p:"password" v:"required#Password cannot be empty" dc:"Password"`
	Email    string `p:"email" v:"required|email#Email cannot be empty|Invalid email format" dc:"Email address"`
	Lang     string `p:"lang" d:"en" dc:"Preferred language"`
}

// RegisterRes defines the response for user registration
type RegisterRes struct {
	api_v1.StandardRes
	Data struct {
		AccountId int64 `json:"accountId" dc:"New account ID"`
	} `json:"data"`
}

// CurrentUserReq defines the request for getting current user info
type CurrentUserReq struct {
	g.Meta `path:"/current-user" method:"get" tags:"Authentication" summary:"Get current user info" sm:"Get current user info"`
}

// CurrentUserRes defines the response for getting current user info
type CurrentUserRes struct {
	api_v1.StandardRes
	Data struct {
		Account struct {
			Id        int64  `json:"id" dc:"Account ID"`
			Username  string `json:"username" dc:"Username"`
			Email     string `json:"email" dc:"Email address"`
			Status    int    `json:"status" dc:"Account status"`
			Lang      string `json:"lang" dc:"Preferred language"`
			CreatedAt string `json:"createdAt" dc:"Creation time"`
		} `json:"account" dc:"Account information"`
		Roles       []string `json:"roles" dc:"User roles"`
		Permissions []string `json:"permissions" dc:"User permissions"`
	} `json:"data"`
}
