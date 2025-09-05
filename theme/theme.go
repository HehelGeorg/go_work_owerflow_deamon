// Package theme предоставляет функционал для применения тем оформления,
// включающих обои, цветовые палитры и фоновое звуковое сопровождение на основе конфигурации
// Для рабочего стола hyprland
package theme

import (
	"fmt"
	"os"
	"os/exec"
)

// Theme описывает правила активации обоев, звука и палитры цветов.
// Все поля являются опциональными; пустые значения будут проигнорированы.
type Theme struct {
	Wallpaper     string `toml:"wallpaper"`
	PaletteScript string `toml:"palette_script"`
	Sound         string `toml:"sound"`
}

func (t *Theme) IsEmpty() bool {
	return t.Wallpaper == "" && t.PaletteScript == "" && t.Sound == ""
}

// ApplyWallpaper активирует обои на основе правил.
// Для применения обой используется путь, указанный в поле Wallpaper.
func (t *Theme) ApplyWallpaper() error {
	if t.Wallpaper == "" {
		return nil
	}
	cmd := exec.Command("swww", "img", t.Wallpaper)
	fmt.Printf("Выполнение: %s\n", cmd.String())
	return cmd.Run()
}

// ApplyPalette активирует цветовую палитру, запуская скрипт.
// Путь к скрипту определяется полем PaletteScript
func (t *Theme) ApplyPalette() error {
	if t.PaletteScript == "" {
		return nil
	}

	cmd := exec.Command("bash", t.PaletteScript)
	fmt.Printf("Выполнение: %s\n", cmd.String())
	return cmd.Run()
}

// PlaySound начинает фоновое зацикленное воспроизведение аудиофайла.
// Путь к файлу определяется полем Sound.
//
// Возвращает процесс mpv, что позволяет вызывающему коду управлять им
// (например, получить PID, отправить дополнительные сигналы или принудительно завершить).
// Если поле Sound пустое, возвращается (nil, nil).
// В случае ошибки запуска возвращается (nil, error).
//
// Внимание: вызывающий код несёт ответственность за завершение процесса
// с помощью process.Kill() или process.Wait() по окончании работы.
func (t *Theme) PlaySound() (*os.Process, error) {

	if t.Sound == "" {
		return nil, nil
	}

	cmd := exec.Command("mpv", "--loop", "yes", t.Sound)

	fmt.Printf("Выполнение: %s\n", cmd.String())

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start sound player: %w", err)
	}

	return cmd.Process, nil
}
