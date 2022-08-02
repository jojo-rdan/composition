package invoice

import (
	"github.com/jojo-rdan/composition/pkg/customer"
	"github.com/jojo-rdan/composition/pkg/invoiceitem"
)

// Los corchetes implican que la relación es de uno a muchos
// esto basicamente está abriendo un slice de tipo Item
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}

// New returns a new invoice
func New(country string, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items}
}
