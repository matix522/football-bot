package new_game

import "github.com/bwmarrin/discordgo"

var Command = discordgo.ApplicationCommand{
	Name:        "new_game",
	Description: "Create new game",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "game_time",
			Description: "Time of the game; Date and time in dd-MM-yyTHH:mm format",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "cancel_deadline",
			Description: "Last chance to cancel the game. Date and time in dd-MM-yyTHH:mm format",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "rsvp_before",
			Description: "Number of days before the event is scheduled, to ping members for rsvp",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Required:    true,
		},
	},
}
