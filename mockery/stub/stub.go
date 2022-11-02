package stub

type Warehouse interface {
	StockNums() int
}

type ProdMgmt struct {
	w         Warehouse
	threshold int
}

func (pm *ProdMgmt) IsAlmostFull() bool {
	return pm.w.StockNums() > pm.threshold
}
