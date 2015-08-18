package thumbor

type Builder struct {
	Server   string
	Secret   string
	Image    string
	Commands AppliedCommands
}

type AppliedCommands struct {
	Trim      string
	Crop      string
	Resize    string
	HAlign    string
	VAlign    string
	SmartCrop bool
	Filters   []string
	Meta      bool
}

func NewBuilder(server, secret string) Builder {
	return Builder{
		Server: server,
		Secret: secret,
	}
}
