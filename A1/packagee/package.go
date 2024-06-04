package  packagee

type Person struct {
	Name     string
	Surname string
	Age int
}

func (p *Person) PrintName() string {
	return p.Name
}

func (p *Person) PrintSurname() string {
	return p.Surname
}