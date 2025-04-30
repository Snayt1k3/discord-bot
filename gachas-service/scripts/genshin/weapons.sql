INSERT INTO weapons (
    id,
    name,
    type,
    rarity,
    base_atk,
    sub_stat,
    sub_value,
    passive,
) VALUES
(1, "Redhorn Stonethresher", "Claymore", 5, 542, "CRIT DMG", 88.2, "DEF is increased by 28%. Normal and Charged Attack DMG is increased by 40% of DEF." ) -- Redhorn Stonethresher (Itto signature)
(2, "Serpent Spine", "Claymore", 4, 510, "CRIT Rate", 27.6, "Every 4s a character is on the field, they will deal 6% more DMG and take 3% more DMG. This effect has a maximum of 5 stacks and will not be reset if the character leaves the field, but will be reduced by 1 stack when the character takes DMG."),
(3, "Skyward Pride", "Claymore", 5, 674, "Energy Recharge", 36.8, "Increases all DMG by 8%. After using an Elemental Burst, a vacuum blade that does 80% of ATK as DMG to opponents along its path will be created when Normal or Charged Attacks hit. Lasts for 20s or 8 vacuum blades."),
(4, "Whiteblind", "Claymore", 4, 510, "DEF%", 51.7, "On hit, Normal or Charged Attacks increase ATK and DEF by 6% for 6s. Max 4 stacks. This effect can only occur once every 0.5s."),
(5, "Kagura's Verity", "Catalyst", 5, 608, "CRIT DMG", 66.2, "Gains the Kagura Dance effect when using an Elemental Skill, causing the Elemental Skill DMG of the character wielding this weapon to increase by 12% for 16s. Max 3 stacks. This character will gain 12% All Elemental DMG Bonus when they possess 3 stacks."),
(6, "Lost Prayer to the Sacred Winds", "Catalyst", 5, 608, "CRIT Rate", 33,1, "Increases Movement SPD by 10%. When in battle, gain an 8% Elemental DMG Bonus every 4s. Max 4 stacks. Lasts until the character falls or leaves combat."),
(7, "Surf's Up", "Catalyst", 5, 542, "CRIT DMG", 88.2, "Max HP increased by 20%. Once every 15s, for the 14s after using an Elemental Skill: Gain 4 Scorching Summer stacks. Each stack increases Normal Attack DMG by 12%. For the duration of the effect, once every 1.5s, lose 1 stack after a Normal Attack hits an opponent; once every 1.5s, gain 1 stack after triggering a Vaporize reaction on an opponent. Max 4 Scorching Summer stacks." ),
(8, "Tome of the Eternal Flow", "Catalyst" 5, 542, "CRIT DMG", 88.2, "HP is increased by 16%. When current HP increases or decreases, Charged Attack DMG will be increased by 14% for 4s. Max 3 stacks. This effect can be triggered once every 0.3s. When the character has 3 stacks or a third stack's duration refreshes, 8 Energy will be restored. This Energy restoration effect can be triggered once every 12s."),
(9, "Cashflow Supervision", "Catalyst", 5, 674, "CRIT Rate", 22.1, "ATK is increased by 16%. When current HP increases or decreases, Normal Attack DMG will be increased by 16% and Charged Attack DMG will be increased by 14% for 4s. Max 3 stacks. This effect can be triggered once every 0.3s. When the wielder has 3 stacks, ATK SPD will be increased by 8%."),
(10, "Skyward Atlas", "Catalyst", 5, 674, "ATK%", 33.1, "Increases Elemental DMG Bonus by 12%. Normal Attack hits have a 50% chance to earn the favor of the clouds, which actively seek out nearby opponents to attack for 15s, dealing 160% ATK DMG. Can only occur once every 30s.")
(11, "Memory of Dust", "Catalyst", 5, 608, "ATK%", 49.6, "Increases Shield Strength by 20%. Scoring hits on opponents increases ATK by 4% for 8s. Max 5 stacks. Can only occur once every 0.3s. While protected by a shield, this ATK increase effect is increased by 100%.")
(12, "Tulaytullah's Remembrance", "Catalyst", 5, 674, "CRIT DMG", 44.1, "Normal Attack SPD is increased by 10%. After the wielder unleashes an Elemental Skill, Normal Attack DMG will increase by 4.8% every second for 14s. After hitting an opponent with a Normal Attack during this duration, Normal Attack DMG will be increased by 9.6%. This increase can be triggered once every 0.3s. The maximum Normal Attack DMG increase per single duration of the overall effect is 48%. The effect will be removed when the wielder leaves the field, and using the Elemental Skill again will reset all DMG buffs."),
(13, "The Widsith", "Catalyst", 4, 510, "CRIT DMG", 55.1, "When the character takes the field, they will gain a random theme song for 10s. This can only occur once every 30s. Recitative: ATK is increased by 60%. Aria: Increases all Elemental DMG by 48%. Interlude: Elemental Mastery is increased by 240."),
(14, "Solar Pearl", "Catalyst", 4, 510, "CRIT Rate", 27.6, "Normal Attack hits increase Elemental Skill and Elemental Burst DMG by 20% for 6s. Likewise, Elemental Skill or Elemental Burst hits increase Normal Attack DMG by 20% for 6s."),
(15, "Black Tassel", "Polearm", 3, 354, "HP%", 46.9, "Increases DMG against slimes by 40%."),
(16, "Favonius Lance", "Polearm", 4, 565,"Energy Recharge", 30.6, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s." )
(17, "Peak Patrol Song", "Sword", 5, 542, "DEF%", 82.7, "Gain 'Ode to Flowers' after Normal or Plunging Attacks hit an opponent: DEF increases by 8% and gain a 10% All Elemental DMG Bonus for 6s. Max 2 stacks. Can trigger once per 0.1s. When this effect reaches 2 stacks or the 2nd stack's duration is refreshed, increase all nearby party members' All Elemental DMG Bonus by 8% for every 1,000 DEF the equipping character has, up to a maximum of 25.6%, for 15s."),
(18, "Uraku Misugiri", "Sword", 5, 542, "CRIT DMG", 88.2, "Normal Attack DMG is increased by 16% and Elemental Skill DMG is increased by 24%. After a nearby active character deals Geo DMG, the aforementioned effects increase by 100% for 15s. Additionally, the wielder's DEF is increased by 20%."),
(19, "Cinnabar Spindle", "Sword", 4, 454, "DEF%", 69, "Elemental Skill DMG is increased by 40% of DEF. The effect will be triggered no more than once every 1.5s and will be cleared 0.1s after the Elemental Skill deals DMG."),
(20, "Harbinger of Dawn", "Sword", 3, 401, "CRIT DMG", 46.9, "When HP is above 90%, increases CRIT Rate by 14%."),
(21, "Favonius Warbow", "Bow", 4, 454, "Energy Recharge", 61.3, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s."),
(22, "Verdict", "Claymore", 5, 674, "CRIT Rate", 22.1, "Increases ATK by 20%. When party members obtain Elemental Shards from Crystallize reactions, the equipping character will gain 1 Seal, increasing Elemental Skill DMG by 18%. The Seal lasts for 15s, and the equipper may have up to 2 Seals at once. All of the equipper's Seals will disappear 0.2s after their Elemental Skill deals DMG."),
(23, "A Thousand Blazing Suns", "Claymore", 5, 741, "CRIT Rate", 11, "Gain the 'Scorching Brilliance' effect when using an Elemental Skill or Burst: CRIT DMG increased by 20% and ATK increased by 28% for 6s. This effect can trigger once every 10s. While a 'Scorching Brilliance' instance is active, its duration is increased by 2s after Normal or Charged attacks deal Elemental DMG. This effect can trigger once every second, and the max duration increase is 6s. Additionally, when the equipping character is in the Nightsoul's Blessing state, 'Scorching Brilliance' effects are increased by 75%, and its duration will not count down when the equipping character is off-field.")
(24, "Primordial Jade Cutter", "Sword", 5, 542, "CRIT Rate", 44.1, "HP increased by 20%. Additionally, provides an ATK Bonus based on 1.2% of the wielder's Max HP.")
(25, "Favonius Sword","Sword", 4, 454, "Energy Recharge", 61.3, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s."),
(26, "Hunter's Path","Bow", 5, 542, "CRIT Rate", 44.1, "Gain 12% All Elemental DMG Bonus. Obtain the Tireless Hunt effect after hitting an opponent with a Charged Attack. This effect increases Charged Attack DMG by 160% of Elemental Mastery. This effect will be removed after 12 Charged Attacks or 10s. Only 1 instance of Tireless Hunt can be gained every 12s."),
(27, "Astral Vulture's Crimson Plumage" "Bow", 5, 608, "CRIT DMG", 66.2, "For 12s after triggering a Swirl reaction, ATK increases by 24%. In addition, when 1/2 or more characters in the party are of a different Elemental Type from the equipping character, the DMG dealt by the equipping character's Charged Attacks is increased by 20%/48% and Elemental Burst DMG dealt is increased by 10%/24%.")
(28, "Aqua Simulacra","Bow", 5, 542, "CRIT DMG", 88.2, "HP is increased by 16%. When there are opponents nearby, the DMG dealt by the wielder of this weapon is increased by 20%. This will take effect whether the character is on-field or not.")
(29, "Skyward Harp", "Bow", 5, 674, "CRIT Rate", 22.1, "Increases CRIT DMG by 20%. Hits have a 60% chance to inflict a small AoE attack, dealing 125% Physical ATK DMG. Can only occur once every 4s.")
(30, "Slingshot", "Bow", 3, 354, "CRIT Rate", 31.2, "If a Normal or Charged Attack hits a target within 0.3s of being fired, increases DMG by 36%. Otherwise, decreases DMG by 10%.")
(31, "Elegy for the End","Bow", 5, 608, "Energy Recharge", 55.1, "A part of the 'Millennial Movement' that wanders amidst the winds. Increases Elemental Mastery by 60. When the Elemental Skills or Elemental Bursts of the character wielding this weapon hit opponents, that character gains a Sigil of Remembrance. This effect can be triggered once every 0.2s and can be triggered even if said character is not on the field. When you possess 4 Sigils of Remembrance, all of them will be consumed and all nearby party members will obtain the 'Millennial Movement: Farewell Song' effect for 12s. 'Millennial Movement: Farewell Song' increases Elemental Mastery by 100 and increases ATK by 20%. Once this effect is triggered, you will not gain Sigils of Remembrance for 20s. Of the many effects of the 'Millennial Movement,' buffs of the same type will not stack.")
(32, "Sacrificial Bow", "Bow", 4, 565, "Energy Recharge", 30.6, "After damaging an opponent with an Elemental Skill, the skill has a 40% chance to end its own CD. Can only occur once every 30s.")
(33, "A Thousand Floating Dreams", "Catalyst", 5, 542, "Elemental Mastery", 265, "Party members other than the equipping character will provide the equipping character with buffs based on whether their Elemental Type is the same as the latter or not. If their Elemental Types are the same, increase Elemental Mastery by 32. If not, increase the equipping character's DMG Bonus from their Elemental Type by 10%. Each of the aforementioned effects can have up to 3 stacks. Additionally, all nearby party members other than the equipping character will have their Elemental Mastery increased by 40. Multiple such effects from multiple such weapons can stack."),
(34, "Sunny Morning Sleep-In", "Catalyst", 5, 542, "Elemental Mastery", 265, "Elemental Mastery increases by 120 for 6s after triggering Swirl. Elemental Mastery increases by 96 for 9s after the wielder's Elemental Skill hits an opponent. Elemental Mastery increases by 32 for 30s after the wielder's Elemental Burst hits an opponent."),
(35, "Sacrificial Fragments", "Catalyst", 4, 454, "Elemental Mastery", 221, "After damaging an opponent with an Elemental Skill, the skill has a 40% chance to end its own CD. Can only occur once every 30s."),
(36, "Dialogues of the Desert Sages", "Polearm", 4, 510, "HP%", 41.3, "When the wielder performs healing, restore 8 Energy. This effect can be triggered once every 10s and can occur even when the character is not on the field.")
(37, "Light of Foliar Incision ", "Sword", 5, 542, "CRIT DMG", 88.2, "CRIT Rate is increased by 4%. After Normal Attacks deal Elemental DMG, the Foliar Incision effect will be obtained, increasing DMG dealt by Normal Attacks and Elemental Skills by 120% of Elemental Mastery. This effect will disappear after 28 DMG instances or 12s. You can obtain Foliar Incision once every 12s.")
(38, "Mistsplitter Reforged", "Sword", 5, 674, "CRIT DMG", 44.1, "Gain a 12% Elemental DMG Bonus for all elements and receive the might of the Mistsplitter's Emblem. At stack levels 1/2/3, the Mistsplitter's Emblem provides a 8/16/28% Elemental DMG Bonus for the character's Elemental Type. The character will obtain 1 stack of Mistsplitter's Emblem in each of the following scenarios: Normal Attack deals Elemental DMG (stack lasts 5s), casting Elemental Burst (stack lasts 10s); Energy is less than 100% (stack disappears when Energy is full). Each stack's duration is calculated independently.")
(39, "Iron Sting", "Sword", 4, 510, "Elemental Mastery", 165, "Dealing Elemental DMG increases all DMG by 6% for 6s. Max 2 stacks. Can occur once every 1s.")
(40, "Favonius Greatsword","Claymore", 4, 454, "Energy Recharge", 61.3, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s.")
(41, "Jadefall's Splendor", "Catalyst", 5, 608, "HP%", 49.6, "For 3s after using an Elemental Burst or creating a shield, the equipping character can gain the Primordial Jade Regalia effect: Restore 4.5 Energy every 2.5s, and gain 0.3% Elemental DMG Bonus for their corresponding Elemental Type for every 1,000 Max HP they possess, up to 12%. Primordial Jade Regalia will still take effect even if the equipping character is not on the field.")
(42, "Prototype Amber", "Catalyst", 4, 510, "HP%", 41.3, "Using an Elemental Burst regenerates 4 Energy every 2s for 6s. All party members will regenerate 4% HP every 2s for this duration.")
(43, "Favonius Codex", "Catalyst", 4, 510, "Energy Recharge", 45.9, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s."),
(44, "Lumidouce Elegy", "Polearm", 5, 608, "CRIT Rate", 33.1, "ATK increased by 15%. After the equipping character triggers Burning on an opponent or deals Dendro DMG to Burning opponents, the DMG dealt is increased by 18%. This effect lasts for 8s, max 2 stacks. When 2 stacks are reached or when the duration is refreshed at 2 stacks, restore 12 Energy. Energy can be restored this way once every 12s. The 2 aforementioned effects can be triggered even when the character is off-field.")
(45, "Calamity Queller", "Polearm", 5, 741, "ATK%", 16.5, "Gain 12% All Elemental DMG Bonus. Obtain Consummation for 20s after using an Elemental Skill, causing ATK to increase by 3.2% per second. This ATK increase has a maximum of 6 stacks. When the character equipped with this weapon is not on the field, Consummation's ATK increase is doubled."),
(46, "Skyward Spine", "Polearm", 5, 674, "Energy Recharge", 36.8, "Increases CRIT Rate by 8% and increases Normal ATK SPD by 12%. Additionally, Normal and Charged Attacks hits on opponents have a 50% chance to trigger a vacuum blade that deals 40% of ATK as DMG in a small AoE. This effect can occur no more than once every 2s.")
(47, "Fang of the Mountain King", "Claymore", 5, 741, "CRIT Rate", 11, "Gain 1 stack of Canopy's Favor after hitting an opponent with an Elemental Skill. This can be triggered once every 0.5s. After a nearby party member triggers a Burning or Burgeon reaction, the equipping character will gain 3 stacks. This effect can be triggered once every 2s and can be triggered even when the triggering party member is off-field. Canopy's Favor: Elemental Skill and Burst DMG is increased by 10% for 6s. Max 6 stacks. Each stack is counted independently.")
(48, "Earth Shaker", "Claymore", 4, 565, "ATK%", 27.6, "After a party member triggers a Pyro-related reaction, the equipping character's Elemental Skill DMG is increased by 16% for 8s. This effect can be triggered even when the triggering party member is not on the field.")


INSERT INTO build_weapons (
    build_id,
    weapon_id,
    priority
) VALUES

-- Geo
(1, 1, 1), (1, 2, 2), (1, 3, 3), (1, 4, 4),  -- Noelle
(2, 5, 1), (2, 6, 2), (2, 7, 3), (2, 8, 4), (2, 9, 5), (2, 10, 5), (2, 11, 6), (2, 12, 7), (2, 13, 8), (2, 14, 9),  -- Ningguang
(3, 15, 1), (3, 16, 2),  -- Zhongli
(4, 17, 1),  (4, 18, 2), (4, 19, 3), (4, 20, 4),  -- Albedo
(5, 21, 1),  -- Gorou
(6, 1, 1), (6, 2, 2), (6, 3, 3), (6, 4, 4),  -- Itto
(7, 16, 1),  -- Yun Jin
(8, 22, 1), (8, 23, 2), (8, 2, 3),  -- Navia
(9, 18, 1), (9, 24, 2), (9, 20, 3),  -- Chiori
(10, 16, 1), -- Kachina
(11, 17, 1), (11, 25, 2), -- Xilonen

-- Dendro
(12, 26, 1), (12, 27, 2), (12, 28, 3),(12, 29, 4), (12, 30, 5),  -- Tighnari
(13, 31, 1), (13, 32, 2), (13, 21, 3), -- Collei
(14, 33, 1), (14, 34, 2), (14, 5, 4), (14, 35, 5)  -- Nahida
(15, 16, 1), (15, 36, 2), -- yaoyao
(16, 37, 1), (16, 24, 2), (16, 38, 3),  (16, 39, 4), -- Alhaitam
(17, 40, 1), -- Kaveh
(18, 41, 1), (18, 42, 2), (18, 43, 3) -- BaiZhu
(19, 25, 1), -- Kirara
(20, 44, 1), (20, 45, 2), (20, 46, 3), (20, 16, 4),-- Emilie
(21, 47, 1), (21, 2, 2), (21, 28, 3), -- Kinich