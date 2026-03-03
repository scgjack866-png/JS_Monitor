package menu

import (
	"OperationAndMonitoring/model/entity"
	"OperationAndMonitoring/model/vo"
	db "OperationAndMonitoring/utils"
)

func RecurRoutes(parentID uint64, menus []entity.Menu) []vo.RouteVO {
	var list []vo.RouteVO
	for _, menu := range menus {
		if menu.MenuType == 4 {
			continue
		}
		if parentID == menu.ParentId {
			var routeVO vo.RouteVO
			var meta vo.Meta
			if menu.MenuType == 1 {
				routeVO.Name = menu.Path
			}
			routeVO.Path = menu.Path
			routeVO.Redirect = menu.Redirect
			routeVO.Component = menu.Component
			meta.Title = menu.Name
			meta.Icon = menu.Icon
			meta.Roles = GetRoleCodeFormMenu(menu.ID)
			meta.Hidden = false == menu.Visible
			routeVO.Meta = meta
			routeVO.Children = RecurRoutes(menu.ID, menus)
			list = append(list, routeVO)
		}
	}
	return list
}

func GetRoleCodeFormMenu(menuID uint64) []string {
	var roles []entity.Role
	var rolesCode []string
	d := entity.Menu{
		ID: menuID,
	}

	db.Test(&d, &roles, "Menus", "Roles")

	for _, role := range roles {
		rolesCode = append(rolesCode, role.Code)
		return rolesCode
	}
	return rolesCode
}

func GetRoles(userIDU uint64) []string {
	var roles []entity.Role
	var rolesName []string
	d := entity.User{
		ID: userIDU,
	}

	db.Test(&d, &roles, "Users", "Roles")

	for _, role := range roles {
		rolesName = append(rolesName, role.Name)
		return rolesName
	}
	return rolesName
}
