package main

import (
	"bot/commands"
	"bot/db"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om optionMap) {
	om = make(optionMap)
	for _, opt := range options {
		om[opt.Name] = opt
		fmt.Println("%s", opt)
	}
	return
}

func interactionAuthor(i *discordgo.Interaction) *discordgo.User {
	if i.Member != nil {
		return i.Member.User
	}
	return i.User
}

func handleEcho(s *discordgo.Session, i *discordgo.InteractionCreate, opts optionMap) {
	fmt.Println("%s", opts["author"])

	builder := new(strings.Builder)
	if v, ok := opts["author"]; ok && v.BoolValue() {
		author := interactionAuthor(i.Interaction)
		builder.WriteString("**" + author.String() + "** says: ")
	}
	builder.WriteString(opts["message"].StringValue())

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})

	if err != nil {
		log.Panicf("could not respond to interaction: %s", err)
	}
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	session, err := discordgo.New("Bot " + os.Getenv("DISORD_BOT_AUTH_TOKEN"))

	if err != nil {
		log.Fatal("Error creating Discord session %s", err)
		return
	}

	_, err = sql.Open("postgres", db.ConnectionString())

	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()
		if data.Name != "echo" {
			return
		}

		handleEcho(s, i, parseOptions(data.Options))
	})
	_, err = session.ApplicationCommandBulkOverwrite("1244611273576939594", "837325346160246795", commands.Commands)

	if err != nil {
		log.Fatalf("could not register commands: %s", err)
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	err = session.Open()
	if err != nil {
		log.Fatalf("could not open session: %s", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()
	if err != nil {
		log.Printf("could not close session gracefully: %s", err)
	}
	log.Fatal("Discord connected")

}
