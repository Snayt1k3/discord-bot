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
(1, "Redhorn Stonethresher", "Claymore", 5, 542, "CRIT DMG", 88.2, "DEF is increased by 28%. Normal and Charged Attack DMG is increased by 40% of DEF." )
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
(49, "Sacrificial Sword", "Sword", 4, 454, "Energy Recharge", 61.3, "After damaging an opponent with an Elemental Skill, the skill has a 40% chance to end its own CD. Can only occur once every 30s.")
(50, "Haran Geppaku Futsu", "Sword", 5, 608, "CRIT Rate", 33.1, "Obtain 12% All Elemental DMG Bonus. When other nearby party members use Elemental Skills, the character equipping this weapon will gain 1 Wavespike stack. Max 2 stacks. This effect can be triggered once every 0.3s. When the character equipping this weapon uses an Elemental Skill, all stacks of Wavespike will be consumed to gain Rippling Upheaval: each stack of Wavespike consumed will increase Normal Attack DMG by 20% for 8s.")
(51, "Summit Shaper", "Sword", 5, 608, "ATK%", 49.6, "Increases Shield Strength by 20%. Scoring hits on opponents increases ATK by 4% for 8s. Max 5 stacks. Can only occur once every 0.3s. While protected by a shield, this ATK increase effect is increased by 100%.")
(52, "Amenoma Kageuchi", "Sword", 4, 454, "ATK%", 55.1, "After casting an Elemental Skill, gain 1 Succession Seed. This effect can be triggered once every 5s. The Succession Seed lasts for 30s. Up to 3 Succession Seeds may exist simultaneously. After using an Elemental Burst, all Succession Seeds are consumed and after 2s, the character regenerates 6 Energy for each seed consumed.")
(53, "Wolf's Gravestone", "Claymore", 5, 608, "ATK%", 49.6, "Increases ATK by 20%. On hit, attacks against opponents with less than 30% HP increase all party members' ATK by 40% for 12s. Can only occur once every 30s.")
(54, "Beacon of the Reed Sea", "Claymore", 5, 608, "CRIT Rate", 33.1, "After the character's Elemental Skill hits an opponent, their ATK will be increased by 20% for 8s. After the character takes DMG, their ATK will be increased by 20% for 8s. The 2 aforementioned effects can be triggered even when the character is not on the field. Additionally, when not protected by a shield, the character's Max HP will be increased by 32%.")
(55, "The Unforged", "Claymore", 5, 608, "ATK%", 49.6, "Increases Shield Strength by 20%. Scoring hits on opponents increases ATK by 4% for 8s. Max 5 stacks. Can only occur once every 0.3s. While protected by a shield, this ATK increase effect is increased by 100%.")
(56, "Sacrificial Greatsword", "Claymore", 4, 565, "Energy Recharge", 30.6, "After damaging an opponent with an Elemental Skill, the skill has a 40% chance to end its own CD. Can only occur once every 30s.")
(57, "The First Great Magic", "Bow", 5, 608, "CRIT DMG", 66.2, "DMG dealt by Charged Attacks increased by 16%. For every party member with the same Elemental Type as the wielder (including the wielder themselves), gain 1 Gimmick stack. For every party member with a different Elemental Type from the wielder, gain 1 Theatrics stack. When the wielder has 1/2/3 or more Gimmick stacks, ATK will be increased by 16%/32%/48%. When the wielder has 1/2/3 or more Theatrics stacks, Movement SPD will be increased by 4%/7%/10%.")
(58, "Amos' Bow", "Bow", 5, 608, "ATK%", 49.6, "Increases Normal and Charged Attack DMG by 12%. After a Normal or Charged Attack is fired, DMG dealt increases by a further 8% every 0.1s the arrow is in the air for up to 5 times."),
(59, "Flower-Wreathed Feathers", "Bow", 4, 510, "ATK%", 41.3, "Decreases Gliding Stamina consumption by 15%. When using Aimed Shots, the DMG dealt by Charged Attacks increases by 6% every 0.5s. This effect can stack up to 6 times and will be removed 10s after leaving Aiming Mode."),
(60, "Hamayumi", "Bow", 4, 454, "ATK%", 55.1, "Increases Normal Attack DMG by 16% and Charged Attack DMG by 12%. When the equipping character's Energy reaches 100%, this effect is increased by 100%.")
(61, "Song of Broken Pines","Claymore", 5, 741, "Physical DMG Bonus", 20.7, 'A part of the "Millennial Movement" that wanders amidst the winds. Increases ATK by 16%, and when Normal or Charged Attacks hit opponents, the character gains a Sigil of Whispers. This effect can be triggered once every 0.3s. When you possess 4 Sigils of Whispers, all of them will be consumed and all nearby party members will obtain the "Millennial Movement: Banner-Hymn" effect for 12s. "Millennial Movement: Banner-Hymn" increases Normal ATK SPD by 12% and increases ATK by 20%. Once this effect is triggered, you will not gain Sigils of Whispers for 20s. Of the many effects of the "Millennial Movement," buffs of the same type will not stack.')
(62, "Prototype Archaic", "Claymore", 4, 565, "ATK%", 27.6, "On hit, Normal or Charged Attacks have a 50% chance to deal an additional 240% ATK DMG to opponents within a small AoE. Can only occur once every 15s.")
(63, "Thundering Pulse", "Bow", 5, 608, "CRIT DMG", 66.2, "Increases ATK by 20% and grants the might of the Thunder Emblem. At stack levels 1/2/3, the Thunder Emblem increases Normal Attack DMG by 12/24/40%. The character will obtain 1 stack of Thunder Emblem in each of the following scenarios: Normal Attack deals DMG (stack lasts 5s), casting Elemental Skill (stack lasts 10s); Energy is less than 100% (stack disappears when Energy is full). Each stack's duration is calculated independently."),
(64, "Polar Star", "Bow", 5, 608, "CRIT Rate", 33.1, "Elemental Skill and Elemental Burst DMG increased by 12%. After a Normal Attack, Charged Attack, Elemental Skill or Elemental Burst hits an opponent, 1 stack of Ashen Nightstar will be gained for 12s. When 1/2/3/4 stacks of Ashen Nightstar are present, ATK is increased by 10/20/30/48%. The stack of Ashen Nightstar created by the Normal Attack, Charged Attack, Elemental Skill or Elemental Burst will be counted independently of the others.")
(65, "The Stringless", "Bow", 4, 510, "Elemental Mastery", 165, "Increases Elemental Skill and Elemental Burst DMG by 24%.")
(66, "Vortex Vanquisher", "Polearm", 5, 608, "ATK%", 49.6, "Increases Shield Strength by 20%. Scoring hits on opponents increases ATK by 4% for 8s. Max 5 stacks. Can only occur once every 0.3s. While protected by a shield, this ATK increase effect is increased by 100%.")
(67, "Engulfing Lightning", "Polearm", 5, 608, "Energy Recharge", 55.1, "ATK increased by 28% of Energy Recharge over the base 100%. You can gain a maximum bonus of 80% ATK. Gain 30% Energy Recharge for 12s after using an Elemental Burst.")
(68, "Key of Khaj-Nisut", "Sword", 5, 542, "HP%", 66.2, "HP increased by 20%. When an Elemental Skill hits opponents, you gain the Grand Hymn effect for 20s. This effect increases the equipping character's Elemental Mastery by 0.12% of their Max HP. This effect can trigger once every 0.3s. Max 3 stacks. When this effect gains 3 stacks, or when the third stack's duration is refreshed, the Elemental Mastery of all nearby party members will be increased by 0.2% of the equipping character's max HP for 20s.")
(69, "The Dockhand's Assistant", "Sword", 4, 510, "HP%", 41.3, "When the wielder is healed or heals others, they will gain a Stoic's Symbol that lasts 30s, up to a maximum of 3 Symbols. When using their Elemental Skill or Burst, all Symbols will be consumed and the Roused effect will be granted for 10s. For each Symbol consumed, gain 40 Elemental Mastery, and 2s after the effect occurs, 2 Energy per Symbol consumed will be restored for said character. The Roused effect can be triggered once every 15s, and Symbols can be gained even when the character is not on the field.")
(70, "Mappa Mere", "Catalyst", 4, 565, "Elemental Mastery", 110, "Triggering an Elemental reaction grants a 8% Elemental DMG Bonus for 10s. Max 2 stacks.")
(71, "Thrilling Tales of Dragon Slayers", "Catalyst", 3, 401, "HP%", 35.2, "When switching characters, the new character taking the field has their ATK increased by 24% for 10s. This effect can only occur once every 20s.")
(72, "Starcaller's Watch", "Catalyst", 5, 542, "Elemental Mastery", 265, "Increases Elemental Mastery by 100. Gain the 'Mirror of Night' effect within 15s after the equipping character creates a shield: The current active party member deals 28% increased DMG to nearby opponents. You can gain the 'Mirror of Night' effect once every 14s.")
(73, "Staff of the Scarlet Sands ", "Polearm", 5, 542, "CRIT Rate", 44.1, "The equipping character gains 52% of their Elemental Mastery as bonus ATK. When an Elemental Skill hits opponents, the Dream of the Scarlet Sands effect will be gained for 10s: The equipping character will gain 28% of their Elemental Mastery as bonus ATK. Max 3 stacks.")
(74, "Staff of Homa", "Polearm", 5, 608, "CRIT DMG", 66.2, "HP increased by 20%. Additionally, provides an ATK Bonus based on 0.8% of the wielder's Max HP. When the wielder's HP is less than 50%, this ATK Bonus is increased by an additional 1% of Max HP."),
(75, "Primordial Jade Winged-Spear", "Polearm", 5, 674, "CRIT Rate", 22.1, "On hit, increases ATK by 3.2% for 6s. Max 7 stacks. This effect can only occur once every 0.3s. While in possession of the maximum possible stacks, DMG dealt is increased by 12%.")
(76, "The Catch", "Polearm", 4, 510, "Energy Recharge", 45.9, "Increases Elemental Burst DMG by 16% and Elemental Burst CRIT Rate by 6%.")
(77, "Aquila Favonia", "Sword", 5, 674, "Physical DMG Bonus", 41.3, "ATK is increased by 20%. Triggers on taking DMG: the soul of the Falcon of the West awakens, holding the banner of resistance aloft, regenerating HP equal to 100% of ATK and dealing 200% of ATK as DMG to surrounding opponents. This effect can only occur once every 15s.")
(78, "Absolution", "Sword", 5, 674, "CRIT DMG", 44.1, "CRIT DMG increased by 20%. Increasing the value of a Bond of Life increases the DMG the equipping character deals by 16% for 6s. Max 3 stacks.")
(79, "Freedom-Sworn", "Sword", 5, 608, "Elemental Mastery", 198, 'A part of the "Millennial Movement" that wanders amidst the winds. Increases DMG by 10%. When the character wielding this weapon triggers Elemental Reactions, they gain a Sigil of Rebellion. This effect can be triggered once every 0.5s and can be triggered even if said character is not on the field. When you possess 2 Sigils of Rebellion, all of them will be consumed and all nearby party members will obtain "Millennial Movement: Song of Resistance" for 12s. "Millennial Movement: Song of Resistance" increases Normal, Charged and Plunging Attack DMG by 16% and increases ATK by 20%. Once this effect is triggered, you will not gain Sigils of Rebellion for 20s. Of the many effects of the "Millennial Movement," buffs of the same type will not stack.')
(80, "The Alley Flash", "Sword", 4, 620, "Elemental Mastery", 55, "Increases DMG dealt by the character equipping this weapon by 12%. Taking DMG disables this effect for 5s.")
(81, "Sapwood Blade", "Sword", 4, 565, "Energy Recharge", 30.6, "After triggering Burning, Quicken, Aggravate, Spread, Bloom, Hyperbloom, or Burgeon, a Leaf of Consciousness will be created around the character for a maximum of 10s. When picked up, the Leaf will grant the character 60 Elemental Mastery for 12s. Only 1 Leaf can be generated this way every 20s. This effect can still be triggered if the character is not on the field. The Leaf of Consciousness' effect cannot stack.")
(82, "Favonius Greatsword", "Claymore", 4, 454, "Energy Recharge", 61.3, "CRIT Hits have a 60% chance to generate a small amount of Elemental Particles, which will regenerate 6 Energy for the character. Can only occur once every 12s.")
(83, "Ballad of the Fjords", "Polearm", 4, 510, "CRIT Rate", 27.6, "When there are at least 3 different Elemental Types in your party, Elemental Mastery will be increased by 120.")
(84, "Dragon's Bane", "Polearm", 4, 454, "Elemental Mastery", 221, "Increases DMG against opponents affected by Hydro or Pyro by 20%.")
(85, "Vivid Notions", "Catalyst", 5, 674, "CRIT DMG", 44.1, "ATK is increased by 28%. When you use a Plunging Attack, you will gain the Dawn's First Hue effect: Plunging Attack CRIT DMG is increased by 28%. When you use an Elemental Skill or Burst, you will gain the Twilight's Splendor effect: Plunging Attack CRIT DMG is increased by 40%. The two effects above each last for 15s, and will be canceled 0.1s after the ground impact hits a target.")
(86, "Flowing Purity", "Catalyst", 4, 565, "ATK%", 27.6, "When using an Elemental Skill, All Elemental DMG Bonus will be increased by 8% for 15s, and a Bond of Life worth 24% of Max HP will be granted. This effect can be triggered once every 10s. When the Bond of Life is cleared, every 1,000 HP cleared in the process will provide 2% All Elemental DMG Bonus, up to a maximum of 12%. This effect lasts 15s.")
(87, "Rust", "Bow", 4, 510, "ATK%", 41.3, "Increases Normal Attack DMG by 40% but decreases Charged Attack DMG by 10%.")
(88, "Prototype Crescent", "Bow", 4, 510, "ATK%", 41.3, "Charged Attack hits on weak points increase Movement SPD by 10% and ATK by 36% for 10s.")
(89, "Fruitful Hook", "Claymore", 4, 565, "ATK%", 27.6, "Increase Plunging Attack CRIT Rate by 16%; After a Plunging Attack hits an opponent, Normal, Charged, and Plunging Attack DMG increased by 16% for 10s.")
(90, "Crimson Moon's Semblance", "Polearm", 5, 674, "CRIT Rate", 22.1, "Grants a Bond of Life equal to 25% of Max HP when a Charged Attack hits an opponent. This effect can be triggered up to once every 14s. In addition, when the equipping character has a Bond of Life, they gain a 12% DMG Bonus; if the value of the Bond of Life is greater than or equal to 30% of Max HP, then gain an additional 24% DMG Bonus.")
(91, "Deathmatch", "Polearm", 4, 454, "CRIT Rate", 36.8, "If there are at least 2 opponents nearby, ATK is increased by 16% and DEF is increased by 16%. If there are fewer than 2 opponents nearby, ATK is increased by 24%.")
(92, "White Tassel", "Polearm", 3, 401, "CRIT Rate", 23.4, "Increases Normal Attack DMG by 24%.")
(93, "The Black Sword", "Sword", 4, 510, "CRIT Rate", 27.6, "Increases DMG dealt by Normal and Charged Attacks by 20%. Additionally, regenerates 60% of ATK as HP when Normal and Charged Attacks score a CRIT Hit. This effect can occur once every 5s.")
(94, "Sacrificial Jade", "Catalyst", 4, 454, "CRIT Rate", 36.8, "When not on the field for more than 5s, Max HP will be increased by 32% and Elemental Mastery will be increased by 40. These effects will be canceled after the wielder has been on the field for 10s.")
(95, "Splendor of Tranquil Waters", "Sword", 5, 542, "CRIT DMG", 88.2, "When the equipping character's current HP increases or decreases, Elemental Skill DMG dealt will be increased by 8% for 6s. Max 3 stacks. This effect can be triggered once every 0.2s. When other party members' current HP increases or decreases, the equipping character's Max HP will be increased by 14% for 6s. Max 2 stacks. This effect can be triggered once every 0.2s. The aforementioned effects can be triggered even if the wielder is off-field.")
(96, "Fleuve Cendre Ferryman", "Sword", 4, 510, "Energy Recharge", 45.9, "Increases Elemental Skill CRIT Rate by 8%. Additionally, increases Energy Recharge by 16% for 5s after using an Elemental Skill.")
(97, "Silvershower Heartstrings", "Bow", 5, 542, "HP%", 66.2, "The equipping character can gain the Remedy effect. When they possess 1/2/3 Remedy stacks, Max HP will increase by 12%/24%/40%. 1 stack may be gained when the following conditions are met: 1 stack for 25s when using an Elemental Skill; 1 stack for 25s when the value of a Bond of Life value increases; 1 stack for 20s for performing healing. Stacks can still be triggered when the equipping character is not on the field. Each stack's duration is counted independently. In addition, when 3 stacks are active, Elemental Burst CRIT Rate will be increased by 28%. This effect will be canceled 4s after falling under 3 stacks.")
(98, "Recurve Bow", "Bow", 3, 354, "HP%", 46.9, "Defeating an opponent restores 8% HP."),
(99, "Ring of Yaxche", "Catalyst", 4, 510, "HP%", 41.3, "Using an Elemental Skill grants the Jade-Forged Crown effect: Every 1,000 Max HP will increase the Normal Attack DMG dealt by the equipping character by 0.6% for 10s. Normal Attack DMG can be increased this way by a maximum of 16%.")

;
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

-- Cryo
(22, 49, 1), (22, 25, 2) -- qiqi
(23, 38, 1), (23, 24, 2), (23, 37, 3), (23, 50, 4,) (23, 51, 5), (23, 52, 6), -- Kaeya
(24, 53, 1), (24, 2, 2), (24, 1, 3), (24, 54, 4), (24, 55, 5), (24, 3, 6), (24, 56, 7)  -- Chongyun
(25, 21, 1), (25, 32, 2), (25, 31, 3), -- Diona
(26, 27, 1), (26, 26, 2), (26, 57, 3), (26, 28, 4), (26, 58, 5), (26, 59, 6), (26, 60, 7) -- Ganyu
(27, 16, 1), -- Rosaria
(28, 61, 1), (28, 53, 2), (28, 3, 3), (28, 2, 4), (28, 62, 5), -- Eula
(29, 38, 1), (29, 24, 2), (29, 52, 3), -- Ayaka
(30, 63, 1), (30, 64, 2), (30, 29, 3), (30, 65, 4), -- Aloy
(31, 45, 1), (31, 67, 2), (31, 66, 3), (31, 46, 4), (31, 16, 5), -- Shenhe
(32, 68, 1), (32, 69, 2), (32, 24, 3), (32, 25, 4),  -- Layla
(33, 16, 1, 34, 36, 2), -- Mika
(34, 3, 1), (34, 23, 2), (34, 61, 3), (34, 53, 4), (34, 2, 5), (34, 62, 6) -- Freminet
(35, 9, 1), (35, 12, 2), (35, 33, 3), (35, 6, 4), (35, 8, 5), (35, 10, 6), (35, 70, 7)  -- Wriothesley
(36, 42, 1), (36, 43, 2), (36, 71, 3), -- Charlotte
(37, 72, 1), (37, 71, 2), (37, 43, 3), -- Citlali

-- Pyro 
(38, 73, 1), (38, 44, 2), (38, 74, 3), (38, 67, 4), (38, 75, 5), (38, 76, 6) -- Xiangling
(39, 6, 1), (39, 7, 2), (39, 8, 3), (39, 9, 4), (39, 5, 5), (39, 12, 6), (39, 13, 7) --Klee
(40, 54, 1), (40, 1, 2), (40, 23, 3), (40, 22, 4), (40, 2, 5), (40, 47, 6), (40, 53, 7), (40, 55, 8), (40, 62, 9) -- Diluc
(41, 38, 1), (41, 77, 2), (41, 78, 3), (41, 79, 4), (41, 17, 5), (41, 80, 6), (41, 81, 7) --Bennett
(42, 31, 1), (42, 21, 1), -- Amber 
(43, 56, 1), (43, 4, 2), (43, 82, 3), -- Xinyan
(44, 74, 1), (44, 73, 2), (44, 83, 3), (44, 44, 4), (44, 84, 4) -- Hutao
(45, 7, 1), (45, 85, 2), (45, 33, 3), (45, 8, 4), (45, 9, 5), (45, 86, 6) -- Yanfei
(46, 63, 1), (46, 28, 2), (46, 57, 3), (46, 64, 4), (46, 26, 5), (46, 29, 6), (46, 87, 7), (46, 30, 8) -- Yoimiya
(47, 16, 1), (47, 15, 2),  -- Thoma
(48, 54, 1), (48, 53, 2), (48, 2, 3), (48, 40, 4) -- Dehya
(49, 57, 1), (49, 27, 2), (49, 28, 3), (49, 88, 4) -- lyney
(50, 16, 1), (50, 36, 2), (50, 15, 3), -- Chevreuse
(51, 23, 1), (51, 2, 2), (51, 1, 3), (51, 22, 4), (51, 54, 5), (51, 89, 6), -- Gaming
(52, 90, 1), (52, 75, 2), (52, 73, 3), (52, 74, 4), (52, 91, 5), (52, 92, 6), -- Arlecchino
(53, 23, 1), (52, 53, 2), (52, 2, 3), (52, 48, 4) -- Mavuika

-- Hydro
(54, 49, 1), (54, 25, 2), -- Xingqiu
(55, 13, 1), (55, 71, 2), (55, 43, 3), -- Mona
(56, 71, 2), (56, 42, 2), (56, 43, 3), -- Barbara
(57, 64, 1), (57, 29, 2), (57, 87, 3), (57, 63, 4) -- Tartaglia
(58, 71, 1), (58, 42, 2), -- Kokomi
(59, 50, 1), (59, 93, 2), (59, 52, 3), -- Ayato
(60, 28, 1), (60, 31, 2), (60, 21, 3), -- Yelan
(61, 16, 1), -- Candace
(62, 68, 1), (62, 39, 2), -- Nilou
(63, 8, 1), (63, 94, 2), (63, 42, 3) -- Neuvillette
(64, 95, 1), (64, 96, 2), (64, 25, 3) -- Furina
(65, 97, 1), (65, 98, 2), (65, 32, 3), (65, 21, 4), -- Sigewinne
(66, 7, 1), (66, 94, 2), (66, 13, 3), (66, 99, 4), -- Mualani