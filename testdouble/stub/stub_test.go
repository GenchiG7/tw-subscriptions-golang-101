package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stub struct {
}

func (s *Stub) StockNums() int {
	return 10
}

func TestProdMgmt(t *testing.T) {

	stub := &Stub{}
	pm := ProdMgmt{stub, 11}

	actual := pm.IsAlmostFull()

	assert.Equal(t, false, actual)

}
