package main

import (
	"fmt"
)

//1
func task1() {
	a := map[string]int{"хлеб": 100, "молоко": 150, "сыр": 300}
	b, c := a["молоко"]
	if c {
		fmt.Println(b)
	} else {
		fmt.Println("Товар отсутствует")
	}

	b, c = a["сок"]
	if c {
		fmt.Println(b)
	} else {
		fmt.Println("Товар отсутствует")
	}
}

//2
func divide(a int, b int) (int, interface{}) {
	if b == 0 {
		return 0, "Ошибка: деление на ноль невозможно"
	}
	return a / b, nil
}

//3
func checkAge(a int) error {
	if a < 0 || a > 120 {
		return errors.New("указан некорректный возраст")
	}
	return nil
}

//4
func validatePassword(a string) (bool, error) {
	if len(a) < 6 {
		return false, errors.New("пароль слишком короткий")
	}
	return true, nil
}

//5
type invalidRadiusError struct{}

func (a invalidRadiusError) Error() string {
	return "радиус не может быть отрицательным"
}

//6
func calculateCircleArea(a int) (float64, error) {
	if a < 0 {
		return 0, invalidRadiusError{}
	}
	return 3.14 * float64(a) * float64(a), nil
}

//7
func findUser(a []string, b string) (int, error) {
	for c, d := range a {
		if d == b {
			return c, nil
		}
	}
	return -1, errors.New("пользователь не найден")
}

func main() {
	task1()

	a, b := divide(10, 2)
	fmt.Println(a, b)
	a, b = divide(10, 0)
	fmt.Println(a, b)

	c := checkAge(25)
	fmt.Println(c)
	c = checkAge(-1)
	fmt.Println(c)

	d, e := validatePassword("abc")
	fmt.Println(d, e)
	d, e = validatePassword("abcdef")
	fmt.Println(d, e)

	f := invalidRadiusError{}
	fmt.Println(f.Error())

	g, h := calculateCircleArea(-2)
	fmt.Println(g, h)

	i, j := findUser([]string{"Аня", "Борис", "Вася"}, "Борис")
	fmt.Println(i, j)
	i, j = findUser([]string{"Аня", "Борис", "Вася"}, "Оля")
	fmt.Println(i, j)
}
