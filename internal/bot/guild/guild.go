package guild

import (
	"os"
)

const GUILD_GLOBAL string = ""
var GUILD_DEV string = os.Getenv("R2SRV")

func GetGuild(devmode bool) string {
	if (devmode) {
		return GUILD_DEV
	}

	return GUILD_GLOBAL
}
