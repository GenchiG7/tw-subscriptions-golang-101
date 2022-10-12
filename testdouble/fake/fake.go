package stub

import "math/rand"

type User struct {
	ID   int
	Name string
}
type UserRepo interface {
	Persis(u User)
	Find(ID int) User
}

type ProdMgmt struct {
	repo UserRepo
}

func (pm *ProdMgmt) CreateUser(name string) int {
	id := rand.Intn(100)
	u := User{id, name}
	pm.repo.Persis(u)

	return id
}

func (pm *ProdMgmt) FindByID(id int) User {
	return pm.repo.Find(id)
}
