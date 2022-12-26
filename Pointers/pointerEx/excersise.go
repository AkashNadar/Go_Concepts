package main

import "fmt"

type Product struct {
	item        string
	securityTag bool
}

func changeSecurity(product *Product) {
	if product.securityTag {
		product.securityTag = false
	} else {
		product.securityTag = true
	}
}

func activate(security *bool) {
	*security = true
}

func deactivate(security *bool) {
	*security = false
	// fmt.Println(*security)
}

func checkout(products []Product) {
	for i := 0; i < len(products); i++ {
		deactivate(&products[i].securityTag)
	}

}

func main() {
	products := []Product{
		{"hall", true},
		{"kitchen", true},
		{"bathroom", true},
		{"storage", true},
	}
	fmt.Println(products)

	changeSecurity(&products[1])
	fmt.Println(products)

	checkout(products)
	fmt.Println(products)
}
