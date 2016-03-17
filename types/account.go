package types

type AccountType int

const (
	Visitor AccountType = iota
	Facebook
)

type Account struct {
	Id         int64
	SecondId   string
	Channel    AccountType
	UniqueCode string
	Name       string
	Password   string
}
