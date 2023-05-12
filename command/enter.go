package command

import "blog/global"

func Handle(option []string) {
	command := option[0]
	global.Logger.Info("1111")
	switch command {
	case "migration":
		Migrations()
	}
}
