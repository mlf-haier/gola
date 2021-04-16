package uri

//URI
const (
	Base = "/ola/v1/device"

	//Wi-Fi设备接入认证使用
	Auth     = Base + "/auth"
	Certify  = Base + "/certify"
	Register = Base + "/register"

	//通用接入认证使用
	Activate     = Base + "/activate"
	Login        = Base + "/login"
	Heartbeat    = Base + "/heartbeat"
	TokenRefresh = Base + "/tokenRefresh"
	Logout       = Base + "/logout"
	Reset        = Base + "/reset"

	//模型使用
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
