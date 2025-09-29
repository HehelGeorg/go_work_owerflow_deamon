package paramexe
import(
	"flag"
    "fmt"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
)


// InitConfig() устанавливает флаг -config
// по которому можно изменить путь к файлу конфигурации
// Если файл конфигурации, установленный через -config не найден,
// Либо он не использовался вообще, то файл конфигурации
// ПО УМОЛЧАНИЮ БУДЕТ В САМОЙ ПАПКЕ ПРОЕКТА
// если имели место быть ошибки, возвращается пустая строка и !nil ошибка
func InitConfig() (string, error) {


	// Установка флага config
	configFlag := flag.String("config", "", "Путь к конфигурационному файлу")
	flag.Parse()

	configPath := *configFlag

	if configPath == "" {
		// Если флаг не указан, используем config.toml рядом с программой
		exePath, err := os.Executable()
		if err != nil {
			return "", fmt.Errorf("ошибка получения пути: %v", err)
		}
		configPath = filepath.Join(filepath.Dir(exePath), "config.toml")
	}

	// Установка дополнительных ограничений viper для файла конфигурации
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")

	
	// Проверяем был ли загружен конфиг
	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("ошибка загрузки конфига: %v", err)
	}

	// Возвращаем строку подключения для дальнейшего unmarshall/десериализации
	return viper.ConfigFileUsed(), nil
}