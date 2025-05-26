package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

//func merge(arr1, arr2 []int) []int {
//	var i, j = 0, 0
//	var c = []int{}
//	for i < len(arr1) || j < len(arr2) {
//		if i < len(arr1) && (j >= len(arr2) || arr1[i] < arr2[j]) {
//			c = append(c, arr1[i])
//			i++
//		} else {
//			c = append(c, arr2[j])
//			j++
//		}
//	}
//	return c
//}
//
//func merge_sort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//	mid := len(arr) / 2
//	var arr1 = merge_sort(arr[:mid])
//	var arr2 = merge_sort(arr[mid:])
//	return merge(arr1, arr2)
//}

func main() {
	//var arr = []int{5, 2, -1, 0, 5, 16, 27, 8, 9, 10}
	//var sortedArr = merge_sort(arr)
	//fmt.Println(arr)
	//fmt.Println(sortedArr)

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
	resultLabel := widget.NewLabel("Здесь будет результат сортировки")
	sortButton := widget.NewButton("Отсортировать", func() {
		inputText := input.Text

		numStrings := strings.Split(inputText, " ")
		var numbers []int
		for _, s := range numStrings {
			num, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				resultLabel.SetText("ERR")
				return
			}
			numbers = append(numbers, num)
		}
	}) // sortButton

	myWindow.SetContent(container.NewVBox(
		mainLabel,
		input,
		sortButton,
		resultLabel,
	))

	myWindow.ShowAndRun()
}
