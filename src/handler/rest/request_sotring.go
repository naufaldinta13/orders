package rest

import (
	"sort"

	"github.com/env-io/validate"
)

type sortingRequest struct {
	Numbers []int  `json:"numbers"`
	Sort    string `json:"sort"`
}

func (r *sortingRequest) Validate() *validate.Response {
	v := validate.NewResponse()

	return v
}

func (r *sortingRequest) Messages() map[string]string {
	return map[string]string{}
}

func (r *sortingRequest) Execute() (mx []int, e error) {

	sort.Ints(r.Numbers)

	if r.Sort == "desc" {
		for i := len(r.Numbers); i > 0; i-- {
			mx = append(mx, r.Numbers[i-1])
		}
	}

	if r.Sort == "asc" {
		for i := 0; i < len(r.Numbers); i++ {
			mx = append(mx, r.Numbers[i])
		}
	}

	return
}
