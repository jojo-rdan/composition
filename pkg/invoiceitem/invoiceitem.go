package invoiceitem

// Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new item, kinda constructor
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}
