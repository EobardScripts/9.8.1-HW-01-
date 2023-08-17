package generator

import (
	"math/rand"
)

const (
	Male   string = "Мужчина"
	Female string = "Женщина"
)

// Случайное число
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// Случайный пол
func randGender() string {
	r := RandInt(0, 2)
	if r == 0 {
		return Male
	} else {
		return Female
	}
}

// Случайное имя, фамилия и пол
func GetRandomName() string {
	var result string
	maleNames := []string{"Серафим", "Пётр", "Марк", "Даниил", "Сергей", "Руслан", "Лев", "Тимофей", "Кирилл", "Ярослав", "Матвей", "Богдан", "Макар", "Роман", "Валерий", "Арсений", "Денис", "Тимур", "Никита", "Константин", "Владимир", "Владислав", "Альберт", "Артемий", "Иван", "Дмитрий", "Станислав", "Павел", "Артём", "Ян", "Илья", "Али", "Глеб", "Георгий", "Герман", "Игорь", "Александр", "Андрей", "Савва", "Фёдор", "Максим", "Даниэль", "Егор", "Евгений", "Василий", "Тихон", "Михаил", "Алексей", "Филипп", "Данила"}
	maleLastNames := []string{"Антонов", "Сорокин", "Захаров", "Киселев", "Голубев", "Васильев", "Березин", "Тарасов", "Матвеев", "Васильев", "Леонов", "Макаров", "Демидов", "Климов", "Кузьмин", "Мартынов", "Куликов", "Долгов", "Астафьев", "Кудрявцев", "Галкин", "Быков", "Акимов", "Олейников", "Кондрашов", "Баранов", "Шишкин", "Рогов", "Щербаков", "Захаров", "Михайлов", "Карасев", "Зорин", "Сидоров", "Ковалев", "Мельников", "Головин", "Горбачев", "Максимов", "Тихомиров", "Мельников", "Прохоров", "Никитин", "Орлов", "Волков", "Михеев", "Иванов", "Климов", "Федоров", "Макаров"}
	femaleNames := []string{"Анастасия", "Яна", "Ева", "Валерия", "Ирина", "Василиса", "София", "Анна", "Надежда", "Алиса", "Олеся", "Дарья", "Маргарита", "Елизавета", "Полина", "Екатерина", "Амина", "Варвара", "Ярослава", "Мария", "Виктория", "Ясмина", "Александра", "Софья", "Вероника", "Ангелина", "Сафия", "Ника", "Кира", "Мирослава", "Ольга", "Ульяна", "Милана", "Наталья", "Эмилия", "Аврора", "Мия", "Алёна", "Злата", "Ксения", "Арина", "Мадина", "Амелия", "Вера", "Лидия", "Мила", "Фатима", "Таисия", "Алия", "Майя"}
	femaleLastNames := []string{"Михайлова", "Щукина", "Беляева", "Иванова", "Ларионова", "Иванова", "Орлова", "Савельева", "Чумакова", "Егорова", "Савина", "Одинцова", "Соколова", "Шаповалова", "Громова", "Панфилова", "Елисеева", "Савина", "Исаева", "Кузнецова", "Краснова", "Исаева", "Суворова", "Пономарева", "Смирнова", "Терентьева", "Захарова", "Чижова", "Тарасова", "Васильева", "Соловьева", "Бурова", "Ковалева", "Бондарева", "Андреева", "Жданова", "Павлова", "Давыдова", "Золотова", "Соболева", "Тихомирова", "Винокурова", "Матвеева", "Борисова", "Волкова", "Панкратова", "Крылова", "Егорова", "Ерофеева", "Иванова"}

	switch randGender() {
	case Male:
		result += maleNames[RandInt(0, len(maleNames))] + ";" + maleLastNames[RandInt(0, len(maleLastNames))] + ";" + Male
	case Female:
		result += femaleNames[RandInt(0, len(femaleNames))] + ";" + femaleLastNames[RandInt(0, len(femaleLastNames))] + ";" + Female
	}

	return result
}
