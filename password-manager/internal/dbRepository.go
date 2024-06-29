package internal

var (
	dbFileName = "whale.db"
)

//https://gosamples.dev/sqlite-intro/

type Repository interface {
	Migrate() error
	Create(password PasswordNodePrimitive) (*PasswordNodePrimitive, error)
	All() ([]PasswordNodePrimitive, error)
	GetPrimitiveByAccount(account string) (*PasswordNodePrimitive, error)
	GetPrimitiveByDomain(domain string) (*PasswordNodePrimitive, error)
	GetPrimitiveByAccountDomain(account string, domain string) (*PasswordNodePrimitive, error)
	GetByAccount(account string) (*PasswordNode, error)
	GetByDomain(domain string) (*PasswordNode, error)
	GetByAccountDomain(account string, domain string) (*PasswordNode, error)
	Update(id int, updated PasswordNodePrimitive)
	Delete(id int)
}
