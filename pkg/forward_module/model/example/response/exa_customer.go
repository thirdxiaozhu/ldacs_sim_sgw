package response

import "ldacs_sim_sgw/pkg/forward_module/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
