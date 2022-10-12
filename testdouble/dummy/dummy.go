package dummy

type Product struct {
	ID int
}

type ProdMgmt struct {
	prods []Product
}

func (pm *ProdMgmt) Len() int {
	return len(pm.prods)
}
