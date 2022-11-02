package spy

type Invoicer interface {
	Issue()
}

type ProdMgmt struct {
	i         Invoicer
	stockNums int
}

func (pm *ProdMgmt) Sell(amount int) {
	pm.stockNums -= amount

	for i := 0; i < amount; i++ {
		pm.i.Issue()
	}
}
