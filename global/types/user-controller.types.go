package types

import "github.com/arifseft/go-actions/utils"

type GetNameResponse struct {
	utils.Response
	Result string `json:"result"`
}

type GetBiodataResult struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetBiodataResponse struct {
	utils.Response
	Result GetBiodataResult `json:"result"`
}

type CreateUserResult struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type CreateUser struct {
	utils.Response
	Result CreateUserResult `json:"result"`
}
