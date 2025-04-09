package genshin

import (
	"bot/internal/discord"
	dtoDiscord "bot/internal/dto/discord"

	"github.com/bwmarrin/discordgo"
)

func showCharacterAscension(data dtoDiscord.HandlerData) error {
	embed := &discordgo.MessageEmbed{
		Title:       "⬆️ Ascension Materials — Raiden Shogun",
		Description: "Materials required to fully ascend Raiden Shogun to Lv. 90, including Talent Level-Up materials.",
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🟣 Vajrada Amethyst",
				Value:  "- Sliver ×1\n- Fragment ×9\n- Chunk ×9\n- Gemstone ×6",
				Inline: true,
			},
			{
				Name:   "⚡ Storm Beads",
				Value:  "- Total: ×46\n(Dropped by **Thunder Manifestation**)",
				Inline: true,
			},
			{
				Name:   "🍇 Amakumo Fruit",
				Value:  "- Total: ×168\n(Found on **Seirai Island**)",
				Inline: true,
			},
			{
				Name:   "🗡️ Handguards",
				Value:  "- Old ×18\n- Kageuchi ×30\n- Famed ×36\n(Dropped by Nobushi)",
				Inline: true,
			},
			{
				Name:   "💰 Mora",
				Value:  "- Total: 420,000",
				Inline: true,
			},
			{
				Name:  "� Ascension Levels",
				Value: "20 ➜ 40 ➜ 50 ➜ 60 ➜ 70 ➜ 80 ➜ 90",
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Full Ascension & Talent Materials",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}

func showCharacterArtifacts(data dtoDiscord.HandlerData) error {
	embed := &discordgo.MessageEmbed{
		Title:       "🛡️ Raiden Shogun — Artifact Guide",
		Description: "Top artifact sets for different Raiden Shogun builds.\nChoose based on your team and role preferences.",
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "⚡ Burst DPS (Main Build)",
				Value: "**4x Emblem of Severed Fate**\n" +
					"> Core meta build — boosts Burst DMG based on Energy Recharge.\n" +
					"**Main Stats:** ER / Electro DMG / Crit\n" +
					"**Substats:** Crit Rate > Crit DMG > ER > ATK%",
			},
			{
				Name: "💠 Mixed Support",
				Value: "**2x Emblem + 2x Noblesse Oblige**\n" +
					"> Balanced hybrid for Burst + team support.\n" +
					"**Main Stats:** ER / Electro DMG / Crit or ATK%\n" +
					"**Substats:** Crit > ER > ATK%",
			},
			{
				Name: "🔋 Energy Battery (Team Utility)",
				Value: "**4x The Exile** *(low AR option)*\n" +
					"> Great in early game, boosts Energy for team.\n" +
					"**Main Stats:** ER / ATK% / Crit Rate\n" +
					"**Substats:** ER > ATK% > Crit Rate",
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Artifact Sets Overview",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil

}

func showCharacterWeapons(data dtoDiscord.HandlerData) error {
	embed := &discordgo.MessageEmbed{
		Title:       "⚔️ Raiden Shogun — Weapon Guide",
		Description: "Recommended weapons for Raiden Shogun depending on your build and availability.",
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "💜 **Engulfing Lightning** (5★)",
				Value: "Signature weapon.\n- Huge Energy Recharge and ATK boost synergy.\n- Best-in-slot for Burst DPS build.",
			},
			{
				Name:  "🌀 **The Catch** (4★)",
				Value: "Best free-to-play option.\n- Boosts Elemental Burst DMG & CRIT Rate.\n- Pairs perfectly with Emblem set.",
			},
			{
				Name:  "🔱 **Staff of Homa / Skyward Spine** (5★)",
				Value: "- Homa: CRIT DMG and HP boost — solid stat stick.\n- Skyward: High ER and some CRIT Rate.",
			},
			{
				Name:  "🔷 **Wavebreaker's Fin / Favonius Lance** (4★)",
				Value: "- Wavebreaker: High Burst DMG scaling.\n- Favonius: Utility polearm with team energy gen.",
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Weapon Recommendations",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}

func showCharacterComps(data dtoDiscord.HandlerData) error {
	embed := &discordgo.MessageEmbed{
		Title:       "👥 Raiden Shogun — Best Team Compositions",
		Description: "Raiden Shogun works best in teams that focus on energy generation and maximizing Electro DMG. Below are the optimal team compositions.",
		Color:       0x9b59b6,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "⚡ **Raiden Shogun + Beidou + Xingqiu + Bennett**",
				Value: "A powerful Electro-focused team with **Beidou** and **Xingqiu** providing additional damage reduction and elemental reaction support.\n- Raiden boosts the team's overall damage through Energy Recharge.",
			},
			{
				Name:  "💥 **Raiden Shogun + Yae Miko + Fischl + Kazuha**",
				Value: "Raiden Shogun shines in this team by triggering Electro-related reactions. **Kazuha** provides grouping and elemental buffs while **Yae Miko** and **Fischl** provide continuous Electro application.",
			},
			{
				Name:  "💫 **Raiden Shogun + Zhongli + Ganyu + Albedo**",
				Value: "A versatile team with great synergy. **Zhongli** provides shields, while **Ganyu** and **Albedo** focus on strong Cryo and Geo DPS. This setup helps Raiden keep her Elemental Burst uptime high.",
			},
			{
				Name:  "🌪️ **Raiden Shogun + Kokomi + Childe + Sucrose**",
				Value: "This team works by using **Kokomi** and **Childe** to apply Hydro, which triggers powerful Electro-Charged reactions with Raiden's attacks. **Sucrose** boosts elemental mastery for higher reaction damage.",
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/d2/96/83/d29683ce9223109447fb6a57ef9f7e3a.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Team Composition Guide",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}

func showCharacterTalents(data dtoDiscord.HandlerData) error {
	embed := &discordgo.MessageEmbed{
		Title:       "📘 Talent Materials — Raiden Shogun",
		Description: "Resources required to level up all three of Raiden Shogun's talents to Lv. 10.",
		Color:       0xad44d9,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://i.pinimg.com/736x/77/97/d7/7797d737a3a35630f6ce321b1a00fc20.jpg",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📚 Talent Books — Light Series",
				Value:  "- Teachings of Light ×9\n- Guide to Light ×63\n- Philosophies of Light ×114",
				Inline: true,
			},
			{
				Name:   "🗡️ Handguards",
				Value:  "- Old ×18\n- Kageuchi ×66\n- Famed ×93\n(Dropped by Nobushi)",
				Inline: true,
			},
			{
				Name:   "🔥 Weekly Boss Material",
				Value:  "- Molten Moment ×18\n(Dropped by **La Signora**)",
				Inline: true,
			},
			{
				Name:   "👑 Crown of Insight",
				Value:  "- Total: ×3 (for maxing all 3 talents)",
				Inline: true,
			},
			{
				Name:   "💰 Mora",
				Value:  "- Total: 4,950,000",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.pinimg.com/736x/0b/18/e8/0b18e8acbf645b7b227689f33785d5c3.jpg",
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Raiden Shogun • Talent Level-Up Costs",
		},
	}

	components := genshinButtons("shogun")

	data.Session.InteractionRespond(data.Event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	discord.EditMessage(data.Session, &discordgo.MessageEdit{
		Embeds:     &[]*discordgo.MessageEmbed{embed},
		Components: &components,
		Channel:    data.Event.ChannelID,
		ID:         data.Event.Message.ID,
	})

	return nil
}
