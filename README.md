# Merge Sort Application

Приложение с графическим интерфейсом для сортировки чисел методом слияния (Merge Sort)

<p align="center">
  <img src="https://s.iimg.su/s/31/9vEndflQtaivI1QLiC77H5Vx3nPoZ4fv6Ju2hO8H.png" alt="Скриншот приложения" width="600">
</p>

## Особенности

- 🚀 Сортировка чисел методом слияния
- 🖥️ Простой и интуитивно понятный интерфейс
- 📋 Ввод данных вручную или вставка из буфера обмена
- 📤 Копирование результатов в буфер обмена
- 🔄 Автоматический перенос длинных результатов
- 🌍 Поддержка Windows, macOS и Linux

## Требования

- [Go](https://golang.org/dl/) версии 1.16+
- Системные зависимости:
  - **Windows**: без дополнительных требований
  - **macOS**: `xcode-select --install`
  - **Linux**: 
    ```
    sudo apt install gcc libgl1-mesa-dev xorg-dev  # Debian/Ubuntu
    sudo dnf install gcc mesa-libGL-devel libX11-devel  # Fedora
    ```

## Установка и запуск

1. Клонируйте репозиторий:
```
git clone https://github.com/your-username/merge-sort-app.git
cd merge-sort-app
```
2. Установите зависимости:
```
go mod download
```
3. Запустите приложение:
```
go run main.go
```
Или соберите бинарник:
```
go build -o merge-sort-app
./merge-sort-app         # Linux/macOS
.\merge-sort-app.exe     # Windows
```

## Использование
1. Введите числа через пробел в поле ввода

2. Нажмите "Отсортировать" для выполнения сортировки

- Дополнительные функции:
  - "Вставить из буфера" - вставляет данные из буфера обмена
  - "Копировать результат" - копирует отсортированные числа в буфер

## Ключевые функции

```
// Сортировка слиянием
func mergeSort(arr []float32) []float32 {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}

// Слияние двух отсортированных массивов
func merge(left, right []float32) []float32 {
    result := make([]float32, 0, len(left)+len(right))
    i, j := 0, 0
    
    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }
    
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}
```

## Контакнтая информация

Telegram: @ruseg557

Email: KadriRF@mpei.ru
