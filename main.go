package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func merge(arr1, arr2 []float32) []float32 {
	var i, j = 0, 0
	var c []float32
	for i < len(arr1) || j < len(arr2) {
		if i < len(arr1) && (j >= len(arr2) || arr1[i] < arr2[j]) {
			c = append(c, arr1[i])
			i++
		} else {
			c = append(c, arr2[j])
			j++
		}
	}
	return c
}

func mergeSort(arr []float32) []float32 {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	var arr1 = mergeSort(arr[:mid])
	var arr2 = mergeSort(arr[mid:])
	return merge(arr1, arr2)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Merge sort application")

	// Настройки Окна
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	myWindow.SetContent(widget.NewLabel("Hello!"))

	// Содержимое окна
	mainLabel := widget.NewLabel("Сортировка методом слияния")
	input := widget.NewEntry()
	input.SetPlaceHolder("Введите значения через пробел")
	errLabel := widget.NewLabel("Здесь будут сообщения об ошибках")
	resLabel := widget.NewLabel("Здесь будет результат сортировки")

	// Кнопка сортировки
	sortButton := widget.NewButton("Отсортировать", func() {
		inputText := input.Text

		numStrings := strings.Split(inputText, " ")
		var numbers []float32
		var flagErr = false
		if inputText == "" { // Проверка на пустоту ввода
			errLabel.SetText("Ничего не введнено")
			return
		}

		for _, s := range numStrings { // Обработка ввода
			num, err := strconv.ParseFloat(strings.TrimSpace(s), 32)
			if err == nil {
				numbers = append(numbers, float32(num))
			}
			if err != nil { // Запоминаем, есть ли недопустимые символы
				flagErr = true
			}
		} // for

		if flagErr { // Вывод сообщения об ошибках (кроме пустоты)
			errLabel.SetText("Недопустимые символы не обнаружены")
		} else {
			errLabel.SetText("Недопустимые символы. Программа их проигнорировала")
		}

		var resArr = mergeSort(numbers)

		if resArr != nil { // Проверка на существование результата
			resLabel.SetText(fmt.Sprint(resArr))
		} else {
			resLabel.SetText("Здесь будет результат сортировки")
		}

	}) // sortButton

	myWindow.SetContent(container.NewVBox(
		mainLabel,
		input,
		sortButton,
		errLabel,
		resLabel,
	))

	myWindow.ShowAndRun()
}
