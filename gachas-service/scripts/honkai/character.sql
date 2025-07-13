INSERT INTO honkai_resources (id, uncommon, rare, epic, created_at, updated_at)
VALUES 
    (1, 'Extinguished Core', 'Glimmering Core', 'Squirming Core', NOW(), NOW()),
    (2, 'Thief''s Instinct', 'Usurper''s Scheme', 'Conqueror''s Will', NOW(), NOW()),
    (3, 'Silvermane Badge', 'Silvermane Insignia', 'Silvermane Medal', NOW(), NOW()),
    (4, 'Ancient Part', 'Ancient Spindle', 'Ancient Engine', NOW(), NOW()),
    (5, 'Immortal Scionette', 'Immortal Aeroblossom', 'Immortal Lumintwig', NOW(), NOW()),
    (6, 'Artifex''s Module', 'Artifex''s Cogwheel', 'Artifex''s Gyreheart', NOW(), NOW()),
    (7, 'Dream Collection Component', 'Dream Flow Valve', 'Dream Making Engine', NOW(), NOW()),
    (8, 'Tatters of Thought', 'Fragments of Impression', 'Shards of Desires', NOW(), NOW()),
    (9, 'Fear-Stomped Flesh', 'Courage-Torn Chest', 'Glory-Aspersed Torso', NOW(), NOW()),
    (10, 'Ethereal Omen', 'Echoing Wail', 'Eternal Lament', NOW(), NOW()),
    (11, 'Shattered Blade', 'Lifeless Blade', 'Worldbreaker Blade', NOW(), NOW()),
    (12, 'Borisin Teeth', 'Lupitoxin Sawteeth', 'Moon Rage Fang', NOW(), NOW()),
    (13, 'Arrow of the Beast Hunter', 'Arrow of the Demon Slayer', 'Arrow of the Starchaser', NOW(), NOW()),
    (14, 'Meteoric Bullet', 'Destined Expiration', 'Countertemporal Shot', NOW(), NOW()),
    (15, 'Key of Inspiration', 'Key of Knowledge', 'Key of Wisdom', NOW(), NOW()),
    (16, 'Rough Sketch', 'Dynamic Outlining', 'Exquisite Colored Draft', NOW(), NOW()),
    (17, 'Endurance of Bronze', 'Oath of Steel', 'Safeguard of Amber', NOW(), NOW()),
    (18, 'Scattered Stardust', 'Crystal Meteorites', 'Divine Amber', NOW(), NOW()),
    (19, 'Obsidian of Dread', 'Obsidian of Desolation', 'Obsidian of Obsession', NOW(), NOW()),
    (20, 'Fiery Spirit', 'Starfire Essence', 'Heaven Incinerator', NOW(), NOW()),
    (21, 'Harmonic Tune', 'Ancestral Hymn', 'Stellaris Symphony', NOW(), NOW()),
    (22, 'Firmament Note', 'Celestial Section', 'Heavenly Melody', NOW(), NOW()),
    (23, 'Seed of Abundance', 'Sprout of Life', 'Flower of Eternity', NOW(), NOW()),
    (24, 'Alien Tree Seed', 'Nourishing Honey', 'Myriad Fruit', NOW(), NOW()),
    (25, 'Bīja of Consciousness', 'Seedling of Manas', 'Flower of Ālaya', NOW(), NOW());

INSERT INTO honkai_ascensions (id, boss_material, resource_id, created_at, updated_at)
VALUES 
    (1, 'Lightning Crown of the Past Shadow', 1, NOW(), NOW()), -- Arlan, Bailu
    (2, 'Endotherm Chitin', 3, NOW(), NOW()), -- Asta
    (3, 'Storm Eye', 3, NOW(), NOW()), -- Bronya
    (4, 'Broken Teeth of Iron Wolf', 4, NOW(), NOW()), -- Clara, Natasha, Luka
    (5, 'Storm Eye', 1, NOW(), NOW()), -- Dan heng
    (6, 'Safeguard of Amber', 17, NOW(), NOW()), -- Gepard
    (7, 'Horn of Snow', 1, NOW(), NOW()), -- Herta, Pela
    (8, 'Endotherm Chitin', 1, NOW(), NOW()), -- Himeko
    (9, 'Endotherm Chitin', 4, NOW(), NOW()), -- Hook
    (10, 'Shape Shifter''s Lightning Staff', 5, NOW(), NOW()), -- Jing Yuan
    (11, 'Horn of Snow', 2, NOW(), NOW()), -- March 7th
    (12, 'Void Cast Iron', 2, NOW(), NOW()), -- Qingque, Seele
    (13, 'Storm Eye', 4, NOW(), NOW()), -- Sampo
    (14, 'Lightning Crown of the Past Shadow', 3, NOW(), NOW()), -- Serval
    (15, 'Broken Teeth of Iron Wolf', 6, NOW(), NOW()), -- Sushang
    (16, 'Lightning Crown of the Past Shadow', 5, NOW(), NOW()), -- Tingyun
    (17, 'Enigmatic Ectostella', 2, NOW(), NOW()), -- Trailblazer
    (18, 'Golden Crown of the Past Shadow', 3, NOW(), NOW()), -- Welt
    (19, 'Gelid Chitin', 2, NOW(), NOW()), -- Yanging
    (20, 'Golden Crown of the Past Shadow', 6, NOW(), NOW()), -- Luocha, Yukong
    (21, 'Void Cast Iron', 4, NOW(), NOW()), -- Silver Wolf
    (22, 'Ascendant Debris', 5, NOW(), NOW()), -- Blade
    (23, 'Shape Shifter''s Lightning Staff', 2, NOW(), NOW()), -- Kafka
    (24, 'Broken Teeth of Iron Wolf', 4, NOW(), NOW()), -- Luka
    (25, 'Suppressing Edict', 5, NOW(), NOW()), -- Imbibitor Lunae
    (26, 'Nail of the Ape', 6, NOW(), NOW()), -- Fuxuan
    (27, "Nail of the Ape", 1, NOW(), NOW()), -- Lynx, Xuyei  
    (28, "Searing Steel Blade", 6, NOW(), NOW()), -- Guinafen
    (29, "Gelid Chitin", 5, NOW(), NOW()), -- Jingliu
    (30, "Searing Steel Blade", 3, NOW(), NOW()), -- Topaz
    (31, "Netherworld Token", 1, NOW(), NOW()), -- Argenti
    (32, "Netherworld Token", 6, NOW(), NOW()), -- Hanya
    (33, "Ascendant Debris", 5, NOW(), NOW()), -- HuoHuo
    (34, "Suppressing Edict", 2, NOW(), NOW()), -- Dr. Ratio
    (35, "Gelid Chitin", 5, NOW(), NOW()), -- Ruan Mei
    (36, "Ascendant Debris", 1, NOW(), NOW()), -- Black Swan
    (37, "Dream Fridge", 4, NOW(), NOW()), -- Misha
    (38, "Dream Flamer", 8, NOW(), NOW()), -- Sparkle
    (39, "Shape Shifter's Lightning Staff", 4, NOW(), NOW()), -- Acheron
    (40, "Suppressing Edict", 8, NOW(), NOW()), -- Aventurine
    (41, "Raging Heart", 4, NOW(), NOW()), -- Gallagher
    (42, "IPC Work Permit", 8, NOW(), NOW()), -- Boothill
    (43, "IPC Work Permit", 4, NOW(), NOW()), -- Robin
    (44, "Raging Heart", 8, NOW(), NOW()), -- Firefly
    (45, "Dream Flamer", 4, NOW(), NOW()), -- Jade
    (46, "Raging Heart", 5, NOW(), NOW()), -- Jiaoqiu, Lingsha
    (47, "IPC Work Permit", 6, NOW(), NOW()), -- Yunli
    (48, "A Glass of the Besotted Era", 6, NOW(), NOW()), -- Feixiao
    (49, "Nail of the Beast Coffin", 6, NOW(), NOW()), -- Moze
    (50, "Chordal Mirage", 4, NOW(), NOW()), -- Rappa
    (51, "Searing Steel Blade", 5, NOW(), NOW()), -- Fugue
    (52, "Chordal Mirage", 8, NOW(), NOW()), -- Sunday
    (53, "Nail of the Beast Coffin", 9, NOW(), NOW()), -- Aglaea
    (54, "Dream Fridge", 1, NOW(), NOW()), -- The Herta
    (55, "Harbinger of Strife", 9, NOW(), NOW()), -- Mydei
    (56, "Darkveil Moonlight", 9, NOW(), NOW()), -- Tribbie
    (57, "A Glass of the Besotted Era", 10, NOW(), NOW()), -- Anaxa
    (58, "Darkveil Moonlight", 10, NOW(), NOW()), -- Castorice, Cipher
    (59, "A Glass of the Besotted Era", 9, NOW(), NOW()), -- Hyacine
    (60, "Darkveil Moonlight", 1, NOW(), NOW()), -- Archer
    (61, "Invasive Clot", 10, NOW(), NOW()), -- Phainon
    (62, "A Glass of the Besotted Era", 2, NOW(), NOW()); -- Saber
    
    
    



INSERT INTO honkai_traces (id, dungeon_resource_id, mob_resource_id, boss_material, created_at, updated_at)
VALUES 
    (1, 11, 1, 'Destroyer''s Final Road', NOW(), NOW()),
    (2, 21, 3, 'Destroyer''s Final Road', NOW(), NOW()),
    (3, 23, 1, 'Guardian''s Lament', NOW(), NOW()),
    (4, 21, 3, 'Guardian''s Lament', NOW(), NOW()),
    (5, 11, 4, 'Guardian''s Lament', NOW(), NOW()),
    (6, 13, 1, 'Destroyer''s Final Road', NOW(), NOW()),
    (7, 17, 3, 'Guardian''s Lament', NOW(), NOW()),
    (8, 15, 1, 'Destroyer''s Final Road', NOW(), NOW()),
    (9, 15, 5, 'Guardian''s Lament', NOW(), NOW()),
    (10, 17, 2, 'Destroyer''s Final Road', NOW(), NOW()),
    (11, 23, 4, 'Guardian''s Lament', NOW(), NOW()),
    (12, 19, 1, 'Guardian''s Lament', NOW(), NOW()),
    (13, 15, 2, 'Guardian''s Lament', NOW(), NOW()),
    (14, 19, 4, 'Guardian''s Lament', NOW(), NOW()),
    (15, 13, 2, 'Guardian''s Lament', NOW(), NOW()),
    (16, 15, 3, 'Guardian''s Lament', NOW(), NOW()),
    (17, 13, 6, 'Guardian''s Lament', NOW(), NOW()),
    (18, 21, 5, 'Destroyer''s Final Road', NOW(), NOW()),
    (19, 11, 2, 'Destroyer''s Final Road', NOW(), NOW()),
    (20, 17, 2, 'Destroyer''s Final Road', NOW(), NOW()),
    (21, 19, 3, 'Destroyer''s Final Road', NOW(), NOW()),
    (22, 23, 6, 'Guardian''s Lament', NOW(), NOW()),
    (23, 19, 4, 'Destroyer''s Final Road', NOW(), NOW()),
    (24, 21, 6, 'Destroyer''s Final Road', NOW(), NOW()),
    (25, 11, 5, 'Regret of Infinite Ochema', NOW(), NOW()),
    (26, 19, 2, 'Regret of Infinite Ochema', NOW(), NOW()),
    (27, 19, 4, 'Regret of Infinite Ochema', NOW(), NOW()),
    (28, 11, 5, 'Regret of Infinite Ochema', NOW(), NOW()),
    (29, 17, 6, 'Regret of Infinite Ochema', NOW(), NOW()),
    (30, 23, 1, 'Regret of Infinite Ochema', NOW(), NOW()),
    (31, 19, 6, 'Regret of Infinite Ochema', NOW(), NOW()),
    (32, 13, 3, 'Regret of Infinite Ochema', NOW(), NOW()),
    (33, 15, 1, 'Regret of Infinite Ochema', NOW(), NOW()),
    (34, 21, 6, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (35, 23, 5, 'Regret of Infinite Ochema', NOW(), NOW()),
    (36, 13, 2, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (37, 21, 5, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (38, 11, 1, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (39, 20, 1, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (40, 12, 4, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (41, 22, 8, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (42, 20, 4, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (43, 18, 8, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (44, 24, 4, 'Past Evils of the Borehole Planet Disaster', NOW(), NOW()),
    (45, 14, 8, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (46, 22, 4, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (47, 22, 2, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (48, 12, 8, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (49, 16, 4, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (50, 20, 5, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (51, 14, 2, 'Destroyer''s Final Road', NOW(), NOW()),
    (52, 12, 6, 'Regret of Infinite Ochema', NOW(), NOW()),
    (53, 14, 6, 'Regret of Infinite Ochema', NOW(), NOW()),
    (54, 24, 5, 'Auspice Sliver', NOW(), NOW()),
    (55, 14, 5, 'Auspice Sliver', NOW(), NOW()),
    (56, 20, 5, 'Regret of Infinite Ochema', NOW(), NOW()),
    (57, 22, 8, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (58, 25, 9, 'Auspice Sliver', NOW(), NOW()),
    (59, 16, 1, 'Auspice Sliver', NOW(), NOW()),
    (60, 25, 2, 'Auspice Sliver', NOW(), NOW()),
    (61, 12, 9, 'Auspice Sliver', NOW(), NOW()),
    (62, 22, 9, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (63, 16, 10, 'Lost Echo of the Shared Wish', NOW(), NOW()),
    (64, 25, 10, 'Auspice Sliver', NOW(), NOW()),
    (65, 20, 10, 'Daythunder Anamnesis',  NOW(), NOW()),
    (66, 25, 9, 'Lost Echo of the Shared Wish',  NOW(), NOW()),
    (67, 14, 1, 'Guardian''s Lament',  NOW(), NOW()),
    (68, 12, 10, 'Daythunder Anamnesis',  NOW(), NOW()),
    (69, 12, 2, 'Destroyer''s Final Road',  NOW(), NOW());

    

INSERT INTO honkai_characters (id, name, rarity, path, attribute, ascension_id, traces_id, created_at, updated_at)
VALUES
    (1, 'Arlan', 4, 'Destruction', 'Lightning', 1, 1, NOW(), NOW()),
    (2, 'Asta', 4, 'Harmony', 'Fire', 2, 2, NOW(), NOW()),
    (3, 'Bailu', 5, 'Abundance', 'Lightning', 1, 3, NULL, NOW(), NOW()),
    (4, 'Bronya', 5, 'Harmony', 'Wind', 3, 4, NOW(), NOW()),
    (5, 'Clara', 5, 'Destruction', 'Physical', 4, 5, NOW(), NOW()),
    (6, 'Dan Heng', 4, 'The Hunt', 'Wind', 5, 6, NOW(), NOW()),
    (7, 'Gepard', 5, 'Preservation', 'Ice', 6, 7, NOW(), NOW()),
    (8, 'Herta', 4, 'Erudition', 'Ice', 7, 8, NOW(), NOW()),
    (9, 'Himeko', 5, 'Erudition', 'Fire', 8, 8, NOW(), NOW()),
    (10, 'Hook', 4, 'Destruction', 'Fire', 9, 5, NOW(), NOW()),
    (11, 'Jing Yuan', 5, 'Erudition', 'Lightning', 10, 9, NOW(), NOW()),
    (12, 'March 7th', 4, 'Preservation', 'Ice', 11, 10, NOW(), NOW()),
    (13, 'Natasha', 4, 'Abundance', 'Physical', 4, 11, NOW(), NOW()),
    (14, 'Pela', 4, 'Nihility', 'Ice', 7, 12, NOW(), NOW()),
    (15, 'Qingque', 4, 'Erudition', 'Quantum', 12, 13, NOW(), NOW()),
    (16, 'Sampo', 4, 'Nihility', 'Wind', 13, 14, NOW(), NOW()),
    (17, 'Seele', 5, 'The Hunt', 'Quantum', 12, 15, NOW(), NOW()),
    (18, 'Serval', 4, 'Erudition', 'Lightning', 14, 16, NOW(), NOW()),
    (19, 'Sushang', 4, 'The Hunt', 'Physical', 15, 17, NOW(), NOW()),
    (20, 'Tingyun', 4, 'Harmony', 'Lightning', 16, 18, NOW(), NOW()),
    (21, 'Trailblazer', 5, 'Destruction', 'Physical', 17, 19, NOW(), NOW()),
    (22, 'Trailblazer', 5, 'Preservation', 'Fire', 17, 20, NOW(), NOW()),
    (23, 'Welt', 5, 'Nihility', 'Imaginary', 18, 21, NOW(), NOW()),
    (24, 'Yanqing', 5, 'The Hunt', 'Ice', 19, 15, NOW(), NOW()),
    (25, 'Luocha', 5, 'Abundance', 'Imaginary', 20, 22, NOW(), NOW()),
    (26, 'Silver Wolf', 5, 'Nihility', 'Quantum', 21, 23, NOW(), NOW()),
    (27, 'Yukong', 4, 'Harmony', 'Imaginary', 20, 24, NOW(), NOW()),
    (28, 'Blade', 5, 'Destruction', 'Wind', 21, 25, NOW(), NOW()),
    (29, 'Kafka', 5, 'Nihility', 'Lightning', 23, 26, NOW(), NOW()),
    (30, 'Luka', 4, 'Nihility', 'Physical', 4, 27, NOW(), NOW()),
    (31, 'Dan Heng • Imbibitor Lunae', 5, 'Destruction', 'Imaginary', 25, 28, NOW(), NOW()),
    (32, 'Fu Xuan', 5, 'Preservation', 'Quantum', 26, 29, NOW(), NOW()),
    (33, 'Lynx', 4, 'Abundance', 'Quantum', 27, 30, NOW(), NOW()),
    (34, 'Guinaifen', 4, 'Nihility', 'Fire', 28, 31, NOW(), NOW()),
    (35, 'Jingliu', 5, 'Destruction', 'Ice', 29, 28, NOW(), NOW()),
    (36, 'Topaz & Numby', 5, 'The Hunt', 'Fire', 30, 32, NOW(), NOW()),
    (37, 'Argenti', 5, 'Erudition', 'Physical', 31, 33, NOW(), NOW()),
    (38, 'Hanya', 4, 'Harmony', 'Physical', 32, 34, NOW(), NOW()),
    (39, 'Huohuo', 5, 'Abundance', 'Wind', 33, 35, NOW(), NOW()),
    (40, 'Dr. Ratio', 5, 'The Hunt', 'Imaginary', 34, 36, NOW(), NOW()),
    (41, 'Ruan Mei', 5, 'Harmony', 'Ice', 35, 37, NOW(), NOW()),
    (42, 'Xueyi', 4, 'Destruction', 'Quantum', 27, 38, NOW(), NOW()),
    (43, 'Black Swan', 5, 'Nihility', 'Wind', 36, 39, NOW(), NOW()),
    (44, 'Misha', 4, 'Destruction', 'Ice', 37, 40, NOW(), NOW()),
    (45, 'Sparkle', 5, 'Harmony', 'Quantum', 38, 41, NOW(), NOW()),
    (46, 'Acheron', 5, 'Nihility', 'Lightning', 39, 42, NOW(), NOW()),
    (47, 'Aventurine', 5, 'Preservation', 'Imaginary', 40, 43, NOW(), NOW()),
    (48, 'Gallagher', 4, 'Abundance', 'Fire', 41, 44, NOW(), NOW()),
    (49, 'Boothill', 5, 'The Hunt', 'Physical', 42, 45, NOW(), NOW()),
    (50, 'Robin', 5, 'Harmony', 'Physical', 43, 46, NOW(), NOW()),
    (51, 'Trailblazer', 5, 'Harmony', 'Imaginary', 17, 47, NOW(), NOW()),
    (52, 'Firefly', 5, 'Destruction', 'Fire', 44, 48, NOW(), NOW()),
    (53, 'Jade', 5, 'Erudition', 'Quantum', 45, 49, NOW(), NOW()),
    (54, 'Jiaoqiu', 5, 'Nihility', 'Fire', 46, 50, NOW(), NOW()),
    (55, 'March 7th', 4, 'The Hunt', 'Imaginary', 11, 51, NOW(), NOW()),
    (56, 'Yunli', 5, 'Destruction', 'Physical', 47, 52, NOW(), NOW()),
    (57, 'Feixiao', 5, 'The Hunt', 'Wind', 48, 53, NOW(), NOW()),
    (58, 'Lingsha', 5, 'Abundance', 'Fire', 46, 54, NOW(), NOW()),
    (59, 'Moze', 4, 'The Hunt', 'Lightning', 49, 55, NOW(), NOW()),
    (60, 'Rappa', 5, 'Erudition', 'Imaginary', 50, 49, NOW(), NOW()),
    (61, 'Fugue', 5, 'Nihility', 'Fire', 51, 50, NOW(), NOW()),
    (62, 'Sunday', 5, 'Harmony', 'Imaginary', 52, 57, NOW(), NOW()),
    (63, 'Aglaea', 5, 'Remembrance', 'Lightning', 53, 58, NOW(), NOW()),
    (64, 'The Herta', 5, 'Erudition', 'Ice', 54, 59, NOW(), NOW()),
    (65, 'Trailblazer', 5, 'Remembrance', 'Ice', 17, 60, NOW(), NOW()),
    (66, 'Mydei', 5, 'Destruction', 'Imaginary', 55, 61, NOW(), NOW()),
    (67, 'Tribbie', 5, 'Harmony', 'Quantum', 56, 62, NOW(), NOW()),
    (68, 'Anaxa', 5, 'Erudition', 'Wind', 57, 63, NOW(), NOW()),
    (69, 'Castorice', 5, 'Remembrance', 'Quantum', 58, 64, NOW(), NOW()),
    (70, 'Cipher', 5, 'Nihility', 'Quantum', 58, 65, NOW(), NOW()),
    (71, 'Hyacine', 5, 'Remembrance', 'Wind', 59, 66, NOW(), NOW()),
    (72, 'Archer', 5, 'Hunt', 'Quantum', 60, 67, NOW(), NOW()),
    (73, 'Phainon', 5, 'Destruction', 'Physical', 61, 68, NOW(), NOW()),
    (74, 'Saber', 5, 'Destruction', 'Wind', 62, 69, NOW(), NOW());