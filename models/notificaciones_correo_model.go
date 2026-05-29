package models

type Destinatarios struct {
	ToAddresses []string `json:"ToAddresses"`
}

type Asunto struct {
	Data string `json:"Data"`
}

type Mensaje struct {
	Body        interface{}   `json:"Body"`
	Subject     Asunto        `json:"Subject"`
	Attachments []interface{} `json:"Attachments"`
}

type Correo struct {
	Destination Destinatarios `json:"Destination"`
	Message     Mensaje       `json:"Message"`
	SourceEmail string        `json:"SourceEmail"`
	SourceName  string        `json:"SourceName"`
}
