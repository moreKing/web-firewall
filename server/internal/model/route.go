package model

type Meta struct {
	Title      string `json:"title,omitempty"`
	I18NKey    string `json:"i18nKey,omitempty"`
	Icon       string `json:"icon,omitempty"`
	Order      int    `json:"order,omitempty"`
	HideInMenu bool   `json:"hideInMenu,omitempty"`
	MultiTab   bool   `json:"multiTab,omitempty"`
	Constant   bool   `json:"constant,omitempty"`
	ActiveMenu string `json:"activeMenu,omitempty"`
}

type Route struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	Component string   `json:"component"`
	Meta      Meta     `json:"meta,omitempty"`
	Children  *[]Route `json:"children,omitempty"`
}
