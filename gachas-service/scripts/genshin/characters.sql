INSERT INTO common_materials (id, common, uncommon, rare, created_at, updated_at)
VALUES 
    (1, 'Damaged Mask', 'Stained Mask', 'Ominous Mask', NOW(), NOW()),
    (2, 'Divining Scroll', 'Sealed Scroll', 'Forbidden Curse Scroll', NOW(), NOW()),
    (3, 'Slime Condensate', 'Slime Secretions', 'Slime Concentrate', NOW(), NOW()),
    (4, 'Treasure Hoarder Insignia', 'Silver Raven Insignia', 'Golden Raven Insignia', NOW(), NOW()),
    (5, 'Recruit''s Insignia', 'Sergeant''s Insignia', 'Lieutenant''s Insignia', NOW(), NOW()),
    (6, 'Whopperflower Nectar', 'Shimmering Nectar', 'Energy Nectar', NOW(), NOW()),
    (7, 'Heavy Horn', 'Black Bronze Horn', 'Black Crystal Horn', NOW(), NOW()),
    (8, 'Old Handguard', 'Kageuchi Handguard', 'Famed Handguard', NOW(), NOW()),
    (9, 'Spectral Husk', 'Spectral Heart', 'Spectral Nucleus', NOW(), NOW()),
    (10, 'Fungal Spores', 'Luminescent Pollen', 'Crystalline Cyst Dust', NOW(), NOW()),
    (11, 'Faded Red Satin', 'Trimmed Red Silk', 'Rich Red Brocade', NOW(), NOW()),
    (12, 'Meshing Gear', 'Mechanical Spur Gear', 'Artificed Dynamic Gear', NOW(), NOW()),
    (13, 'Transoceanic Pearl', 'Transoceanic Chunk', 'Xenochromatic Crystal', NOW(), NOW()),
    (14, 'Hunter''s Sacrificial Knife', 'Agent''s Sacrificial Knife', 'Inspector''s Sacrificial Knife', NOW(), NOW()),
    (15, 'Sentry''s Wooden Whistle', 'Warrior''s Metal Whistle', 'Saurian-Crowned Warrior''s Golden Whistle', NOW(), NOW()),
    (16, 'Juvenile Fang', 'Seasoned Fang', 'Tyrant''s Fang', NOW(), NOW()),
    (17, 'Firm Arrowhead', 'Sharp Arrowhead', 'Weathered Arrowhead', NOW(), NOW());

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
    -- Patch 1.0
    (1, 'Dvalin''s Sigh', 1, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Amber 
    (2, 'Ring of Boreas', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Barbara
    (3, 'Dvalin''s Sigh', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Beidou
    (4, 'Dvalin''s Plume', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Bennett
    (5, 'Dvalin''s Sigh', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Chongyun
    (6, 'Dvalin''s Plume', 2, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),  -- Diluc
    (7, 'Spirit Locket of Boreas', 3, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Fischl
    (8, 'Dvalin''s Plume', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Jean
    (9, 'Spirit Locket of Boreas', 3, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()),  -- Kaeya
    (10, 'Ring of Boreas', 4, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Keqing
    (11, 'Ring of Boreas', 1, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Klee
    (12, 'Dvalin''s Claw', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Lisa
    (13, 'Ring of Boreas', 2, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Mona
    (14, 'Spirit Locket of Boreas', 4, 'Q(Burst) > NA(Normal Attack) > E(Skill)', NOW(), NOW()),  -- Ningguang
    (15, 'Dvalin''s Claw', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Noelle
    (16, 'Tail of Boreas', 4, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Qiqi
    (17, 'Dvalin''s Claw', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Razor
    (18, 'Spirit Locket of Boreas', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Sucrose
    (19, 'Tail of Boreas', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Venti
    (20, 'Dvalin''s Claw', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Xiangling
    (21, 'Dvalin''s Claw', 6, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Xingqiu

    -- Patch: 1.1
    (22, 'Shard of a Foul Legacy', 1, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()),  -- Diona
    (23, 'Shard of a Foul Legacy', 1, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Tartaglia
    (24, 'Tusk of Monoceros Caeli', 6, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Xinyan
    (25, 'Tusk of Monoceros Caeli', 6, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Zhongli

    -- Patch: 1.2
    (26, 'Tusk of Monoceros Caeli', 3, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()),  -- Albedo
    (27, 'Shadow of the Warrior', 5, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),  -- Ganyu

    -- Patch: 1.3
    (28, 'Shard of a Foul Legacy', 5, 'NA(Normal Attack) > E(Skill) > Q(Burst)', NOW(), NOW()),  -- Hu Tao
    (29, 'Shadow of the Warrior', 5, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Xiao

    -- Patch: 1.4
    (30, 'Shadow of the Warrior', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Rosaria

    -- Patch: 1.5
    (31, 'Dragon Lord''s Crown', 2, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Eula
    (32, 'Bloodjade Branch', 6, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()),  -- Yanfei

    -- Patch: 1.7
    (33, 'Gilded Scale', 5, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Kazuha

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
    (96, 'Eroded Scale-Feather', 18, 'NA(Normal Attack) > Q(Burst) > E(Skill)', NOW(), NOW()), -- Varesa

    -- Patch: 5.6
    (97, 'Ascended Sample: Rook', 18, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Ifa
    (98, 'Eroded Horn', 13, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()), -- Escoffier

    -- Patch: 5.7
    (99, 'Eroded Scale-Feather', 3, 'Q(Burst) > E(Skill) > NA(Normal Attack)', NOW(), NOW()), -- Dahlia
    (100, 'Ascended Sample: Knight', 16, 'E(Skill) > Q(Burst) > NA(Normal Attack)', NOW(), NOW()); -- Skirk



INSERT INTO ascension_materials (id, local_specialty, boss_drops, gem, created_at, updated_at)
VALUES
    (1, 'Beryl Conch', 'Tourbillon Device', 'Shivada Jade', NOW(), NOW()), -- 1 — Charlotte
    (2, 'Lumidouce Bell', 'Fontemer Unihorn', 'Agnidus Agate', NOW(), NOW()), -- 2 — Chevreuse
    (3, 'Lumitoile', 'Fontemer Unihorn', 'Vajrada Amethyst', NOW(), NOW()), -- 3 — Clorinde
    (4, 'Lakelight Lily', 'Fragment of a Golden Melody', 'Nagadus Emerald', NOW(), NOW()), -- 4 — Emilie
    (5, 'Romaritime Flower', 'Artificed Spare Clockwork Component – Coppelius', 'Shivada Jade', NOW(), NOW()), -- 5 — Freminet
    (6, 'Lakelight Lily', 'Water That Failed to Transcend', 'Varunada Lazurite', NOW(), NOW()), -- 6 — Furina
    (7, 'Lumidouce Bell', 'Artificed Spare Clockwork Component — Coppelia', 'Vayuda Turquoise', NOW(), NOW()), -- 7 — Lynette
    (8, 'Rainbow Rose', 'Emperor’s Resolution', 'Agnidus Agate', NOW(), NOW()), -- 8 — Lyney
    (9, 'Spring of the First Dewdrop', 'Artificed Spare Clockwork Component — Coppelius', 'Prithiva Topaz', NOW(), NOW()), -- 9 — Navia
    (10, 'Lumitoile', 'Fontemer Unihorn', 'Varunada Lazurite', NOW(), NOW()), -- 10 — Neuvillette
    (11, 'Romaritime Flower', 'Water That Failed to Transcend', 'Varunada Lazurite', NOW(), NOW()), -- 11 — Sigewinne
    (12, 'Subdetection Unit', 'Tourbillon Device', 'Shivada Jade', NOW(), NOW()), -- 12 — Wriothesley
    (13, 'Onikabuto', 'Riftborn Regalia', 'Prithiva Topaz', NOW(), NOW()), -- 13 — Arataki Itto
    (14, 'Dendrobium', 'Artificed Spare Clockwork Component — Coppelia', 'Prithiva Topaz', NOW(), NOW()), -- 14 — Chiori
    (15, 'Sango Pearl', 'Perpetual Heart', 'Prithiva Topaz', NOW(), NOW()), -- 15 — Gorou
    (16, 'Sea Ganoderma', 'Marionette Core', 'Vayuda Turquoise', NOW(), NOW()), -- 16 — Kaedehara Kazuha
    (17, 'Sakura Bloom', 'Perpetual Heart', 'Shivada Jade', NOW(), NOW()), -- 17 — Kamisato Ayaka
    (18, 'Sakura Bloom', 'Dew of Repudiation', 'Shivada Jade', NOW(), NOW()), -- 18 — Kamisato Ayato
    (19, 'Amakumo Fruit', 'Evergloom Ring', 'Nagadus Emerald', NOW(), NOW()), -- 19 — Kirara
    (20, 'Dendrobium', 'Storm Beads', 'Vajrada Amethyst', NOW(), NOW()), -- 20 — Kujou Sara
    (21, 'Naku Weed', 'Runic Fang', 'Vajrada Amethyst', NOW(), NOW()), -- 21 — Kuki Shinobu
    (22, 'Amakumo Fruit', 'Storm Beads', 'Vajrada Amethyst', NOW(), NOW()), -- 22 — Raiden Shogun
    (23, 'Sango Pearl', 'Dew of Repudiation', 'Varunada Lazurite', NOW(), NOW()), -- 23 — Sangonomiya Kokomi
    (24, 'Crystal Marrow', 'Marionette Core', 'Vayuda Turquoise', NOW(), NOW()), -- 24 — Sayu
    (25, 'Onikabuto', 'Runic Fang', 'Vayuda Turquoise', NOW(), NOW()), -- 25 — Shikanoin Heizou
    (26, 'Fluorescent Fungus', 'Smoldering Pearl', 'Agnidus Agate', NOW(), NOW()), -- 26 — Thoma
    (27, 'Sea Ganoderma', 'Dragonheir’s False Fin', 'Vajrada Amethyst', NOW(), NOW()), -- 27 — Yae Miko
    (28, 'Naku Weed', 'Smoldering Pearl', 'Agnidus Agate', NOW(), NOW()), -- 28 — Yoimiya
    (29, 'Sea Ganoderma', 'Talisman of the Enigmatic Land', 'Vayuda Turquoise', NOW(), NOW()), -- 29 — Yumemizuki Mizuki
    (30, 'Sand Grease Pupa', 'Pseudo-Stamens', 'Nagadus Emerald', NOW(), NOW()), -- 30 — Alhaitham
    (31, 'Henna Berry', 'Light Guiding Tetrahedron', 'Varunada Lazurite', NOW(), NOW()), -- 31 — Candace
    (32, 'Rukkhashava Mushroom', 'Majestic Hooked Beak', 'Nagadus Emerald', NOW(), NOW()), -- 32 — Collei
    (33, 'Scarab', 'Thunderclap Fruitcore', 'Vajrada Amethyst', NOW(), NOW()), -- 33 — Cyno
    (34, 'Sand Grease Pupa', 'Light Guiding Tetrahedron', 'Agnidus Agate', NOW(), NOW()), -- 34 — Dehya
    (35, 'Kalpalata Lotus', 'Thunderclap Fruitcore', 'Vajrada Amethyst', NOW(), NOW()), -- 35 — Dori
    (36, 'Henna Berry', 'Light Guiding Tetrahedron', 'Vayuda Turquoise', NOW(), NOW()), -- 36 — Faruzan
    (37, 'Mourning Flower', 'Quelled Creeper', 'Nagadus Emerald', NOW(), NOW()), -- 37 — Kaveh
    (38, 'Nilotpala Lotus', 'Perpetual Caliber', 'Shivada Jade', NOW(), NOW()), -- 38 — Layla
    (39, 'Kalpalata Lotus', 'Quelled Creeper', 'Nagadus Emerald', NOW(), NOW()), -- 39 — Nahida
    (40, 'Padisarah', 'Perpetual Caliber', 'Varunada Lazurite', NOW(), NOW()), -- 40 — Nilou
    (41, 'Trishiraite', 'Cloudseam Scale', 'Vajrada Amethyst', NOW(), NOW()), -- 41 — Sethos
    (42, 'Nilotpala Lotus', 'Majestic Hooked Beak', 'Nagadus Emerald', NOW(), NOW()), -- 42 — Tighnari
    (43, 'Rukkhashava Mushroom', 'Perpetual Caliber', 'Vayuda Turquoise', NOW(), NOW()), -- 43 — Wanderer
    (44, 'Violetgrass', 'Evergloom Ring', 'Nagadus Emerald', NOW(), NOW()), -- 44 — Baizhu
    (45, 'Noctilucous Jade', 'Lightning Prism', 'Vajrada Amethyst', NOW(), NOW()), -- 45 — Beidou
    (46, 'Cor Lapis', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 46 — Chongyun
    (47, 'Starconch', 'Emperor’s Resolution', 'Agnidus Agate', NOW(), NOW()), -- 47 — Gaming
    (48, 'Qingxin', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 48 — Ganyu
    (49, 'Silk Flower', 'Juvenile Jade', 'Agnidus Agate', NOW(), NOW()), -- 49 — Hu Tao
    (50, 'Cor Lapis', 'Lightning Prism', 'Vajrada Amethyst', NOW(), NOW()), -- 50 — Keqing
    (51, 'Clearwater Jade', 'Gold-Inscribed Secret Source Core', 'Vayuda Turquoise', NOW(), NOW()), -- 51 — Lan Yan
    (52, 'Glaze Lily', 'Basalt Pillar', 'Prithiva Topaz', NOW(), NOW()), -- 52 — Ningguang
    (53, 'Violetgrass', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 53 — Qiqi
    (54, 'Qingxin', 'Dragonheir’s False Fin', 'Shivada Jade', NOW(), NOW()), -- 54 — Shenhe
    (55, 'Jueyun Chili', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 55 — Xiangling
    (56, 'Pluie Lotus', 'Cloudseam Scale', 'Vayuda Turquoise', NOW(), NOW()), -- 56 — Xianyun
    (57, 'Qingxin', 'Juvenile Jade', 'Vayuda Turquoise', NOW(), NOW()), -- 57 — Xiao
    (58, 'Silk Flower', 'Cleansing Heart', 'Varunada Lazurite', NOW(), NOW()), -- 58 — Xingqiu
    (59, 'Violetgrass', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 59 — Xinyan
    (60, 'Noctilucous Jade', 'Juvenile Jade', 'Agnidus Agate', NOW(), NOW()), -- 60 — Yanfei
    (61, 'Jueyun Chili', 'Quelled Creeper', 'Nagadus Emerald', NOW(), NOW()), -- 61 — Yaoyao
    (62, 'Starconch', 'Runic Fang', 'Varunada Lazurite', NOW(), NOW()), -- 62 — Yelan
    (63, 'Glaze Lily', 'Riftborn Regalia', 'Prithiva Topaz', NOW(), NOW()), -- 63 — Yun Jin
    (64, 'Cor Lapis', 'Basalt Pillar', 'Prithiva Topaz', NOW(), NOW()), -- 64 — Zhongli
    (65, 'Cecilia', 'Basalt Pillar', 'Prithiva Topaz', NOW(), NOW()), -- 65 — Albedo
    (66, 'Small Lamp Grass', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 66 — Amber
    (67, 'Philanemo Mushroom', 'Cleansing Heart', 'Varunada Lazurite', NOW(), NOW()), -- 67 — Barbara
    (68, 'Windwheel Aster', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 68 — Bennett
    (69, 'Small Lamp Grass', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 69 — Diluc
    (70, 'Calla Lily', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 70 — Diona
    (71, 'Dandelion Seed', 'Crystalline Bloom', 'Shivada Jade', NOW(), NOW()), -- 71 — Eula
    (72, 'Small Lamp Grass', 'Lightning Prism', 'Vajrada Amethyst', NOW(), NOW()), -- 72 — Fischl
    (73, 'Dandelion Seed', 'Hurricane Seed', 'Vayuda Turquoise', NOW(), NOW()), -- 73 — Jean
    (74, 'Calla Lily', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 74 — Kaeya
    (75, 'Philanemo Mushroom', 'Everflame Seed', 'Agnidus Agate', NOW(), NOW()), -- 75 — Klee
    (76, 'Valberry', 'Lightning Prism', 'Vajrada Amethyst', NOW(), NOW()), -- 76 — Lisa
    (77, 'Wolfhook', 'Pseudostamens', 'Shivada Jade', NOW(), NOW()), -- 77 — Mika
    (78, 'Philanemo Mushroom', 'Cleansing Heart', 'Varunada Lazurite', NOW(), NOW()), -- 78 — Mona
    (79, 'Valberry', 'Basalt Pillar', 'Prithiva Topaz', NOW(), NOW()), -- 79 — Noelle
    (80, 'Wolfhook', 'Lightning Prism', 'Vajrada Amethyst', NOW(), NOW()), -- 80 — Razor
    (81, 'Valberry', 'Hoarfrost Core', 'Shivada Jade', NOW(), NOW()), -- 81 — Rosaria
    (82, 'Windwheel Aster', 'Hurricane Seed', 'Vayuda Turquoise', NOW(), NOW()), -- 82 — Sucrose
    (83, 'Cecilia', 'Hurricane Seed', 'Vayuda Turquoise', NOW(), NOW()), -- 83 — Venti
    (84, 'Rainbow Rose', 'Fragment of a Golden Melody', 'Agnidus Agate', NOW(), NOW()), -- 84 — Arlecchino
    (85, 'Starconch', 'Cleansing Heart', 'Varunada Lazurite', NOW(), NOW()), -- 85 — Tartaglia (Childe)
    (86, 'Withering Purpurbloom', 'Ensnaring Gaze', 'Vayuda Turquoise', NOW(), NOW()), -- 86 — Chasca
    (87, 'Quenepa Berry', 'Talisman of the Enigmatic Land', 'Shivada Jade', NOW(), NOW()), -- 87 — Citlali
    (88, 'Dracolite', 'Ensnaring Gaze', 'Vajrada Amethyst', NOW(), NOW()), -- 88 — Iansan
    (89, 'Quenepa Berry', 'Overripe Flamegrapefruit', 'Prithiva Topaz', NOW(), NOW()), -- 89 — Kachina
    (90, 'Saurian Claw Succulent', 'Overripe Flamegrapefruit', 'Nagadus Emerald', NOW(), NOW()), -- 90 — Kinich
    (91, 'Withering Purpurbloom', 'Gold-Inscribed Secret Source Core', 'Agnidus Agate', NOW(), NOW()), -- 91 — Mavuika
    (92, 'Sprayfeather Gill', 'Mark of the Binding Blessing', 'Varunada Lazurite', NOW(), NOW()), -- 92 — Mualani
    (93, 'Glowing Hornshroom', 'Mark of the Binding Blessing', 'Vajrada Amethyst', NOW(), NOW()), -- 93 — Ororon
    (94, 'Skysplit Gembloom', 'Sparkless Statue Core', 'Vajrada Amethyst', NOW(), NOW()), -- 94 — Varesa
    (95, 'Brilliant Chrysanthemum', 'Gold-Inscribed Secret Source Core', 'Prithiva Topaz', NOW(), NOW()), -- 95 — Xilonen
    (96, 'Crystal Marrow', 'Crystalline Bloom', 'Shivada Jade', NOW(), NOW()), -- 96 — Aloy
    (97, 'Saurian Claw Succulent', 'Sparkless Statue Core', 'Vayuda Turquoise', NOW(), NOW()), -- 97 — Ifa
    (98, 'Beryl Conch', 'Secret Source Airflow Accumulator', 'Shivada Jade', NOW(), NOW()), -- 98 — Escoffier
    (99, 'Calla Lily', 'Secret Source Airflow Accumulator', 'Varunada Lazurite', NOW(), NOW()), -- 99 — Dahlia
    (100, 'Skysplit Gembloom', 'Ensnaring Gaze', 'Shivada Jade', NOW(), NOW()), -- 100 — Skirk


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
    (1, 'Amber', 'Pyro', 'Bow', 'ATK%', 'Mondstadt', 4, 66, 1, 17, NOW(), NOW()),
    (2, 'Barbara', 'Hydro', 'Catalyst', 'HP%', 'Mondstadt', 4, 67, 2, 2, NOW(), NOW()),
    (3, 'Beidou', 'Electro', 'Claymore', 'Electro DMG Bonus%', 'Liyue', 4, 45, 3, 4, NOW(), NOW()),
    (4, 'Bennett', 'Pyro', 'Sword', 'Energy Recharge%', 'Mondstadt', 4, 68, 4, 4, NOW(), NOW()),
    (5, 'Chongyun', 'Cryo', 'Claymore', 'ATK%', 'Liyue', 4, 46, 5, 1, NOW(), NOW()),
    (6, 'Diluc', 'Pyro', 'Claymore', 'CRIT Rate%', 'Mondstadt', 5, 69, 6, 5, NOW(), NOW()),
    (7, 'Fischl', 'Electro', 'Bow', 'ATK%', 'Mondstadt', 4, 72, 7, 17, NOW(), NOW()),
    (8, 'Jean', 'Anemo', 'Sword', 'Healing Bonus%', 'Mondstadt', 5, 73, 8, 1, NOW(), NOW()),
    (9, 'Kaeya', 'Cryo', 'Sword', 'Energy Recharge%', 'Mondstadt', 4, 74, 9, 4, NOW(), NOW()),
    (10, 'Keqing', 'Electro', 'Sword', 'CRIT DMG%', 'Liyue', 5, 50, 10, 6, NOW(), NOW()),
    (11, 'Klee', 'Pyro', 'Catalyst', 'Pyro DMG Bonus%', 'Mondstadt', 5, 75, 11, 2, NOW(), NOW()),
    (12, 'Lisa', 'Electro', 'Catalyst', 'Elemental Mastery', 'Mondstadt', 4, 76, 12, 3, NOW(), NOW()),
    (13, 'Mona', 'Hydro', 'Catalyst', 'Energy Recharge%', 'Mondstadt', 5, 78, 13, 6, NOW(), NOW()),
    (14, 'Ningguang', 'Geo', 'Catalyst', 'Geo DMG Bonus%', 'Liyue', 4, 52, 14, 5, NOW(), NOW()),
    (15, 'Noelle', 'Geo', 'Claymore', 'DEF%', 'Mondstadt', 4, 79, 15, 1, NOW(), NOW()),
    (16, 'Qiqi', 'Cryo', 'Sword', 'Healing Bonus%', 'Liyue', 5, 53, 16, 2, NOW(), NOW()),
    (17, 'Razor', 'Electro', 'Claymore', 'Physical DMG Bonus%', 'Mondstadt', 4, 80, 17, 1, NOW(), NOW()),
    (18, 'Sucrose', 'Anemo', 'Catalyst', 'Anemo DMG Bonus%', 'Mondstadt', 4, 82, 18, 6, NOW(), NOW()),
    (19, 'Venti', 'Anemo', 'Bow', 'Energy Recharge%', 'Mondstadt', 5, 83, 19, 3, NOW(), NOW()),
    (20, 'Xiangling', 'Pyro', 'Polearm', 'Elemental Mastery', 'Liyue', 4, 55, 20, 3, NOW(), NOW()),
    (21, 'Xingqiu', 'Hydro', 'Sword', 'ATK%', 'Liyue', 4, 58, 21, 1, NOW(), NOW()),

    -- Patch: 1.1
    (22, 'Diona', 'Cryo', 'Bow', 'Cryo DMG Bonus%', 'Mondstadt', 4, 70, 22, 17, NOW(), NOW()),
    (23, 'Tartaglia', 'Hydro', 'Bow', 'Hydro DMG Bonus%', 'Snezhnaya', 5, 85, 23, 5, NOW(), NOW()),
    (24, 'Xinyan', 'Pyro', 'Claymore', 'ATK%', 'Liyue', 4, 59, 24, 4, NOW(), NOW()),
    (25, 'Zhongli', 'Geo', 'Polearm', 'Geo DMG Bonus%', 'Liyue', 5, 64, 25, 3, NOW(), NOW()),

    -- Patch: 1.2
    (26, 'Albedo', 'Geo', 'Polearm', 'Geo DMG Bonus%', 'Mondstadt', 5, 65, 26, 2, NOW(), NOW()),
    (27, 'Ganyu', 'Cryo', 'Bow', 'CRIT DMG%', 'Liyue', 5, 48, 27, 6, NOW(), NOW()),

    -- Patch: 1.3
    (28, 'Hu Tao', 'Pyro', 'Polearm', 'CRIT DMG%', 'Liyue', 5, 49, 28, 6, NOW(), NOW()),
    (29, 'Xiao', 'Anemo', 'Polearm', 'CRIT Rate%', 'Liyue', 5, 57, 29, 3, NOW(), NOW()),

    -- Patch: 1.4
    (30, 'Rosaria', 'Cryo', 'Polearm', 'ATK%', 'Mondstadt', 4, 81, 30, 5, NOW(), NOW()),

    -- Patch: 1.5
    (31, 'Eula', 'Cryo', 'Claymore', 'CRIT DMG%', 'Mondstadt', 5, 71, 31, 1, NOW(), NOW()),
    (32, 'Yanfei', 'Pyro', 'Catalyst', 'Pyro DMG Bonus%', 'Liyue', 4, 60, 32, 4, NOW(), NOW()),

    -- Patch: 1.6
    (33, 'Kaedehara Kazuha', 'Anemo', 'Sword', 'Elemental Mastery', 'Inazuma', 5, 16, 33, 4, NOW(), NOW()),

    -- Patch: 2.0
    (34, 'Kamisato Ayaka', 'Cryo', 'Sword', 'CRIT DMG%', 'Inazuma', 5, 17, 34, 8, NOW(), NOW()),
    (35, 'Sayu', 'Anemo', 'Claymore', 'Elemental Mastery', 'Inazuma', 4, 24, 35, 6, NOW(), NOW()),
    (36, 'Yoimiya', 'Pyro', 'Bow', 'CRIT Rate%', 'Inazuma', 5, 28, 36, 2, NOW(), NOW()),

    -- Patch: 2.1
    (37, 'Aloy', 'Cryo', 'Bow', 'Cryo DMG Bonus%', 'No region', 5, 96, 37, 9, NOW(), NOW()),
    (38, 'Kujou Sara', 'Electro', 'Bow', 'ATK%', 'Inazuma', 4, 20, 38, 1, NOW(), NOW()),
    (39, 'Raiden Shogun', 'Electro', 'Polearm', 'Energy Recharge%', 'Inazuma', 5, 22, 39, 8, NOW(), NOW()),
    (40, 'Sangonomiya Kokomi', 'Hydro', 'Catalyst', 'Hydro DMG Bonus%', 'Inazuma', 5, 23, 40, 9, NOW(), NOW()),

    -- Patch: 2.2
    (41, 'Thoma', 'Pyro', 'Polearm', 'ATK%', 'Inazuma', 4, 26, 41, 4, NOW(), NOW()),

    -- Patch: 2.3
    (42, 'Arataki Itto', 'Geo', 'Claymore', 'CRIT Rate%', 'Inazuma', 5, 13, 42, 3, NOW(), NOW()),
    (43, 'Gorou', 'Geo', 'Bow', 'Geo DMG Bonus%', 'Inazuma', 4, 15, 43, 9, NOW(), NOW()),

    -- Patch: 2.4
    (44, 'Shenhe', 'Cryo', 'Polearm', 'ATK%', 'Liyue', 5, 54, 44, 6, NOW(), NOW()),
    (45, 'Yun Jin', 'Geo', 'Polearm', 'Energy Recharge%', 'Liyue', 4, 63, 45, 1, NOW(), NOW()),

    -- Patch: 2.5
    (46, 'Yae Miko', 'Electro', 'Catalyst', 'CRIT Rate%', 'Inazuma', 5, 27, 46, 8, NOW(), NOW()),

    -- Patch: 2.6
    (47, 'Kamisato Ayato', 'Hydro', 'Sword', 'CRIT DMG%', 'Inazuma', 5, 18, 47, 8, NOW(), NOW()),

    -- Patch: 2.7
    (48, 'Kuki Shinobu', 'Electro', 'Sword', 'HP%', 'Inazuma', 4, 21, 48, 9, NOW(), NOW()),
    (49, 'Yelan', 'Hydro', 'Bow', 'CRIT Rate%', 'Liyue', 5, 62, 49, 5, NOW(), NOW()),

    -- Patch: 2.8
    (50, 'Shikanoin Heizou', 'Anemo', 'Catalyst', 'Anemo DMG Bonus%', 'Inazuma', 4, 25, 50, 4, NOW(), NOW()),

    -- Patch: 3.0
    (51, 'Collei', 'Dendro', 'Bow', 'ATK%', 'Sumeru', 4, 32, 51, 17, NOW(), NOW()),
    (52, 'Dori', 'Electro', 'Claymore', 'HP%', 'Sumeru', 4, 35, 52, 11, NOW(), NOW()),
    (53, 'Tighnari', 'Dendro', 'Bow', 'Dendro DMG Bonus%', 'Sumeru', 5, 42, 53, 4, NOW(), NOW()),

    -- Patch: 3.1
    (54, 'Candace', 'Hydro', 'Polearm', 'HP%', 'Sumeru', 4, 31, 54, 11, NOW(), NOW()),
    (55, 'Cyno', 'Electro', 'Polearm', 'CRIT DMG%', 'Sumeru', 5, 33, 55, 2, NOW(), NOW()),
    (56, 'Nilou', 'Hydro', 'Sword', 'HP%', 'Sumeru', 5, 40, 56, 10, NOW(), NOW()),

    -- Patch: 3.2
    (57, 'Layla', 'Cryo', 'Sword', 'HP%', 'Sumeru', 4, 38, 57, 2, NOW(), NOW()),
    (58, 'Nahida', 'Dendro', 'Catalyst', 'Elemental Mastery', 'Sumeru', 5, 39, 58, 10, NOW(), NOW()),

    -- Patch: 3.3
    (59, 'Faruzan', 'Anemo', 'Bow', 'ATK%', 'Sumeru', 4, 36, 59, 11, NOW(), NOW()),
    (60, 'Wanderer', 'Anemo', 'Catalyst', 'CRIT Rate%', 'Sumeru', 5, 43, 60, 8, NOW(), NOW()),

    -- Patch: 3.4
    (61, 'Alhaitham', 'Dendro', 'Sword', 'Dendro DMG Bonus%', 'Sumeru', 5, 30, 61, 11, NOW(), NOW()),
    (62, 'Yaoyao', 'Dendro', 'Polearm', 'HP%', 'Liyue', 4, 61, 62, 3, NOW(), NOW()),

    -- Patch: 3.5
    (63, 'Dehya', 'Pyro', 'Claymore', 'HP%', 'Sumeru', 5, 34, 63, 11, NOW(), NOW()),
    (64, 'Mika', 'Cryo', 'Polearm', 'HP%', 'Mondstadt', 4, 77, 64, 5, NOW(), NOW()),

    -- Patch: 3.6
    (65, 'Baizhu', 'Dendro', 'Catalyst', 'HP%', 'Liyue', 5, 44, 65, 10, NOW(), NOW()),
    (66, 'Kaveh', 'Dendro', 'Claymore', 'HP%', 'Sumeru', 4, 37, 66, 10, NOW(), NOW()),

    -- Patch: 3.7
    (67, 'Kirara', 'Dendro', 'Sword', 'HP%', 'Inazuma', 4, 19, 67, 9, NOW(), NOW()),

    -- Patch: 4.0
    (68, 'Freminet', 'Cryo', 'Claymore', 'ATK%', 'Fontaine', 4, 5, 68, 13, NOW(), NOW()),
    (69, 'Lynette', 'Anemo', 'Sword', 'Anemo DMG Bonus%', 'Fontaine', 4, 7, 69, 12, NOW(), NOW()),
    (70, 'Lyney', 'Pyro', 'Bow', 'CRIT Rate%', 'Fontaine', 5, 8, 70, 5, NOW(), NOW()),

    -- Patch: 4.1
    (71, 'Neuvillette', 'Hydro', 'Catalyst', 'CRIT DMG%', 'Fontaine', 5, 10, 71, 13, NOW(), NOW()),
    (72, 'Wriothesley', 'Cryo', 'Catalyst', 'CRIT DMG%', 'Fontaine', 5, 12, 72, 12, NOW(), NOW()),

    -- Patch: 4.2
    (73, 'Charlotte', 'Cryo', 'Catalyst', 'ATK%', 'Fontaine', 4, 1, 73, 12, NOW(), NOW()),
    (74, 'Furina', 'Hydro', 'Sword', 'CRIT Rate%', 'Fontaine', 5, 6, 74, 6, NOW(), NOW()),

    -- Patch: 4.3
    (75, 'Chevreuse', 'Pyro', 'Polearm', 'HP%', 'Fontaine', 4, 2, 75, 12, NOW(), NOW()),
    (76, 'Navia', 'Geo', 'Claymore', 'CRIT DMG%', 'Fontaine', 5, 9, 76, 13, NOW(), NOW()),

    -- Patch: 4.4
    (77, 'Gaming', 'Pyro', 'Claymore', 'ATK%', 'Liyue', 4, 47, 77, 3, NOW(), NOW()),
    (78, 'Xianyun', 'Anemo', 'Catalyst', 'ATK%', 'Liyue', 5, 56, 78, 2, NOW(), NOW()),

    -- Patch: 4.5
    (79, 'Chiori', 'Geo', 'Sword', 'CRIT Rate%', 'Inazuma', 5, 14, 79, 9, NOW(), NOW()),

    -- Patch: 4.6
    (80, 'Arlecchino', 'Pyro', 'Polearm', 'CRIT DMG%', 'Snezhnaya', 5, 84, 80, 5, NOW(), NOW()),

    -- Patch: 4.7
    (81, 'Clorinde', 'Electro', 'Sword', 'CRIT Rate%', 'Fontaine', 5, 3, 81, 13, NOW(), NOW()),
    (82, 'Sethos', 'Electro', 'Bow', 'Elemental Mastery', 'Sumeru', 4, 41, 82, 11, NOW(), NOW()),
    (83, 'Sigewinne', 'Hydro', 'Bow', 'HP%', 'Fontaine', 5, 11, 83, 13, NOW(), NOW()),

    -- Patch: 4.8
    (84, 'Emilie', 'Dendro', 'Polearm', 'CRIT DMG%', 'Fontaine', 5, 4, 84, 12, NOW(), NOW()),

    -- Patch: 5.0
    (85, 'Kachina', 'Geo', 'Polearm', 'Geo DMG Bonus%', 'Natlan', 4, 89, 85, 15, NOW(), NOW()),
    (86, 'Kinich', 'Dendro', 'Claymore', 'CRIT DMG%', 'Natlan', 5, 90, 86, 16, NOW(), NOW()),
    (87, 'Mualani', 'Hydro', 'Catalyst', 'CRIT Rate%', 'Natlan', 5, 92, 87, 15, NOW(), NOW()),

    -- Patch: 5.1
    (88, 'Xilonen', 'Geo', 'Sword', 'DEF%', 'Natlan', 5, 95, 88, 15, NOW(), NOW()),

    -- Patch: 5.2
    (89, 'Chasca', 'Anemo', 'Bow', 'CRIT Rate%', 'Natlan', 5, 86, 89, 16, NOW(), NOW()),
    (90, 'Ororon', 'Electro', 'Bow', 'ATK%', 'Natlan', 4, 93, 90, 16, NOW(), NOW()),

    -- Patch: 5.3
    (91, 'Citlali', 'Cryo', 'Catalyst', 'Elemental Mastery', 'Natlan', 5, 87, 91, 16, NOW(), NOW()),
    (92, 'Lan Yan', 'Anemo', 'Catalyst', 'ATK%', 'Liyue', 4, 51, 92, 6, NOW(), NOW()),
    (93, 'Mavuika', 'Pyro', 'Claymore', 'CRIT DMG%', 'Natlan', 5, 91, 93, 15, NOW(), NOW()),

    -- Patch: 5.4
    (94, 'Yumemizuki Mizuki', 'Anemo', 'Catalyst', 'Elemental Mastery', 'Inazuma', 5, 29, 94, 8, NOW(), NOW()),

    -- Patch: 5.5
    (95, 'Iansan', 'Electro', 'Polearm', 'ATK%', 'Natlan', 4, 88, 95, 15, NOW(), NOW()),
    (96, 'Varesa', 'Electro', 'Catalyst', 'CRIT Rate%', 'Natlan', 5, 94, 96, 16, NOW(), NOW()),

    -- Patch: 5.6
    (97, 'Ifa', 'Anemo', 'Catalyst', 'Elemental Mastery', 'Natlan', 4, 97, 97, 16, NOW(), NOW()),
    (98, 'Escoffier', 'Cryo', 'Polearm', 'CRIT Rate%', 'Fontaine', 5, 98, 98, 12, NOW(), NOW()),

    -- Patch: 5.7
    (99, 'Dahlia', 'Hydro', 'Sword', 'ATK%', 'Mondstadt', 4, 99, 99, 17, NOW(), NOW()),
    (100, 'Skirk', 'Cryo', 'Sword', 'ATK%', '?', 5, 100, 100, 12, NOW(), NOW()),
