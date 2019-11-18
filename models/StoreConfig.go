package models

type StoreConfig struct {
	AskPin      *PinResponse
	Activate    *ActivePinResponse
	Aria2Config *Aria2Config
	Address     string
	Token     string
}
