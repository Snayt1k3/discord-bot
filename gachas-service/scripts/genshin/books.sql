INSERT INTO books (id, common, uncommon, rare, weekdays, created_at, updated_at)
VALUES 
-- Мондштадт
(1, 'Teachings of Freedom', 'Guide to Freedom', 'Philosophies of Freedom', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(2, 'Teachings of Resistance', 'Guide to Resistance', 'Philosophies of Resistance', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(3, 'Teachings of Ballad', 'Guide to Ballad', 'Philosophies of Ballad', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW()),

-- Ли Юэ
(4, 'Teachings of Prosperity', 'Guide to Prosperity', 'Philosophies of Prosperity', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(5, 'Teachings of Diligence', 'Guide to Diligence', 'Philosophies of Diligence', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(6, 'Teachings of Gold', 'Guide to Gold', 'Philosophies of Gold', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW()),

-- Инадзума
(7, 'Teachings of Transience', 'Guide to Transience', 'Philosophies of Transience', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(8, 'Teachings of Elegance', 'Guide to Elegance', 'Philosophies of Elegance', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(9, 'Teachings of Light', 'Guide to Light', 'Philosophies of Light', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW()),

-- Сумеру
(10, 'Teachings of Admonition', 'Guide to Admonition', 'Philosophies of Admonition', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(11, 'Teachings of Ingenuity', 'Guide to Ingenuity', 'Philosophies of Ingenuity', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(12, 'Teachings of Praxis', 'Guide to Praxis', 'Philosophies of Praxis', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW()),

-- Фонтейн
(13, 'Teachings of Justice', 'Guide to Justice', 'Philosophies of Justice', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(14, 'Teachings of Equity', 'Guide to Equity', 'Philosophies of Equity', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(15, 'Teachings of Order', 'Guide to Order', 'Philosophies of Order', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW()),

-- Натлан
(16, 'Teachings of Contention', 'Guide to Contention', 'Philosophies of Contention', ARRAY['Monday', 'Thursday', 'Sunday'], NOW(), NOW()),
(17, 'Teachings of Kindling', 'Guide to Kindling', 'Philosophies of Kindling', ARRAY['Tuesday', 'Friday', 'Sunday'], NOW(), NOW()),
(18, 'Teachings of Conflict', 'Guide to Conflict', 'Philosophies of Conflict', ARRAY['Wednesday', 'Saturday', 'Sunday'], NOW(), NOW());

INSERT INTO talents (id, boss_drops, books_id, talent_priority, created_at, updated_at)
VALUES 

-- Amber 
(1, "Dvalin's Sigh", 1, "", NOW(), NOW()),
-- Barbara
(2, "Ring of Boreas", 1, "", NOW(), NOW())
-- Beidou
(3, "Dvalin's Sigh", 6, "", NOW(), NOW())
-- Bennett
(4, "Dvalin's Plume", 2, "", NOW(), NOW())
-- Chonguyn
(5, "Dvalin's Sigh", 5, "", NOW(), NOW())
-- Diluc
(6, "Dvalin's Plume", 2, "", NOW(), NOW())
-- Fischl
(7, "Spirit Locket of Boreas", 3, "", NOW(), NOW())
-- Jean
(8, "Dvalin's Plume", 2, "", NOW(), NOW())
-- Kaeya
(9, "Spirit Locket of Boreas", 3, "", NOW(), NOW())
-- Keqing
(10, "Ring of Boreas", 4, "", NOW(), NOW())
-- Klee
(11, "Ring of Boreas", 1, "", NOW(), NOW())
-- Lisa
(12, "Dvalin's Claw", 3, "", NOW(), NOW())
-- Mona
(13, "Ring of Boreas", 2, "", NOW(), NOW())
-- Ningguang
(14, "Spirit Locket of Boreas", 4, "", NOW(), NOW())
-- Noelle
(15, "Dvalin's Claw", 2, "", NOW(), NOW())
-- Qiqi
(16, "Tail of Boreas", 4, "", NOW(), NOW())
-- Razor
(17, "Dvalin's Claw", 2, "", NOW(), NOW())
-- Sucrose
(18, "Spirit Locket of Boreas", 1, "", NOW(), NOW())
-- Venti
(19, "Tail of Boreas", 3, "", NOW(), NOW())
-- Xiangling
(20, "Dvalin's Claw", 5, "", NOW(), NOW())
-- Xingqiu
(21, "Dvalin's Claw", 6, "", NOW(), NOW())
-- Diona
(22, "Shard of a Foul Legacy", 1, "", NOW(), NOW())
-- Tartaglia
(23, "Shard of a Foul Legacy", 1, "", NOW(), NOW())
-- Xinyan
(24, "Tusk of Mo­no­ceros Caeli", 6, "", NOW(), NOW())
-- Zhongli
(25, "Tusk of Mo­no­ceros Caeli", 6, "", NOW(), NOW())
-- Albedo
(26, "Tusk of Mo­no­ceros Caeli", 3, "", NOW(), NOW())
-- Ganyu
(27, "Shadow of the Warrior", 5, "", NOW(), NOW())
-- Hu Tao
(28, "Shard of a Foul Legacy", 5, "", NOW(), NOW())
-- Xiao
(29, "Shadow of the Warrior", 5, "", NOW(), NOW())
-- Rosaria
(30, "Shadow of the Warrior", 3, "", NOW(), NOW())
-- Eula
(31, "Drag­on Lord's Crown", 2, "", NOW(), NOW())
-- YanFei
(32, "Blood­jade Branch", 6, "", NOW(), NOW())
-- Kazuha
(33, "Gil­ded Scale", 5, "", NOW(), NOW())

-- Patch: 2.0
(34, "Blood­jade Branch", 8, "", NOW(), NOW()) -- Ayaka
(35, "Gil­ded Scale", 9, "", NOW(), NOW()) -- Sayu
(36, "Drag­on Lord's Crown", 7, "", NOW(), NOW()) -- Yoimiya

-- Patch: 2.1
(37, "Molten Moment", 1, "", NOW(), NOW()) -- Aloy
(38, "Ashen Heart", 8, "", NOW(), NOW()) -- Kujou Sara
(39, "Molten Moment", 9, "", NOW(), NOW()) -- Raiden Shogun
(40, "Hell­fire But­ter­fly", 7, "", NOW(), NOW()) -- Kokomi

-- Patch: 2.2
(41, "Hell­fire But­ter­fly", 7, "", NOW(), NOW()) -- Thoma

-- Patch: 2.3
(42, "Ashen Heart", 8, "", NOW(), NOW()) -- Itto
(43, "Molten Moment", 9, "", NOW(), NOW()) -- Gorou

-- Patch: 2.4
(44, "Hell­fire But­ter­fly", 4, "", NOW(), NOW()) -- Shenhe
(45, "Ashen Heart", 5, "", NOW(), NOW()) -- Yun Jin

-- Patch: 2.5
(46, "The Meaning of Aeons", 9, "", NOW(), NOW()) -- Yae miko

-- Patch: 2.6
(47, "Mudra of the Malefic General", 8, "", NOW(), NOW()) -- Ayato

-- Patch: 2.7
(48, "Tears of the Calami­tous God", 8, "", NOW(), NOW()) -- Kuki
(49, "Gil­ded Scale", 4, "", NOW(), NOW()) -- Yelan

-- Patch: 2.8
(50, "The Meaning of Aeons", 7, "", NOW(), NOW()) -- Heizou

-- Patch: 3.0
(51, "Tears of the Calami­tous God", 12, "", NOW(), NOW()) -- Collei
(52, "Blood­jade Branch", 11, "", NOW(), NOW()) -- Dori
(53, "The Meaning of Aeons", 10, "", NOW(), NOW()) -- Tighnari

-- Patch: 3.1
(54, "Tears of the Calami­tous God", 10, "", NOW(), NOW()) -- Candace
(55, "Mudra of the Malefic General", 10, "", NOW(), NOW()) -- Cyno
(56, "Tears of the Calami­tous God", 12, "", NOW(), NOW()) -- Nilou

-- Patch: 3.2
(57, "Mirror of Mushin", 11, "", NOW(), NOW()) -- Layla
(58, "Puppet Strings", 11, "", NOW(), NOW()) -- Nahida

-- Patch: 3.3
(59, "Puppet Strings", 10, "", NOW(), NOW()) -- Faruzan
(60, "Daka's Bell", 12, "", NOW(), NOW()) -- Wanderer

-- Patch: 3.4
(61, "Mirror of Mushin", 11, "", NOW(), NOW()) -- Alhaitham
(62, "Daka's Bell", 5, "", NOW(), NOW()) -- Yaoyao

-- Patch: 3.5
(63, "Puppet Strings", 12, "", NOW(), NOW()) -- Dehya
(64, "Mirror of Mushin", 3, "", NOW(), NOW()) -- Mika

-- Patch: 3.6
(65, "World­span Fern", 6, "", NOW(), NOW()) -- Baizhu
(66, "Primor­dial Green­bloom", 11, "", NOW(), NOW()) -- Kaveh

-- Patch: 3.7
(67, "Ever­amber", 7, "", NOW(), NOW()) -- Kirara

-- Patch: 4.0
(68, "World­span Fern", 13, "", NOW(), NOW()) -- Freminet
(69, "Ever­amber", 15, "", NOW(), NOW()) -- lynette
(70, "Primor­dial Green­bloom", 14, "", NOW(), NOW()) -- lyney

-- Patch: 4.1
(71, "Ever­amber", 14, "", NOW(), NOW()) -- Neuvillette
(72, "Primor­dial Green­bloom", 15, "", NOW(), NOW()) -- Wriothesley

-- Patch: 4.2
(73, "Light­less Silk String", 13, "", NOW(), NOW()), -- Charlotte
(74, "Light­less Mass", 13, "", NOW(), NOW()), -- Furina

-- Patch: 4.3
(75, "Light­less Eye of the Mael­strom", 15, "", NOW(), NOW()), -- Chevreuse
(76, "Light­less Silk String", 14, "", NOW(), NOW()), -- Navia

-- Patch: 4.4
(77, "Light­less Mass", 4, "", NOW(), NOW()), -- Gaming
(78, "Light­less Eye of the Mael­strom", 6, "", NOW(), NOW()), -- Xianyun

-- Patch: 4.5
(79, "Light­less Silk String", 9, "", NOW(), NOW()), -- Chiori

-- Patch: 4.6
(80, "Fading Candle", 15, "", NOW(), NOW()), -- Arlecchino

-- Patch: 4.7
(81, "Ever­amber", 13, "", NOW(), NOW()), -- Clorinde
(82, "Daka's Bell", 12, "", NOW(), NOW()), -- Sethos
(83, "Light­less Eye of the Mael­strom", 14, "", NOW(), NOW()), -- Sigewinne

-- Patch: 4.8
(84, "Silken Feather", 15, "", NOW(), NOW()) -- Emilie

-- Patch: 5.0
(85, "Fading Candle", 18, "", NOW(), NOW()), -- Kachina
(86, "De­nial and Judg­ment", 17, "", NOW(), NOW()), -- Kinich
(87, "Light­less Mass", 16, "", NOW(), NOW()), -- Mualani

-- Patch: 5.1
(88, "Mirror of Mushin", 17, "", NOW(), NOW()), -- Xilonen

-- Patch: 5.2
(89, "Silken Feather", 18, "", NOW(), NOW()), -- Chasca
(90, "Light­less Silk String", 17, "", NOW(), NOW()), -- Ororon

-- Patch: 5.3
(91, "De­nial and Judg­ment", 17, "", NOW(), NOW()), -- Citlali
(92, "Eroded Sunfire", 5, "", NOW(), NOW()), -- Lan Yan
(93, "Eroded Horn", 16, "", NOW(), NOW()), -- Mavuika

-- Patch: 5.4
(94, "Fading Candle", 7, "", NOW(), NOW()), -- Yumemizuki Mizuki

-- Patch: 5.5
(95, "De­nial and Judg­ment", 16, "", NOW(), NOW()), -- Iansan
(96, "Eroded Scale-Feather", 18, "", NOW(), NOW()) -- Varesa