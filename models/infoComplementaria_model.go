package models

type InfoComplementaria struct {
	Id                int    `json:"Id"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
}

type IdWrapper struct {
	Id float64 `json:"Id"`
}

type IdentificacionTercero struct {
	Numero             string    `json:"Numero"`
	DigitoVerificacion int       `json:"DigitoVerificacion"`
	TipoDocumentoId    IdWrapper `json:"TipoDocumentoId"`
	TerceroId          IdWrapper `json:"TerceroId"`
	Activo             bool      `json:"Activo"`
}
