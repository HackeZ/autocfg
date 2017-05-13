package autocfg

type Config struct {
	ENV    string `json:"env"`
	Module string `json:"module"`
}

// DependModule
type DependModule struct {
	Name    string
	Address string
	Timeout int
}
