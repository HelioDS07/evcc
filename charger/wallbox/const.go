package wallbox

type Status int

const (
	WAITING Status = iota
	CHARGING
	READY
	PAUSED
	SCHEDULED
	DISCHARGING
	ERROR
	DISCONNECTED
	LOCKED
	UPDATING
)

const (
	ActionResume = 1
	ActionPause  = 2
)