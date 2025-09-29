package daemon_package

import (
	"fmt"
	. "go_work_oewflow_daemon/gRPC/commands"
	"os"

	
)




type push struct{
	response CommandResponse
	request CommandRequest
}

func(p *push) WithOperationResult(succes bool){
	p.response.Success = succes
}

func(p *push) WithMessage( message string ){
	p.response.Message = message
}

func(p *push) WithArgs(args map[string]string){
	p.response.Payload = args
}

func(p *push) WithRequest(request CommandRequest){
	p.request = request
}


func(p *push) PushResponse() {


	 if p.request.ReplyChan == nil {
        return // Это асинхронный запрос, не нужно отвечать
    }
	p.request.ReplyChan <- p.response

}

func (d *Daemon) RunwallpaperDynamicWorker() {
	select {


	// Обработка запросов
	case request := <- d.GrpcRouter.Out["Wallpaper"]:



		switch request.Action{

		// Смена обоев
		case "change-walpapper":
			// Менять обои
			// В разработке
			buildPush := push{}
			buildPush.WithRequest(request)
			buildPush.WithOperationResult(true)
			buildPush.WithMessage("Walppaper complete changed")
			buildPush.PushResponse()

		}






	// Обработка сигналов 
	case themeId := <-d.MultiplexingChans["wallpaper"]:

		theme := d.Themes[themeId]

		err := theme.ApplyWallpaper()

		if err != nil {
			d.errChan <- fmt.Errorf(" DYNAMIC_WALLPAPER_WORKER: Ошибка при попытке установке обоев: \n %v \n ", err)
			return
		}

	case <-d.ctx.Done():
		fmt.Printf(" DYNAMIC_WALLPAPER_WORKER: Завершение работы")
		// Добавлю освобождение ресурсов
	}
}

func (d *Daemon) RunPalleteDynamicWorker() {
	select {





		// Обработка запросов
	case request := <- d.GrpcRouter.Out["pallete"]:



		switch request.Action{

		// Смена палитры
		case "change-pallete":
			// Менять палитру
			// В разработке
			buildPush := push{}
			buildPush.WithRequest(request)
			buildPush.WithOperationResult(true)
			buildPush.WithMessage("Pallete complete changed")
			buildPush.PushResponse()

		}




	case themeId := <-d.MultiplexingChans["pallete"]:

		theme := d.Themes[themeId]

		err := theme.ApplyPalette()

		if err != nil {
			d.errChan <- fmt.Errorf(" DYNAMIC_PALLETE_WORKER: Ошибка при попытке установке палитры: \n %v \n ", err)
			return
		}

	case <-d.ctx.Done():
		fmt.Printf(" DYNAMIC_PALLETE_WORKER: Завершение работы")
		// Добавлю освобождение ресурсов
	}

}

func (d *Daemon) RunSoundDynamicWorker() {

	var themeId int

	var _ *os.Process

	select {

			// Обработка запросов
	case request := <- d.GrpcRouter.Out["pallete"]:



		switch request.Action{

		// Смена звук
		case "change-sound":
			// Менять палитру
			// В разработке
			buildPush := push{}
			buildPush.WithRequest(request)
			buildPush.WithOperationResult(true)
			buildPush.WithMessage("Sound complete changed")
			buildPush.PushResponse()


		case "change-volume-sound":

			//Менять громкость звука
			//В разработке 
			buildPush := push{}
			buildPush.WithRequest(request)
			buildPush.WithOperationResult(true)
			buildPush.WithMessage("Sound volume complete changed")
			buildPush.PushResponse()

		}




	case themeId = <-d.MultiplexingChans["sound"]:

		theme := d.Themes[themeId]

		var err error

		_, err = theme.PlaySound()

		if err != nil {
			d.errChan <- fmt.Errorf(" DYNAMIC_WALLPAPER_WORKER: Ошибка при попытке установке обоев: \n %v \n ", err)
			return
		}

	case <-d.ctx.Done():
		fmt.Printf(" DYNAMIC_WALLPAPER_WORKER: Завершение работы")
		// Добавлю освобождение ресурсов
	}

}
