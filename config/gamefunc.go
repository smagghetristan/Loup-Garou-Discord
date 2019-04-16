package config

type CannotFindError interface {
	Error() string
}

func (p Player) Infect() {
	p.Infected = true
}

func GetRoleByName(RoleName string) (Role, bool) {
	for i := range AllRoles {
		if AllRoles[i].Name == RoleName {
			return AllRoles[i], false
		}
	}
	return Role{}, true
}
