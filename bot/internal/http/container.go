package http

type Container struct {
	Log         *Log
	Settings    *Settings
	Interaction *Interaction
	Roles       *Roles
	Welcome     *Welcome
	Moderation  *AutoMode
}

func NewContainer() *Container {
	return &Container{
		Log:         NewLog(),
		Settings:    NewSettings(),
		Interaction: NewInteraction(),
		Welcome:     NewWelcome(),
		Moderation:  NewAutoMode(),
		Roles:       NewRoles(),
	}
}
