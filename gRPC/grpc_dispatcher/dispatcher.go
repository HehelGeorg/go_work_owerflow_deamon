package grpcrouter
import (
	"context"
	. "go_work_oewflow_daemon/gRPC/commands"
	"log"
	"sync"
)


const (
	// Размер буфера для входящей очереди gRPC-команд.
	DispatcherInCapacity = 20 
)

// Dispatcher - маршрутизатор команд между gRPC-сервером и воркерами (акторами).
type Dispatcher struct {
	ctx context.Context
	wg  sync.WaitGroup

	// IN: Единый входящий канал для ВСЕХ gRPC-команд. Это основная очередь.
	In chan CommandRequest
	// OUT: Карта каналов для пересылки команд конкретным воркерам.
	Out map[string]chan CommandRequest
}

// NewDispatcher создает новый экземпляр Диспетчера.
func NewDispatcher(ctx context.Context) *Dispatcher {
	return &Dispatcher{
		ctx: ctx,
		In:  make(chan CommandRequest, DispatcherInCapacity),
		Out: make(map[string]chan CommandRequest),
	}
}



// Регистрация воркера на получение запросов
func (d *Dispatcher) RegisterWorker(workerName string, workerInChan chan CommandRequest) {
	if _, exists := d.Out[workerName]; exists {
		log.Fatalf("Диспетчер: Воркер '%s' уже зарегистрирован!", workerName)
	}
	d.Out[workerName] = workerInChan
}

// Run запускает горутину маршрутизатора.
func (d *Dispatcher) Run() {
	// Добавляем горутину в очередь,  при ctx.done удаляем
	d.wg.Add(1)

	go func() {

		// Когда сработает <- ctx.done горутина безопасно свернется
		defer d.wg.Done()
		log.Println("Диспетчер: Запущен и готов к маршрутизации команд.")


		//центральный цикл
		for {

			select {

			// Если появился новый запрос
			case command, ok := <-d.In:



				// Входящий канал закрыт?
				if !ok {
					
					log.Println("Диспетчер: Входящий канал закрыт. Завершаю работу.")
					return
				}
				
								
				
				// Мащрутизация

				// Ищем канал, искомого воркера
				targetChan, exists := d.Out[command.TargetWorker]

				// Канал воркера существует?
				if !exists {
					log.Printf("Диспетчер: Неизвестный воркер '%s'. Игнорирую команду.", command.TargetWorker)
					continue
				}

				// Асинхронная отправка команды воркеру.
				select {

				// Отправляем воркеру запрос
				case targetChan <- command:
					// Команда успешно отправлена воркеру.
					log.Printf("Диспетчер: Команда '%s' маршрутизирована воркеру '%s'.", command.Action, command.TargetWorker)

				// Контекст отменен
				case <-d.ctx.Done():
					// Контекст отменен во время попытки отправки.
					log.Println("Диспетчер: Контекст отменен. Прерываю отправку.")
					return

				// Прочие случаи
				default:
					log.Printf("Диспетчер: Канал воркера '%s' переполнен! Воркер не успевает. Команда пропущена.", command.TargetWorker)
				}



			// Если контекст отменен	
			case <-d.ctx.Done():
				// Главный контекст отменен.
				log.Println("Диспетчер: Контекст отменен. Завершаю работу.")
				return
			}
		}
	}()
}

// Wait ожидает завершения горутины Диспетчера.
func (d *Dispatcher) Wait() {
	d.wg.Wait()
}