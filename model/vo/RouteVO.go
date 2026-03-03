package vo

type RouteVO struct {
	Path      string    `json:"path"`
	Component string    `json:"component"`
	Redirect  string    `json:"redirect"`
	Name      string    `json:"name"`
	Meta      Meta      `json:"meta"`
	Children  []RouteVO `json:"children,omitempty"`
}

type Meta struct {
	Title     string   `json:"title"`
	Icon      string   `json:"icon"`
	Hidden    bool     `json:"hidden"`
	Roles     []string `json:"roles"`
	KeepAlive bool     `json:"keepAlive"`
}
