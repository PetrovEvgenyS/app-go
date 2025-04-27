package main

import "fmt"

// bookmarkMap определяет тип для хранения закладок в формате "название" - "URL"
type bookmarkMap map[string]string

func main() {
	// Инициализация пустого списка закладок
	bookmarks := bookmarkMap{}
	fmt.Println("__ Приложение для закладок __")

Menu: // Бесконечный цикл для отображения меню и обработки выбора пользователя
	for {
		// Получение выбора пользователя из меню
		variant := getMenu()
		switch variant {
		case 1:
			// Печать всех закладок
			fmt.Println("Закладки:")
			printBookmarks(bookmarks)
		case 2:
			// Добавление новой закладки
			addBookmark(bookmarks)
		case 3:
			// Удаление закладки
			deleteBookmark(bookmarks)
		case 4:
			// Выход из приложения
			fmt.Println("Выход из приложения.")
			break Menu // Выход из бесконечного цикла
		default:
			// Обработка неверного ввода
			fmt.Println("Неверный вариант. Попробуйте снова.")
		}
	}
}

// getMenu отображает меню и возвращает выбранный пользователем вариант
func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

// printBookmarks выводит список всех закладок
func printBookmarks(bookmarks bookmarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Закладки отсутствуют.")
		return
	}
	for key, value := range bookmarks {
		fmt.Println(key, " : ", value)
	}
}

// addBookmark добавляет новую закладку в список
func addBookmark(bookmarks bookmarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название закладки: ")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите URL закладки: ")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
	fmt.Println("Закладка добавлена.")
}

// deleteBookmark удаляет закладку из списка по названию
func deleteBookmark(bookmarks bookmarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название закладки для удаления: ")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
