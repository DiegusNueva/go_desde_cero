package invoice

import (
	"github.com/DiegusNueva/composition/pkg/customer"
	"github.com/DiegusNueva/composition/pkg/invoiceitem"
)

// Invoice is a struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer // relación de uno a uno
	items   invoiceitem.Items // relación de uno a muchos
}

// New returns a new Invoice
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
