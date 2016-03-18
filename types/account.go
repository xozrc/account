package types

type AccountType int

const (
	Visitor AccountType = iota
	Facebook
)

func (at AccountType) Equal(ati int) bool {
	if int(at) == ati {
		return true
	}
	return false
}

type Account struct {
	Id          int64
	SecondId    string
	AccountType int
	UniqueCode  string
	Name        string
}

func NewAccount(id int64, secondId string, accountType int, uniqueCode string, name string) (acc *Account) {
	acc = &Account{
		Id:          id,
		SecondId:    secondId,
		AccountType: accountType,
		UniqueCode:  uniqueCode,
		Name:        name,
	}
	return
}
