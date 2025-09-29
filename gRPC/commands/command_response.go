package commands

type CommandResponse struct {
    Success bool
    Message string
    Payload map[string]string 
}