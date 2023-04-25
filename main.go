package main

import (
	"fmt"
)

func main() {
	fmt.Println("Баллы ЕГЭ.")
	fmt.Println("Давайте узнаем, поступите ли Вы в институт иностранных языков?")

	fmt.Println("Введите результат экзамена по русскому языку:")
	var examRussian int
	fmt.Scan(&examRussian)

	fmt.Println("Введите результат экзамена по математике:")
	var examMath int
	fmt.Scan(&examMath)

	fmt.Println("Введите результат экзамена по английскому языку:")
	var examEnglish int
	fmt.Scan(&examEnglish)

	//переменная проходного балла в институт. создана для удобства изменения проходного балла.
	requiredAmount := 275

	//набранная сумма баллов пользователя за три экзамена
	userGenAmount := examRussian + examMath + examEnglish

	if userGenAmount >= requiredAmount {
		fmt.Println("Необходимая сумма баллов для поступления:", requiredAmount)
		fmt.Println("Количество набранных баллов:", userGenAmount)
		fmt.Println("Поздравляем! Вы поступили!")
	} else if userGenAmount > 0 {
		fmt.Println("Необходимая сумма баллов для поступления:", requiredAmount)
		fmt.Println("Количество набранных баллов:", userGenAmount)
		fmt.Println("Сожалеем, но Вы не поступили.")
	} else {
		fmt.Println("Вы ввели странные значения. Попробуйте снова.")
	}

}
