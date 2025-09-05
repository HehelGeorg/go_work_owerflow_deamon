package workspacelistener

import (
	"bufio"
	"fmt"
	. "go_work_oewflow_daemon/theme"
)

// RunWorkerspaceListener на основе выбранного workspace
// изменяет тему рабочего сзцола
//
// Детали реализации:
// Подключается к сокету hyprland
// Слушает события, пытаясь найти событие о смене рабочего стола
// Находя его, парсит его номер
// Относительно этого номер вызвает методы темы, которая находится пот таким же номером
func RunWorkerspaceListener(themes map[int]Theme, SIG chan<- int) {

	// Подключение к сокету
	conn, err := connectToHyprlandSocket()
	if err != nil {
		fmt.Printf("Критическая ошибка: не удалось подключиться к сокету Hyprland: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Успешно подключено к сокету Hyprland. Ожидание событий...")

	// Создание сканера на основе подключения сокета

	scanner := bufio.NewScanner(conn)

	// Тело

	for scanner.Scan() {

		event := scanner.Text()

		// Получаем Id рабочего стола,
		workspaceId, ok := parseWorkspaceID(event)

		if !ok {
			continue // игнорирование событий, не связанных с переключением рабочего стола
		}

		// если событие связано, получаем его из карты за O(1)

		theme, themeExists := themes[workspaceId]

		// пропуск рабочих столов, не указанных в карте
		if !themeExists {
			fmt.Printf("Тема для стола %d не найдена, ничего не делаем.\n", workspaceId)
			continue
		}

		// Активация темы

		// Активация обоев
		if err := theme.ApplyWallpaper(); err != nil {
			fmt.Printf("Ошибка смены обоев: %v\n", err)
		}

		//Активация палитры
		if err := theme.ApplyPalette(); err != nil {
			fmt.Printf("Ошибка смены палитры %v\n", err)
		}

		//Отправка сигнала:
		SIG <- workspaceId

	}

}
