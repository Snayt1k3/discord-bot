package handlers

import (
	"api-gateway/internal/interfaces"
	pb "api-gateway/proto"
	"google.golang.org/grpc"
)

// Clients holds gRPC clients for different services
type Clients struct {
	Settings pb.SettingsServiceClient
	Welcome  pb.WelcomeServiceClient
	Log      pb.LogServiceClient
	AutoMode pb.AutoModServiceClient
	Roles    pb.RolesServiceClient
}

// NewClients initializes and returns a Clients struct with gRPC clients
func NewClients(cc grpc.ClientConnInterface) *Clients {
	return &Clients{
		Settings: pb.NewSettingsServiceClient(cc),
		Welcome:  pb.NewWelcomeServiceClient(cc),
		Log:      pb.NewLogServiceClient(cc),
		AutoMode: pb.NewAutoModServiceClient(cc),
		Roles:    pb.NewRolesServiceClient(cc),
	}
}

// Holds all handler groups
type Handlers struct {
	Settings *SettingsHandlers
	Roles    *RolesHandlers
	Welcome  *WelcomeHandlers
	Automode *AutoModeHandlers
	Log      *LogHandlers
}

// NewHandlers initializes and returns a Handlers struct with all handler groups
func NewHandlers(cc grpc.ClientConnInterface, cache interfaces.RedisInterface) *Handlers {
	return &Handlers{
		Settings: NewSettingsHandlers(cc, cache),
		Roles:    NewRolesHandlers(cc),
		Welcome:  NewWelcomeHandlers(cc),
		Automode: NewAutoModeHandlers(cc),
		Log:      NewLogHandlers(cc),
	}
}
