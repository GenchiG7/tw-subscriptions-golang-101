package stub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fake struct {
	m map[int]User
}

func (f *Fake) Persis(u User) {
	f.m[u.ID] = u
}

func (f *Fake) Find(ID int) User {
	return f.m[ID]
}

func TestProdMgmt(t *testing.T) {

	fake := &Fake{make(map[int]User)}
	pm := ProdMgmt{fake}

	name := "g7"
	id := pm.CreateUser(name)
	actual := pm.FindByID(id)

	assert.Equal(t, name, actual.Name)

}
