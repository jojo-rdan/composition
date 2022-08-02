package main

import (
	"fmt"

	"github.com/jojo-rdan/composition/pkg/customer"
	"github.com/jojo-rdan/composition/pkg/invoice"
	"github.com/jojo-rdan/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Bogotá",
		customer.New("Jordan", "Calle evergreen", "12345"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso de go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.23),
			invoiceitem.New(3, "Curso de testing con go", 90.00),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
