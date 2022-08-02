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
	items   invoiceitem.Items
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}

// New returns a new invoice
func New(country string, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items}
}
