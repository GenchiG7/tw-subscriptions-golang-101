package dummy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProdMgmt(t *testing.T) {
	dummyProd1 := Product{}
	dummyProd2 := Product{}

	pm := ProdMgmt{[]Product{dummyProd1, dummyProd2}}

	actual := pm.Len()

	assert.Equal(t, 2, actual)

}
