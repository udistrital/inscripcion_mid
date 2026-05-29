package utils

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/utils_oas/request"
)

func SendTemplatedEmail(inputemailtemplated map[string]interface{}) (result error) {
	var resultadoPost map[string]interface{}
	if errSendTemplatedEmail := request.SendJsonEscapeUnicode(beego.AppConfig.String("notificacionService")+"email/enviar_templated_email", "POST", &resultadoPost, inputemailtemplated); errSendTemplatedEmail == nil {

		return nil
	} else {
		result = errSendTemplatedEmail
	}
	return result
}

func SendEmail(inputMail map[string]interface{}) (result error) {
	// Envio de mail
	logs.Info("Entra al envio del servicio de notificaciones")
	logs.Info(beego.AppConfig.String("notificacionService") + "email/enviar_email")
	fmt.Println(inputMail)
	fmt.Println("")
	var resultadoPost map[string]interface{}
	errSendEmail := request.SendJsonEscapeUnicode(beego.AppConfig.String("notificacionService")+"email/enviar_email", "POST", &resultadoPost, inputMail)
	if errSendEmail == nil {
		logs.Info("Todo correcto")
		fmt.Println(resultadoPost)
		fmt.Println("")
		return nil
	} else {
		result = errSendEmail
		logs.Info(result)
	}
	return result
}

func SendNotificacionCambioEstadoSolicitud(data map[string]interface{}, email string) (result error) {
	// Armado de objeto
	logs.Info("Entra al servicio de armado del mail")

	mail := map[string]interface{}{
		"Destination": map[string]interface{}{
			"ToAddresses": []string{email},
		},
		"Message": map[string]interface{}{
			"Body": data,
			"Subject": map[string]interface{}{
				"Data": "Novedad en inscripción SGA",
			},
			"Attachments": []interface{}{},
		},
		"SourceEmail": "notificacionessga@udistrital.edu.co",
		"SourceName":  "Notificaciones inscripciones",
	}
	fmt.Println(mail)
	fmt.Println("")

	return SendEmail(mail)
}

func SendNotificationInscripcionSolicitud(data map[string]interface{}, email string) (result error) {
	var toAddresses []string
	var destinations []map[string]interface{}

	destination := map[string]interface{}{
		"Destination": map[string]interface{}{
			"ToAddresses": append(toAddresses, email),
		},
		"ReplacementTemplateData": data,
	}

	fecha_actual := time.Now()
	m := map[string]interface{}{
		"dia":    fecha_actual.Day(),
		"mes":    GetNombreMes(fecha_actual.Month()),
		"anio":   fecha_actual.Year(),
		"nombre": "",
		"estado": "inscripción solicitada",
	}

	dataEmail := map[string]interface{}{
		"Source":              "Notificacion <notificacionessga@udistrital.edu.co>",
		"Template":            "TEST_SGA_inscripcion-cambio-estado",
		"Destinations":        append(destinations, destination),
		"DefaultTemplateData": m,
	}

	return SendTemplatedEmail(dataEmail)
}

func SendNotificationInscripcionComprobante(data map[string]interface{}, email string, attachments []map[string]interface{}) (result error) {
	var toAddresses []string
	var destinations []map[string]interface{}

	destination := map[string]interface{}{
		"Destination": map[string]interface{}{
			"ToAddresses": append(toAddresses, email),
		},
		"ReplacementTemplateData": data,
		"Attachments":             attachments,
	}

	fecha_actual := time.Now()
	m := map[string]interface{}{
		"dia":     fecha_actual.Day(),
		"mes":     GetNombreMes(fecha_actual.Month()),
		"anio":    fecha_actual.Year(),
		"nombre":  "",
		"periodo": "solicitado",
	}

	dataEmail := map[string]interface{}{
		"Source":              "Notificacion <notificaciones_sga@udistrital.edu.co>",
		"Template":            "TEST_SGA_inscripcion-pago",
		"Destinations":        append(destinations, destination),
		"DefaultTemplateData": m,
	}

	return SendTemplatedEmail(dataEmail)
}
