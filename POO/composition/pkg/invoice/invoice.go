package invoice

import "github.com/DiegusNueva/composition/pkg/customer"

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
