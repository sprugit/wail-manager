package internal

type PasswordNodePrimitive struct {
	ID       int
	ACCOUNT  string
	DOMAIND  string
	PASSWORD string
	PARENT   int
}

type PasswordNode struct {
	ID       int
	ACCOUNT  string
	DOMAIND  string
	PASSWORD string
	CHILD    *PasswordNode
}
