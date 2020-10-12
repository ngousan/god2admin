package authUserMenus

type AuthUserMenus struct {
	Title    string          `json:"title"`
	Icon     string          `json:"icon"`
	Path     string          `json:"path"`
	MenuId   int             `json:"menuId"`
	ParentId int             `json:"parentId"`
	Children []AuthUserMenus `json:"children" gorm:"-"`
}
