package human

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func New(name, lastName string, age int, gender string, crimes int) *Man {
	m := Man{
		Name:     name,
		LastName: lastName,
		Age:      age,
		Gender:   gender,
		Crimes:   crimes,
	}

	return &m
}

func (m *Man) GetName() string {
	return m.Name
}

func (m *Man) GetLastName() string {
	return m.LastName
}

func (m *Man) GetAge() int {
	return m.Age
}

func (m *Man) GetGender() string {
	return m.Gender
}

func (m *Man) GetCrimes() int {
	return m.Crimes
}

func (m *Man) AllInfo() string {
	return fmt.Sprintf("Подозреваемый: %s %s\nВозраст: %d\nПол: %s\nСовершенных преступлений: %d", m.LastName, m.Name, m.Age, m.Gender, m.Crimes)
}
