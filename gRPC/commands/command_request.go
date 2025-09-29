package commands


type CommandRequest struct{
	TargetWorker string
	Action string
	Args map[string]string

	ReplyChan chan CommandResponse


}
