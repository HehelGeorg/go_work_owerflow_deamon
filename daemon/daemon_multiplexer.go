package daemon_package

import "fmt"

func(d *Daemon) RunMultiplexer() {

	select{

	case themeId, ok := <- d.MainChan:
		
		if !ok {
			fmt.Printf(" MULTIPLEXER: входной недоступен или закрыт, завершает работу \n")
			return 
		}

		for  _ , plexingChan := range d.MultiplexingChans{
			select{

			case plexingChan <- themeId:
					
			default:
				fmt.Printf(" MUlTIPLEXR: канал заполнен, приступаю к очищению\n")
				for len(plexingChan) > 0 {
					<- plexingChan
				}
				plexingChan <- themeId
			}

		}



	
	case <- d.ctx.Done():
		fmt.Printf("\n MULTIPLEXER: завершает работу\n")
		return
	}


	



}