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

INSERT INTO talent_materials (id, boss_drops, books_id, talent_priority, created_at, updated_at)
VALUES 

-- Amber 
(1, 'Dvalin''s Sigh', 1, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Barbara
(2, 'Ring of Boreas', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Beidou
(3, 'Dvalin''s Sigh', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Bennett
(4, 'Dvalin''s Plume', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Chongyun
(5, 'Dvalin''s Sigh', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Diluc
(6, 'Dvalin''s Plume', 2, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),
-- Fischl
(7, 'Spirit Locket of Boreas', 3, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Jean
(8, 'Dvalin''s Plume', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Kaeya
(9, 'Spirit Locket of Boreas', 3, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()),
-- Keqing
(10, 'Ring of Boreas', 4, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Klee
(11, 'Ring of Boreas', 1, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Lisa
(12, 'Dvalin''s Claw', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Mona
(13, 'Ring of Boreas', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Ningguang
(14, 'Spirit Locket of Boreas', 4, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()),
-- Noelle
(15, 'Dvalin''s Claw', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Qiqi
(16, 'Tail of Boreas', 4, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Razor
(17, 'Dvalin''s Claw', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Sucrose
(18, 'Spirit Locket of Boreas', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Venti
(19, 'Tail of Boreas', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Xiangling
(20, 'Dvalin''s Claw', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Xingqiu
-- Xingqiu
(21, 'Dvalin''s Claw', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Diona
(22, 'Shard of a Foul Legacy', 1, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Tartaglia
(23, 'Shard of a Foul Legacy', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Xinyan
(24, 'Tusk of Monoceros Caeli', 6, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Zhongli
(25, 'Tusk of Monoceros Caeli', 6, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Albedo
(26, 'Tusk of Monoceros Caeli', 3, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),
-- Ganyu
(27, 'Shadow of the Warrior', 5, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),
-- Hu Tao
(28, 'Shard of a Foul Legacy', 5, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),
-- Xiao
(29, 'Shadow of the Warrior', 5, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Rosaria
(30, 'Shadow of the Warrior', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Eula
(31, 'Dragon Lord''s Crown', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Yanfei
(32, 'Bloodjade Branch', 6, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),
-- Kazuha
(33, 'Gilded Scale', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),
-- Patch: 2.0
(34, 'Bloodjade Branch', 8, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()), -- Ayaka
(35, 'Gilded Scale', 9, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Sayu
(36, 'Dragon Lord''s Crown', 7, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()), -- Yoimiya
-- Patch: 2.1
(37, 'Molten Moment', 1, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Aloy
(38, 'Ashen Heart', 8, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kujou Sara
(39, 'Molten Moment', 9, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Raiden Shogun
(40, 'Hellfire Butterfly', 7, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kokomi
-- Patch: 2.2
(41, 'Hellfire Butterfly', 7, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Thoma

-- Patch: 2.3
(42, 'Ashen Heart', 8, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()), -- Itto
(43, 'Molten Moment', 9, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Gorou

-- Patch: 2.4
(44, 'Hellfire Butterfly', 4, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Shenhe
(45, 'Ashen Heart', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Yun Jin

-- Patch: 2.5
(46, 'The Meaning of Aeons', 9, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Yae Miko

-- Patch: 2.6
(47, 'Mudra of the Malefic General', 8, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Ayato

-- Patch: 2.7
(48, 'Tears of the Calamitous God', 8, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kuki
(49, 'Gilded Scale', 4, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Yelan

-- Patch: 2.8
(50, 'The Meaning of Aeons', 7, 'E(Skill) > NA(Normal Attack) > Q(Burst)', NOW(), NOW()), -- Heizou

-- Patch: 3.0
(51, 'Tears of the Calamitous God', 12, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Collei
(52, 'Bloodjade Branch', 11, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Dori
(53, 'The Meaning of Aeons', 10, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()), -- Tighnari

-- Patch: 3.1
(54, 'Tears of the Calamitous God', 10, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Candace
(55, 'Mudra of the Malefic General', 10, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Cyno
(56, 'Tears of the Calamitous God', 12, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Nilou

-- Patch: 3.2
(57, 'Mirror of Mushin', 11, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Layla
(58, 'Puppet Strings', 11, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Nahida

-- Patch: 3.3
(59, 'Puppet Strings', 10, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Faruzan
(60, 'Daka''s Bell', 12, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()), -- Wanderer

-- Patch: 3.4
(61, 'Mirror of Mushin', 11, 'E(Skill) > NA(Normal Attack) > Q(Burst)', NOW(), NOW()), -- Alhaitham
(62, 'Daka''s Bell', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Yaoyao

-- Patch: 3.5
(63, 'Puppet Strings', 12, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Dehya
(64, 'Mirror of Mushin', 3, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Mika

-- Patch: 3.6
(65, 'World­span Fern', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Baizhu
(66, 'Primor­dial Green­bloom', 11, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()), -- Kaveh

-- Patch: 3.7
(67, 'Ever­amber', 7, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kirara

-- Patch: 4.0
(68, 'World­span Fern', 13, 'E(Skill) > NA(Normal Attack) > Q(Burst)', NOW(), NOW()), -- Freminet
(69, 'Ever­amber', 15, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- lynette
(70, 'Primor­dial Green­bloom', 14, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()), -- lyney

-- Patch: 4.1
(71, 'Ever­amber', 14, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()), -- Neuvillette
(72, 'Primor­dial Green­bloom', 15, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()), -- Wriothesley

-- Patch: 4.2
(73, 'Light­less Silk String', 13, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Charlotte
(74, 'Light­less Mass', 13, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Furina

-- Patch: 4.3
(75, 'Light­less Eye of the Mael­strom', 15, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Chevreuse
(76, 'Light­less Silk String', 14, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Navia

-- Patch: 4.4
(77, 'Light­less Mass', 4, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Gaming
(78, 'Light­less Eye of the Mael­strom', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Xianyun

-- Patch: 4.5
(79, 'Light­less Silk String', 9, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Chiori

-- Patch: 4.6
(80, 'Fading Candle', 15, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()), -- Arlecchino

-- Patch: 4.7
(81, 'Ever­amber', 13, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Clorinde
(82, 'Daka''s Bell', 12, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()), -- Sethos
(83, 'Light­less Eye of the Mael­strom', 14, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Sigewinne

-- Patch: 4.8
(84, 'Silken Feather', 15, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Emilie

-- Patch: 5.0
(85, 'Fading Candle', 18, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kachina
(86, 'De­nial and Judg­ment', 17, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Kinich
(87, 'Light­less Mass', 16, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Mualani

-- Patch: 5.1
(88, 'Mirror of Mushin', 17, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Xilonen

-- Patch: 5.2
(89, 'Silken Feather', 18, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Chasca
(90, 'Light­less Silk String', 17, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Ororon

-- Patch: 5.3
(91, 'De­nial and Judg­ment', 17, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Citlali
(92, 'Eroded Sunfire', 5, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Lan Yan
(93, 'Eroded Horn', 16, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Mavuika

-- Patch: 5.4
(94, 'Fading Candle', 7, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Yumemizuki Mizuki

-- Patch: 5.5
(95, 'De­nial and Judg­ment', 16, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Iansan
(96, 'Eroded Scale-Feather', 18, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()); -- Varesa