package thumbor

import (
	"fmt"
	"strings"
)

func (c *AppliedCommands) String() string {
	var commands []string

	appendCommand := func(command string) {
		if len(command) > 0 {
			commands = append(commands, command)
		}
	}

	if c.Meta {
		commands = append(commands, "meta")
	}

	appendCommand(c.Trim)
	appendCommand(c.Crop)
	appendCommand(c.Resize)
	appendCommand(c.HAlign)
	appendCommand(c.VAlign)

	if c.SmartCrop {
		commands = append(commands, "smart")
	}

	if len(c.Filters) > 0 {
		commands = append(commands, fmt.Sprintf("filters:%s", strings.Join(c.Filters, ":")))
	}

	return strings.Join(commands, "/")
}

func (b *Builder) String() string {
	url := fmt.Sprintf("%s/%s/", b.Server, b.Secret)

	if commands := b.Commands.String(); len(commands) != 0 {
		url += commands + "/"
	}

	url += b.Image

	return url
}
