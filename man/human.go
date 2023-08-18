package human

import "fmt"

type Man struct {
	name     string
	lastName string
	age      int
	gender   string
	crimes   int
}

func New(name, lastName string, age int, gender string, crimes int) *Man {
	m := Man{
		name:     name,
		lastName: lastName,
		age:      age,
		gender:   gender,
		crimes:   crimes,
	}

	return &m
}

func (m *Man) Name() string {
	return m.name
}

func (m *Man) LastName() string {
	return m.lastName
}

func (m *Man) Age() int {
	return m.age
}

func (m *Man) Gender() string {
	return m.gender
}

func (m *Man) Crimes() int {
	return m.crimes
}

func (m *Man) AllInfo() string {
	return fmt.Sprintf("Подозреваемый: %s %s\nВозраст: %d\nПол: %s\nСовершенных преступлений: %d", m.lastName, m.name, m.age, m.gender, m.crimes)
}
