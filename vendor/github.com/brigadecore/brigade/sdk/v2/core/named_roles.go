package core

import (
	libAuthz "github.com/brigadecore/brigade/sdk/v2/lib/authz"
	"github.com/brigadecore/brigade/sdk/v2/system"
)

const (
	// Core-specific, system-level roles...

	// RoleNameEventCreator is the name of a system-level Role that enables
	// principals to create Events for all Projects-- provided the Events have a
	// specific value in the Source field. This is useful for Event gateways,
	// which should be able to create Events for all Projects, but should NOT be
	// able to impersonate other gateways.
	RoleNameEventCreator libAuthz.RoleName = "EVENT_CREATOR"

	// RoleNameProjectCreator is the name of a system-level Role that enables
	// principals to create new Projects.
	RoleNameProjectCreator libAuthz.RoleName = "PROJECT_CREATOR"

	// Core-specific, project-level roles...

	// RoleNameProjectAdmin is the name of a project-level Role that enables a
	// principal to manage a specific Project.
	RoleNameProjectAdmin libAuthz.RoleName = "ADMIN"

	// RoleNameProjectDeveloper is the name of a project-level Role that enables a
	// principal to update a specific project.
	RoleNameProjectDeveloper libAuthz.RoleName = "DEVELOPER"

	// RoleNameProjectUser is the name of project-level Role that enables a
	// principal to create and manage Events for a specific Project.
	RoleNameProjectUser libAuthz.RoleName = "USER"
)

// Core-specific, system-level roles...

// RoleEventCreator returns a system-level Role that enables principals to
// create Events for all Projects-- provided the Events have a value in the
// Source field that matches the value in this Role's Scope field. This is
// useful for Event gateways, which should be able to create Events for all
// Projects, but should NOT be able to impersonate other gateways.
func RoleEventCreator(eventSource string) libAuthz.Role {
	return libAuthz.Role{
		Type:  system.RoleTypeSystem,
		Name:  RoleNameEventCreator,
		Scope: eventSource,
	}
}

// RoleProjectCreator returns a system-level Role that enables principals to
// create new Projects.
func RoleProjectCreator() libAuthz.Role {
	return libAuthz.Role{
		Type: system.RoleTypeSystem,
		Name: RoleNameProjectCreator,
	}
}

// Core-specific, project-level roles...

// RoleProjectAdmin returns a project-level Role that enables a principal to
// manage the Project whose ID matches the value of the Scope field. If the
// value of the Scope field is RoleScopeGlobal ("*"), then the Role is unbounded
// and enables a principal to manage all Projects.
func RoleProjectAdmin(projectID string) libAuthz.Role {
	return libAuthz.Role{
		Type:  RoleTypeProject,
		Name:  RoleNameProjectAdmin,
		Scope: projectID,
	}
}

// RoleProjectDeveloper returns a project-level Role that enables a principal to
// update the Project whose ID matches the value of the Scope field. If the
// value of the Scope field is RoleScopeGlobal ("*"), then the Role is unbounded
// and enables a principal to update all Projects.
func RoleProjectDeveloper(projectID string) libAuthz.Role {
	return libAuthz.Role{
		Type:  RoleTypeProject,
		Name:  RoleNameProjectDeveloper,
		Scope: projectID,
	}
}

// RoleProjectUser returns a project-level Role that enables a principal to
// create and manage Events for the Project whose ID matches the value of the
// Scope field. If the value of the Scope field is RoleScopeGlobal ("*"), then
// the Role is unbounded and enables a principal to create and manage Events for
// all Projects.
func RoleProjectUser(projectID string) libAuthz.Role {
	return libAuthz.Role{
		Type:  RoleTypeProject,
		Name:  RoleNameProjectUser,
		Scope: projectID,
	}
}
