// Package theme предоставляет функционал для применения тем оформления,
// включающих обои, цветовые палитры и фоновое звуковое сопровождение на основе конфигурации
// Для рабочего стола hyprland
package theme

import (
	"context"
	"fmt"
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

// PlaySoundLoop начинает фоновое зацикленное воспроизведение аудиофайла.
// Путь к файлу определяется полем Sound.
//
// Возвращает функцию stop, вызов которой завершит воспроизведение,
// и ошибку, если воспроизведение не удалось начать.
// Функция stop безопасна для многократного вызова и может быть вызвана из любого goroutine.
func (t *Theme) PlaySound() (stopFunc func(), err error) {

	if t.Sound == "" {
		return nil, nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "mpv", "--loop", "yes", t.Sound)

	fmt.Printf("Выполнение: %s\n", cmd.String())

	if err := cmd.Start(); err != nil {
		cancel() // Важно: освободить контекст в случае ошибки
		return nil, fmt.Errorf("failed to start sound player: %w", err)
	}

	// Возвращаем функцию, которая отменяет контекст и ожидает завершения процесса.
	stop := func() {
		cancel()
		_ = cmd.Wait() // Игнорируем ошибку, так как мы просто останавливаем воспроизведение
	}

	return stop, nil
}
