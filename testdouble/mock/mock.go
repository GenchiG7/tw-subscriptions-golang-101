package mock

type Prod struct {
	Type int
}
type PriceHelper interface {
	GetPrice(prodType int) int
}

type ProdMgmt struct {
	p PriceHelper
}

func (pm *ProdMgmt) TotalPrice(prods []Prod) int {

	total := 0
	for _, prod := range prods {
		total += pm.p.GetPrice(prod.Type)
	}

	return total
}
