package thumbor

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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

func generateSignature(args, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(args))
	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

func (b *Builder) String() string {
	var path string
	if path = b.Commands.String(); len(path) > 0 {
		path += "/"
	}
	path += b.Image
	signature := "unsafe"
	if b.Secret != "" {
		signature = generateSignature(path, b.Secret)
	}
	return fmt.Sprintf("%s/%s/%s", b.Server, signature, path)
}
