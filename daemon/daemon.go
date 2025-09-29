package daemon_package
import(
"context"
. "go_work_oewflow_daemon/daemon/theme"
. "go_work_oewflow_daemon/gRPC/grpc_dispatcher"
)

type Daemon struct{
	Themes map[int]Theme
	MainChan chan int 
	MultiplexingChans map[string]chan int 
	GrpcRouter *Dispatcher
	errChan chan error
	ctx context.Context
}



func(d *Daemon) Run_Daemon(){

}


