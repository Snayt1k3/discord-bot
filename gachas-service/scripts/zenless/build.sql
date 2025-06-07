INSERT INTO zenless_weapons (id, name, rarity, type, base_atk, sub_stat, sub_value, passive, created_at, updated_at, deleted_at)
VALUES 
    (1, 'Blazing Laurel', 'S', 'Stun', 713, 'Impact', 18, 'Upon launching a Quick Assist or Perfect Assist, the equipper''s Impact increases by 25–40% for 8s. When the equipper hits with a Basic Attack, apply Wilt for 30s (up to 20 stacks). Each stack increases CRIT DMG of Ice/Fire attacks by 1.5–2.4%. Only one instance can be active squad-wide.', NOW(), NOW(), NULL),
    (2, 'Deep Sea Visitor', 'S', 'Attack', 713, 'Crit Rate', 24, 'Increases Ice DMG by 25–50%. Basic Attacks increase CRIT Rate by 10–20% for 8s. Dash Attacks increase CRIT Rate by another 10–20% for 15s. Durations are tracked separately.', NOW(), NOW(), NULL),
    (3, 'Elegant Vanity', 'S', 'Support', 713, 'Attack', 30, 'Quick/Chain/Assist entries grant 5–7 Energy (1/5s). Consuming 25+ Energy boosts squad DMG by 10–16%, stacks 2x, lasts 20s, refreshes on trigger. Only one instance can exist per squad.', NOW(), NOW(), NULL),
    (4, 'Flamemaker Shaker', 'S', 'Anomaly', 713, 'Attack', 30, 'Off-field: Energy Regen +0.6–1.2/s. EX/Assist hits boost DMG by 3.5–7% (max 10 stacks, 6s). Off-field: stack effect is doubled. At 5+ stacks, gain +50–100 Anomaly Proficiency for 6s.', NOW(), NOW(), NULL),
    (5, 'Flight of Fancy', 'S', 'Anomaly', 713, 'Anomaly Proficiency', 90, 'Increases Anomaly Buildup Rate by 40–64%. Dealing Ether DMG grants +20–32 Anomaly Proficiency (5s), up to 6 stacks. Triggers every 0.5s, refreshes duration.', NOW(), NOW(), NULL),
    (6, 'Fusion Compiler', 'S', 'Anomaly', 684, 'PEN Ratio', 24, 'ATK +12–24%. Using Special/EX Attack grants +25–50 Anomaly Proficiency for 8s, stacking up to 3 times. Stacks have independent durations.', NOW(), NOW(), NULL),
    (7, 'Hailstorm Shrine', 'S', 'Anomaly', 743, 'Crit Rate', 24, 'CRIT DMG +50–80%. On EX Special Attack or squad applying Attribute Anomaly, gain +20–32% Ice DMG (max 2 stacks, 15s each). Stacks tracked separately.', NOW(), NOW(), NULL),
    (8, 'Heartstring Nocturne', 'S', 'Attack', 713, 'Crit Rate', 24, 'CRIT DMG +50–80%. On entering field, Chain or Ultimate, gain 1 Heartstring stack (max 2). Each stack allows Chain/Ultimate to ignore 12.5–20% Fire RES, lasts 30s. Refreshes on trigger.', NOW(), NOW(), NULL),
    (9, 'Hellfire Gears', 'S', 'Stun', 684, 'Impact', 18, 'Off-field: Energy Regen +0.6–1.2/s. Using EX Special Attack boosts Impact by 10–20% for 10s (max 2 stacks, tracked separately).', NOW(), NOW(), NULL),
    (10, 'Ice-Jade Teapot', 'S', 'Stun', 713, 'Impact', 18, 'Basic Attacks grant Tea-riffic: +0.7–1.4% Impact per stack (max 30, 8s each). At 15+ stacks, squad DMG +20–32% for 10s. Same effects don’t stack.', NOW(), NOW(), NULL),
    (11, 'Myriad Eclipse', 'S', 'Attack', 713, 'Crit Rate', 24, 'CRIT DMG +45–72%. EX/Chain/Ultimate Ice DMG grants Absolute Zero Death Sentence (3s): ignore 25–40% DEF.', NOW(), NOW(), NULL),
    (12, 'Qingming Birdcage', 'S', 'Rupture', 743, 'HP', 30, 'CRIT Rate +20–32%. EX Special Attack grants Qingming Companion (max 2 stacks, 15s): +8–12.8% Ether DMG & +10–16% Ether Sheer DMG. Gain 2 stacks on entering combat.', NOW(), NOW(), NULL),
    (13, 'Riot Suppressor Mark VI', 'S', 'Attack', 713, 'Crit Damage', 48, 'CRIT Rate +15%. EX Special Attack grants 8 Charge stacks. Basic Ether DMG consumes 1 stack to boost DMG by 35%.', NOW(), NOW(), NULL),
    (14, 'Severed Innocence', 'S', 'Attack', 713, 'Crit DMG', 48, 'CRIT DMG +30–48%. Hitting with skills grants stack (+10–16% CRIT DMG, max 3, 30s each). At 3 stacks: +20–32% Electric DMG.', NOW(), NOW(), NULL),
    (15, 'Sharpened Stinger', 'S', 'Anomaly', 713, 'Anomaly Proficiency', 90, 'Dash Attack grants Predatory Instinct (+12–24% Physical DMG, max 3 stacks, 10s). Perfect Dodge/entering combat gives 3 stacks. Max stacks: +40–80% Anomaly Buildup Rate.', NOW(), NOW(), NULL),
    (16, 'Spectral Gaze', 'S', 'Stun', 713, 'Crit Rate', 24, 'Aftershock with Electric DMG reduces enemy DEF by 25–40% for 5s. If off-field, gain Spirit Lock (max 3): +4–6.4% Impact per stack (12s each). At 3 stacks: extra +8–12.8% Impact.', NOW(), NOW(), NULL),
    (17, 'Steel Cushion', 'S', 'Attack', 684, 'Crit Rate', 24, 'Physical DMG +20–40%. DMG +25–50% when attacking from behind.', NOW(), NOW(), NULL),
    (18, 'The Brimstone', 'S', 'Attack', 684, 'Attack', 30, 'Basic/Dash/Dodge Attacks grant +3.5–7% ATK for 8s (max 8 stacks, trigger per 0.5s). Duration tracked separately.', NOW(), NOW(), NULL),
    (19, 'The Restrained', 'S', 'Stun', 684, 'Impact', 18, 'On hit: Basic ATK DMG & Daze +6–12% for 8s (max 5 stacks, once per skill). Duration tracked separately.', NOW(), NOW(), NULL),
    (20, 'Timeweaver', 'S', 'Anomaly', 713, 'Attack', 30, 'Electric Anomaly Buildup +30–50%. Special/EX hitting anomalous enemies: +75–115 Anomaly Prof. (15s). At ≥375 Prof.: +25–35% Disorder DMG.', NOW(), NOW(), NULL),
    (21, 'Tusks of Fury', 'S', 'Defence', 713, 'Impact', 18, 'Shield value +30–60%. On squad Interrupt/Perfect Dodge: squad DMG +18–36%, Daze +12–24% for 20s. Same effects don’t stack.', NOW(), NOW(), NULL),
    (22, 'Weeping Cradle', 'S', 'Support', 684, 'PEN Ratio', 24, 'Off-field Energy Regen +0.6–1.2/s. Attacks boost squad DMG vs. target by 10–20% (3s), growing +1.7–3.3%/0.5s (max +19.8%). Duration refreshes only.', NOW(), NOW(), NULL),
    (23, 'Zanshin Herb Case', 'S', 'Attack', 713, 'Crit Damage', 48, 'CRIT Rate +10–16%. Dash Electric DMG +40–64%. On squad Anomaly/Stun: extra CRIT Rate +10–16% for 15s.', NOW(), NOW(), NULL),
    (24, 'Bashful Demon', 'A', 'Support', 624, 'Attack', 25, 'Ice DMG +15–24%. EX Special boosts squad ATK by 2–3.2% (12s, 4 stacks). Refreshes duration. Effects don’t stack.', NOW(), NOW(), NULL),
    (25, 'Big Cylinder', 'A', 'Defence', 624, 'Defense', 40, 'DMG taken −7.5–12%. After being hit, next attack CRITs and deals +600–960% DEF as bonus DMG (CD: 7.5s).', NOW(), NOW(), NULL),
    (26, 'Box Cutter', 'A', 'Stun', 624, 'Impact', 15, 'Aftershock boosts Physical DMG +15–24% and Daze +10–16% for 10s.', NOW(), NOW(), NULL),
    (27, 'Bunny Band', 'A', 'Defence', 594, 'Defense', 40, 'Max HP +8–12.8%. When shielded, ATK +10–16%.', NOW(), NOW(), NULL),
    (28, 'Cannon Rotor', 'A', 'Attack', 594, 'Crit Rate', 20, 'ATK +7.5–12%. CRITs deal +200% ATK as bonus DMG (CD: 8–6s).', NOW(), NOW(), NULL),
    (29, 'Demara Battery Mark II', 'A', 'Stun', 624, 'Impact', 15, 'Electric DMG +15–24%. On Dodge Counter or Assist hit: Energy Gen Rate +18–27.5% for 8s.', NOW(), NOW(), NULL),
    (30, 'Drill Rig - Red Axis', 'A', 'Attack', 624, 'Energy Regen', 50, 'EX or Chain Attack: Electric Basic/Dash ATK DMG +50–80% for 10s (CD: 15s).', NOW(), NOW(), NULL),
    (31, 'Electro-Lip Gloss', 'A', 'Anomaly', 594, 'Anomaly Proficiency', 75, 'If enemies have Anomaly: ATK +10–16%, deal +15–25% more DMG to them.', NOW(), NOW(), NULL),
    (32, 'Gilded Blossom', 'A', 'Attack', 594, 'Attack', 25, 'ATK +6–9.6%. EX Special Attack DMG +15–24%.', NOW(), NOW(), NULL),
    (33, 'Housekeeper', 'A', 'Attack', 624, 'Attack', 25, 'Off-field: Energy Regen +0.45–0.72/s. On EX hit: Physical DMG +3–4.8% (15 stacks, 1s, refreshable).', NOW(), NOW(), NULL),
    (34, 'Kaboom the Cannon', 'A', 'Support', 624, 'Energy Regen', 50, 'On any hit: squad ATK +2.5–4% (8s, max 4 stacks). Each unit can provide 1 stack.', NOW(), NOW(), NULL),
    (35, 'Marcato Desire', 'A', 'Attack', 594, 'Crit Rate', 20, 'On EX/Chain hit: ATK +6–9.6% for 8s. Vs Anomalous target: bonus +6–9.6%.', NOW(), NOW(), NULL),
    (36, 'Original Transmorpher', 'A', 'Defence', 594, 'HP', 25, 'Max HP +8–12.5%. When attacked: Impact +10–16% for 12s.', NOW(), NOW(), NULL),
    (37, 'Peacekeeper - Specialized', 'A', 'Defence', 624, 'Attack', 25, 'While shielded: Energy Regen +0.4/s. EX & Assist Follow-up Buildup +36%.', NOW(), NOW(), NULL),
    (38, 'Precious Fossilized Core', 'A', 'Stun', 594, 'Impact', 15, 'If target HP ≥50%: Daze +10–16%. If HP ≥75%: bonus +10–16%.', NOW(), NOW(), NULL),
    (39, 'Puzzle Sphere', 'A', 'Rupture', 594, 'Attack', 25, 'On EX: CRIT DMG +16–25.6% (12s). If target HP <50%: EX DMG +20–32%.', NOW(), NOW(), NULL),
    (40, 'Radiowave Journey', 'A', 'Rupture', 594, 'HP', 25, 'On Chain/Ultimate: gain stack (+80–128 Sheer Force, 12s, 3 stacks max, tracked separately).', NOW(), NOW(), NULL),
    (41, 'Rainforest Gourmet', 'A', 'Anomaly', 594, 'Anomaly Proficiency', 75, 'Every 10 Energy spent: ATK +2.5–4% (10s, max 10 stacks, separate durations).', NOW(), NOW(), NULL),
    (42, 'Roaring Ride', 'A', 'Anomaly', 624, 'Attack', 25, 'On EX hit: random 5s buff (ATK +8–12.8%, Anomaly Prof. +40–64, or Buildup Rate +25–40%). No stacking, can overlap, refreshes duration.', NOW(), NOW(), NULL),
    (44, 'Six Shooter', 'A', 'Stun', 594, 'Impact', 15, 'Gains a Charge stack every 3s (max 6). On EX Special Attack, consume all stacks; each stack increases Daze inflicted by 4–6.4%.', NOW(), NOW(), NULL),
    (45, 'Slice of Time', 'A', 'Support', 594, 'PEN Ratio', 20, 'Squad members'' Dodge Counter, EX, Assist, or Chain Attacks generate Decibels and Energy for equipper, triggering independently with 12s cooldown each.', NOW(), NOW(), NULL),
    (46, 'Spring Embrace', 'A', 'Defence', 594, 'Attack', 25, 'DMG taken reduced by 7.5–12%. When attacked, Energy Regen +10–16% for 12s, transferred to new on-field character on swap.', NOW(), NOW(), NULL),
    (47, 'Starlight Engine', 'A', 'Attack', 594, 'Attack', 25, 'Launching Dodge Counter or Quick Assist increases ATK by 12–19.2% for 12s.', NOW(), NOW(), NULL),
    (48, 'Starlight Engine Replica', 'A', 'Attack', 624, 'Attack', 25, 'Physical DMG +36–57.5% for 8s after hitting enemy ≥6m away with Basic or Dash Attack.', NOW(), NOW(), NULL),
    (49, 'Steam Oven', 'A', 'Stun', 594, 'Energy Regen', 50, 'For every 10 Energy accumulated, Impact +2–3.2% (max 8 stacks). After Energy spent, bonus lasts 8s.', NOW(), NOW(), NULL),
    (50, 'Street Superstar', 'A', 'Attack', 594, 'Attack', 25, 'On squad Chain Attack, gain Charge stack (max 3). On Ultimate, consume stacks to increase skill DMG by 15–24% per stack.', NOW(), NOW(), NULL),
    (51, 'The Vault', 'A', 'Support', 624, 'Energy Regen', 50, 'Dealing Ether DMG with EX, Chain, or Ultimate boosts squad DMG vs target by 15–24% and Energy Regen +0.5–0.8/s for 2s.', NOW(), NOW(), NULL),
    (52, 'Tremor Trigram Vessel', 'A', 'Defence', 624, 'Attack', 25, 'EX Special and Ultimate DMG +25–40%. When squad member takes DMG or heals, gain 2–3.2 Energy (CD 5s).', NOW(), NOW(), NULL),
    (53, 'Unfettered Game Ball', 'A', 'Support', 594, 'Energy Regen', 50, 'On Attribute Counter trigger, all squad members'' CRIT Rate vs struck enemy +12–20% for 12s. Effects do not stack.', NOW(), NOW(), NULL),
    (54, 'Weeping Gemini', 'A', 'Anomaly', 594, 'Attack', 25, 'When squad inflicts Attribute Anomaly, gain Anomaly Proficiency buff +30–46 (max 4 stacks), lost on Stun or target defeat.', NOW(), NOW(), NULL),
    (55, '[Cinder] Cobalt', 'A', 'Rupture', 475, 'HP', 20, 'On combat start or switch-in, ATK +7.2–11.5% for 10s (CD 20s).', NOW(), NOW(), NULL),
    (56, '[Identity] Base', 'B', 'Defence', 475, 'Defense', 32, 'When attacked, DEF +20–32% for 8s.', NOW(), NOW(), NULL),
    (57, '[Identity] Inflection', 'B', 'Defence', 475, 'Defense', 32, 'When attacked, reduces attacker''s DMG by 6–10% for 12s.', NOW(), NOW(), NULL),
    (58, '[Lunar] Decrescent', 'B', 'Attack', 475, 'Attack', 20, 'Launching a Chain Attack or Ultimate increases the equipper''s DMG by 15–25% for 6s.', NOW(), NOW(), NULL),
    (59, '[Lunar] Noviluna', 'B', 'Attack', 475, 'Crit Rate', 16, 'Launching an EX Special Attack generates 3–5 Energy for the equipper. 12s cooldown.', NOW(), NOW(), NULL),
    (60, '[Lunar] Pleniluna', 'B', 'Attack', 475, 'Attack', 20, 'Basic, Dash, and Dodge Counter DMG increased by 12–20%.', NOW(), NOW(), NULL),
    (61, '[Magnetic Storm] Alpha', 'B', 'Anomaly', 475, 'Attack', 20, 'Accumulating Anomaly Buildup increases Anomaly Mastery by 25–40 for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (62, '[Magnetic Storm] Bravo', 'B', 'Anomaly', 356, 'Anomaly Proficiency', NULL, 'Accumulating Anomaly Buildup increases Anomaly Proficiency by 25–40 for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (63, '[Magnetic Storm] Charlie', 'B', 'Anomaly', 475, 'PEN Ratio', 16, 'When squad inflicts Attribute Anomaly, generates 3.5–5.5 Energy. 12s cooldown.', NOW(), NOW(), NULL),
    (64, '[Reverb] Mark I', 'B', 'Support', 475, 'Attack', 20, 'Launching an EX Special Attack increases all squad members'' Impact by 8–12% for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (65, '[Reverb] Mark II', 'B', 'Support', 475, 'Energy Regen', 40, 'Launching EX Special or Chain Attack increases all squad members'' Anomaly Mastery and Proficiency by 10–16 for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (66, '[Reverb] Mark III', 'B', 'Support', 475, 'HP', 20, 'Launching Chain Attack or Ultimate increases all squad members'' ATK by 8–12% for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (67, '[Vortex] Arrow', 'B', 'Stun', 475, 'Impact', 12, 'Attacks inflict 8–12% more Daze on main target.', NOW(), NOW(), NULL),
    (68, '[Vortex] Hatchet', 'B', 'Stun', 475, 'Energy Regen', 16, 'On entering combat or switch-in, Impact +9–13% for 10s. 20s cooldown.', NOW(), NOW(), NULL),
    (69, '[Vortex] Revolver', 'B', 'Stun', 475, 'Attack', 20, 'EX Special Attacks inflict 10–16% more Daze.', NOW(), NOW(), NULL);
INSERT INTO zenless_discs (id, name, two_piece_bonus, four_piece_bonus, created_at, updated_at, deleted_at)
VALUES 
    (1, 'Astral Voice', 'ATK +10%.', 'Whenever any squad member enters the field using a Quick Assist, all squad members gain 1 stack of Astral, up to a maximum of 3 stacks, and lasting 15s. Repeated triggers reset the duration. Each stack of Astral increases the DMG dealt by the character entering the field using a Quick Assist by 8%. Only one of this effect can be active at a time in the same squad.', NOW(), NOW(), NULL),
    (2, 'Branch & Blade Song', 'CRIT DMG +16%.', 'When Anomaly Mastery exceeds or equals 115 points, the equipper''s CRIT damage increases by 30%. When any squad member applies Freeze or triggers the Shatter effect on an enemy, the equipper''s CRIT Rate increases by 12%, lasting 15s.', NOW(), NOW(), NULL),
    (3, 'Chaos Jazz', 'Anomaly Proficiency +30.', 'Fire DMG and Electric DMG are increased by 15%. When off-field, the damage caused by EX Special Attacks and Assist Attacks is increased by 20%. When the character switches back onto the field, this buff continues for 5s. The lasting effect can be triggered once every 7.5s.', NOW(), NOW(), NULL),
    (4, 'Chaotic Metal', 'Ether DMG +10%.', 'The equipper''s CRIT DMG increases by 20%. When any character in the squad triggers Corruption''s additional DMG, this effect further increases by 5.5% for 8s, stacking up to 6 times. Repeated triggers reset the duration.', NOW(), NOW(), NULL),
    (5, 'Fanged Metal', 'Physical DMG +10%.', 'Whenever a squad member inflicts Assault on an enemy, the equipper deals 35% increased damage to the target for 12s.', NOW(), NOW(), NULL),
    (6, 'Freedom Blues', 'Anomaly Proficiency +30.', 'When an EX Special Attack hits an enemy, reduce the targets Anomaly Buildup RES to equippers attribute by 20% for 8s. This effect does not stack with others of the same attribute.', NOW(), NOW(), NULL),
    (7, 'Hormone Punk', 'ATK +10%.', 'Upon entering combat or switching in, its user''s ATK increased by 25% for 10s. This effect can be triggered once every 20s.', NOW(), NOW(), NULL),
    (8, 'Inferno Metal', 'Fire DMG +10%.', 'Upon hitting a Burning enemy, the equipper''s CRIT Rate is increased by 28% for 8s.', NOW(), NOW(), NULL),
    (9, 'King of the Summit', 'Increases Daze of attacks by 6%.', 'When the equipper is a Stun character and uses an EX Special Attack or Chain Attack, increases CRIT DMG of all squad members by 15%, and when the equipper''s CRIT Rate is more than or equal to 50%, further increases CRIT DMG by 15%, lasting 15s. Repeated triggers reset the duration. Passive effects of the same name do not stack.', NOW(), NOW(), NULL),
    (10, 'Phaethon''s Melody', 'Anomaly Mastery +8%.', 'When any squad member uses an EX Special Attack, the equipper''s Anomaly Proficiency increases by 45 for 8s. If the character using the EX Special Attack is not the equipper, the equipper''s Ether DMG is increased by 25%.', NOW(), NOW(), NULL),
    (11, 'Polar Metal', 'Ice DMG +10%.','Increases the DMG of Basic Attack and Dash Attack by 20%. When any squad member inflicts Freeze or Shatter, this effect increases by an additional 20% for 12s.', NOW(), NOW(), NULL),
    (12, 'Proto Punk', 'Shield Effect +15%.', 'When any squad member triggers a Defensive Assist or Evasive Assist, all squad members deal 15% increased DMG, lasting 10s. Passive effects of the same name do not stack.', NOW(), NOW(), NULL),
    (13, 'Puffer Electro', 'PEN Ratio +8%.', 'Ultimate DMG increases by 20%. Launching an Ultimate increases the equipper''s ATK by 15% for 12s.', NOW(), NOW(), NULL),
    (14, 'Shadow Harmony', 'The DMG of Aftershocks and Dash Attacks is increased by 15%.', 'Upon hitting an enemy with an Aftershock or Dash Attack, if the DMG dealt aligns with the equipper''s attribute, the equipper gains 1 stack of a buff effect, at most once per use of a skill. For each stack, the equipper''s ATK increases by 4%, and CRIT Rate increases by 4%. The effect can stack up to 3 times and lasts for 15s. Repeated triggers reset the duration.', NOW(), NOW(), NULL),
    (15, 'Shockstar Disco', 'Impact +6%.', 'Basic Attacks, Dash Attacks, and Dodge Counters, inflict 20% more Daze upon the main target.', NOW(), NOW(), NULL),
    (16, 'Soul Rock', 'DEF +16%.', 'Upon receiving an enemy attack and losing HP, the equipper takes 40% less DMG for 2.5s. This effect can trigger once every 15s.', NOW(), NOW(), NULL),
    (17, 'Swing Jazz', 'Energy Regen +20%.', 'Launching a Chain Attack or Ultimate increases all squad members'' DMG by 15% for 12s. Passive effects of the same name do not stack.', NOW(), NOW(), NULL),
    (18, 'Thunder Metal', 'Electric DMG +10%.', 'As long as an enemy in combat is Shocked, the equipper''s ATK is increased by 28%.', NOW(), NOW(), NULL),
    (19, 'Woodpecker Electro', 'CRIT Rate +8%.', 'Triggering a critical hit with a Basic Attack, Dodge Counter or EX Special Attack increases the equipper''s ATK by 9% for 6s. The buff duration for different skills are calculated separately.', NOW(), NOW(), NULL),
    (20, 'Yunkui Tales', 'HP +10%.', 'When using [EX Special Attack], [Chain Attack], or [Ultimate], CRIT Rate increases by 4%, stacking up to 3 times and lasting 15s. Repeated triggers reset the duration. When having 3 stacks of this effect, Sheer DMG increases by 10%.', NOW(), NOW(), NULL);

INSERT INTO zenless_stats (id, fourth_disc, fifth_disc, sixth_disc, sub_stats_priority, created_at, updated_at, deleted_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

INSERT INTO zenless_discs_presets (id, four_piece_set_id, two_piece_set_id, created_at, updated_at, deleted_at)
VALUES (?, ?, ?, ?, ?, ?, ?);


INSERT INTO zenless_builds (id, character_id, stats_id, created_at, updated_at, deleted_at)
VALUES (?, ?, ?, ?, ?, ?);

INSERT INTO zenless_build_weapon (build_id, weapon_id, priority)
VALUES (?, ?, ?);

INSERT INTO zenless_build_disc (build_id, disc_id, priority)
VALUES (?, ?, ?);