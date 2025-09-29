package daemon_package

import (
	"fmt"
	"net"
	"os"
)

// ConnectToHyprlandSocket Подключается  к соекту Hyprland
// возвращает объект сокета для дальнейшего использовани
func ConnectToHyprlandSocket() (net.Conn, error) {
	signature := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if signature == "" {
		return nil, fmt.Errorf("переменная HYPRLAND_INSTANCE_SIGNATURE не установлена")
	}
	socketPath := fmt.Sprintf("/tmp/hypr/%s/.socket2.sock", signature)
	return net.Dial("unix", socketPath)
}
