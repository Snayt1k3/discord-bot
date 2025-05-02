INSERT INTO characters (
    id,
    name,
    element,
    weapon_type,
    base_stat,
    region,
    rarity,
    ascension_id,
    talents_id,
    common_materials_id,
    created_at,
    updated_at
)
VALUES
-- Patch: 1.0
(1, "Amber", "Pyro", "Bow", "ATK%", "Mondstadt", 4, 66, 1, 17, NOW(), NOW()),
(2, "Barbara", "Hydro", "Catalyst", "HP%", "Mondstadt", 4, 67, 2, 2, NOW(), NOW())
(3, "Beidou", "Electro", "Claymore", "Electro DMG Bonus%", "Liyue", 4, 45, 3, 4, NOW(), NOW())
(4, "Bennett", "Pyro", "Sword", "Energy Recharge%", "Mondstadt", 4, 68, 4, 4, NOW(), NOW())
(5, "Chongyun", "Cryo", "Claymore", "ATK%", "Liyue", 4, 46, 5, 1, NOW(), NOW())
(6, "Diluc", "Pyro", "Claymore", "CRIT Rate%", "Mondstadt", 5, 69, 6, 5, NOW(), NOW())
(7, "Fischl", "Electro", "Bow", "ATK%", "Mondstadt", 4, 72, 7, 17, NOW(), NOW())
(8, "Jean", "Anemo", "Sword", "Healing Bonus%", "Mondstadt", 5, 73, 8, 1, NOW(), NOW())
(9, "Kaeya", "Cryo", "Sword", "Energy Recharge%", "Mondstadt", 4, 74, 9, 4, NOW(), NOW())
(10, "Keqing", "Electro", "Sword", "CRIT DMG%", "Liyue", 5, 50, 10, 6, NOW(), NOW())
(11, "Klee", "Pyro", "Catalyst", "Pyro DMG Bonus%", "Mondstadt", 5, 75, 11, 2, NOW(), NOW())
(12, "Lisa", "Electro", "Catalyst", "Elemental Mastery", "Mondstadt", 4, 76, 12, 3, NOW(), NOW())
(13, "Mona", "Hydro", "Catalyst", "Energy Recharge%", "Mondstadt", 5, 78, 13, 6, NOW(), NOW())
(14, "Ningguang", "Geo", "Catalyst", "Geo DMG Bonus%", "Liyue", 4, 52, 14, 5, NOW(), NOW())
(15, "Noelle", "Geo", "Claymore", "DEF%", "Mondstadt", 4, 79, 15, 1, NOW(), NOW())
(16, "Qiqi", "Cryo", "Sword", "Healing Bonus%", "Liyue", 5, 53, 16, 2, NOW(), NOW())
(17, "Razor", "Electro", "Claymore", "Physical DMG Bonus%", "Mondstadt", 4, 80, 17, 1, NOW(), NOW())
(18, "Sucrose", "Anemo", "Catalyst", "Anemo DMG Bonus%", "Mondstadt", 4, 82, 18, 6, NOW(), NOW())
(19, "Venti", "Anemo", "Bow", "Energy Recharge%", "Mondstadt", 5, 83, 19, 3, NOW(), NOW())
(20, "Xiangling", "Pyro", "Polearm", "Elemental Mastery", "Liyue", 4, 55, 20, 3, NOW(), NOW())
(21, "Xingqiu", "Hydro", "Sword", "ATK%", "Liyue", 4, 58, 21, 1, NOW(), NOW())

-- Patch: 1.1
(22, "Diona", "Cryo", "Bow", "Cryo DMG Bonus%", "Mondstadt", 4, 70, 22, 17, NOW(), NOW())
(23, "Tartaglia", "Hydro", "Bow", "Hydro DMG Bonus%", "Snezhnaya", 5, 85, 23, 5, NOW(), NOW())
(24, "Xinyan", "Pyro", "Claymore", "ATK%", "Liyue", 4, 59, 24, 4, NOW(), NOW())
(25, "Zhongli", "Geo", "Polearm", "Geo DMG Bonus%", "Liyue", 5, 64, 25, 3, NOW(), NOW())

-- Patch: 1.2
(26, "Albedo", "Geo", "Polearm", "Geo DMG Bonus%", "Mondstadt", 5, 65, 26, 2, NOW(), NOW())
(27, "Ganyu", "Cryo", "Bow", "CRIT DMG%", "Liyue", 5, 48, 27, 6, NOW(), NOW())

-- Patch: 1.3
(28, "Hu Tao", "Pyro", "Polearm", "CRIT DMG%", "Liyue", 5, 49, 28, 6, NOW(), NOW())
(29, "Xiao", "Anemo", "Polearm", "CRIT Rate%", "Liyue", 5, 57, 29, 3, NOW(), NOW())

-- Patch: 1.4
(30, "Rosaria", "Cryo", "Polearm", "ATK%", "Mondstadt", 4, 81, 30, 5, NOW(), NOW())

-- Patch: 1.5
(31, "Eula", "Cryo", "Claymore", "CRIT DMG%", "Mondstadt", 5, 71, 31, 1, NOW(), NOW())
(32, "Yanfei", "Pyro", "Catalyst", "Pyro DMG Bonus%", " Liyue", 4, 60, 32, 4, NOW(), NOW())

-- Patch: 1.6
(33, "Kaedehara Kazuha", "Anemo", "Sword", "Elemental Mastery", "Inazuma", 5, 16, 33, 4, NOW(), NOW())

-- Patch: 2.0
(34, "Kamisato Ayaka", "Cryo", "Sword", "CRIT DMG%", "Inazuma", 5, 17, 34, 8, NOW(), NOW())
(35, "Sayu", "Anemo", "Claymore", "Elemental Mastery", "Inazuma", 4, 24, 35, 6, NOW(), NOW())
(36, "Yoimiya", "Pyro", "Bow", "CRIT Rate%", "Inazuma", 5, 28, 36, 2, NOW(), NOW())

-- Patch: 2.1
(37, "Aloy", "Cryo", "Bow", "Cryo DMG Bonus%", "No region", 5, 96, 37, 9, NOW(), NOW())
(38, "Kujou Sara", "Electro", "Bow", "ATK%", "Inazuma", 4, 20, 38, 1, NOW(), NOW())
(39, "Raiden Shogun", "Electro", "Polearm", "Energy Recharge%", "Inazuma", 5, 22, 39, 8, NOW(), NOW())
(40, "Sangonomiya Kokomi", "Hydro", "Catalyst", "Hydro DMG Bonus%", "Inazuma", 5, 23, 40, 9, NOW(), NOW())

-- Patch: 2.2
(41, "Thoma", "Pyro", "Polearm", "ATK%", "Inazuma", 4, 26, 41, 4, NOW(), NOW())

-- Patch: 2.3
(42, "Arataki Itto", "Geo", "Claymore", "CRIT Rate%", "Inazuma", 5, 13, 42, 3, NOW(), NOW())
(43, "Gorou", "Geo", "Bow", "Geo DMG Bonus%", "Inazuma", 4, 15, 43, 9, NOW(), NOW())

-- Patch: 2.4
(44, "Shenhe", "Cryo", "Polearm", "ATK%", "Liyue", 5, 54, 44, 6, NOW(), NOW())
(45, "Yun Jin", "Geo", "Polearm", "Energy Recharge%", "Liyue", 4, 63, 45, 1, NOW(), NOW())

-- Patch: 2.5
(46, "Yae Miko", "Electro", "Catalyst", "CRIT Rate%", "Inazuma", 5, 27, 46, 8, NOW(), NOW())

-- Patch: 2.6
(47, "Kamisato Ayato", "Hydro", "Sword", "CRIT DMG%", "Inazuma", 5, 18, 47, 8, NOW(), NOW())

-- Patch: 2.7
(48, "Kuki Shinobu", "Electro", "Sword", "HP%", "Inazuma", 4, 21, 48, 9, NOW(), NOW())
(49, "Yelan", "Hydro", "Bow", "CRIT Rate%", "Liyue", 5, 62, 49, 5, NOW(), NOW())

-- Patch: 2.8
(50, "Shikanoin Heizou", "Anemo", "Catalyst", "Anemo DMG Bonus%", "Inazuma", 4, 25, 50, 4, NOW(), NOW())

-- Patch: 3.0
(51, "Collei", "Dendro", "Bow", "ATK%", "Sumeru", 4, 32, 51, 17, NOW(), NOW())
(52, "Dori", "Electro", "Claymore", "HP%", "Sumeru", 4, 35, 52, 11, NOW(), NOW())
(53, "Tighnari", "Dendro", "Bow", "Dendro DMG Bonus%", "Sumeru", 5, 42, 53, 4, NOW(), NOW())

-- Patch: 3.1
(54, "Candace", "Hydro", "Polearm", "HP%", "Sumeru", 4, 31, 54, 11, NOW(), NOW())
(55, "Cyno", "Electro", "Polearm", "CRIT DMG%", "Sumeru", 5, 33, 55, 2, NOW(), NOW())
(56, "Nilou", "Hydro", "Sword", "HP%", "Sumeru", 5, 40, 56, 10, NOW(), NOW())

-- Patch: 3.2
(57, "Layla", "Cryo", "Sword", "HP%", "Sumeru", 4, 38, 57, 2, NOW(), NOW())
(58, "Nahida", "Dendro", "Catalyst", "Elemental Mastery", "Sumeru", 5, 39, 58, 10, NOW(), NOW())

-- Patch: 3.3
(59, "Faruzan", "Anemo", "Bow", "ATK%", "Sumeru", 4, 36, 59, 11, NOW(), NOW())
(60, "Wanderer", "Anemo", "Catalyst", "CRIT Rate%", "Sumeru", 5, 43, 60, 8, NOW(), NOW())

-- Patch: 3.4
(61, "Alhaitham", "Dendro", "Sword", "Dendro DMG Bonus%", "Sumeru", 5, 30, 61, 11, NOW(), NOW())
(62, "Yaoyao", "Dendro", "Polearm", "HP%", "Liyue", 4, 61, 62, 3, NOW(), NOW())

-- Patch: 3.5
(63, "Dehya", "Pyro", "Claymore", "HP%", "Sumeru", 5, 34, 63, 11, NOW(), NOW())
(64, "Mika", "Cryo", "Polearm", "HP%", "Mondstadt", 4, 77, 64, 5, NOW(), NOW())

-- Patch: 3.6
(65, "Baizhu", "Dendro", "Catalyst", "HP%", "Liyue", 5, 44, 65, 10, NOW(), NOW())
(66, "Kaveh", "Dendro", "Claymore", "HP%", "Sumeru", 4, 37, 66, 10, NOW(), NOW())

-- Patch: 3.7
(67, "Kirara", "Dendro", "Sword", "HP%", "Inazuma", 4, 19, 67, 9, NOW(), NOW())

-- Patch: 4.0
(68, "Freminet", "Cryo", "Claymore", "ATK%", "Fontaine", 4, 5, 68, 13, NOW(), NOW())
(69, "Lynette", "Anemo", "Sword", "Anemo DMG Bonus%", "Fontaine", 4, 7, 69, 12, NOW(), NOW())
(70, "Lyney", "Pyro", "Bow", "CRIT Rate%", "Fontaine", 5, 8, 70, 5, NOW(), NOW())

-- Patch: 4.1
(71, "Neuvillette", "Hydro", "Catalyst", "CRIT DMG%", "Fontaine", 5, 10, 71, 13, NOW(), NOW())
(72, "Wriothesley", "Cryo", "Catalyst", "CRIT DMG%", "Fontaine", 5, 12, 72, 12, NOW(), NOW())

-- Patch: 4.2
(73, "Charlotte", "Cryo", "Catalyst", "ATK%", "Fontaine", 4, 1, 73, 12, NOW(), NOW()),
(74, "Furina", "Hydro", "Sword", "CRIT Rate%", "Fontaine", 5, 6, 74, 6, NOW(), NOW()),

-- Patch: 4.3
(75, "Chevreuse", "Pyro", "Polearm", "HP%", "Fontaine", 4, 2, 75, 12, NOW(), NOW()),
(76, "Navia", "Geo", "Claymore", "CRIT DMG%", "Fontaine", 5, 9, 76, 13, NOW(), NOW()),

-- Patch: 4.4
(77, "Gaming", "Pyro", "Claymore", "ATK%", "Liyue", 4, 47, 77, 3, NOW(), NOW()),
(78, "Xianyun", "Anemo", "Catalyst", "ATK%", "Liyue", 5, 56, 78, 2, NOW(), NOW()),

-- Patch: 4.5
(79, "Chiori", "Geo", "Sword", "CRIT Rate%", "Inazuma", 5, 14, 79, 9, NOW(), NOW()),

-- Patch: 4.6
(80, "Arlecchino", "Pyro", "Polearm", "CRIT DMG%", "Snezhnaya", 5, 84, 80, 5, NOW(), NOW()),

-- Patch: 4.7
(81, "Clorinde", "Electro", "Sword", "CRIT Rate%", "Fontaine", 5, 3, 81, 13, NOW(), NOW()),
(82, "Sethos", "Electro", "Bow", "Elemental Mastery", "Sumeru", 4, 41, 82, 11, NOW(), NOW()),
(83, "Sigewinne", "Hydro", "Bow", "HP%", "Fontaine", 5, 11, 83, 13, NOW(), NOW()),

-- Patch: 4.8
(84, "Emilie", "Dendro", "Polearm", "CRIT DMG%", "Fontaine", 5, 4, 84, 12, NOW(), NOW())

-- Patch: 5.0
(85, "Kachina", "Geo", "Polearm", "Geo DMG Bonus%", "Natlan", 4, 89, 85, 15, NOW(), NOW()),
(86, "Kinich", "Dendro", "Claymore", "CRIT DMG%", "Natlan", 5, 90, 86, 16, NOW(), NOW()),
(87, "Mualani", "Hydro", "Catalyst", "CRIT Rate%", "Natlan", 5, 92, 87, 15, NOW(), NOW()),

-- Patch: 5.1
(88, "Xilonen", "Geo", "Sword", "DEF%", "Natlan", 5, 95, 88, 15, NOW(), NOW()),

-- Patch: 5.2
(89, "Chasca", "Anemo", "Bow", "CRIT Rate%", "Natlan", 5, 86, 89, 16, NOW(), NOW()),
(90, "Ororon", "Electro", "Bow", "ATK%", "Natlan", 4, 93, 90, 16, NOW(), NOW()),

-- Patch: 5.3
(91, "Citlali", "Cryo", "Catalyst", "Elemental Mastery", "Natlan", 5, 87, 91, 16, NOW(), NOW()),
(92, "Lan Yan", "Anemo", "Catalyst", "ATK%", "Liyue", 4, 51, 92, 6, NOW(), NOW()),
(93, "Mavuika", "Pyro", "Claymore", "CRIT DMG%", "Natlan", 5, 91, 93, 15, NOW(), NOW()),

-- Patch: 5.4
(94, "Yumemizuki Mizuki", "Anemo", "Catalyst", "Elemental Mastery", "Inazuma", 5, 29, 94, 8, NOW(), NOW()),

-- Patch: 5.5
(95, "Iansan", "Electro", "Polearm", "ATK%", "Natlan", 4, 88, 95, 15, NOW(), NOW()),
(96, "Varesa", "Electro", "Catalyst", "CRIT Rate%", "Natlan", 5, 94, 96, 16, NOW(), NOW())
