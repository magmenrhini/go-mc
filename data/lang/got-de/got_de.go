package got_de

import "github.com/Tnze/go-mc/chat"

func init() { chat.SetLanguage(Map) }

var Map = map[string]string{"addServer.add": "Manwu", "addServer.enterIp": "Skalkinassaus Stadinamo", "addServer.enterName": "Skalkinassaus Namo", "addServer.hideAddress": "Affilhai Stadinamo", "addServer.title": "Inmaidei Skalkinassaus Kunþi", "advancement.advancementNotFound": "Unkunþa framgahts: %s", "advancements.adventure.hero_of_the_village.title": "Haliþs Weihsis", "advancements.adventure.ol_betsy.description": "Skiutai krukabuga", "advancements.adventure.two_birds_one_arrow.title": "Twai Fuglos, Aina Arhwazna", "advancements.adventure.voluntary_exile.title": "Silbawili Distaheins", "advancements.husbandry.complete_catalogue.description": "Tamei alla kunja katte!", "advancements.husbandry.complete_catalogue.title": "Alls Katalaugus", "argument.color.invalid": "Unkunþa faiheins '%s'", "argument.entity.notfound.entity": "Ni waihts bigitana", "argument.entity.notfound.player": "Ni laikands bigitans", "argument.entity.options.level.negative": "Hauhei ni sijai wiþra", "argument.entity.options.limit.toosmall": "Marka skal mais þau 1 wisan", "argument.entity.options.unterminated": "Usbeidana andeis mahte", "argument.entity.selector.not_allowed": "Walja andletans nist", "argument.id.invalid": "Unraiht ID", "argument.id.unknown": "Unkunþ ID: %s", "argument.item.tag.disallowed": "Sokeiniwaurda ni andletanda, þatainei waihts", "argument.nbt.array.mixed": "Ni mag analagjan %s in %s", "argument.nbt.expected.key": "Usbeidans lukils", "argument.nbt.expected.value": "Usbeidan wairþ", "argument.nbt.list.mixed": "Ni mag analagjan %s in wikon %s", "argument.player.unknown": "Sa laikands nist", "argument.pos.missing.double": "Usbaþ gatewos", "argument.pos.missing.int": "Usbaþ kubulageinais", "argument.pos.outofworld": "So lageins nist þis fairhwaus!", "argument.pos.unloaded": "So lageins atbaurana nist", "argument.pos3d.incomplete": "Unfull (usbaþ 3 gatewo)", "argument.rotation.incomplete": "Unfull (usbaþ 2 gatewo)", "argument.time.invalid_unit": "Unraiht wairþ", "arguments.block.tag.unknown": "Unkunþ kubusokeiniwaurd '%s'", "arguments.function.unknown": "Unkunþa mahts %s", "arguments.nbtpath.nothing_found": "Ni mahta bigitan stabins galeikondans %s", "arguments.operation.div0": "Ni mag nullu dailjan", "block.minecraft.air": "Luftus", "block.minecraft.bamboo": "Bambu", "block.minecraft.banner": "Fana", "block.minecraft.barrel": "Barils", "block.minecraft.bed": "Badi", "block.minecraft.bell": "Skilla", "block.minecraft.birch_sign": "Bairkos Taikns", "block.minecraft.birch_wall_sign": "Bairkos Waddjaus Taikns", "block.minecraft.blast_furnace": "Blastauhns", "block.minecraft.cake": "Koka", "block.minecraft.campfire": "Bibaurgeins fon", "block.minecraft.cartography_table": "Kartjabiuþs", "block.minecraft.chest": "Arka", "block.minecraft.clay": "Þāho", "block.minecraft.cornflower": "Kaurnabloma", "block.minecraft.dark_oak_sign": "Riqizeinaizos Aikos Taikns", "block.minecraft.dark_oak_wall_sign": "Swarta Aikis Waddjaus Taikns", "block.minecraft.dirt": "Aírþa", "block.minecraft.fire": "Fon", "block.minecraft.furnace": "Aúhns", "block.minecraft.glass": "Glas", "block.minecraft.grass": "Gras", "block.minecraft.gravel": "Griut", "block.minecraft.ice": "Eis", "block.minecraft.jungle_sign": "Widaus Taikns", "block.minecraft.lantern": "Lukarn", "block.minecraft.lily_of_the_valley": "Lilei Dals", "block.minecraft.oak_sign": "Aikos Taikns", "block.minecraft.oak_wall_sign": "Aikos Waddjustaikns", "block.minecraft.potted_cornflower": "Kaurnabloma Aurkja", "block.minecraft.red_sandstone_wall": "Rauda Malmastainis Waddjus", "block.minecraft.redstone_wire": "Redstone striks", "block.minecraft.sand": "Malma", "block.minecraft.sandstone_wall": "Malmastainis Waddjus", "block.minecraft.smithing_table": "Smiþjabiuþs", "block.minecraft.smoker": "Raukja", "block.minecraft.snow": "Snáiws", "block.minecraft.sponge": "Swamms", "block.minecraft.stone": "Stains", "block.minecraft.stone_stairs": "Stainaihos Usstaigos", "block.minecraft.stonecutter": "Stainamaitja", "block.minecraft.sweet_berry_bush": "Sutis basje aihwatundi", "block.minecraft.torch": "Hais", "block.minecraft.water": "Wato", "block.minecraft.wet_sponge": "Nats swamms", "block.minecraft.wither_rose": "Dauþjands rauþabloma", "book.signButton": "Taikn", "chat.cannotSend": "Ni mahta insandjan waurd", "chat.editBox": "gawaurdi", "chat.type.announcement": "[%s] %s", "chat.type.emote": "* %s %s", "chat.type.team.hover": "Sandei waurd du Kunja", "chat.type.team.sent": "-> %s <%s> %s", "chat.type.team.text": "%s <%s> %s", "chat.type.text": "<%s> %s", "chat.type.text.narrate": "%s qiþiþ %s", "commands.ban.failed": "Ni waihts inmaidida. Sa laikands faurdammiþs", "commands.banip.failed": "Ni waihts inmaidida. Þata IP ju ist faurdammiþ", "commands.banip.invalid": "Unraihts IP staþs þau unkunþs laikands", "commands.banip.success": "Faurdammiþ IP %s: %s", "commands.datapack.list.available.none": "Ni sind managizona datahuzda", "commands.datapack.unknown": "Unkunþ datahuzd '%s'", "commands.drop.no_held_items": "Waihts ni mag haban waihtins", "commands.effect.clear.everything.failed": "Mundrei ni habaiþ toja du gatairan", "commands.execute.conditional.fail": "Gakust ni ustauh", "commands.execute.conditional.pass": "Gakust ustauh", "commands.execute.conditional.pass_count": "Gakust ustauh, raþjo: %s", "commands.give.success.multiple": "Gaf %s %s du %s laikandam", "commands.give.success.single": "Gaf %s %s du %s", "commands.locate.failed": "Ni mahta þana strukturu bigitan nehwa", "commands.op.failed": "Ni waihts inmaidida. Sa laikands ju waurstwa ist", "commands.pardon.failed": "Ni waihts inmaidida. Faurdammiþs nist", "commands.pardonip.failed": "Ni waihts inmaidida. Þata IP nist faurdammiþ", "commands.pardonip.invalid": "Unraiht IP stadinamo", "commands.particle.failed": "Þata daililo ni was anasiun du hwamma", "commands.playsound.failed": "Drunjus ufarfairra ist du hausjan", "commands.recipe.give.failed": "Ni niujos laiseinos galaisida", "commands.recipe.take.failed": "Ni laiseinos ufarmaudidos sind", "commands.replaceitem.block.failed": "Mundrikubus nist gahabands", "commands.replaceitem.entity.failed": "Ni mahta %s atbairan in staþ %s", "commands.save.alreadyOff": "Gafastan ju ni taujiþ", "commands.save.alreadyOn": "Gafastan ju taujiþ", "commands.save.failed": "Ni mahta laik gafastan (istu ganoh rum ana mesa?)", "commands.save.saving": "Gafastada laiks (magi lagg mel wairþan!)", "commands.save.success": "Gafastaida laik", "commands.scoreboard.objectives.add.duplicate": "Mundrei ju ist miþ þamma namin", "commands.scoreboard.objectives.display.alreadyEmpty": "Ni waihts inmaidida. So mundrisiuns juþan laus ist", "commands.scoreboard.objectives.list.empty": "Ni sind mundreins", "commands.setblock.failed": "Ni mahta lagjan kubu", "commands.summon.failed": "Ni mahta haitan waiht", "commands.tag.add.failed": "Mundrei jaþþe habaiþ ju sokeiniwaurd jaþþe habaiþ ufarfilu sokeiniwaurde", "commands.tag.remove.failed": "Mundrei þo sokeiniwaurd ni habaiþ", "commands.team.add.duplicate": "Ju ist kuni miþ þamma namin", "commands.team.add.success": "Gaskapan kuni %s", "commands.team.empty.unchanged": "Ni waihts inmaidida. Þata kuni juþþan ist full", "commands.team.join.success.single": "Anaaiauk %s du kunja %s", "commands.team.list.members.success": "Kuni %s habaiþ gadailans: %s", "commands.team.list.teams.empty": "Ni sind kunja", "commands.team.list.teams.success": "Sind %s kunja: %s", "commands.team.option.color.unchanged": "Ni waihts inmaidida. Þata kuni ju habaiþ faihein", "commands.team.option.deathMessageVisibility.unchanged": "Ni waihts inmaidida. Dauþaus waurdis anasiuns ju ist þata wairþ", "commands.team.option.name.unchanged": "Ni waihts inmaidida. Þata kuni ju habaiþ þata namo", "commands.team.option.nametagVisibility.unchanged": "Ni waihts inmaidida. Namasokeiniwaurdis anasiuns ju ist þata wairþ", "commands.team.remove.success": "Usnuman kuni %s", "commands.teammsg.failed.noteam": "Skuld ist þus wisan ana kunja du insandjan waurd kuni þeinata", "commands.weather.set.clear": "Lagei wiþr du bairhtamma", "commands.weather.set.rain": "Lagei wiþr du rigna", "commands.weather.set.thunder": "Lagei wiþr du rigna jah þunara", "commands.whitelist.remove.failed": "Laikands nist hweitawikon", "commands.worldborder.center.failed": "Ni waihts inmaidida. Fairhwaus marka ju ist jainar", "commands.worldborder.set.failed.nochange": "Ni waihts inmaidida. Fairhwaus marka ju ist þizos mikileins", "connect.authorizing": "Atgaggando...", "connect.connecting": "Gawidando skalkinassau...", "connect.failed": "Ni mahta gawidan skalkinassau", "connect.joining": "Gamaineiþ fairhwu...", "container.barrel": "Barils", "container.blast_furnace": "Blastauhns", "container.cartography_table": "Kartjabiuþs", "container.loom": "Waibjastols", "container.smoker": "Raukja", "container.stonecutter": "Stainamaitja", "controls.reset": "Aftrasatei", "controls.resetAll": "Aftrasatei Haubidilons", "createWorld.customize.buffet.generator": "Walei timrjahaidu", "createWorld.customize.buffet.generatortype": "Fairhwutimrjands:", "createWorld.customize.custom.baseSize": "Diupeins Gasuleinais Mikilei", "createWorld.customize.custom.confirm1": "Ufarmeleiþ andwairþa þeina", "createWorld.customize.custom.confirm2": "lageinos jah ni magt aftra usniman.", "createWorld.customize.custom.confirmTitle": "Hwota!", "createWorld.customize.custom.defaults": "Bi biuhtja", "createWorld.customize.custom.depthNoiseScaleExponent": "Diupeins Drunjaus Uslageins", "createWorld.customize.custom.depthNoiseScaleX": "Diupei Drunjumikileins X", "createWorld.customize.custom.depthNoiseScaleZ": "Diupei Drunjumikileins Z", "createWorld.customize.custom.heightScale": "Hauheins Mikilei", "createWorld.customize.custom.mainNoiseScaleX": "Gamaina Drunjumikilei X", "createWorld.customize.custom.mainNoiseScaleY": "Gamaina Drunjumikilei Y", "createWorld.customize.custom.mainNoiseScaleZ": "Gamaina Drunjumikilei Z", "createWorld.customize.custom.next": "Iftums Laufs", "createWorld.customize.custom.preset.caveChaos": "Hulandja Untewos", "createWorld.customize.custom.preset.goodLuck": "Filu Awi", "createWorld.customize.custom.preset.isleLand": "Hulmaland", "createWorld.customize.custom.preset.waterWorld": "Fairƕus Watins", "createWorld.customize.custom.presets": "Fauralageinos", "createWorld.customize.custom.prev": "Aftums Laufs", "createWorld.customize.custom.riverSize": "Mikilei aƕos", "createWorld.customize.custom.seaLevel": "Mareins Hauheins", "createWorld.customize.custom.useCaves": "Hulundjos", "createWorld.customize.custom.useMonuments": "Okaianos Gaminþja", "createWorld.customize.custom.useOceanRuins": "Okaianaus Gabrukos", "createWorld.customize.custom.useTemples": "Alhs", "createWorld.customize.custom.useVillages": "Weihsos", "createWorld.customize.custom.useWaterLakes": "Watalagjus", "createWorld.customize.flat.height": "Hauhiþa", "createWorld.customize.flat.layer": "%s", "createWorld.customize.flat.layer.top": "Þata Hauhisto - %s", "createWorld.customize.flat.removeLayer": "Gatairai Gaþrask", "createWorld.customize.preset.desert": "Auþida", "createWorld.customize.preset.overworld": "Ufarfairhwus", "createWorld.customize.preset.snowy_kingdom": "Þiudinassus Snaiweins", "createWorld.customize.preset.the_void": "Þairko", "createWorld.customize.preset.water_world": "Watafairhwus", "createWorld.customize.presets": "Fauralageinos", "createWorld.customize.presets.select": "Brukei Fauralageinais", "createWorld.customize.presets.title": "Walei Fauralagein", "demo.day.2": "Dags Anþar", "demo.day.3": "Dags Þridja", "demo.day.4": "Dags Fidurþa", "demo.day.5": "Þata ist aftumists dags þeins!", "demo.day.warning": "Hweila þeina ist haldis faurgaggana!", "demo.help.buy": "Bugei Nu!", "demo.help.later": "Þairhwis Laikan!", "demo.help.movementMouse": "Atsaihw miþ musa", "demo.remainingTime": "Hweila aflifnandei: %s", "disconnect.closed": "Gawiss lukana", "disconnect.endOfStream": "Andeis flodaus", "disconnect.genericReason": "%s", "disconnect.loginFailed": "Ni mahta atgaggan", "disconnect.loginFailedInfo": "Ni mahta atgaggan: %s", "disconnect.lost": "Gawiss Fralaus", "effect.minecraft.bad_omen": "Ubila Taikns", "effect.minecraft.hero_of_the_village": "Haliþs Weihsis", "entity.minecraft.fox": "Fauho", "entity.minecraft.panda": "Pandabaira", "entity.minecraft.trader_llama": "Lewjins Lama", "entity.minecraft.villager.mason": "Timrja", "entity.minecraft.villager.weaponsmith": "Wepnasmiþja", "entity.minecraft.wandering_trader": "Farands Lewja", "event.minecraft.raid": "Ufardrus", "event.minecraft.raid.defeat": "Biutands", "event.minecraft.raid.raiders_remaining": "Ufardrusjans bileiband: %s", "event.minecraft.raid.victory": "Sigis", "filled_map.locked": "Lukana", "generator.default": "Bi biuhtja", "gui.all": "All", "gui.back": "Gawandei", "gui.cancel": "Afqiþan", "gui.done": "Taujan ist", "gui.down": "Dalaþ", "gui.narrate.button": "%s haubidilo", "gui.narrate.editBox": "%s arka inmaideinais: %s", "gui.narrate.slider": "%s sliupja", "gui.no": "Ne", "gui.none": "Ni ainhun", "gui.ok": "Waila", "gui.proceed": "Framgagg", "gui.recipebook.moreRecipes": "Attekai af Taihswon faur mais", "gui.recipebook.toggleRecipes.all": "Bandwjando all", "gui.recipebook.toggleRecipes.blastable": "Bandwei blastamahteiga", "gui.recipebook.toggleRecipes.craftable": "S𐌷𐍉𐍅 c𐍂𐌰𐍆𐍄", "gui.recipebook.toggleRecipes.smokable": "Bandwei raukamahteiga", "gui.toMenu": "Ibuks du skalkinassaus wikon", "gui.toTitle": "Gawandei du haubidaskairma", "gui.up": "Iup", "gui.yes": "Ja", "item.minecraft.arrow": "Arƕazna", "item.minecraft.black_dye": "Faheins Swarta", "item.minecraft.blue_dye": "Faheins Blewa", "item.minecraft.bow": "Buga", "item.minecraft.brown_dye": "Faheins Bruna", "item.minecraft.cat_spawn_egg": "Kattis Addi Gaskaftais", "item.minecraft.coal": "Hauri", "item.minecraft.crossbow": "Krukabuga", "item.minecraft.flower_banner_pattern.desc": "Blomins Siuns", "item.minecraft.fox_spawn_egg": "Fauhons Addi Gaskaftais", "item.minecraft.green_dye": "Faheins Groni", "item.minecraft.mojang_banner_pattern.desc": "Waihts", "item.minecraft.panda_spawn_egg": "Pandos Addi Gaskaftais", "item.minecraft.pillager_spawn_egg": "Wilwins Addi Gaskaftais", "item.minecraft.ravager_spawn_egg": "Qistjins Addi Gaskaftais", "item.minecraft.red_dye": "Faheins Rauda", "item.minecraft.skull_banner_pattern.desc": "Hwairneins Siuns", "item.minecraft.sweet_berries": "Sutizona Basja", "item.minecraft.trader_llama_spawn_egg": "Lewjins Lamos Addi Gaskaftais", "item.minecraft.wandering_trader_spawn_egg": "Farandins Lewjins Addi Gaskaftais", "item.minecraft.white_dye": "Faheins Ƕeita", "item.minecraft.yellow_dye": "Faheins Gilwa", "jigsaw_block.final_state": "Wairþiþ:", "key.chat": "Uslukai Gawaurdi", "key.jump": "Laikai", "lanServer.otherPlayers": "Lageinos Anþaraim Laikandam", "lanServer.scanning": "Sokjanda laikeis ana ganatja þeinamma stada", "lanServer.start": "Anastodei LAN Fairhwu", "lanServer.title": "LAN Fairhwus", "language.code": "got", "language.name": "Gutrazda", "language.region": "Dakia jah Heispanja", "lectern.take_book": "Nim Bokos", "mcoServer.title": "Meinakraftais Anaganatjis Fairhwus", "menu.convertingLevel": "Gaskeirjando fairƕu", "menu.generatingLevel": "Skapjando fairƕu", "menu.generatingTerrain": "Timrjando land", "menu.loadingLevel": "Atbairiþ fairhwu", "menu.multiplayer": "Managlaikands", "menu.online": "Meinakraftais Midjungardeis", "menu.options": "Mahteis...", "menu.paused": "Laiks swibana", "menu.preparingSpawn": "Fauramanwjands anastodeinistaþ: %s%%", "menu.quit": "Afleiþ Laik", "menu.reportBugs": "Merei Airzeins", "menu.returnToGame": "Ibuks du Laikai", "menu.returnToMenu": "Gafastai jah Afleiþ du Ufarmelja", "menu.savingChunks": "Gafastjando dailins", "menu.savingLevel": "Gafastando fairƕu", "menu.sendFeedback": "Gif Mun", "menu.shareToLan": "Uslukan du LAN", "menu.singleplayer": "Ainalaikands", "menu.working": "Waurkjando...", "merchant.current_level": "Lewjins andwairþs griþs", "merchant.deprecated": "Haimjans gafulljan andstalda ni mais twaim sinþam ana dag.", "merchant.level.1": "Duginnands", "merchant.level.2": "Siponeis", "merchant.level.3": "Gasinþja", "merchant.level.4": "Froþs", "merchant.level.5": "Frauja", "merchant.next_level": "Lewjins aftuma griþs", "merchant.trades": "Kaupa", "multiplayer.connect": "Gawidai", "multiplayer.disconnect.duplicate_login": "Atiddjes fram anþaramma stada", "multiplayer.disconnect.flying": "Fliugan mahts nist ana þamma skalkinassau", "multiplayer.disconnect.generic": "Ungagahaftiþ", "multiplayer.disconnect.idling": "Ni gawagides ufarlaggata mel!", "multiplayer.disconnect.illegal_characters": "Unraihtos bokos in gawaurdja", "multiplayer.disconnect.name_taken": "Þata namo ju numan ist", "multiplayer.disconnect.not_whitelisted": "Ni is hweitawikon þis skalkinassaus!", "multiplayer.disconnect.server_full": "Sa skalkinassus ist fulls!", "multiplayer.disconnect.server_shutdown": "Skalkinassus lukans ist", "multiplayer.disconnect.slow_login": "Ufarlagg mel þaurfta du atgaggan", "multiplayer.disconnect.unexpected_query_response": "Unbeidana biuhta data fram gasta", "multiplayer.disconnect.unverified_username": "Ni mahta gasigljan atgagganamo!", "multiplayer.downloadingTerrain": "Atbairando land...", "multiplayer.ipinfo": "Melei IP skalkinassaus du gawidan:", "multiplayer.player.joined.renamed": "%s (airis kunþs swa %s) gamainida in laikai", "multiplayer.player.left": "%s aflaiþ laik", "multiplayer.status.and_more": "... jah %s mais ...", "multiplayer.status.cannot_connect": "Ni mag gawidan skalkinassau", "multiplayer.status.cannot_resolve": "Ni mag lausjan gastinamo", "multiplayer.status.client_out_of_date": "Gasts ufaralþeis!", "multiplayer.status.finished": "Manwu", "multiplayer.status.no_connection": "(ni gawiss)", "multiplayer.status.old": "Fairni", "multiplayer.status.unknown": "???", "multiplayer.stopSleeping": "Afleiþ Ligr", "multiplayer.title": "Laikai Managlaikand", "narrator.button.accessibility": "Atgahtei", "narrator.button.difficulty_lock": "Unazeteins luk", "narrator.button.difficulty_lock.locked": "Lukan ist", "narrator.button.difficulty_lock.unlocked": "Uslukai", "narrator.button.language": "Razda", "narrator.controls.bound": "%s gagahaftiþ ist þairh %s", "narrator.controls.unbound": "%s gagahaftiþ nist", "narrator.joining": "Gamainjands", "narrator.loading": "Atbairando: %s", "narrator.loading.done": "Taujan ist", "narrator.screen.title": "Haubidaskairms", "narrator.select": "Gawaliþ: %s", "narrator.select.world": "Gawalida %s, laikans niujaba: %s, %s, %s, usmereins: %s", "narrator.toast.disabled": "Gataihands Nist", "narrator.toast.enabled": "Gataihands Ist", "optimizeWorld.confirm.title": "Gabotei fairƕu", "optimizeWorld.stage.failed": "Ni was mahts! :(", "optimizeWorld.stage.finished": "Ustiuhands...", "options.accessibility.text_background": "Boko Hindargrundus", "options.accessibility.text_background.chat": "Gawaurdi", "options.accessibility.text_background.everywhere": "Hwaruh", "options.accessibility.text_background_opacity": "Boko Hindargrundaus Gaskadweins", "options.accessibility.title": "Atgahteins Lageinos...", "options.ao.off": "AF", "options.attack.crosshair": "Krukahrusk", "options.chat.color": "Farwos", "options.chat.links": "Ganatjis Gawisseis", "options.chat.opacity": "Gaskadweins Meleinais Gawaurdjis", "options.chat.title": "Lageinos gawaurdje...", "options.chat.visibility": "Gawaurdi", "options.chat.visibility.full": "Bandwiþ", "options.chat.visibility.hidden": "Affilhan", "options.chat.visibility.system": "Þatainei Gagrefteis", "options.clouds.fast": "Hrusk", "options.difficulty.easy": "Azet", "options.difficulty.hard": "Aglu", "options.difficulty.hardcore": "Sleidi", "options.difficulty.normal": "Gamain", "options.difficulty.peaceful": "Gawairþeig", "options.fov.min": "Gamain", "options.fullscreen.current": "Andwairþ", "options.gamma": "Bairhtei", "options.gamma.max": "Bairht", "options.graphics.fast": "Hrusk", "options.guiScale.auto": "Silba", "options.hidden": "Affilhan", "options.language": "Razda...", "options.languageWarning": "Mahts ist ei razdo gaskeireinos ni sind 100%% godos", "options.mainHand.left": "Af Hleidumein", "options.mainHand.right": "Af Taihswon", "options.mipmapLevels": "Mipmap Hauhiþos", "options.modelPart.cape": "Hakuls", "options.modelPart.hat": "Hatta", "options.mouse_settings": "Musis Lageinos...", "options.mouse_settings.title": "Musis Lageinos", "options.music": "Saggws", "options.narrator.all": "Gataihiþ all", "options.narrator.chat": "Gataihiþ gawaurdi", "options.narrator.notavailable": "Nist", "options.narrator.off": "AF", "options.narrator.system": "Gataihiþ swstaima", "options.off": "AF", "options.on": "ANA", "options.particles": "Daililons", "options.particles.all": "All", "options.renderClouds": "Milhmans", "options.sensitivity.max": "ABRABASPRAUTO!!!", "options.snooper.title": "Gif uns data!", "options.sound": "Drunjus", "options.sounds": "Saggws & Drunjus...", "options.sounds.title": "Saggws & Drunjumahteis", "options.title": "Mahteis", "options.video": "Filmalageinos...", "options.videoTitle": "Filmalageinos", "options.visible": "Bandwiþs", "parsing.long.expected": "Usbeidan lagg", "parsing.long.invalid": "Unraiht lagg '%s'", "recipe.notFound": "Unkunþa laiseins: %s", "selectServer.add": "Anaaiauk Skalkinassu", "selectServer.defaultName": "Meinakraftais Skalkinassus", "selectServer.delete": "Gatairai", "selectServer.deleteButton": "Gatairai", "selectServer.deleteQuestion": "Arniba wileis þana skalkinassu gatairan?", "selectServer.deleteWarning": "'%s' fraliusada in aiwin! (Lagg mel!)", "selectServer.edit": "Inmaidei", "selectServer.empty": "laus", "selectServer.hiddenAddress": "(Affilhan)", "selectServer.refresh": "Ananiwei", "selectServer.select": "Gamainei Skalkinassu", "selectServer.title": "Walei Skalkinassu", "selectWorld.backupEraseCache": "Auþjan ganasiþ kunþi", "selectWorld.backupJoinConfirmButton": "B𐌰c𐌺𐌿𐍀 𐌰𐌽𐌳 𐌻𐍉𐌰𐌳", "selectWorld.backupJoinSkipButton": "Wait ƕa tauja!", "selectWorld.backupQuestion.customized": "C𐌿𐍃𐍄𐍉𐌼𐌹𐌶𐌴𐌳 𐍅𐍉𐍂𐌻𐌳𐍃 𐌰𐍂𐌴 𐌽𐍉 𐌻𐍉𐌽𐌲𐌴𐍂 𐍃𐌿𐍀𐍀𐍉𐍂𐍄𐌴𐌳", "selectWorld.backupWarning.customized": "U𐌽𐍆𐍉𐍂𐍄𐌿𐌽𐌰𐍄𐌴𐌻𐌸, 𐍅𐌴 𐌳𐍉 𐌽𐍉𐍄 𐍃𐌿𐍀𐍀𐍉𐍂𐍄 c𐌿𐍃𐍄𐍉𐌼𐌹𐌶𐌴𐌳 𐍅𐍉𐍂𐌻𐌳𐍃 𐌹𐌽 𐍄𐌷𐌹𐍃 𐍈𐌴𐍂𐍃𐌹𐍉𐌽 𐍉𐍆 m𐌹𐌽𐌴c𐍂𐌰𐍆𐍄.  W𐌴 c𐌰𐌽 𐍃𐍄𐌹𐌻𐌻 𐌻𐍉𐌰𐌳 𐍄𐌷𐌹𐍃 𐍅𐍉𐍂𐌻𐌳 𐌰𐌽𐌳 𐌺𐌴𐌴𐍀 𐌴𐍈𐌴𐍂𐌸𐍄𐌷𐌹𐌽𐌲 𐍄𐌷𐌴 𐍅𐌰𐌸 𐌹𐍄 𐍅𐌰𐍃, 𐌱𐌿𐍄 𐌰𐌽𐌸 𐌽𐌴𐍅𐌻𐌸 𐌲𐌴𐌽𐌴𐍂𐌰𐍄𐌴𐌳 𐍄𐌴𐍂𐍂𐌰𐌹𐌽 𐍅𐌹𐌻𐌻 𐌽𐍉 𐌻𐍉𐌽𐌲𐌴𐍂 𐌱𐌴 c𐌿𐍃𐍄𐍉𐌼𐌹𐌶𐌴𐌳.  W𐌴'𐍂𐌴 𐍃𐍉𐍂𐍂𐌸 𐍆𐍉𐍂 𐍄𐌷𐌴 𐌹𐌽c𐍉𐌽𐍈𐌴𐌽𐌹𐌴𐌽c𐌴!", "selectWorld.bonusItems": "Hauh Huzd:", "selectWorld.cheats": "C𐌷𐌴𐌰𐍄𐍃", "selectWorld.conversion": "Skuld ist gaskeirjan!", "selectWorld.create": "Skapei Fairhwu Niujana", "selectWorld.createDemo": "Laikan niujana unfullana fairƕu", "selectWorld.customizeType": "C𐌿𐍃𐍄𐍉𐌼𐌹𐌶𐌴", "selectWorld.delete": "Gatairai", "selectWorld.deleteButton": "Gatairai", "selectWorld.deleteQuestion": "Arniba wileis þana fairƕu gatairan?", "selectWorld.edit": "Inmaidei", "selectWorld.edit.backup": "Gatauei gafastain", "selectWorld.edit.backupSize": "mikilei: %s MB", "selectWorld.edit.openFolder": "Andaugei waurt fairƕaus", "selectWorld.edit.save": "Gafastai", "selectWorld.edit.title": "Inmaidei fairƕu", "selectWorld.empty": "laus", "selectWorld.enterName": "Namo Fairƕaus", "selectWorld.enterSeed": "Fraiw du skapjand fairƕaus", "selectWorld.futureworld.error.title": "Airzei warþ!", "selectWorld.gameMode.adventure": "Wratodus", "selectWorld.gameMode.creative": "Skapjando", "selectWorld.gameMode.spectator": "Saiƕands", "selectWorld.gameMode.spectator.line1": "Magt saiƕan, ni attekan", "selectWorld.hardcoreMode": "Sleidi:", "selectWorld.hardcoreMode.info": "Fairhwus gatairada þan gaswiltais", "selectWorld.mapFeatures": "Gatimrei Strukturos:", "selectWorld.mapType": "Hiwi Fairhwaus:", "selectWorld.mapType.normal": "Gamain", "selectWorld.moreWorldOptions": "Managizeins Fairhwumahteis...", "selectWorld.newWorld": "Niujis Fairƕus", "selectWorld.recreate": "Aftra-Skapei", "selectWorld.recreate.customized.text": "Inmaididai fairhwjus ni juþan stodidai sind in þizai usmereinai Meinakraftais. Mageima sokjan du aftragaskapjan ita miþ fraiwa samamma jah wistim, akei landa-biuhtja fraliusanda. Habais uns faurqiþana unazetein!", "selectWorld.recreate.customized.title": "Inmaididai Fairhwjus ni juþan stodidai sind", "selectWorld.recreate.error.title": "Airzei warþ!", "selectWorld.resultFolder": "Gafastada in:", "selectWorld.search": "sokei fairhwuns", "selectWorld.seedInfo": "Gafastai laus du niman unkunþata fraiw", "selectWorld.title": "Gawalei fairƕu", "selectWorld.tooltip.fromNewerVersion1": "Sa fairƕus gafastans ist in niujizein usmereinai,", "selectWorld.tooltip.fromNewerVersion2": "atbairan þana fairƕu magi aglons giban!", "selectWorld.tooltip.snapshot1": "D𐍉𐌽'𐍄 𐍆𐍉𐍂𐌲𐌴𐍄 𐍄𐍉 𐌱𐌰c𐌺𐌿𐍀 𐍄𐌷𐌹𐍃 𐍅𐍉𐍂𐌻𐌳", "selectWorld.tooltip.snapshot2": "𐌱𐌴𐍆𐍉𐍂𐌴 𐌸𐍉𐌿 𐌻𐍉𐌰𐌳 𐌹𐍄 𐌹𐌽 𐍄𐌷𐌹𐍃 𐍃𐌽𐌰𐍀𐍃𐌷𐍉𐍄.", "selectWorld.tooltip.unsupported": "Sa fairƕus nist juþan stodiþs jah magt þana þatainei laikan in %s", "selectWorld.unable_to_load": "Ni mahta atbairan fairhwuns", "selectWorld.version": "Usmereins:", "selectWorld.versionJoinButton": "L𐍉𐌰𐌳 a𐌽𐌸𐍅𐌰𐌸", "selectWorld.versionQuestion": "D𐍉 𐌸𐍉𐌿 𐍂𐌴𐌰𐌻𐌻𐌸 𐍅𐌰𐌽𐍄 𐍄𐍉 𐌻𐍉𐌰𐌳 𐍄𐌷𐌹𐍃 𐍅𐍉𐍂𐌻𐌳?", "selectWorld.versionUnknown": "unkunþ", "selectWorld.world": "Fairƕus", "soundCategory.block": "Kubjus", "soundCategory.music": "Saggws", "soundCategory.player": "Laikands", "soundCategory.voice": "Stibna/Qiss", "soundCategory.weather": "Wiþr", "spectatorMenu.next_page": "Iftums Laufs", "spectatorMenu.previous_page": "Aftums Laufs", "spectatorMenu.team_teleport": "Fairrabair du Kunjagadailin", "spectatorMenu.team_teleport.prompt": "Walei kuni du fairrabairan du", "spectatorMenu.teleport": "Fairrabair du Laikand", "spectatorMenu.teleport.prompt": "Walei laikand du fairrabairan du", "structure_block.integrity.seed": "Strukturos Fraiw", "subtitles.block.smoker.smoke": "Raukja raukiþ", "subtitles.entity.fox.aggro": "Fauho hatizoþ", "subtitles.entity.fox.ambient": "Fauho peipoþ", "subtitles.entity.fox.bite": "Fauho beitiþ", "subtitles.entity.fox.death": "Fauho gaswiltiþ", "subtitles.entity.fox.eat": "Fauho itiþ", "subtitles.entity.fox.hurt": "Fauho gawinniþ", "subtitles.entity.fox.screech": "Fauho kreitiþ", "subtitles.entity.fox.sleep": "Fauho slepiþ", "subtitles.entity.fox.spit": "Fauhso speiwiþ", "subtitles.entity.panda.bite": "Pandabaira beitiþ", "subtitles.entity.panda.death": "Pandabaira gaswiltiþ", "subtitles.entity.panda.eat": "Pandabaira itiþ", "subtitles.entity.panda.hurt": "Pandabaira gawinniþ", "subtitles.entity.panda.step": "Panda trimpiþ", "subtitles.entity.ravager.attack": "Qistja beitiþ", "subtitles.entity.ravager.death": "Qistja gaswiltiþ", "subtitles.entity.ravager.hurt": "Qistja gasleiþjada", "subtitles.entity.ravager.roar": "Qistja auhjoþ", "subtitles.entity.ravager.step": "Qistja trimpiþ", "subtitles.entity.ravager.stunned": "Qistja gamarzjada", "subtitles.entity.villager.work_butcher": "Skilja arbaideiþ", "subtitles.entity.villager.work_cartographer": "Kartja arbaideiþ", "subtitles.entity.villager.work_cleric": "Gudja arbaideiþ", "subtitles.entity.villager.work_farmer": "Waurtja arbaideiþ", "subtitles.entity.villager.work_fisherman": "Fiskja arbaideiþ", "subtitles.entity.villager.work_librarian": "Bokahusja arbaideiþ", "subtitles.entity.villager.work_mason": "Timrja arbaideiþ", "subtitles.entity.villager.work_shepherd": "Hairdeis arbaidiþ", "subtitles.entity.villager.work_toolsmith": "Tolasmiþja arbaideiþ", "subtitles.entity.villager.work_weaponsmith": "Wepnasmiþja arbaideiþ", "subtitles.entity.wandering_trader.death": "Farands Lewja gaswiltiþ", "subtitles.entity.wandering_trader.hurt": "Farands Lewja gasleiþjada", "subtitles.entity.wandering_trader.no": "Farands Lewja unwereiþ", "subtitles.entity.wandering_trader.trade": "Farands Lewja leweiþ", "subtitles.entity.wandering_trader.yes": "Farands Lewja miþqiþiþ", "subtitles.entity.witch.celebrate": "Haljaruna hlasjiþ", "subtitles.item.berries.pick": "Basja paupond", "subtitles.item.book.page_turn": "Boka gawagjiþ", "subtitles.item.crop.plant": "Fraiw saian", "subtitles.item.crossbow.load": "Krukabuga manweiþ", "subtitles.item.crossbow.shoot": "Krukabuga skiutiþ", "subtitles.item.nether_wart.plant": "Fraiw saian", "team.notFound": "Unkunþ kuni '%s'", "translation.test.args": "%s %s", "translation.test.complex": "Fauralageins, %s%2s aftra %s jah %1s þata aftumist %s jah %1s jah aftra!", "translation.test.invalid": "fagino %", "translation.test.invalid2": "fagino %s", "translation.test.none": "Hails, fairƕau!", "translation.test.world": "fairƕus"}