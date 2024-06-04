package commands

import (
	"bot/commands/echo"
	"bot/commands/new_game"

	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	&echo.Command,
	&new_game.Command,
}
