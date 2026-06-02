package valueobject

type PoliticaGeofence string

var (
	APENAS_DENTRO_RAIO PoliticaGeofence = "APENAS_DENTRO_RAIO"
	QUALQUER_LUGAR_TAG PoliticaGeofence = "QUALQUER_LUGAR_TAG"
	QUALQUER_LUGAR PoliticaGeofence = "QUALQUER_LUGAR"
)