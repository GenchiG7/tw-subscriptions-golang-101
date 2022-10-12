package spy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Spy struct {
	record int
}

func (s *Spy) Issue() {
	s.record++
}

func TestProdMgmt(t *testing.T) {

	spy := &Spy{}
	pm := ProdMgmt{spy, 11}

	pm.Sell(2)

	assert.Equal(t, 2, spy.record)

}
