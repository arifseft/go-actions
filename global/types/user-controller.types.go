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
