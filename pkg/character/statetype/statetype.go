package statetype

type Type string

const (
	Idle      Type = "Idle"
	Walking   Type = "Walking"
	Attacking Type = "Attacking"
	StandBy   Type = "StandBy"
)
