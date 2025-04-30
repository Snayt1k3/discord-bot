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
(65, 18, "HEAL SUPPORT", NOW(), NOW()) -- Baizhu
(67, 19, "Support", NOW(), NOW()) -- Kirara
(84, 20, "OFF-FIELD DPS", NOW(), NOW()) -- Emilie
(85, 21, "DPS", NOW(), NOW()) -- Kinich
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
;