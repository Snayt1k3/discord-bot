package handlers

import (
	"bot/internal/discord"
	"fmt"
	"math/rand"
	"time"
	"github.com/bwmarrin/discordgo"
)

var (
	// Temporary solution
	channelId       = ""
	messageTemplate = []string{
		"Еще один путешественник прибыл… %v, время на этом сервере течет иначе, чем в мире снаружи. Надеюсь, твое пребывание здесь будет долгим и спокойным.", 
		"Ты присоединился, %v… Люди говорят, что время летит незаметно в хорошей компании. Возможно, ты тоже это почувствуешь.",
		"Еще один человек… %v, люди приходят и уходят, но, возможно, ты задержишься здесь дольше, чем кажется.",
		"Когда-то здесь уже были другие… Но каждое новое знакомство — это возможность узнать что-то новое. Добро пожаловать, %v.",
		"Добро пожаловать, %v. В этом месте нет конца пути, только новые встречи. Оставайся, если хочешь.",
	}
)	


func OnNewMemberJoin(s *discordgo.Session, u *discordgo.GuildMemberAdd) {
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	randomIndex := randGen.Intn(len(messageTemplate))
	formattedMessage := fmt.Sprintf(messageTemplate[randomIndex], u.Member.Nick)

	discord.SendChannelMessage(channelId, formattedMessage)
}
