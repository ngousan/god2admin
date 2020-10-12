package authUserRouters

type AuthUserRouters struct {
	Name      string                  `json:"name"`
	Path      string                  `json:"path"`
	Component string                  `json:"component"`
	Redirect  AuthUserRoutersRedirect `json:"redirect"`
	Meta      AuthUserRoutersMeta     `json:"meta"`
	Hidden    bool                    `json:"hidden"`
	Group     string                  `json:"group"`
	Children  []AuthUserRouters       `json:"children"`
}

type AuthUserRoutersMeta struct {
	Title string `json:"title"`
	Auth  bool   `json:"auth"`
	Cache bool   `json:"cache"`
}

type AuthUserRoutersRedirect struct {
	Name string `json:"name"`
}

type AuthUserRoutersView struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Title     string `json:"title"`
	Auth      bool   `json:"auth"`
	Cache     bool   `json:"cache"`
	Hidden    bool   `json:"hidden"`
	Group     string `json:"group"`
}
