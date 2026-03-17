package models

import (
	"encoding/json"
	"strings"
	"time"
)

// CustomTime handles multiple date formats
type CustomTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler interface
func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)

	if str == "null" || str == "" {
		ct.Time = time.Time{}
		return nil
	}

	// List of possible time formats
	formats := []string{
		"2006-01-02T15:04:05Z07:00",              // RFC3339
		"2006-01-02T15:04:05.000Z",               // RFC3339 with milliseconds
		"2006-01-02T15:04:05",                    // ISO format without timezone
		"2006-01-02 15:04:05.999999 -0700 MST",   // Custom format with timezone
		"2006-01-02 15:04:05.999999 +0000 +0000", // Your specific format
		"2006-01-02 15:04:05",                    // Simple datetime
		"2006-01-02",                             // Date only
	}

	var err error
	for _, format := range formats {
		ct.Time, err = time.Parse(format, str)
		if err == nil {
			return nil
		}
	}

	return err
}

// MarshalJSON implements json.Marshaler interface
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ct.Time.Format(time.RFC3339))
}

// TipoContribuyenteId model
type TipoContribuyenteId struct {
	Activo            bool       `json:"Activo"`
	CodigoAbreviacion string     `json:"CodigoAbreviacion"`
	Descripcion       string     `json:"Descripcion"`
	FechaCreacion     CustomTime `json:"FechaCreacion"`
	FechaModificacion CustomTime `json:"FechaModificacion"`
	Id                int        `json:"Id"`
	Nombre            string     `json:"Nombre"`
}

// TerceroId model
type TerceroId struct {
	Activo              bool                `json:"Activo"`
	FechaCreacion       CustomTime          `json:"FechaCreacion"`
	FechaModificacion   CustomTime          `json:"FechaModificacion"`
	FechaNacimiento     CustomTime          `json:"FechaNacimiento"`
	Id                  int                 `json:"Id"`
	LugarOrigen         int                 `json:"LugarOrigen"`
	NombreCompleto      string              `json:"NombreCompleto"`
	PrimerApellido      string              `json:"PrimerApellido"`
	PrimerNombre        string              `json:"PrimerNombre"`
	SegundoApellido     string              `json:"SegundoApellido"`
	SegundoNombre       string              `json:"SegundoNombre"`
	TipoContribuyenteId TipoContribuyenteId `json:"TipoContribuyenteId"`
	UsuarioWSO2         string              `json:"UsuarioWSO2"`
}

// TipoDocumentoId model
type TipoDocumentoId struct {
	Activo            bool       `json:"Activo"`
	CodigoAbreviacion string     `json:"CodigoAbreviacion"`
	Descripcion       string     `json:"Descripcion"`
	FechaCreacion     CustomTime `json:"FechaCreacion"`
	FechaModificacion CustomTime `json:"FechaModificacion"`
	Id                int        `json:"Id"`
	Nombre            string     `json:"Nombre"`
	NumeroOrden       int        `json:"NumeroOrden"`
}

// DatosIdentificacion model
type DatosIdentificacion struct {
	Activo             bool            `json:"Activo"`
	CiudadExpedicion   int             `json:"CiudadExpedicion"`
	DigitoVerificacion int             `json:"DigitoVerificacion"`
	DocumentoSoporte   int             `json:"DocumentoSoporte"`
	FechaCreacion      CustomTime      `json:"FechaCreacion"`
	FechaExpedicion    CustomTime      `json:"FechaExpedicion"`
	FechaModificacion  CustomTime      `json:"FechaModificacion"`
	Id                 int             `json:"Id"`
	Numero             string          `json:"Numero"`
	TerceroId          TerceroId       `json:"TerceroId"`
	TipoDocumentoId    TipoDocumentoId `json:"TipoDocumentoId"`
}

// DatosEstudiante Oracle JBPM model
type DatosEstudianteOracle struct {
	Codigo            int64   `json:"codigo,string"`
	Estado            string  `json:"estado"`
	Periodo           string  `json:"periodo"`
	Promedio          float64 `json:"promedio,string"`
	NombreTipoCarrera string  `json:"nombre_tipo_carrera"`
	Nombre            string  `json:"nombre"`
	Pensum            int     `json:"pensum,string"`
	Rendimiento       float64 `json:"rendimiento,string"`
	PorcentajeCursado int     `json:"porcentaje_cursado,string"`
	Carrera           int     `json:"carrera,string"`
	Nivel             string  `json:"nivel"`
	Anio              int     `json:"anio,string"`
	TipoCarrera       int     `json:"tipo_carrera,string"`
}
type EstudianteCollection struct {
	DatosEstudiante []DatosEstudianteOracle `json:"datosEstudiante"`
}
type DatosEstudianteResponse struct {
	EstudianteCollection EstudianteCollection `json:"estudianteCollection"`
}

// DatosCarrerasOracle JBPM model
type DatosProyectoOracle struct {
	ASCraCod      int    `json:"AS_CRA_COD,string"`
	ASCraCodSnies string `json:"AS_CRA_COD_SNIES"`
	ASEstado      string `json:"AS_ESTADO"`
	ASCraNom      string `json:"AS_CRA_NOM"`
}
type ProyectosCollection struct {
	Proyecto []DatosProyectoOracle `json:"proyecto"`
}
type ProyectosResponse struct {
	Proyectos ProyectosCollection `json:"proyectos"`
}

// inscripciones model
type TipoInscripcion struct {
	Id                int    `json:"Id"`
	Nombre            string `json:"Nombre"`
	Descripcion       string `json:"Descripcion"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	Activo            bool   `json:"Activo"`
	NumeroOrden       int    `json:"NumeroOrden"`
	NivelId           int    `json:"NivelId"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
	Especial          bool   `json:"Especial"`
}

type EstadoInscripcion struct {
	Id                int    `json:"Id"`
	Nombre            string `json:"Nombre"`
	Descripcion       string `json:"Descripcion"`
	CodigoAbreviacion string `json:"CodigoAbreviacion"`
	Activo            bool   `json:"Activo"`
	NumeroOrden       int    `json:"NumeroOrden"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
}

type Inscripcion struct {
	Id                  int               `json:"Id"`
	PersonaId           int               `json:"PersonaId"`
	ProgramaAcademicoId int               `json:"ProgramaAcademicoId"`
	ReciboInscripcion   string            `json:"ReciboInscripcion"`
	PeriodoId           int               `json:"PeriodoId"`
	EnfasisId           int               `json:"EnfasisId"`
	NotaFinal           float64           `json:"NotaFinal"`
	AceptaTerminos      bool              `json:"AceptaTerminos"`
	FechaAceptaTerminos string            `json:"FechaAceptaTerminos"`
	Activo              bool              `json:"Activo"`
	FechaCreacion       string            `json:"FechaCreacion"`
	FechaModificacion   string            `json:"FechaModificacion"`
	Credencial          int               `json:"Credencial"`
	Opcion              int               `json:"Opcion"`
	EstadoInscripcion   EstadoInscripcion `json:"EstadoInscripcionId"`
	TipoInscripcion     TipoInscripcion   `json:"TipoInscripcionId"`
	TipoCupo            int               `json:"TipoCupo"`
}

// RecibosOracle model
type ReciboResponse struct {
	ReciboCollection ReciboCollection `json:"reciboCollection"`
}

type ReciboCollection struct {
	Recibo []Recibo `json:"recibo"`
}

type Recibo struct {
	Estado              string `json:"estado"`
	Ano                 int    `json:"ano"`
	Cuota               int    `json:"cuota"`
	Periodo             int    `json:"periodo"`
	FechaPagado         string `json:"fecha_pagado"`
	Secuencia           int    `json:"secuencia"`
	Documento           string `json:"documento"`
	FechaOrdinario      string `json:"fecha_ordinario"`
	Pago                string `json:"pago"`
	Nombre              string `json:"nombre"`
	Fecha               string `json:"fecha"`
	ValorExtraordinario int    `json:"valor_extraordinario"`
	Observaciones       string `json:"observaciones"`
	Carrera             int    `json:"carrera"`
	ValorPagado         int    `json:"valor_pagado"`
	FechaExtraordinario string `json:"fecha_extraordinario"`
	ValorOrdinario      int    `json:"valor_ordinario"`
}
