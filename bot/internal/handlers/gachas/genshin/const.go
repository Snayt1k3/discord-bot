package genshin

var (
	charactersEmojis = map[string]string{}
	weaponsEmojis    = map[string]string{}
	artifactsEmojis  = map[string]string{}
	talentsEmojis    = map[string]string{}
	ascensionEmojis  = map[string]string{}
	images           = map[string]string{}
)


func GetCharacterEmoji(name string) string {
	if emoji, ok := charactersEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetWeaponEmoji(name string) string {
	if emoji, ok := weaponsEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetArtifactEmoji(name string) string {
	if emoji, ok := artifactsEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetTalentEmoji(name string) string {
	if emoji, ok := talentsEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetAscensionEmoji(name string) string {
	if emoji, ok := ascensionEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetImage(name string) string {
	if image, ok := images[name]; ok {
		return image
	}
	return ""
}