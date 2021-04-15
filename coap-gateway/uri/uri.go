package uri

//URI
const (
	Base = "/ola/v1/device"

	//接入认证使用接口
	Auth     = Base + "/auth"
	Certify  = Base + "/certify"
	Register = Base + "/register"

	//模型接口
	Services    = Base + "/services"
	Properities = Base + "/properities"
	Action      = Base + "/actions"
	Event       = Base + "/event"

	ProperitiesControl   = Properities + "/control"
	ProperitiesArray     = Properities + "/array"
	ProperitiesSubscribe = Properities + "/subscribe"
	ProperitiesNotify    = Properities + "/notify"
	EventSubscribe       = Event + "/subscribe"
)
