INSERT INTO zenless_materials (id, common, rare, epic, created_at, updated_at)
VALUES 
  -- Talents Material
  (1, 'Basic Physical Chip', 'Advanced Physical Chip', 'Specialized Physical Chip' ,NOW(), NOW()),
  (2, 'Basic Burn Chip', 'Advanced Burn Chip', 'Specialized Burn Chip', NOW(), NOW()),
  (3, 'Basic Freeze Chip', 'Advanced Freeze Chip', 'Specialized Freeze Chip', NOW(), NOW()),
  (4, 'Basic Shock Chip', 'Advanced Shock Chip', 'Specialized Shock Chip', NOW(), NOW()),
  (5, 'Basic Ether Chip', 'Advanced Ether Chip', 'Specialized Ether Chip', NOW(), NOW()),

  -- Agent Promotion Material
  (6, 'Basic Attack Certification Seal', 'Advanced Attack Certification Seal', 'Pioneer''s Certification Seal', NOW(), NOW()),
  (7, 'Basic Stun Certification Seal', 'Advanced Stun Certification Seal', 'Buster Certification Seal', NOW(), NOW()),
  (8, 'Basic Anomaly Certification Seal', 'Advanced Anomaly Certification Seal', 'Controller Certification Seal', NOW(), NOW()),
  (9, 'Basic Defense Certification Seal', 'Advanced Defense Certification Seal', 'Defender Certification Seal', NOW(), NOW()),
  (10, 'Basic Support Certification Seal', '	Advanced Support Certification Seal', 'Ruler Certification Seal', NOW(), NOW()),
  (11, 'Basic Rupture Certification Seal', 'Advanced Rupture Certification Seal', 'Arbiter Certification Seal', NOW(), NOW()),
  (12, 'Basic Rupture Certification Seal', 'Advanced Rupture Certification Seal', 'Arbiter Certification Seal', NOW(), NOW());

INSERT INTO zenless_nodes (id, hunt_boss_material, expert_boss_material, resource_id, created_at, updated_at)
VALUES 
  (1, 'Finale Dance Shoes', 'Higher Dimensional Data: Destructive Advance', 4, NOW(), NOW()), -- Alexandrina Sebastiane
  (2, 'Ferocious Grip', 'Higher Dimensional Data: Murderous Obituary', 4, NOW(), NOW()), -- Anby Demara
  (3, 'Living Drive', 'Higher Dimensional Data: Destructive Advance', 4, NOW(), NOW()), -- Anton Ivanov
  (4, 'Sycophant''s Refinement', 'Higher Dimensional Data: Thunderous Dragon', 4, NOW(), NOW()), -- Asaba Harumasa
  (5, 'Finale Dance Shoes', 'Higher Dimensional Data: Thunderous Dragon', 5, NOW(), NOW()), -- Astra Yao
  (6, 'Living Drive', 'Higher Dimensional Data: Steel Malice', 2, NOW(), NOW()), -- Ben Bigger
  (7, 'Ferocious Grip', 'Higher Dimensional Data: Ethereal Pursuit', 1, NOW(), NOW()), -- Billy kid
  (8, 'Scarlet Engine', 'Higher Dimensional Data: Stealth Phantom', 2, NOW(), NOW()), -- Burnice White
  (9, 'Scarlet Engine', 'Higher Dimensional Data: Stealth Phantom', 1, NOW(), NOW()), -- Caesar
  (10, 'Finale Dance Shoes', 'Higher Dimensional Data: Crimson Awe', 1, NOW(), NOW()), -- Corin Wickes
  (11, 'Ferocious Grip', 'Higher Dimensional Data: Murderous Obituary', 3, NOW(), NOW()), -- Ellen Joe
  (12, 'Scarlet Engine', 'Higher Dimensional Data: Steel Malice', 2, NOW(), NOW()), -- Evelyn Chevalier
  (13, 'Living Drive', 'Higher Dimensional Data: Destructive Advance', 4, NOW(), NOW()), -- Grace Howard
  (14, 'Sycophant''s Refinement', 'Higher Dimensional Data: Thunderous Dragon', 3, NOW(), NOW()), -- Hoshimi Miyabi
  (15, 'Finale Dance Shoes', 'Higher Dimensional Data: Murderous Obituary', 3, NOW(), NOW()), -- Vlad
  (16, 'Ferocious Grip', 'Higher Dimensional Data: Falling Fist', 1, NOW(), NOW()), -- Jane Doe
  (17, 'Living Drive', 'Higher Dimensional Data: Steel Malice', 2, NOW(), NOW()), -- Koleda Belobog
  (18, 'Scarlet Engine', 'Higher Dimensional Data: Crimson Awe', 2, NOW(), NOW()), -- Lighter
  (19, 'Ferocious Grip', 'Higher Dimensional Data: Steel Malice', 2, NOW(), NOW()), -- Luciana de Montefio
  (20, 'Ferocious Grip', 'Higher Dimensional Data: Crimson Awe', 1, NOW(), NOW()), -- Nekomiya Mana
  (21, 'Ferocious Grip', 'Higher Dimensional Data: Murderous Obituary', 5, NOW(), NOW()), -- Nicole Demara
  (22, 'Finale Dance Shoes', 'Higher Dimensional Data: Crimson Awe', 1, NOW(), NOW()), -- Piper Wheel
  (23, 'Scarlet Engine', 'Higher Dimensional Data: Stealth Phantom', 1, NOW(), NOW()), -- Pulchra Fellini
  (24, 'Living Drive', 'Higher Dimensional Data: Ethereal Pursuit', 4, NOW(), NOW()), -- Qingyi
  (25, 'Living Drive', 'Higher Dimensional Data: Falling Fist', 4, NOW(), NOW()), -- Seth Lowell
  (26, 'Living Drive', 'Higher Dimensional Data: Mortal Cleave', 4, NOW(), NOW()), -- Soldier 0 - Anby
  (27, 'Finale Dance Shoes', 'Higher Dimensional Data: Destructive Advance', 2, NOW(), NOW()), -- Soldier 11
  (28, 'Finale Dance Shoes', 'Higher Dimensional Data: Murderous Obituary', 3, NOW(), NOW()), -- Soukaku
  (29, 'Ferocious Grip', 'Higher Dimensional Data: Mortal Cleave', 4, NOW(), NOW()), -- Trigger
  (30, 'Living Drive', 'Higher Dimensional Data: Destructive Advance', 4, NOW(), NOW()), -- Tsukishiro Yanagi
  (31, 'Sycophant''s Refinement', 'Higher Dimensional Data: Thunderous Dragon', 5, NOW(), NOW()), -- Vivian Banshee
  (32, 'Finale Dance Shoes', 'Higher Dimensional Data: Ethereal Pursuit', 3, NOW(), NOW()), -- Von Lycaon
  (33, 'Living Drive', 'Higher Dimensional Data: Ethereal Pursuit', 5, NOW(), NOW()), -- Zhu Yuan
  (34, 'Exuvia of Refinement', 'Higher Dimensional Data: Miasmic Elytron', 5, NOW(), NOW()), -- Yixuan
  (35, 'Exuvia of Refinement', 'Higher Dimensional Data: Miasmic Elytron', 1, NOW(), NOW()), -- Pan Yinhu
  (36, 'Exuvia of Refinement', 'Higher Dimensional Data: Miasmic Elytron', 2, NOW(), NOW()), -- Ju Fufu
  -- (2, '', '', 2, NOW(), NOW()),


INSERT INTO zenless_characters (id, name, specialty, rank, attribute, faction, ascension_id, nodes_id, created_at, updated_at)
VALUES 
  (1, 'Alexandrina Sebastiane', 'Support', 'S', 'Electric', 'Victoria Housekeeping Co.', 10, 1, NOW(), NOW()),
  (2, 'Anby Demara', 'Stun', 'A', 'Electric', 'Cunning Hares', 7, 2, NOW(), NOW()),
  (3, 'Anton Ivanov', 'Attack', 'A', 'Electric', 'Belobog Heavy Industries', 6, 3, NOW(), NOW()),
  (4, 'Asaba Harumasa', 'Attack', 'S', 'Electric', 'Hollow Special Operations Section 6', 6, 4, NOW(), NOW()),
  (5, 'Astra Yao', 'Support', 'S', 'Ether', 'Stars of Lyra', 10, 5, NOW(), NOW()),
  (6, 'Ben Bigger', 'Defence', 'A', 'Fire', 'Belobog Heavy Industries', 9, 6, NOW(), NOW()),
  (7, 'Billy Kid', 'Attack', 'A', 'Physical', 'Cunning Hares', 6, 7, NOW(), NOW()),
  (8, 'Burnice White', 'Anomaly', 'S', 'Fire', 'Sons of Calydon', 8, 8, NOW(), NOW()),
  (9, 'Caesar King', 'Defence', 'S', 'Physical', 'Sons of Calydon', 9, 9, NOW(), NOW()),
  (10, 'Corin Wickes', 'Attack', 'A', 'Physical', 'Victoria Housekeeping Co.', 6, 10, NOW(), NOW()),
  (11, 'Ellen Joe', 'Attack', 'S', 'Ice', 'Victoria Housekeeping Co.', 6, 11, NOW(), NOW()),
  (12, 'Evelyn Chevalier', 'Attack', 'S', 'Fire', 'Stars of Lyra', 6, 12, NOW(), NOW()),
  (13, 'Grace Howard', 'Anomaly', 'S', 'Electric', 'Belobog Heavy Industries', 8, 13, NOW(), NOW()),
  (14, 'Hoshimi Miyabi', 'Anomaly', 'S', 'Frost', 'Hollow Special Operations Section 6', 8, 14, NOW(), NOW()),
  (15, 'Hugo Vlad', 'Attack', 'S', 'Ice', 'Mockingbird', 6, 15, NOW(), NOW()),
  (16, 'Jane Doe', 'Anomaly', 'S', 'Physical', 'Criminal Investigation Special Response Team', 8, 16, NOW(), NOW()),
  (17, 'Koleda Belobog', 'Stun', 'S', 'Fire', 'Belobog Heavy Industries', 7, 17, NOW(), NOW()),
  (18, 'Lighter', 'Stun', 'S', 'Fire', 'Sons of Calydon', 7, 18, NOW(), NOW()),
  (19, 'Luciana de Montefio', 'Support', 'A', 'Fire', 'Sons of Calydon', 10, 19, NOW(), NOW()),
  (20, 'Nekomiya Mana', 'Attack', 'S', 'Physical', 'Cunning Hares', 6, 20, NOW(), NOW()),
  (21, 'Nicole Demara', 'Support', 'A', 'Ether', 'Cunning Hares', 10, 21, NOW(), NOW()), 
  (22, 'Piper Wheel', 'Anomaly', 'A', 'Physical', 'Sons of Calydon', 8, 22, NOW(), NOW()),
  (23, 'Pulchra Fellini', 'Stun', 'A', 'Physical', 'Sons of Calydon', 7, 23, NOW(), NOW()),
  (24, 'Qingyi', 'Stun', 'S', 'Electric', 'Criminal Investigation Special Response Team', 7, 24, NOW(), NOW()),
  (25, 'Seth Lowell', 'Defence', 'A', 'Electric', 'Criminal Investigation Special Response Team', 9, 25, NOW(), NOW()),
  (26, 'Soldier 0 - Anby', 'Attack', 'S', 'Electric', 'Defense Force - Silver Squad', 6, 26, NOW(), NOW()),
  (27, 'Soldier 11', 'Attack', 'S', 'Fire', 'Obol Squad', 6, 27, NOW(), NOW()),
  (28, 'Soukaku', 'Support', 'A', 'Ice', ' Hollow Special Operations Section 6', 10, 28, NOW(), NOW()),
  (29, 'Trigger', 'Stun', 'S', 'Electric', 'Obol Squad', 7, 29, NOW(), NOW()),
  (30, 'Tsukishiro Yanagi', 'Anomaly', 'S', 'Electric', 'Hollow Special Operations Section 6', 8, 30, NOW(), NOW()),
  (31, 'Vivian Banshee', 'Anomaly', 'S', 'Ether', 'Mockingbird', 8, 31, NOW(), NOW()),
  (32, 'Von Lycaon', 'Stun', 'S', 'Ice', 'Victoria Housekeeping Co.', 7, 32, NOW(), NOW()), 
  (33, 'Zhu Yuan', 'Attack', 'S', 'Ether', 'Criminal Investigation Special Response Team', 6, 33, NOW(), NOW());
  (34, 'Yixuan', 'Rupture', 'S', 'Auric Ink', 'Yunkui Summit', 12, 34, NOW(), NOW()),
  (35, 'Pan Yinhu', 'Defence', 'A', 'Physical', 'Yunkui Summit', 9, 35, NOW(), NOW()),
  (36, 'Ju Fufu', 'Stun', 'S', 'Fire', 'Yunkui Summit', 7, 36, NOW(), NOW()),
  -- (1, '', '', '', '', '', 1, 1, NOW(), NOW()),




