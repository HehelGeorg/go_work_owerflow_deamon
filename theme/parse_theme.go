package theme

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	WorkSpaces map[string]Theme `toml:"workspaces"`
}

// Parse читает и парсит TOML-конфигурацию из указанного файла.
//
// Ожидает, что конфигурация содержит раздел [workspaces] с ключами в виде строковых
// представлений номеров рабочих столов (например, "1", "2").
//
// Возвращает карту, где ключ — целочисленный номер рабочего стола (от 1 до 9),
// а значение — соответствующая тема. Пустые темы пропускаются.
// Ключи, которые не могут быть преобразованы в числа от 1 до 9, игнорируются.
func Parse(configPath string) (map[int]Theme, error) {
	var config Config

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, fmt.Errorf("ошибка чтения или парсинга config.toml: %w", err)
	}

	themesMap := make(map[int]Theme)

	for wsStr, theme := range config.WorkSpaces {

		if theme.IsEmpty() {
			continue
		}

		var wsInt int

		// Преобразуем ключ "1", "2" и т.д. в число.
		if _, err := fmt.Sscanf(wsStr, "%d", &wsInt); err != nil {
			fmt.Printf("Предупреждение: неверный ключ рабочего стола в конфиге: %s\n", wsStr)
			continue
		}

		if wsInt > 0 && wsInt < 10 {
			themesMap[wsInt] = theme
		}

	}

	return themesMap, nil
}
