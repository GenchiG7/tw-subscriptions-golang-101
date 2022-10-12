package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Mock struct {
	recordMap map[int]int
	returnMap map[int]int
}

func (m *Mock) Expect(prodType, price, numsOfCall int) {
	m.returnMap[prodType] = price
	m.recordMap[prodType] = numsOfCall
}

func (m *Mock) GetPrice(prodType int) int {
	m.recordMap[prodType]--
	return m.returnMap[prodType]
}

func TestProdMgmt(t *testing.T) {

	m := &Mock{make(map[int]int), make(map[int]int)}
	m.Expect(1, 2, 1)
	m.Expect(2, 4, 1)

	pm := ProdMgmt{m}

	actual := pm.TotalPrice([]Prod{{1}, {2}})

	assert.Equal(t, 6, actual)
	for _, v := range m.recordMap {
		assert.Equal(t, 0, v)
	}

}
