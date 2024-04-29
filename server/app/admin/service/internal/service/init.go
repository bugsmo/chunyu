//go:build wireinject
// +build wireinject

package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthenticationService,
	NewUserService,
	NewRoleService,
	NewPositionService,
	NewOrganizationService,
	NewMenuService,
	NewDictService,
	NewDictDetailService,
	NewRouterService,
	NewTaskService,
)
