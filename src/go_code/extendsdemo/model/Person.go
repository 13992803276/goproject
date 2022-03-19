package model

type Person struct {
	Id      int16
	address string
}

func (p *Person) GetAddress() string {
	return p.address
}

func (p *Person) SetAddress(address string) {
	p.address = address
}
