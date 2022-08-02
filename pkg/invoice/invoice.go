package invoice

import (
	"github.com/jojo-rdan/composition/tree/master/pkg/customer"
	"github.com/jojo-rdan/composition/tree/master/pkg/invoiceitem"
)

//Los corchetes implican que la relación es de uno a muchos
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}
