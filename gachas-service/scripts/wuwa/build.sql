INSERT INTO wuwa_weapons (id, name, weapon_type, rarity, passive, base_atk, sub_stat, sub_value, created_at, updated_at, deleted_at)
VALUES
    (1, 'Abyss Surges', 'Gauntlets', 5, 'Increases Energy Regen by 12.8%. When hitting a target with Resonance Skill, increases Basic Attack DMG Bonus by 10%, lasting for 8s. When hitting a target with Basic Attacks, increases Resonance Skill DMG Bonus by 10%, lasting for 8s.', 587, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (2, 'Ages of Harvest', 'Broadblade', 5, 'Grants 12% Attribute DMG Bonus. Casting Intro Skill gives the equipper Ageless Marking, which grants 24% Resonance Skill DMG Bonus for 12s. Casting Resonance Skill gives the equipper Ethereal Endowment, which grants 24% Resonance Skill DMG Bonus for 12s.', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (3, 'Blazing Brilliance', 'Sword', 5, 'ATK increased by 12%. The equipper gains 1 stack of Searing Feather upon dealing damage (trigger: 0.5s), and gains 5 stacks upon casting Resonance Skill. Each stack grants 4% Resonance Skill DMG Bonus (up to 14). All stacks expire in 12s.', 587, 'CRIT DMG', 48.6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (4, 'Blazing Justice', 'Gauntlets', 5, 'Increases ATK by 12%. Casting Basic Attack grants: ignores 8% DEF and Amplifies Spectro Frazzle DMG by 50% for 6s. Recasting resets duration.', 587, 'CRIT DMG', 48.6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (5, 'Bloodpact''s Pledge', 'Sword', 5, 'Providing Healing increases Resonance Skill DMG by 10% for 6s. When Rover: Aero casts Unbound Flow, Aero DMG by nearby Resonators is Amplified by 10% for 30s.', 588, 'Energy Reg.', 38.9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (6, 'Cosmic Ripples', 'Rectifier', 5, 'Increases Energy Regen by 12.8%. Hitting with Basic Attacks increases Basic ATK DMG Bonus by 3.2%, stacking 5 times (8s duration, trigger: 0.5s).', 500, 'ATK', 54.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (7, 'Emerald of Genesis', 'Sword', 5, 'Increases Energy Regen by 12.8%. When Resonance Skill is released, increases ATK by 6%, stacking up to 2 times (10s).', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (8, 'Luminous Hymn', 'Rectifier', 5, 'Increase ATK by 12%. Hitting Spectro Frazzle grants 14% Basic and Heavy ATK DMG Bonus, stacking 3 times (6s). Outro Skill Amplifies Spectro Frazzle DMG by 30% for 30s. Cannot stack.', 500, 'CRIT Rate', 36.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (9, 'Lustrous Razor', 'Broadblade', 5, 'Increases Energy Regen by 12.8%. Using Resonance Skill increases Resonance Liberation DMG by 7%, stacking up to 3 (12s).', 587, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (10, 'Red Spring', 'Sword', 5, 'Increase ATK by 12%. Basic ATK grants 10% Basic ATK DMG Bonus for 14s, stacking up to 3 (1s CD). Consuming Concerto Energy gives 40% Basic DMG Bonus for 10s. Ends when switched out.', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (11, 'Rime-Draped Sprouts', 'Rectifier', 5, 'Increases ATK by 12%. Resonance Skill gives 12% Basic ATK DMG Bonus for 6s, stacking 3 times. Outro Skill with 3 stacks gives 52% off-field Basic ATK DMG Bonus for 27s.', 500, 'CRIT DMG', 72.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (12, 'Static Mist', 'Pistols', 5, 'Increases Energy Regen by 12.8%. Outro Skill gives switched-in Resonator 10% ATK (14s). Max 1 stack.', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (13, 'Stellar Symphony', 'Rectifier', 5, 'Increase HP by 12%. Casting Resonance Liberation restores 8 Concerto Energy (20s CD). Healing with Resonance Skill increases nearby ATK by 14% for 30s. Cannot stack.', 412, 'Energy Regen', 77.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (14, 'Stringmaster', 'Rectifier', 5, 'Increases DMG Bonus by 12%. Resonance Skill DMG increases ATK by 12%, stacking 2 times (5s). Off-field Resonator gains 12% additional ATK.', 500, 'CRIT Rate', 36.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (15, 'The Last Dance', 'Pistols', 5, 'Increases ATK by 12%. Every time Intro Skill or Resonance Liberation is cast, Resonance Skill DMG Bonus increases by 48% for 5s.', 500, 'CRIT DMG', 72.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (16, 'Tragicomedy', 'Gauntlets', 5, 'Increases ATK by 12%. Every time Basic Attack or Intro Skill is cast, Heavy Attack DMG Bonus increases by 48% for 3s.', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (17, 'Unflickering Valor', 'Sword', 5, 'Increase Crit. Rate by 8%. Casting Resonance Liberation gives 24% Basic Attack DMG Bonus for 10s. Dealing Basic Attack DMG gives 24% Basic Attack DMG Bonus for 4s.', 415, 'Energy Reg.', 77.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (18, 'Verdant Summit', 'Broadblade', 5, 'Increases the DMG Bonus of all Resonance Attributes by 12%. Every time Intro Skill or Resonance Liberation is cast, increases Heavy Attack DMG Bonus by 24%, stacking up to 2 time(s). This effect lasts for 14s.', 587, 'CRIT DMG', 48.6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (19, 'Verity''s Handle', 'Gauntlets', 5, 'Gain 12% Attribute DMG Bonus. When using Resonance Liberation, the wielder gains 48% Resonance Liberation DMG Bonus for 8s. This effect can be extended by 5 seconds each time Resonance Skills are cast, up to 3 times.', 587, 'CRIT Rate', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (20, 'Whispers of Sirens', 'Rectifier', 5, 'Increases ATK by 12%. Casting Echo Skill within 10s after casting Intro Skill or Basic Attacks grants 1 stacks of Gentle Dream. Echoes with the same name can only trigger this effect once, stacking up to 2 times, lasting for 10s. When reaching 2 stacks, casting Echo Skill no longer resets the duration of this effect. This effect activates up to once per 10s. Switching to another Resonator ends this effect early. With 1 stacks: Grants 40% Basic Attack DMG Bonus. With 2 stacks: Ignores 12% of the target''s Havoc RES.', 500, 'CRIT DMG', 72.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (21, 'Amity Accord', 'Gauntlets', 4, 'When Intro Skill is released, increases Resonance Liberation DMG Bonus by 20%, lasting for 15s.', 337, 'DEF', 61.5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (22, 'Augment', 'Rectifier', 4, 'When Resonance Liberation is released, increases the caster''s ATK by 15%, lasting for 15s.', 412, 'CRIT Rate', 20.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (23, 'Autumntrace', 'Broadblade', 4, 'Increases ATK by 4% upon dealing Basic Attack DMG or Heavy Attack DMG, stacking up to 5 time(s). This effect lasts for 7s and can be triggered 1 time(s) every 1s.', 412, 'CRIT Rate', 20.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (24, 'Broadblade#41', 'Broadblade', 4, 'When the Resonator''s HP is above 80%, increases ATK by 12%. When the Resonator''s HP is below 40%, restores their HP by 5% for dealing Basic Attack DMG or Heavy Attack DMG. This effect can be triggered 1 time(s) every 8s.', 412, 'Energy Reg.', 32.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (25, 'Cadenza', 'Pistols', 4, 'When Resonance Skill is released, restores Concerto Energy by 8. This effect can be triggered 1 time every 20s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (26, 'Call of the Abyss', 'Rectifier', 4, 'Casting Resonance Liberation increases the Resonator''s Healing Bonus by 16% for 15s.', 338, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (27, 'Celestial Spiral', 'Gauntlets', 4, 'Casting the Resonance Skill grants 6 Resonance Energy and increases ATK by 10%, lasting for 16s. This effect can be triggered once every 20s.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (28, 'Comet Flare', 'Rectifier', 4, 'When hitting a target with Basic Attacks or Heavy Attacks, increases Healing Bonus by 3%, stacking up to 3 time(s). This effect lasts for 8s and can be triggered 1 time(s) every 0.6s.', 412, 'HP', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (29, 'Commando of Conviction', 'Sword', 4, 'When Intro Skill is released, increases ATK by 15%, lasting for 15s.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (30, 'Dauntless Evernight', 'Broadblade', 4, 'When Intro Skill is released, increases ATK by 8% and DEF by 15%, lasting for 15s.', 337, 'DEF', 61.5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (31, 'Discord', 'Broadblade', 4, 'When Resonance Skill is released, restores Concerto Energy by 8. This effect can be triggered 1 time every 20s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (32, 'Endless Collapse', 'Sword', 4, 'Casting the Resonance Skill grants 6 Resonance Energy and increases ATK by 10%, lasting for 16s. This effect can be triggered once every 20s.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (33, 'Fables of Wisdom', 'Sword', 4, 'Dealing DMG to enemies with Negative Statuses increases the wielder''s ATK by 4% for 10s. This effect can be triggered 1 time per second, stackable up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (34, 'Fusion Accretion', 'Rectifier', 4, 'Casting the Resonance Skill grants 6 Resonance Energy and increases ATK by 10%, lasting for 16s. This effect can be triggered once every 20s.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (35, 'Gauntlets#21D', 'Gauntlets', 4, 'When the Resonator dashes or dodges, increases ATK by 8%. Increases Counter Attack DMG by 50%, lasting for 8s. When Counter Attack is performed, restores the Resonator''s HP by 5%. This effect can be triggered 1 time(s) every 6s.', 387, 'Energy Reg.', 38.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (36, 'Helios Cleaver', 'Broadblade', 4, 'Within 12s after Resonance Skill is cast, increases ATK by 3% every 2s, stacking up to 4 time(s). This effect can be triggered 1 time(s) every 12s. When the number of stacks reaches 4, all stacks will be reset within 6s.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (37, 'Hollow Mirage', 'Gauntlets', 4, 'When Resonance Liberation is released, grants 3 stack(s) of Iron Armor. Each stack increases ATK and DEF by 3%, stacking up to 3 time(s). When the Resonator takes damage, reduces the number of stacks by 1.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (38, 'Jinzhou Keeper', 'Rectifier', 4, 'When Intro Skill is released, increases the caster''s ATK by 8% and HP by 10%, lasting for 15s.', 387, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (39, 'Legend of Drunken Hero', 'Gauntlets', 4, 'Dealing DMG to enemies with Negative Statuses increases the wielder''s ATK by 4% for 10s. This effect can be triggered 1 time per second, stackable up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (40, 'Lumingloss', 'Sword', 4, 'When Resonance Skill is released, increases Basic Attack DMG and Heavy Attack DMG by 20%, stacking up to 1 time(s). This effect lasts for 10s and can be triggered 1 time(s) every 1s.', 387, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (41, 'Lunar Cutter', 'Sword', 4, 'Equipped Resonator gains 6 stack(s) of Oath upon entering the battlefield. Each stack increases ATK by 2%, up to 6 stacks. This effect can be triggered 1 time(s) every 12s. Oath reduces by 1 stack(s) every 2s. Equipped Resonator gains an additional 6 stack(s) of Oath upon defeating an enemy.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (42, 'Marcato', 'Gauntlets', 4, 'When Resonance Skill is released, restores Concerto Energy by 8. This effect can be triggered 1 time every 20s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (43, 'Meditations on Mercy', 'Broadblade', 4, 'Dealing DMG to enemies with Negative Statuses increases the wielder''s ATK by 4% for 10s. This effect can be triggered 1 time per second, stackable up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (44, 'Novaburst', 'Pistols', 4, 'When the Resonator dashes or dodges, increases ATK by 4%, stacking up to 3 time(s). This effect lasts for 8s.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (45, 'Ocean''s Gift', 'Rectifier', 4, 'Dealing DMG to enemies with Spectro Frazzle increases the wielder''s Spectro DMG by 6%, gaining 1 stack per second for 6s, stacking up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (46, 'Overture', 'Sword', 4, 'When Resonance Skill is released, restores Concerto Energy by 8. This effect can be triggered 1 time every 20s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (47, 'Pistols#26', 'Pistols', 4, 'When the Resonator takes no damage, increases ATK by 6% every 5s, stacking up to 2 time(s). This effect lasts for 8s. When the Resonator takes damage, reduces the number of stacks by 1 and restores their HP by 5%.', 387, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (48, 'Rectifier#25', 'Rectifier', 4, 'When Resonance Skill is released, if the Resonator''s HP is below 60%, restores their HP by 5%. This effect can be triggered 1 time(s) every 8s; if the Resonator''s HP is above 60%, increases ATK by 12%, lasting for 10s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (49, 'Relativistic Jet', 'Pistols', 4, 'Casting the Resonance Skill grants 6 Resonance Energy and increases ATK by 10%, lasting for 16s. This effect can be triggered once every 20s.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (50, 'Romance in Farewell', 'Pistols', 4, 'Dealing DMG to enemies with Negative Statuses increases the wielder''s ATK by 4% for 10s. This effect can be triggered 1 time per second, stackable up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (51, 'Somnoire Anchor', 'Sword', 4, 'Gain 1 stack of Hiss when dealing damage to the target, with 1 stack generated every 1s. Hiss: each stack increases the wielder''s ATK by 2% for 3s, stacking up to 10 times. Switching off the wielder clears all stacks. Gaining 10 stacks increases the wielder''s Crit. Rate by 6%.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (52, 'Stonard', 'Gauntlets', 4, 'When Resonance Skill is released, increases the caster''s Resonance Liberation DMG Bonus by 18%, lasting for 15s.', 412, 'CRIT Rate', 20.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (53, 'Sword#18', 'Sword', 4, 'When Equipped Resonator''s HP drops below 40%, increases Heavy Attack DMG by 18% and restores HP by 5% upon dealing Heavy Attack DMG. This effect can be triggered 1 time(s) every 8s.', 387, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (54, 'Thunderbolt', 'Pistols', 4, 'When hitting a target with Basic Attacks or Heavy Attacks, increases Resonance Skill DMG Bonus by 7%, stacking up to 3 time(s). This effect lasts for 10s and can be triggered 1 time(s) every 1s.', 387, 'ATK', 36.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (55, 'Undying Flame', 'Pistols', 4, 'When Intro Skill is released, increases Resonance Skill DMG Bonus by 20% for 15s.', 412, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (56, 'Variation', 'Rectifier', 4, 'When Resonance Skill is released, restores Concerto Energy by 8. This effect can be triggered 1 time every 20s.', 337, 'Energy Reg.', 51.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (57, 'Waltz in Masquerade', 'Rectifier', 4, 'Dealing DMG to enemies with Negative Statuses increases the wielder''s ATK by 4% for 10s. This effect can be triggered 1 time per second, stackable up to 4 times.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (58, 'Waning Redshift', 'Broadblade', 4, 'Casting the Resonance Skill grants 6 Resonance Energy and increases ATK by 10%, lasting for 16s. This effect can be triggered once every 20s.', 462, 'ATK', 18.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (59, 'Beguiling Melody', 'Broadblade', 3, 'When Intro Skill is cast, restores 4 Concerto Energy. When Outro Skill is cast, restores 4 Resonance Energy.', 300, 'ATK', 30.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (60, 'Broadblade of Night', 'Broadblade', 3, 'When Intro Skill is released, increases ATK by 8%, lasting for 10s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (61, 'Broadblade of Voyager', 'Broadblade', 3, 'When Resonance Skill is released, restores Resonance Energy by 8. This effect can be triggered 1 time every 20s.', 300, 'Energy Reg.', 32.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (62, 'Gauntlets of Night', 'Gauntlets', 3, 'When Intro Skill is released, increases ATK by 8%, lasting for 10s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (63, 'Gauntlets of Voyager', 'Gauntlets', 3, 'When Resonance Skill is released, restores Resonance Energy by 8. This effect can be triggered 1 time(s) every 20s.', 325, 'DEF', 30.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (64, 'Guardian Broadblade', 'Broadblade', 3, 'The damage of Basic Attack and Heavy Attack is increased by 12%.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (65, 'Guardian Gauntlets', 'Gauntlets', 3, 'Increases Resonance Liberation DMG Bonus by 12%.', 300, 'DEF', 38.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (66, 'Guardian Pistols', 'Pistols', 3, 'Increases Resonance Skill DMG Bonus by 12%.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (67, 'Guardian Rectifier', 'Rectifier', 3, 'Increases Basic Attack and Heavy Attack DMG Bonus by 12%.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (68, 'Guardian Sword', 'Sword', 3, 'Increases Resonance Skill DMG by 12%.', 300, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (69, 'Originite: Type I', 'Broadblade', 3, 'When Resonance Skill is released, restores HP by 3%. This effect can be triggered 1 time(s) every 12s.', 300, 'DEF', 38.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (70, 'Originite: Type II', 'Sword', 3, 'When Resonance Liberation is released, restores HP by 5%. This effect can be triggered 1 time(s) every 20s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (71, 'Originite: Type III', 'Pistols', 3, 'When Counter Attack is released, restores HP by 1.6%. This effect can be triggered 1 time(s) every 6s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (72, 'Originite: Type IV', 'Gauntlets', 3, 'When hitting a target with Basic Attacks, restores HP by 0.5%. This effect can be triggered 1 time(s) every 3s.', 300, 'CRIT DMG', 40.5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (73, 'Originite: Type V', 'Rectifier', 3, 'When Intro Skill is released, restores HP by 5%. This effect can be triggered 1 time(s) every 20s.', 300, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (74, 'Pistols of Night', 'Pistols', 3, 'When Intro Skill is released, increases ATK by 8%, lasting for 10s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (75, 'Pistols of Voyager', 'Pistols', 3, 'When Resonance Skill is released, restores Resonance Energy by 8. This effect can be triggered 1 time(s) every 20s.', 300, 'ATK', 30.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (76, 'Rectifier of Night', 'Rectifier', 3, 'When Intro Skill is released, increases ATK by 8%, lasting for 10s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (77, 'Rectifier of Voyager', 'Rectifier', 3, 'When Resonance Skill is released, restores Resonance Energy by 8. This effect can be triggered 1 time every 20s.', 300, 'Energy Reg.', 32.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (78, 'Sword of Night', 'Sword', 3, 'When Intro Skill is released, increases ATK by 8%, lasting for 10s.', 325, 'ATK', 24.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (79, 'Sword of Voyager', 'Sword', 3, 'When Resonance Skill is released, restores Resonance Energy by 8. This effect can be triggered 1 time every 20s.', 300, 'Energy Reg.', 32.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (80, 'Tyro Broadblade', 'Broadblade', 2, 'Increases ATK by 5%.', 275, 'ATK', 14.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (81, 'Tyro Gauntlets', 'Gauntlets', 2, 'Increases ATK by 5%.', 275, 'ATK', 14.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (82, 'Tyro Pistols', 'Pistols', 2, 'Increases ATK by 5%.', 275, 'ATK', 14.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (83, 'Tyro Rectifier', 'Rectifier', 2, 'Increases ATK by 5%.', 275, 'ATK', 14.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (84, 'Tyro Sword', 'Sword', 2, 'Increases ATK by 5%.', 275, 'ATK', 14.8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (85, 'Training Broadblade', 'Broadblade', 1, 'Increases ATK by 4%.', 250, 'ATK', 11.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (86, 'Training Gauntlets', 'Gauntlets', 1, 'Increases ATK by 4%.', 250, 'ATK', 11.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (87, 'Training Pistols', 'Pistols', 1, 'Increases ATK by 4%.', 250, 'ATK', 11.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (88, 'Training Rectifier', 'Rectifier', 1, 'Increases ATK by 4%.', 250, 'ATK', 11.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (89, 'Training Sword', 'Sword', 1, 'Increases ATK by 4%.', 250, 'ATK', 11.4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (90, 'Woodland Aria', 'Pistols', 5, 'ATK is increased by 12%. Inflicting Aero Erosion on the target gives 24% Aero DMG Bonus for 10s. Hitting targets with Aero Erosion reduces their Aero RES by 10% for 20s. Effects of the same name cannot be stacked.', 500, 'CRIT Rate', 36, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (91, 'Defier''s Thorn', 'Sword', 5, 'Max HP is increased by 12%. 15s after casting Intro Skill or Basic Attacks, ignores 8% of the target''s DEF when dealing damage. If the target has at least 1 stack of Aero Erosion, the DMG taken by the target is Amplified by 20%.');
INSERT INTO wuwa_echoes (id, name, two_piece_bonus, full_set_bonus, created_at, updated_at)
VALUES 
    (1, 'Freezing Frost', 'Glacio DMG increases by 10%.', 'Upon using Basic Attack or Heavy Attack, Glacio DMG increases by 10%, stacking up to three times, lasting for 15s.', NOW(), NOW()),
    (2, 'Molten Rift', 'Fusion DMG% increases by 10%.', 'Upon using Resonance Skill, Fusion DMG increases by 30% for 15s.', NOW(), NOW()),
    (3, 'Void Thunder', 'Electro DMG increases by 10%.', 'Upon using Heavy Attack or Resonance Skill, Electro DMG increases by 15%, stacking up to 2 times, each stack lasting for 15s.', NOW(), NOW()),
    (4, 'Sierra Gale', 'Aero DMG increase by 10%.', 'Upon using Intro Skill, Aero DMG increases by 30% for 15s.', NOW(), NOW()),
    (5, 'Celestial Light', 'Spectro DMG increases by 10%.', 'Upon using Intro Skill, Spectro DMG increases by 30% for 15s.', NOW(), NOW()),
    (6, 'Havoc Eclipsee', 'Havoc DMG increases by 10%.', 'Upon using Basic Attack or Heavy Attack, Havoc DMG increases by 7.5%, stacking up to four times for 15s.', NOW(), NOW()),
    (7, 'Rejuvenating Glow', 'Healing increases by 10%.', 'Upon healing allies, increase ATK of the entire team by 15%, lasting 30s.', NOW(), NOW()),
    (8, 'Moonlit Clouds', 'Energy Regen increases by 10%.', 'Upon using Outro Skill, ATK of the next Resonator increases by 22.5% for 15s.', NOW(), NOW()),
    (9, 'Lingering Tunes', 'ATK increases by 10%.', 'While on the field, ATK increases by 5% every 1.5s, stacking up to 4 times. Outro Skill DMG increases by 60%.', NOW(), NOW()),
    (10, 'Frosty Resolve', 'Resonance Skill DMG +12%.', 'Casting Resonance Skill grants 22.5% Glacio DMG Bonus for 15s and casting Resonance Liberation increases Resonance Skill DMG by 18%, lasting for 5s. This effect stacks up to 2 times.', NOW(), NOW()),
    (11, 'Eternal Radiance', 'Spectro DMG +10%.', 'Inflicting enemies with Spectro Frazzle increases Crit. Rate by 20% for 15s. Attacking enemies with 10 stacks of Spectro Frazzle grants 15% Spectro DMG Bonus for 15s.', NOW(), NOW()),
    (12, 'Midnight Veil', 'Havoc DMG +10%.', 'When Outro Skill is triggered, deal additional 480% Havoc DMG to surrounding enemies, considered Outro Skill DMG, and grant the incoming Resonator 15% Havoc DMG Bonus for 15s.', NOW(), NOW()),
    (13, 'Empyrean Anthem', 'Energy Regen +10%.', 'Increase the Resonator''s Coordinated Attack DMG by 80%. Upon a critical hit of Coordinated Attack, increase the active Resonator''s ATK by 20% for 4s.', NOW(), NOW()),
    (14, 'Tidebreaking Courage', 'Energy Regen +10%.', 'Increase the Resonator''s ATK by 15%. Reaching 250% Energy Regen increases all Attribute DMG by 30% for the Resonator.', NOW(), NOW()),
    (15, 'Gusts of Welkin', 'Aero DMG +10%.', 'Inflicting Aero Erosion upon enemies increases Aero DMG for all Resonators in the team by 15%, and for the Resonator triggering this effect by an additional 15%, lasting for 20s.', NOW(), NOW()),
    (16, 'Flaming Clawprint', 'Fusion DMG +10%.', 'Casting Resonance Liberation grants all Resonators in the team 15% Fusion DMG Bonus and the caster 20% Resonance Liberation DMG Bonus, lasting for 35s.', NOW(), NOW()),
    (17, 'Windward Pilgrimage', 'Aero DMG +10%.', 'Hitting a target with Aero Erosion increases Crit. Rate by 10% and grants 30% Aero DMG Bonus, lasting for 10s.', NOW(), NOW());

INSERT INTO wuwa_stats(id, four_cost_echo_stat, three_cost_echo_stat, one_cost_echo_stat, sub_stats_priority)
VALUES 
    (1, 'CRIT DMG / CRIT Rate', 'Aero DMG / Energy Regen', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Aalto
    (2, 'Healing Bonus / HP%', 'Energy Regen', 'HP%', 'Energy Regeneration (Until Breakpoint) > HP% > Flat HP'), -- Baizhi
    (3, 'CRIT DMG / CRIT Rate', 'Energy Regen / Fusion DMG', 'ATK%', 'Energy Regeneration (No more than 280% in menu or 180% from Echoes/Weapons) > CRIT RATE = CRIT DMG > Basic ATK DMG% > ATK% > Flat ATK'), -- Brant
    (4, 'CRIT DMG / CRIT Rate', 'Electro DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK = Liberation DMG%'), -- Calcharo
    (5, 'CRIT DMG / CRIT Rate', 'Havoc DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Basic DMG% > Flat ATK'), -- Camellya
    (6, 'CRIT DMG / CRIT Rate', 'Havoc DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Satisfied) > CRIT RATE = CRIT DMG > ATK% > Flat ATK = Basic ATK DMG%'), -- Cantarella
    (7, 'CRIT DMG / CRIT Rate', 'Glacio DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Satisfied) >= CRIT RATE = CRIT DMG > ATK% > Resonance Skill DMG > Flat ATK'), -- Carlotta
    (8, 'CRIT DMG / CRIT Rate', 'Fusion DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK = Resonance Skill DMG%'),  -- Changli
    (9, 'CRIT DMG / CRIT Rate', 'Fusion DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Chixia
    (10, 'CRIT DMG / CRIT Rate', 'Havoc DMG', 'ATK%', 'CRIT RATE = CRIT DMG > ATK% > Flat ATK > Heavy DMG% > Resonance Skill DMG%'), -- Danjin
    (11, 'CRIT DMG / CRIT Rate', 'Fusion DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK > Basic DMG%'), -- Encore 
    (12, 'CRIT DMG / CRIT Rate', 'Aero DMG / Energy Regen', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK > Heavy DMG% > Liberation DMG%'), -- Jianxin 
    (13, 'CRIT DMG / CRIT Rate', 'Spectro DMG / ATK', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Resonance Skill DMG% > Flat ATK'), -- Jinhsi
    (14, 'CRIT DMG / CRIT Rate', 'Aero DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK = Heavy DMG%'), -- Jiyan
    (15, 'CRIT DMG / CRIT Rate', 'Glacio DMG / ATK%', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Lingyang
    (16, 'CRIT DMG / CRIT Rate', 'Electro DMG / Energy Regen', 'ATK%', 'Energy Regeneration (Until Satisfied) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Lumi
    (17, 'CRIT DMG / CRIT Rate', 'Fusion DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Mortefi
    (18, 'CRIT DMG / CRIT Rate', 'Spectro DMG', 'ATK%', 'Energy Regeneration (Until Satisfied) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK > Heavy ATK DMG%'), -- Phoebe
    (19, 'CRIT DMG / CRIT Rate', 'Havoc DMG', 'ATK%', 'Energy Regeneration (Until Satisfied) >= CRIT RATE > CRIT DMG > ATK% > Heavy ATK DMG% > Flat ATK'), -- Roccia
    (20, 'CRIT DMG / CRIT Rate', 'Aero DMG', 'ATK%', 'Energy Regeneration (Until Satisfied) > CRIT RATE = CRIT DMG > ATK% > Skill DMG% > Flat ATK'), -- Rover (Aero)
    (21, 'CRIT DMG / CRIT Rate', 'Havoc DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Rover (Havoc)
    (22, 'CRIT DMG / CRIT Rate', 'Spectro DMG', 'ATK%', ' CRIT RATE = CRIT DMG > ATK% > Flat ATK > Liberation DMG%'), -- Rover (Spectro)
    (23, 'CRIT DMG / CRIT Rate', 'Glacio DMG', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Sanhua
    (24, 'DEF%', 'Energy Regen / DEF%', 'DEF%', 'ENERGY Regen% (until breakpoint) > Crit rate / Crit dmg > DEF% > DEF'), -- Taoqi
    (25, 'CRIT DMG / Healing Bonus / HP%', 'Energy Regen / Spectro DMG / HP%', 'HP%', 'Energy Regeneration (Until Breakpoint) CRIT DMG% >= ULT DMG% > HP% > FLAT HP'), -- Sk
    (26, 'Healing Bonus / ATK', 'Energy Regen', 'ATK%', 'Energy Regeneration (Until Breakpoint) > ATK% > Flat ATK'), -- Verina
    (27, 'CRIT DMG / CRIT Rate', 'Electro DMG / ATK', 'ATK%', 'CRIT RATE = CRIT DMG > ATK% > Flat ATK = Liberation DMG%'), -- Xiangli Yao
    (28, 'CRIT DMG / CRIT Rate', 'Aero DMG / Energy Regen', 'ATK%', ' Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK > Liberation DMG% > Basic DMG%'), -- Yangyang
    (29, 'CRIT DMG / CRIT Rate', 'Electro DMG / ATK', 'ATK%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > ATK% > Flat ATK = Resonance Skill DMG%'), -- Yinlin
    (30, 'Healing Bonus', 'Energy Regen / ATK%', 'ATK%', 'Energy Regeneration (Until satisfied) = ATK% > FLAT ATK'), -- Youhu
    (31, 'CRIT DMG / CRIT Rate', 'Electro DMG / DEF', 'DEF%', 'Energy Regeneration (Until Breakpoint) >= CRIT RATE = CRIT DMG > DEF% > Flat DEF'), -- Yuanwu
    (32, 'CRIT DMG / CRIT Rate', 'Spectro DMG', 'ATK%', 'Energy Regeneration (Until Satisfied) >= CRIT RATE = CRIT DMG > ATK% > Heavy ATK DMG% > Flat ATK'), -- Zani
    (33, 'CRIT DMG / CRIT Rate', 'Glacio DMG', 'ATK%', 'Energy Regeneration (Until Comfortable) > CRIT RATE = CRIT DMG > Basic DMG% = ATK% > Flat ATK'), -- Zhezhi
    (34, 'CRIT DMG / CRIT Rate', 'Aero DMG / ATK%', 'ATK%', 'CRIT RATE = CRIT DMG > ATK% > Flat ATK'), -- Ciaccona
    (35, 'CRIT DMG / CRIT Rate', 'Aero DMG / HP%', 'HP%', 'Energy Regen (Until Satisfied) > CRIT RATE = CRIT DMG > HP% = Basic ATK DMG% > Liberation DMG% > Flat HP'); -- Cartethiya
INSERT INTO wuwa_builds(id, character_id, stats_id, best_primary_echo) 
VALUES
    (1, 25, 1, 'Nightmare: Feilian Beringal'), -- Aalto
    (2, 26, 2, 'Fallacy of No Return'), -- Baizhi
    (3, 3, 3, 'Nightmare: Inferno Rider / Dragon of Dirge'), -- Brant
    (4, 4, 4, 'Nightmare: Thundering Mephis / Nightmare: Tempest Mephis') -- Calcharo
    (5, 5, 5, 'Nightmare: Crownless'), -- Camellya
    (6, 6, 6, 'Lorelei'), -- Cantarella
    (7, 7, 7, 'Sentry Construct'), -- Carlotta
    (8, 8, 8, 'Nightmare: Inferno Rider'), -- Changli
    (9, 2, 9, 'Nightmare: Inferno Rider'), -- Chixia
    (10, 27, 10, 'Dreamless / Impermanence Heron'), -- Danjin
    (11, 9, 11, 'Nightmare: Inferno Rider'), -- Encore
    (12, 10, 12, 'Nightmare: Feilian Beringal'), -- Jianxin
    (13, 11, 13, 'Jué'), -- Jinhsi
    (14, 12, 14, 'Nightmare: Feilian Beringal'), -- Jiyan
    (15, 13, 15, 'Mech Abomination / Sentry Construct'), -- Lingyang
    (16, 28, 16, 'Impermanence Heron / Nightmare: Thundering Mephis') -- Lumi
    (17, 29, 17, 'Impermanence Heron / Hecate'), -- Mortefi
    (18, 14, 18, 'Nightmare: Mourning Aix'), -- Phoebe
    (19, 15, 19, 'Nightmare: Impermanence Heron / Impermanence Heron'), -- Roccia
    (20, 16, 20, 'Fallacy of No Return'), -- Rover (Aero) 
    (21, 17, 21, 'Dreamless'), -- Rover (Havoc)
    (22, 18, 22, 'Impermanence Heron / Nightmare: Mourning Aix'), -- Rover (Spectro)
    (23, 30, 23, 'Impermanence Heron'), -- Sanhua
    (24, 31, 24, 'Fallacy of No Return'), -- Taoqi
    (25, 19, 25, 'Fallacy of No Return'), -- Sk
    (26, 20, 26, 'Fallacy of No Return'), -- Verina
    (27, 21, 27, 'Nightmare: Thundering Mephis'), -- Xiangli Yao
    (28, 1, 28, 'Impermanence Heron / Nightmare: Feilian Beringal'), -- Yangyang
    (29, 22, 29, 'Impermanence Heron / Nightmare: Tempest Mephis'), -- Yinlin
    (30, 32, 30, 'Fallacy of No Return / Impermanence Heron'), -- Youhu
    (31, 33, 31, 'Fallacy of No Return / Impermanence Heron / Nightmare: Tempest Mephis'), -- Yuanwu
    (32, 23, 32, 'Capitaneus'), -- Zani
    (33, 24, 33, 'Impermanence Heron / Hecate'), -- Zhezhi
    (34, 34, 34, 'Reminiscence: Fleurdelys'), -- Ciaccona
    (35, 35, 35, 'Reminiscence: Fleurdelys'); -- Cartethiya

INSERT INTO build_wuwa_weapons (wuwa_build_id, wuwa_weapon_id)
VALUES 
    (1, 12), (1, 44), (1, 49), (1, 47), (1, 54), (1, 55), -- Aalto
    (2, 13), (2, 56), (2, 26), (2, 77), (2, 48), -- Baizhi
    (3, 17), (3, 46), (3, 79), -- Brant
    (4, 9), (4, 2), (4, 18), (4, 23), (4, 58), (4, 36), (4, 24), -- Calcharo
    (5, 10), (5, 7), (5, 51), (5, 40), (5, 29), (5, 32), (5, 41), -- Camellya
    (6, 20), (6, 11), (6, 14), (6, 6), (6, 22), (6, 38), -- Cantarella
    (7, 15), (7, 12), (7, 54), (7, 55), (7, 44), (7, 47), -- Carlotta
    (8, 3), (8, 7), (8, 10), (8, 29), (8, 41), -- Changli
    (9, 15), (9, 12), (9, 54), (9, 49), (9, 44), (9, 55), -- Chixia
    (10, 3), (10, 10), (10, 7), (10, 29), (10, 32), (10, 51), -- Danjin
    (11, 14), (11, 11), (11, 6), (11, 22), (11, 34), (11, 38) -- Encore
    (12, 19), (12, 1), (12, 52), (12, 27), (12, 37), -- Jianxin
    (13, 2), (13, 18), (13, 9), (13, 24), (13, 58),  -- Jinhsi
    (14, 18), (14, 2), (14, 9), (14, 23), (14, 58), -- Jiyan
    (15, 19), (15, 1), (15, 27), (15, 37), -- Lingyang
    (16, 2), (16, 18), (16, 9), (16, 23), (16, 58), -- Lumi
    (17, 12), (17, 15), (17, 49), (17, 44), (17, 54), -- Mortefi
    (18, 8), (18, 14), (18, 22), (18, 6), (18, 57), -- Phoebe
    (19, 16), (19, 19), (19, 1), (19, 27), (19, 37), -- Roccia
    (20, 5), (20, 46), (20, 79), -- Rover (Aero)
    (21, 10), (21, 3), (21, 7), (21, 51), (21, 29), (21, 40), -- Rover (Havoc)
    (22, 3), (22, 7), (22, 41), (22, 32), (22, 29), -- Rover (Spectro)
    (23, 3), (23, 10), (23, 7), (23, 29), (23, 32), (23, 41), -- Sanhua
    (24, 31), (24, 30), -- Taoqi
    (25, 13), (25, 56), (25, 48), (25, 77), -- Sk
    (26, 56), (26, 13), (26, 26), (25, 48), (25, 77), -- Verina
    (27, 19), (27, 1), (27, 52), (27, 27), (27, 37), -- Xiangli Yao
    (28, 3), (28, 7), (28, 40), (28, 32), (28, 41), -- Yangyang
    (29, 14), (29, 11), (29, 6), (29, 22), (29, 38), -- Yinlin
    (30, 42), (30, 35), (30, 1), (30, 27), -- Youhu
    (31, 19), (31, 21), (31, 52), (31, 65), (31, 72), -- Yuanwu
    (32, 4), (32, 16), (32, 19), (32, 1), (32, 39), -- Zani
    (33, 11), (33, 14), (33, 6), (33, 22), (33, 38), -- Zhezhi
    (34, 90), (34, 15), (34, 12), (34, 50), (34, 49), -- Ciaccona
    (35, 91), (35, 10), (35, 7), (35, 68); -- Cartethiya

INSERT INTO build_wuwa_echoes (wuwa_build_id, wuwa_echoes_id)
VALUES 
    (1, 4), -- Aalto
    (2, 7), -- baizhi
    (3, 2), (3, 14), -- Brant
    (4, 3), (4, 9), -- Calcharo
    (5, 6), (5, 9), -- Camellya
    (6, 12), (6, 8), -- Cantarella
    (7, 10), (7, 9), -- Carlotta
    (8, 2), -- Changli
    (9, 2),  -- Chixia
    (10, 6), (10, 8), -- Danjin
    (11, 2), -- Encore
    (12, 8), (12, 7), (12, 4) -- Jianxin
    (13, 5), -- Jinhsi
    (14, 4), -- Jiyan
    (15, 9), (15, 10), -- Lingyang
    (16, 8), (16, 3), -- Lumi
    (17, 8), (17, 13), -- Mortefi
    (18, 11),  -- Phoebe
    (19, 12), (19, 8), -- Roccia
    (20, 7), -- Rover (Aero)
    (21, 6), -- Rover (Havoc)
    (22, 8), (22, 11), -- Rover (Spectro)
    (23, 8), -- Sanhua
    (24, 7), -- Taoqi
    (25, 7), -- Sk
    (26, 7), -- Verina
    (27, 3), -- Xiangli Yao
    (28, 8), (28, 4) -- Yangyang
    (29, 8), (29, 3), (29, 13), -- Yinlin
    (30, 7), (30, 8), -- Youhu
    (31, 7), (31, 8), (31, 13), -- Yuanwu
    (32, 11), -- Zani
    (33, 8), (33, 13), -- Zhezhi 
    (34, 15), -- Ciaccona  
    (35, 17); -- Cartethiya


