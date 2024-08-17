package entities

type AuthRole uint

const (
	Staff AuthRole = iota + 1
	Manager
	Owner // Kritik işlemleri sadece bu role sahip kullanıcı yapabilir (hastaneyi silme gibi) kaydı oluşturan hesaptır
)

// TODO Veri tipleri düzenlenecek

type User struct {
	Base
	Contact
	ID       string   `json:"id" gorm:"type:varchar(255);unique"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Password string   `json:"password"`
	Role     AuthRole `json:"role"`
}
