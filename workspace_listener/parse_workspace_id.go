package workspacelistener

import (
	"fmt"
	"strings"
)

// parseWorkspaceID парсит событие сокета hyprland
// в целочисленный id рабочего стола, с которым событие
// ассоциировано
func parseWorkspaceID(event string) (int, bool) {
	if !strings.HasPrefix(event, "workspace>>") {
		return 0, false
	}
	var id int
	_, err := fmt.Sscanf(event, "workspace>>%d", &id)
	if err != nil {
		return 0, false
	}
	return id, true
}
