package conf

var (
	// RPCSvrListener rpc server listener
	RPCSvrListener = struct {
		NetWork string
		Address string
	}{"unix", "rpc.sock"}

	// RPCSvrBinPath rpc server bin path
	RPCSvrBinPath = "./bin/task-rpc"
)
