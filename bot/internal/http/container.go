package http

type Container struct {
	Log         *Log
	Settings    *Settings
	User 		*UserService
	Roles       *Roles
	Welcome     *Welcome
	Moderation  *AutoMode
}

func NewContainer() *Container {
	return &Container{
		Log:         NewLog(),
		Settings:    NewSettings(),
		User: 		 NewUserService(),
		Roles:       NewRoles(),
		Welcome:     NewWelcome(),
		Moderation:  NewAutoMode(),
	}
}
