package daemon_package

import (
	"bufio"
	"fmt"
. "go_work_oewflow_daemon/daemon/workspace_listener"
)

// RunWorkerspaceListener на основе выбранного workspace
// изменяет тему рабочего сзцола
//
// Детали реализации:
// Подключается к сокету hyprland
// Слушает события, пытаясь найти событие о смене рабочего стола
// Находя его, парсит его номер
// Относительно этого номер вызвает методы темы, которая находится пот таким же номером
func(d *Daemon) RunWorkspaceListenerWorker() {

	// Подключение к сокету
	conn, err := ConnectToHyprlandSocket()
	if err != nil {
		fmt.Printf(" MAINWORKER: Критическая ошибка: не удалось подключиться к сокету Hyprland: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println(" MAINWORKER: Успешно подключено к сокету Hyprland. Ожидание событий...\n ")

	// Создание сканера на основе подключения сокета

	scanner := bufio.NewScanner(conn)

	// Тело

	for scanner.Scan() {

		event := scanner.Text()

		// Получаем Id рабочего стола,
		themeId, ok := ParseWorkspaceID(event)

		if !ok {
			continue // игнорирование событий, не связанных с переключением рабочего стола
		}

		_ , ok = d.Themes[themeId]

		if !ok {
			fmt.Printf(" MAINWORKER: Рабочему столу с номером { %v } не была присвоена собственная тема  \n", themeId)
			continue
		}
	
	
		//Отправка сигнала:
		select{

		case d.MainChan <- themeId :
			fmt.Printf(" MAINWORKER: Номер рабочего стола имеет конфигурацию. \nКонфигурация применяется:\n")
			
		case  <- d.ctx.Done() :
			fmt.Printf(" MAINWORKER: Работа главного воркера завершена\n ")
			return
		}


	}

}

