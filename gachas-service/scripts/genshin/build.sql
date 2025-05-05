INSERT INTO builds (
    character_id,
    stats_id,
    name,
    created_at,
    updated_at,
) VALUES

-- Geo Characters Builds
(15, 1, "Main Dps", NOW(), NOW()), -- Noelle
(14, 2, "Main Dps", NOW(), NOW()), -- Ningguang
(25, 3, "Shield Support", NOW(), NOW()), -- Zhongli
(26, 4, "OFF-FIELD DPS", NOW(), NOW()), -- Albedo
(43, 5, "Support", NOW(), NOW()), -- Gorou
(42, 6, "Main Dps", NOW(), NOW()), -- Itto
(45, 7, "Support", NOW(), NOW()), -- Yun Jin
(76, 8, "Main Dps", NOW(), NOW()), -- Navia
(79, 9, "Sub Dps", NOW(), NOW()), -- Chiori
(85, 10, "SUPPORT AND OFF-FIELD DPS", NOW(), NOW()), -- Kachina
(88, 11, "Support", NOW(), NOW()), -- Xilonen

-- Dendro Characters builds
(53, 12, "QUICK SWAP DPS", NOW(), NOW()), -- Tighnari
(51, 13, "Support", NOW(), NOW()), -- Collei
(58, 14, "BUFF SUPPORT & DAMAGE", NOW(), NOW()), -- Nahida
(62, 15, "HEAL SUPPORT", NOW(), NOW()), -- Yaoyao
(61, 16, "SPREAD DPS", NOW(), NOW()), -- Alhaitam
(66, 17, "BLOOM / BURGEON DRIVER", NOW(), NOW()) -- Kaveh
(65, 18, "SUPPORT", NOW(), NOW()) -- Baizhu
(67, 19, "Support", NOW(), NOW()) -- Kirara
(84, 20, "OFF-FIELD DPS", NOW(), NOW()) -- Emilie
(85, 21, "DPS", NOW(), NOW()) -- Kinich

-- Cryo Characters builds
(16, 22, "SUPPORT", NOW(), NOW()) -- Qiqi
(9, 23, "DPS", NOW(), NOW()) -- Kaeya
(5, 24, "SUPPORT", NOW(), NOW()) -- Chongyun
(22, 25, "SUPPORT", NOW(), NOW()) -- Diona
(27, 26, "DPS", NOW(), NOW()), -- Ganyu
(30, 27, "SUPPORT", NOW(), NOW()), -- Rosaria
(31, 28, "DPS", NOW(), NOW()), -- Eula
(34, 29, "DPS", NOW(), NOW()), -- Ayaka
(37, 30, "BURST DPS", NOW(), NOW()), -- Aloy
(44, 31, "BUFF SUPPORT", NOW(), NOW()), -- Shenhe
(57, 32, "SUPPORT", NOW(), NOW()), -- Layla
(64, 33, "SUPPORT", NOW(), NOW()), -- Mika
(68, 34, "PHYSICAL DPS", NOW(), NOW()), -- Freminet
(72, 35, "MELT DPS", NOW(), NOW()), -- Wriothesley
(73, 36, "HEAL SUPPORT", NOW(), NOW()), -- Charlotte
(91, 37, "BUFF SUPPORT", NOW(), NOW()), -- Citlali

-- Pyro characters builds
(20, 38, "OFF-FIELD VAPORIZE DPS", NOW(), NOW()), -- Xiangling
(11, 39, "DPS", NOW(), NOW()), -- Klee
(6, 40, "DPS", NOW(), NOW()), -- Diluc
(4, 41, "SUPPORT", NOW(), NOW()), -- Bennett
(1, 42, "BUFF", NOW(), NOW()), -- Amber
(24, 43, "SHIELD SUPPORT", NOW(), NOW()) -- Xinyan
(28, 44, "DPS", NOW(), NOW()) -- Hu tao
(32, 45, "DPS",  NOW(), NOW()) -- Yanfei
(36, 46, "DPS", NOW(), NOW()), -- Yoimiya
(41, 47, "SHIELD SUPPORT", NOW(), NOW()), -- Thoma
(63, 48, "DPS", NOW(), NOW()), -- Dehya
(70, 49, "DPS", NOW(), NOW()), -- Lyney
(75, 50, "BUFF", NOW(), NOW()), --Chevreuse
(77, 51, "DPS", NOW(), NOW()), -- Gaming
(80, 52, "DPS", NOW(), NOW()), -- Arlecchino
(93, 53, "DPS", NOW(), NOW()), -- Mavuika

-- Hydro characters
(21, 54, "OFF-FIELD DPS", NOW(), NOW()), -- Xingqiu
(13, 55, "SUPPORT", NOW(), NOW()), -- Mona
(2, 56, "HEAL SUPPORT", NOW(), NOW()), -- Barbara
(23, 57, "DPS",  NOW(), NOW()), -- Tartaglia
(40, 58, "SUPPORT", NOW(), NOW()), -- Kokomi
(47, 59, "DPS", NOW(), NOW()), -- Ayato
(49, 60, "OFF-FIELD DPS", NOW(), NOW()), -- Yelan
(54, 61, "SUPPORT", NOW(), NOW()), -- Candace
(56, 62, "BLOOM SUPPORT", NOW(), NOW()), -- Nilou
(71, 63, "DPS", NOW(), NOW()), -- Neuvillette
(74, 64, "SUPPORT", NOW(), NOW()), -- Furina
(83, 65, "HEAL", NOW(), NOW()), -- Sigewinne
(87, 66, "DPS",  NOW(), NOW()), -- Mualani

-- Electro Characters
(17, 67, "DPS", NOW(), NOW()), -- Razor
(12, 68, "AGGRAVATE DPS", NOW(), NOW()), -- Lisa
(10, 69, "AGGRAVATE DPS", NOW(), NOW()), -- keqing
(7, 70, "OFF-FIELD DPS", NOW(), NOW()), -- Fischl
(3, 71, "OFF-FIELD AGGRAVATE DPS", NOW(), NOW()), -- Beidou
(39, 72, "DPS & ENERGY SUPPORT Build", NOW(), NOW()), -- Raiden
(38, 73, "BUFF SUPPORT & DAMAGE", NOW(), NOW()), -- Sara
(46, 74, "OFF-FIELD DPS", NOW(), NOW()), -- Yae Miko
(48, 75, "HYPERBLOOM", NOW(), NOW()), -- Kuki Shinobu
(52, 76, "HEAL SUPPORT", NOW(), NOW()), -- Dori
(55, 77, "DPS", NOW(), NOW()), -- Cyno
(82, 78, "NORMAL ATTACK DPS", NOW(), NOW()), -- Sethos
(81, 79, "DPS", NOW(), NOW()), -- Clorinde
(90, 80, "OFF-FIELD DPS", NOW(), NOW()), -- Ororon
(96, 81, "DPS", NOW(), NOW()), -- Varesa
(95, 82, "BUFF", NOW(), NOW()), -- Iansan

-- Anemo

;

INSERT INTO stats (
    id,
    sands,
    goblet,
    sub_stats_priority,
) VALUES

-- Geo Characters Stats
(1, "DEF% / ATK%", "Geo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Energy Recharge 4.DEF%"), -- Noelle dps
(2, "ATK%", "Geo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge"), -- Ningguang
(3, "HP%", "HP%", "HP% / CRIT Rate", "1.HP% 2.Energy Recharge"), -- Zhongli,
(4, "DEF%", "Geo DMG Bonus / DEF", "CRIT Rate / CRIT DMG / DEF%", "1.CRIT Rate 2.CRIT DMG 3. DEF% 4. ATK% 5. Energy Recharge"), -- Albedo
(5, "Energy Recharge", "Geo DMG Bonus / DEF%", "CRIT Rate / DEF% / Healing Bonus", "1.Energy Recharge 2.DEF% 3.CRIT Rate"), -- Gorou
(6, "DEF%", "Geo DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.DEF% 5.ATK%"), -- Itto
(7, "DEF% / Energy Recharge", "DEF%", "DEF% / CRIT Rate", "1.DEF% 2.Energy Recharge 3.DEF 4.CRIT Rate"), -- Yun Jin
(8, "ATK%", "Geo DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK%"), -- Navia
(9, "DEF%", "Geo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.DEF% 4.ATK% 5.Energy Recharge"), -- Chiori
(10, "DEF% / Energy Recharge", "Geo DMG Bonus / DEF%", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.DEF%"), -- Kachina
(11, "DEF% / Energy Recharge", "DEF%", "DEF% / Healing Bonus / CRIT Rate", "1.Energy Recharge 2.DEF% 3.DEF 4.CRIT Rate"), -- Xilonen

-- Dendro Characters stats
(12, "ATK% / Elemental Mastery", "Dendro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK% 5.Energy Recharge"), -- Tighnari
(13, "Energy Recharge / ATK% / Elemental Mastery", "Dendro DMG Bonus / Elemental Mastery", "CRIT Rate / CRIT DMG / Elemental Mastery", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery"), -- Collei
(14, "Elemental Mastery", "Elemental Mastery / Dendro DMG Bonus", "Elemental Mastery / CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.Crit DMG 4.Elemental Mastery 5.ATK%") -- nahida
(15, "HP% / Energy Recharge", "HP%", "Healing Bonus", "1.Energy Recharge 2.HP% 3.CRIT Rate 4.HP%") -- yaoyao
(16, "Elemental Mastery / ATK%", "Dendro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK%") -- alhaitam
(17, "Energy Recharge / Elemental Mastery", "Elemental Mastery", "Elemental Mastery / CRIT Rate", "1.Energy Recharge 2.Elemental Mastery 3.CRIT Rate") -- Kaveh
(18, "HP% / Energy Recharge", "HP%", "HP% / Healing Bonus", "1.Energy Recharge 2.HP% 3.CRIT Rate") -- Baizhu
(19, "HP% / Energy Recharge", "HP%", "HP% / CRIT Rate", "1.Energy Recharge 2.HP% 3.CRIT Rate") -- Kirara
(20, "ATK%", "Dendro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge") -- Emilie
(21, "ATK%", "Dendro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge") -- Kinich

-- Cryo Characters stats
(22, "ATK% / Energy Recharge", "ATK%", "Healing Bonus / ATK%", "1.ATK% 2.Energy Recharge 3.ATK% 4.CRIT Rate 5.DEF%"), -- Qiqi
(23, "ATK%", "Cryo DMG Bonus", "CRIT DMG", "1.CRIT DMG 2.ATK% 3.Energy Recharge 4.CRIT Rate"), -- Kaeya
(24, "ATK% / Energy Recharge / Elemental Mastery", "ATK% / Energy Recharge / Elemental Mastery", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge") -- Chonguyn
(25, "Energy Recharge / HP%", "HP%", "HP% / Healing Bonus", "1.Energy Recharge 2.HP% 3.CRIT Rate 4.HP%"), -- Diona
(26, "Elemental Mastery / ATK%", "Cryo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK%"), -- Ganyu
(27, "ATK% / Energy Recharge", "Cryo DMG Bonus", "CRIT Rate", "1.CRIT Rate 2.Energy Recharge 3.CRIT DMG 4.ATK%"),
(28, "ATK%", "Physical DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Energy Recharge 4.ATK%"), -- Eula
(29, "ATK%", "Cryo DMG Bonus", "CRIT DMG / ATK%", "1.CRIT DMG 2.ATK% 3.Energy Recharge 4.CRIT Rate") -- Ayaka
(30, "ATK% / Elemental Mastery", "Cryo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge") -- Aloy
(31, "ATK% / Energy Recharge", "ATK%", "ATK%", "1.Energy Recharge 2.ATK% 3.CRIT Rate 4.CRIT DMG") -- Shenhe
(32, "HP% / Energy Recharge", "HP%", "HP% / CRIT Rate", "1.Energy Recharge 2.HP% 3.CRIT Rate 4.HP")
(33, "Energy Recharge / HP%", "HP%", "Healing Bonus / CRIT Rate", "1.Energy Recharge 2.CRIT Rate 3.HP%")
(34, "ATK%", "Physical DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK%"), -- Freminet
(35, "ATK% / Elemental Mastery", "Cryo DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.Crit DMG 3.ATK% 4.Elemental Mastery") -- Wriothesley
(36, "Energy Recharge / ATK%", "ATK%", "Healing Bonus / ATK% / CRIT Rate", "1.Energy Recharge 2.ATK% 3.CRIT Rate 4.ATK% 5.CRIT DMG") -- Charlotte
(37, "Elemental Mastery / Energy Recharge", "Elemental Mastery", "Elemental Mastery / CRIT Rate", "1.Energy Recharge 2.Elemental Mastery 3.CRIT Rate") -- Citlali

-- Pyro stats
(38, "Energy Recharge / ATK% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery") -- Xiangling
(39, "ATK%", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery") -- Klee
(40, "Elemental Mastery / ATK%", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge") -- Diluc
(41, "Energy Recharge / HP%", "HP%", "Healing Bonus / HP%", "1.Energy Recharge 2.HP%") -- Bennett
(42, "Energy Recharge / ATK%", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery") --  Amber
(43, "DEF% / Energy Recharge", "DEF%", "DEF%", "1.Energy Recharge 2.ATK% 3.DEF% 4.CRIT Rate 5.CRIT DMG") -- Xinyan
(44, "HP% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.HP% 5.ATK%") --Hutao
(45, "ATK% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge") -- Yanfei
(46, "ATK% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge") -- Yoimiya
(47, "Energy Recharge / HP%", "HP%", "HP% / CRIT Rate", "1.Energy Recharge 2.HP% 3.CRIT Rate") -- Thoma
(48, "ATK% / Energy Recharge", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery 6.HP%") -- Dehya
(49, "ATK%", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge"), -- Lyney
(50, "HP% / Energy Recharge", "HP%", "HP%", "1.Energy Recharge 2.HP% 3.CRIT Rate") -- Chevreuse
(51, "Elemental Mastery / ATK% / Energy Recharge", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK%") -- Gaming
(52, "ATK% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK%"), -- Arlecchino
(53, "ATK% / Elemental Mastery", "Pyro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Mavuika

-- Hydro
(54, "Energy Recharge / ATK%", "Hydro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge"), -- Xingqiu
(55, "ATK% / Energy Recharge / Elemental Mastery", "Hydro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Energy Recharge 4.ATK% 5.Elemental Mastery"), -- Mona
(56, "HP%", "HP%", "Healing Bonus", "1.HP%"), -- Barbara
(57, "ATK%", "Hydro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge"), -- Tartaglia
(58, "Energy Recharge / HP%", "HP%", "Healing Bonus / HP%", "1.Energy Recharge 2.HP%"), -- Kokomi,\
(59, "ATK%", "Hydro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK%") -- Ayato
(60, "Energy Recharge / HP%", "Hydro DMG Bonus / HP%", "CRIT Rate / CRIT DMG / HP%", "1.Energy Recharge 2.HP% 3.CRIT Rate 4.CRIT DMG"), -- Yelan
(61, "Energy Recharge / HP%", "HP% / Hydro DMG Bonus", "HP% / CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.HP%"), -- Candace
(62, "HP%", "HP%", "HP%", "1.HP% 2.Elemental Mastery 3.Energy Recharge"), -- Nilou
(63, "HP%", "Hydro DMG Bonus / HP%", "CRIT Rate / CRIT DMG / HP%", "1.CRIT Rate 2.Crit DMG 3.HP%"), -- Neuvillette    
(64, "Energy Recharge / HP%", "HP% / Hydro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.HP% 3.CRIT Rate 4.Crit DMG"), -- Furina
(65, "HP%", "HP%", "HP%", "1.HP% 2.Energy Recharge"), -- Sigewinne
(66, "HP% / Elemental Mastery", "Hydro DMG Bonus", "CRIT Rate / CRIT DMG / HP%", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.HP%"), -- Mualani

-- Electro
(67, "ATK%", "Physical DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge"), -- Razor
(68, "ATK% / Elemental Mastery", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Lisa
(69, "ATK% / Elemental Mastery", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Keqing
(70, "ATK%", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Fischl
(71, "ATK% / Energy Recharge", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge"), -- Beidou
(72, "Energy Recharge / ATK%", "Electro DMG Bonus / ATK%", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge 5.Elemental Mastery"), -- Raiden
(73, "Energy Recharge / ATK%", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK%"), -- Sara
(74, "Energy Recharge / ATK%", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Energy Recharge 5.Elemental Mastery"), -- Yae Miko
(75, "Elemental Mastery", "Elemental Mastery", "Elemental Mastery", "1.Elemental Mastery 2.HP% 3.Energy Recharge"), -- Kuki Shinobu
(76, "Energy Recharge / HP% / Elemental Mastery", "HP% / Elemental Mastery", "Healing Bonus / HP% / Elemental Mastery", "1.CRIT Rate 2.Energy Recharge 3.HP% 4.Elemental Mastery")
(77, "Elemental Mastery / ATK%", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Cyno
(78, "Elemental Mastery / Energy Recharge", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.Elemental Mastery 4.ATK%"), -- Sethos
(79, "ATK% / Elemental Mastery", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery 5.Energy Recharge"), -- Clorinde
(80, "Elemental Mastery / Energy Recharge", "Elemental Mastery", "Elemental Mastery", "1.Energy Recharge 2.CRIT Rate 3.CRIT DMG 4.ATK% 5.Elemental Mastery"), -- Ororon
(81, "ATK%", "Electro DMG Bonus", "CRIT Rate / CRIT DMG", "1.CRIT Rate 2.CRIT DMG 3.ATK% 4.Elemental Mastery"), -- Varesa

;
