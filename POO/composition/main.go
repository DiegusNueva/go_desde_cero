package main

import (
	"fmt"

	"github.com/DiegusNueva/composition/pkg/customer"
	"github.com/DiegusNueva/composition/pkg/invoice"
	"github.com/DiegusNueva/composition/pkg/invoiceitem"
)

func main() {

	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Alejandro", "cl 123 # 23-4", "658987421"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.23),
			invoiceitem.New(3, "Curso de testing con Go", 90.00),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)

}
