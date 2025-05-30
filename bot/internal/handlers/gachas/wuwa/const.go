package wuwa

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

// todo: Remove this to config files
// This is a temporary solution
var (
	charactersEmojis = map[string]string{
		"Aalto":           "<:Aalto_Icon:1373955568246591489>",
		"Baizhi":          "<:Baizhi_Icon:1373958674997051393>",
		"Brant":           "<:Brant_Icon:1373958785919356998>",
		"Calcharo":        "<:Calcharo_Icon:1373958838377512970>",
		"Camellya":        "<:Camellya_Icon:1373958906195476531>",
		"Cantarella":      "<:Cantarella_Icon:1373958959462879343>",
		"Carlotta":        "<:Carlotta_Icon:1373959046352339044>",
		"Changli":         "<:Changli_Icon:1373959104716083291>",
		"Chixia":          "<:Chixia_Icon:1373959166473011300>",
		"Danjin":          "<:Danjin_Icon:1373959219870564493>",
		"Encore":          "<:Encore_Icon:1373959311264583680>",
		"Jianxin":         "<:Jianxin_Icon:1373959361600295004>",
		"Jinhsi":          "<:Jinhsi_Icon:1373959477031866409>",
		"Jiyan":           "<:Jiyan_Icon:1373959589242077228>",
		"Lingyang":        "<:Lingyang_Icon:1373959815860064306>",
		"Lumi":            "<:Lumi_Icon:1373959884449775627>",
		"Mortefi":         "<:Mortefi_Icon:1373959947234054224>",
		"Phoebe":          "<:Phoebe_Icon:1373960000938184734>",
		"Roccia":          "<:Roccia_Icon:1373960095150637088>",
		"Rover (Aero)":    "<:Rover_Icon:1373960161928151121>",
		"Rover (Havoc)":   "<:Rover_Icon:1373960161928151121>",
		"Rover (Spectro)": "<:Rover_Icon:1373960161928151121>",
		"Sanhua":          "<:Sanhua_Icon:1373960330811801681>",
		"Shorekeeper":     "<:Shorekeeper_Icon:1373960401896603658>",
		"Taoqi":           "<:Taoqi_Icon:1373960483324821594>",
		"Verina":          "<:Verina_Icon:1373960556746244116>",
		"Xiangli Yao":     "<:Xiangli_Yao_Icon:1373960622382780426>",
		"Yangyang":        "<:Yangyang_Icon:1373960728913907712>",
		"Yinlin":          "<:Yinlin_Icon:1373960987983741039>",
		"Youhu":           "<:Youhu_Icon:1373961044526895135>",
		"Yuanwu":          "<:Yuanwu_Icon:1373961099988176978>",
		"Zani":            "<:Zani_Icon:1373961189293555792>",
		"Zhezhi":          "<:Zhezhi_Icon:1373961251859861504>",
		"Ciaccona":        "<:Ciaccona_Icon:1376852453445730314>",
	}
	weaponsEmojis = map[string]string{}
	echoesEmojis  = map[string]string{
		"Celestial Light":      "<:Icon_Celestial_Light:1374358901398442064>",
		"Empyrean Anthem":      "<:Icon_Empyrean_Anthem:1374359066096308295>",
		"Eternal Radiance":     "<:Icon_Eternal_Radiance:1374359157968339034>",
		"Freezing Frost":       "<:Icon_Freezing_Frost:1374359239405080696>",
		"Frosty Resolve":       "<:Icon_Frosty_Resolve:1374359326092689421>",
		"Gusts of Welkin":      "<:Icon_Gusts_of_Welkin:1374359312390033468>",
		"Lingering Tunes":      "<:Icon_Lingering_Tunes:1374359772211445861>",
		"Midnight Veil":        "<:Icon_Midnight_Veil:1374360030169792594>",
		"Molten Rift":          "<:Icon_Molten_Rift:1374360117985673217>",
		"Moonlit Clouds":       "<:Icon_Moonlit_Clouds:1374360194649292891>",
		"Rejuvenating Glow":    "<:Icon_Rejuvenating_Glow:1374360241080111104>",
		"Sierra Gale":          "<:Icon_Sierra_Gale:1374360290854178847>",
		"Havoc Eclipsee":       "<:Icon_Sunsinking_Eclipse:1374360371330023454>",
		"Tidebreaking Courage": "<:Icon_Tidebreaking_Courage:1374360440632643697>",
		"Void Thunder":         "<:Icon_Void_Thunder:1374360490364375040>",
	}
	images = map[string]map[string]string{
		"Aalto":           {"thumbnail": "https://i.pinimg.com/736x/1c/22/bd/1c22bdeb0a06374b23c10e6699ec8cc8.jpg", "footer": "https://i.pinimg.com/736x/73/3a/a5/733aa53fba9289c49c92b274e70c68fc.jpg"},
		"Baizhi":          {"thumbnail": "https://i.pinimg.com/736x/c5/cd/79/c5cd79f94649ea5c221d3d21c3ea0628.jpg", "footer": "https://i.pinimg.com/736x/c5/4f/57/c54f57ea7a2993a2d2510b0468bb49a7.jpg"},
		"Brant":           {"thumbnail": "https://i.pinimg.com/736x/04/33/9d/04339d0399f2fb8a5d2e855b4d6c7e2a.jpg", "footer": "https://i.pinimg.com/736x/05/e0/d2/05e0d2f8bb2577167ddaea46c56744ab.jpg"},
		"Calcharo":        {"thumbnail": "https://i.pinimg.com/736x/8e/ef/b1/8eefb189f5e41f776997b1f9dad71056.jpg", "footer": "https://i.pinimg.com/736x/5d/cb/aa/5dcbaa163ecc766a9f19c4637db76302.jpg"},
		"Camellya":        {"thumbnail": "https://i.pinimg.com/736x/12/45/70/12457010879f96102820f25e1fbe8f4a.jpg", "footer": "https://i.pinimg.com/736x/ab/b8/54/abb854c476a72c6da415d192f6874138.jpg"},
		"Cantarella":      {"thumbnail": "https://i.pinimg.com/736x/a7/82/0f/a7820fe06c3babd7b91933ec3686c976.jpg", "footer": "https://i.pinimg.com/736x/23/74/e5/2374e566a51ec172db79f6cdf203a3f2.jpg"},
		"Carlotta":        {"thumbnail": "https://i.pinimg.com/736x/fd/9d/e8/fd9de8445a73f49a65f409a31f66f795.jpg", "footer": "https://i.pinimg.com/736x/9d/bf/21/9dbf21235056e3937f468063a019e02b.jpg"},
		"Changli":         {"thumbnail": "https://i.pinimg.com/736x/e2/24/ec/e224ecc2bdb9ef6e6d1c8a69d2669ee4.jpg", "footer": "https://i.pinimg.com/736x/9c/4d/a3/9c4da3dbf159fc94681a7a4458012e35.jpg"},
		"Chixia":          {"thumbnail": "https://i.pinimg.com/736x/77/b3/d7/77b3d7c14912255c4c87001a2ac571d4.jpg", "footer": "https://cdna.artstation.com/p/assets/images/images/074/953/244/large/cisl-25resized.jpg?1713373941"},
		"Danjin":          {"thumbnail": "https://i.pinimg.com/736x/f5/45/38/f54538bb88f4069f9b1f00edf5a06ef0.jpg", "footer": "https://i.pinimg.com/736x/52/a7/15/52a715c680d7b94c7c46719755ce6a10.jpg"},
		"Encore":          {"thumbnail": "https://i.pinimg.com/736x/64/dd/61/64dd615e5d848e267ec8b566d66cc30e.jpg", "footer": "https://i.pinimg.com/736x/5c/5e/75/5c5e75694a78040e5db0ba31f68b5bfe.jpg"},
		"Jianxin":         {"thumbnail": "https://i.pinimg.com/736x/90/9e/b4/909eb41687278866949cc93d57af564e.jpg", "footer": "https://i.pinimg.com/736x/3e/4e/84/3e4e844130982121de90fbe03904ab8f.jpg"},
		"Jinhsi":          {"thumbnail": "https://i.pinimg.com/736x/e3/06/e9/e306e91ef8c48b1205911b03cbc898f5.jpg", "footer": "https://i.pinimg.com/736x/2f/94/db/2f94db97718fa75a6adb3b13010a41ef.jpg"},
		"Jiyan":           {"thumbnail": "https://i.pinimg.com/736x/38/5d/c8/385dc8865964dfe97269e3bce40a07f9.jpg", "footer": "https://i.pinimg.com/736x/de/3e/a2/de3ea2789f67a4c772c31116b83e09c8.jpg"},
		"Lingyang":        {"thumbnail": "https://i.pinimg.com/736x/4f/9a/91/4f9a91dd376b2b3a1ce66f13a68b971f.jpg", "footer": "https://i.pinimg.com/736x/de/93/a7/de93a70f2701319fc1da9d73e37c1b4d.jpg"},
		"Lumi":            {"thumbnail": "https://i.pinimg.com/736x/95/f5/2d/95f52d87db68138f6b1687a62e19d425.jpg", "footer": "https://i.pinimg.com/736x/f9/fd/c4/f9fdc496440bac6065c5c106d812f9b8.jpg"},
		"Mortefi":         {"thumbnail": "https://i.pinimg.com/736x/54/14/e6/5414e673ae85370b3924a7d0ef9e1ea5.jpg", "footer": "https://s1.zerochan.net/Mortefi.600.4183831.jpg"},
		"Phoebe":          {"thumbnail": "https://i.pinimg.com/736x/7f/dd/5f/7fdd5f8c57a954e4cd89a07b83cacc4c.jpg", "footer": "https://i.pinimg.com/736x/e1/6e/a5/e16ea5c4d45b65c33a85e1239310c56b.jpg"},
		"Roccia":          {"thumbnail": "https://i.pinimg.com/736x/d5/10/8c/d5108ccfc03db20696c52d0504edb879.jpg", "footer": "https://i.pinimg.com/736x/00/1d/bc/001dbcc9c3cedece7dfb1941c08fc840.jpg"},
		"Rover (Aero)":    {"thumbnail": "https://i.pinimg.com/736x/6f/bc/84/6fbc84279a56f77e411012a8ff0d7f9c.jpg", "footer": "https://i.pinimg.com/736x/2d/10/2b/2d102b4aca89758e5ebeba8f0dd2b521.jpg"},
		"Rover (Spectro)": {"thumbnail": "https://i.pinimg.com/736x/6f/bc/84/6fbc84279a56f77e411012a8ff0d7f9c.jpg", "footer": "https://i.pinimg.com/736x/2d/10/2b/2d102b4aca89758e5ebeba8f0dd2b521.jpg"},
		"Rover (Havoc)":   {"thumbnail": "https://i.pinimg.com/736x/6f/bc/84/6fbc84279a56f77e411012a8ff0d7f9c.jpg", "footer": "https://i.pinimg.com/736x/2d/10/2b/2d102b4aca89758e5ebeba8f0dd2b521.jpg"},
		"Sanhua":          {"thumbnail": "https://i.pinimg.com/736x/bc/c3/e5/bcc3e52ef8e0a0ee2181e31ea06d62c4.jpg", "footer": "https://i.pinimg.com/736x/df/ed/16/dfed1655ceba276ff03dbdea71d74054.jpg"},
		"Shorekeeper":     {"thumbnail": "https://i.pinimg.com/736x/75/d8/b6/75d8b64be5902e8b99505212b212d0a2.jpg", "footer": "https://i.pinimg.com/736x/7b/ec/38/7bec389ec20b6cd2ec795fa289a075b9.jpg"},
		"Taoqi":           {"thumbnail": "https://i.pinimg.com/736x/77/b7/63/77b763f6a89dd2046e2078969bf1309d.jpg", "footer": "https://image-2.uhdpaper.com/wallpaper/taoqi-wuthering-waves-hd-wallpaper-uhdpaper.com-451@2@b.jpg"},
		"Verina":          {"thumbnail": "https://i.pinimg.com/736x/05/75/66/057566d0cce74aa042d8e92cbfeae0cd.jpg", "footer": "https://i.pinimg.com/736x/d7/05/11/d705113da72d4d6923a8aa62e552da9d.jpg"},
		"Xiangli Yao":     {"thumbnail": "https://i.pinimg.com/736x/64/77/09/6477093da1a8e804088e402d7d07f7e4.jpg", "footer": "https://static.zerochan.net/Xiangli.Yao.full.4360186.jpg"},
		"Yinlin":          {"thumbnail": "https://i.pinimg.com/736x/9a/37/1d/9a371d5dc6db4a268bc8d736ec4cb866.jpg", "footer": "https://i.pinimg.com/736x/c3/18/6b/c3186bf837be7faa14ef93f308df8b25.jpg"},
		"Youhu":           {"thumbnail": "https://i.pinimg.com/736x/98/65/e9/9865e9fed2060077a89f860ceb8fc8f5.jpg", "footer": "https://i.pinimg.com/736x/15/8a/b0/158ab0b440654ad917e0c91c9f32a04e.jpg"},
		"Yuanwu":          {"thumbnail": "https://i.pinimg.com/736x/0c/d5/f5/0cd5f5bda8f380278a78e69df358a141.jpg", "footer": "https://i.pinimg.com/736x/3f/6f/ba/3f6fbad43616c8f788d26e959cdb5297.jpg"},
		"Zani":            {"thumbnail": "https://i.pinimg.com/736x/ed/43/fe/ed43fe9036039c8eaf48e7679667314b.jpg", "footer": "https://i.pinimg.com/736x/1f/f3/5a/1ff35a16a41b1d9dc13113de19c14fad.jpg"},
		"Zhezhi":          {"thumbnail": "https://i.pinimg.com/736x/4d/94/03/4d94030f5aa2b0f9ead157aec7b91088.jpg", "footer": "https://i.pinimg.com/736x/7f/a0/9b/7fa09bcf1e4e8aa28ac97fe42329f559.jpg"},
		"Ciaccona":        {"thumbnail": "https://i.pinimg.com/736x/b5/6c/1f/b56c1f51110c5f02cf623ad80ace4a65.jpg", "footer": "https://i.pinimg.com/736x/19/ee/44/19ee440e74948265657c58f17acc63f6.jpg"},
		"Yangyang":        {"thumbnail": "https://i.pinimg.com/736x/3f/e4/36/3fe436b93f1cc57842cd2c913c697388.jpg", "footer": "https://i.pinimg.com/736x/bf/6c/98/bf6c98addb821db442c04af7bcfecb46.jpg"},
		// "": {"thumbnail": "", "footer": "",},
	}
	weaponTypes = map[string]string{
		"Rectifier":  "<:Rectifier_Icon:1374351384018948116>",
		"Sword":      "<:Sword_Icon:1374351865546014750>",
		"Pistols":    "<:Pistols_Icon:1374351373038125056>",
		"Gauntlets":  "<:Gauntlets_Icon:1374351356806168706>",
		"Broadblade": "<:Broadblade_Icon:1374351340624543815>",
	}
	attributeEmojis = map[string]string{
		"Spectro": "<:Spectro_Icon:1373952798600728586>",
		"Havoc":   "<:Havoc_Icon:1373952790107390114>",
		"Glacio":  "<:Glacio_Icon:1373952780502433844>",
		"Fusion":  "<:Fusion_Icon:1373952766267097170>",
		"Electro": "<:Electro_Icon:1373952758708834335>",
		"Aero":    "<:Aero_Icon:1373952741721899089>",
	}
)

func GetWeaponTypeEmoji(name string) string {
	if emoji, ok := weaponTypes[name]; ok {
		return emoji
	}
	return ""
}

func GetAttributeEmoji(name string) string {
	if emoji, ok := attributeEmojis[name]; ok {
		return emoji
	}
	return ""
}

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

func GetEchoEmoji(name string) string {
	if emoji, ok := echoesEmojis[name]; ok {
		return emoji
	}
	return ""
}

func GetImage(name string) (string, string) {
	if image, ok := images[name]; ok {
		return image["thumbnail"], image["footer"]
	}
	return "", ""
}

func ParseCustomEmoji(emojiStr string) *discordgo.ComponentEmoji {
	re := regexp.MustCompile(`<a?:([a-zA-Z0-9_]+):(\d+)>`)
	matches := re.FindStringSubmatch(emojiStr)

	if len(matches) != 3 {
		return &discordgo.ComponentEmoji{Name: emojiStr}
	}

	return &discordgo.ComponentEmoji{
		Name:     matches[1],
		ID:       matches[2],
		Animated: strings.HasPrefix(emojiStr, "<a:"),
	}
}
