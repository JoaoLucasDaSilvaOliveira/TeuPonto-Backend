package valueobject

type GeofencePolicy string

var (
	JUST_INSIDE_RADIUS GeofencePolicy = "JUST_INSIDE_RADIUS"
	ANYWHERE_USING_TAG GeofencePolicy = "ANYWHERE_USING_TAG"
	ANYWHERE GeofencePolicy = "ANYWHERE"
)