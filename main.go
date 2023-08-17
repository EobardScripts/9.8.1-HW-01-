package main

import (
	"981HW-01/generator"
	"981HW-01/man"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Генерация базы подозреваемых
func createBase() map[string]human.Man {
	count := generator.RandInt(5, 11)
	s := make(map[string]human.Man)
	for i := 0; i <= count; i++ {
		//0 - имя, 1 - фамилия, 2 - пол
		nameData := strings.Split(generator.GetRandomName(), ";")

		m := human.New(nameData[0], nameData[1], generator.RandInt(16, 51), nameData[2], generator.RandInt(0, 21))

		s[fmt.Sprintf("%s %s", m.LastName, m.Name)] = *m
	}
	return s
}

// Выявление подозреваемого с максимальным кол-вом преступлений
func maxCrimes(c []human.Man) int {
	max := 0
	i := 0
	for index, culprit := range c {
		if culprit.GetCrimes() > max {
			max = culprit.GetCrimes()
			i = index
		}
	}

	return i
}

// Поиск подозреваемых в базе
func setCulprit(p map[string]human.Man, s []string) string {
	var culprits []human.Man

	for _, suspect := range s {
		culprit, ok := p[suspect]
		if ok {
			culprits = append(culprits, culprit)
		}
	}

	return culprits[maxCrimes(culprits)].AllInfo()
}

func main() {
	//Для случайных чисел
	rand.Seed(time.Now().UnixNano())
	people := createBase()
	var suspects []string

	//Создание списка для сравнения с подозреваемыми
	for i := 0; i <= 2000; i++ {
		//0 - имя, 1 - фамилия, 2 - пол
		names := strings.Split(generator.GetRandomName(), ";")
		suspects = append(suspects, fmt.Sprintf("%s %s", names[1], names[0]))
	}

	fmt.Println(setCulprit(people, suspects))
}
