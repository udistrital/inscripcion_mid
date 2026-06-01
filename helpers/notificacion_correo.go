package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/inscripcion_mid/models"
	"github.com/udistrital/inscripcion_mid/utils"
	"github.com/udistrital/utils_oas/request"
)

func EnviarNotificacionObservacionInscripcion(dataInscripcion map[string]interface{}) {
	// logs.Info("Entra al envío de correo")
	var tercero models.TerceroId
	errPersona := request.GetJson(beego.AppConfig.String("TercerosService")+"tercero/"+fmt.Sprint(dataInscripcion["PersonaId"]), &tercero)
	logs.Info("datos del tercero")
	fmt.Println(tercero)
	fmt.Println("")
	if errPersona == nil {
		bodyEmail := map[string]interface{}{
			"Html": map[string]interface{}{
				"Data": "<p>Estimado(a) aspirante:</p> <p> Cordial saludo. </p> <p> Le informamos que se ha registrado una novedad en su inscripción <strong>" + dataInscripcion["ReciboInscripcion"].(string) + "</strong>. Uno o varios de los documentos cargados requieren revisión, actualización o corrección. </p> <p> Por favor, ingrese a la plataforma <strong>SGAv2</strong> y consulte las observaciones registradas con el fin de realizar los ajustes solicitados dentro de los plazos establecidos. </p> <p> Agradecemos su atención y pronta gestión. </p> <br> <p>Atentamente, <br> <strong>Universidad Distrital Francisco José de Caldas</strong> <br> Sistema de Gestión Académica - SGAv2 </p>",
			},
			"Text": map[string]interface{}{
				"Data": "Novedad en Inscripción solicitada",
			},
		}
		logs.Info("Cuerpo del correo")
		fmt.Println(bodyEmail)
		fmt.Println(tercero.UsuarioWSO2)
		fmt.Println("")
		utils.SendNotificacionCambioEstadoSolicitud(bodyEmail, tercero.UsuarioWSO2)
	}
}
