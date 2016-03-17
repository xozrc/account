package types

type AccountType int

const (
	Visitor AccountType = iota
	Facebook
)

type Account struct {
	Id         int64
	SecondId   string
	Channel    int
	UniqueCode string
	Name       string
}

func NewAccount(id int64, secondId string, channel int, uniqueCode string, name string) (acc *Account) {
	acc = &Account{
		Id:         id,
		SecondId:   secondId,
		Channel:    channel,
		UniqueCode: uniqueCode,
		Name:       name,
	}
	return
}
