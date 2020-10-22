<template>
  <div class="container">
    <h1> Pokédraft </h1>
    <Picks :playersAndPicks="playersAndPicks"/>
    <Pool :pool="pool" @do-pick="doPick"/>
  </div>
</template>

<script>
// @ is an alias to /src
import Picks from '@/components/Picks.vue'
import Pool from '@/components/Pool.vue'

export default {
  name: 'Draft',
  props: {
    userID: String,
    roomID: String
  },
  components: {
    Picks,
    Pool
  },
  created: function() {
    this.connection = new WebSocket("ws://localhost:8081/ws");

    this.connection.onmessage = function(event) {
      console.log(event);
    }

    this.connection.onopen = function(event) {
      console.log(event)
      console.log("Successfully connected to the echo websocket server...")
    }
  },
  methods: {
    doPick(mon) {
      var message = 'pick|'+ this.roomID + '|' + this.userID + '|' + mon;
      console.log(message);
      this.connection.send(message);
    }
  },
  data: function() {return {
    connection: null,
    playersAndPicks: [
{"name": "Whitter", "picks": ["Clefable", "Aggron-Mega", "Kommo-O", "Slowking", "Tangrowth", "Tornadus", "Cofagrigus", "Haunter", "Electrode", "Thievul", "Flareon"]},
{"name": "Dieu Amphibien", "picks": ["Azumarill", "Steelix-Mega", "Reuniclus", "Salazzle", "Swellow", "Skuntank", "Breloom", "Goodra", "Porygon2", "Arctozolt", "Shedinja"]},
{"name": "Exploziff", "picks": ["Altaria-Mega", "Arcanine", "Nidoqueen", "Mandibuzz", "Empoleon", "Azelf", "Vikavolt", "Gourgeist", "Lycanroc", "Passimian", "Kangaskhan"]},
{"name": "Meracx", "picks": ["Jirachi", "Rotom-Heat", "Manectric-Mega", "Gengar", "Gastrodon", "Weezing-Galar", "Tyrantrum", "Ursaring", "Dusclops", "Clawitzer", "Dustox"]},
{"name": "Patoche", "picks": ["Hydreigon", "Terrakion", "Scizor", "Gligar", "Florges", "Type: Null", "Tangela", "Garbodor", "Eelektross", "Delphox", "Druddigon"]},
{"name": "Ilyes25", "picks": ["Victini", "Aerodactyl-Mega", "Tapu Fini", "Rotom-Mow", "Registeel", "Dragalge", "Audino", "Piloswine", "Spiritomb", "Heracross", "Swirlix"]},
{"name": "BDLC", "picks": ["Necrozma", "Slowbro", "Pangoro", "Beedrill-Mega", "Diggersby", "Bisharp", "Rotom", "Camerupt", "Sneasel", "Corvisquire", "Bewear"]},
{"name": "JK", "picks": ["Cobalion", "Nidoking", "Moltres", "Milotic", "Latios", "Drifblim", "Grimmsnarl", "Eldegoss", "Sudowoodo", "Pincurchin", "Bibarel"]},
{"name": "Chrycks", "picks": ["Pinsir-Mega", "Dragonite", "Tentacruel", "Whimsicott", "Incineroar", "Miltank", "Hitmonchan", "Sandslash", "Raichu-Alola", "Copperajah", "Coalossal"]},
{"name": "zS", "picks": ["Togekiss", "Raikou", "Krookodile", "Nihilego", "Emboar", "Mesprit", "Rotom-Fan", "Mantine", "Froslass", "Lilligant", "Gurdurr"]},
{"name": "Msnt", "picks": ["Heracross-Mega", "Toxapex", "Hippowdon", "Hatterene", "Bronzong", "Thundurus", "Virizion", "Magmortar", "Roselia", "Greedent", "Linoone-Galar"]},
{"name": "Drafeu-kun", "picks": ["Infernape", "Rotom-Wash", "Crobat", "Slurpuff", "Decidueye", "Metagross", "Donphan", "Persian-Alola", "Kyurem", "Meowstic", "Gastly"]},
{"name": "Smouky", "picks": ["Weavile", "Mimikyu", "Magneton", "Silvally", "Ditto", "Flygon", "Starmie", "Cradily", "Golurk", "Hariyama", "Charjabug"]},
{"name": "Cousin", "picks": ["Zapdos", "Suicune", "Ferrothorn", "Sceptile-Mega", "Mienshao", "Comfey", "Rhyperior", "Claydol", "Drapion", "Shiinotic", "Purugly"]},
{"name": "Boulicrok", "picks": ["Salamence", "Slowbro-Mega", "Klefki", "Mudsdale", "Roserade", "Cinderace", "Scrafty", "Jolteon", "Cryogonal", "Misdreavus", "Rampardos"]},
{"name": "Neetox", "picks": ["Rillaboom", "Heatran", "Primarina", "Staraptor", "Ampharos-Mega", "Runerigus", "Dottler", "Hypno", "Lucario", "Weezing", "Sableye"]}
    ],
    pool: [
        {title: "gen1", list: ["Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon", "Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie", "Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill", "Pidgey", "Pidgeotto", "Pidgeot", "Rattata", "Raticate", "Spearow", "Fearow", "Ekans", "Arbok", "Pikachu", "Raichu", "Sandshrew", "Sandslash", "Nidoran♀", "Nidorina", "Nidoqueen", "Nidoran♂", "Nidorino", "Nidoking", "Clefairy", "Clefable", "Vulpix", "Ninetales", "Jigglypuff", "Wigglytuff", "Zubat", "Golbat", "Oddish", "Gloom", "Vileplume", "Paras", "Parasect", "Venonat", "Venomoth", "Diglett", "Dugtrio", "Meowth", "Persian", "Psyduck", "Golduck", "Mankey", "Primeape", "Growlithe", "Arcanine", "Poliwag", "Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam", "Machop", "Machoke", "Machamp", "Bellsprout", "Weepinbell", "Victreebel", "Tentacool", "Tentacruel", "Geodude", "Graveler", "Golem", "Ponyta", "Rapidash", "Slowpoke", "Slowbro", "Magnemite", "Magneton", "Farfetch’d", "Doduo", "Dodrio", "Seel", "Dewgong", "Grimer", "Muk", "Shellder", "Cloyster", "Gastly", "Haunter", "Gengar", "Onix", "Drowzee", "Hypno", "Krabby", "Kingler", "Voltorb", "Electrode", "Exeggcute", "Exeggutor", "Cubone", "Marowak", "Hitmonlee", "Hitmonchan", "Lickitung", "Koffing", "Weezing", "Rhyhorn", "Rhydon", "Chansey", "Tangela", "Kangaskhan", "Horsea", "Seadra", "Goldeen", "Seaking", "Staryu", "Starmie", "Mr. Mime", "Scyther", "Jynx", "Electabuzz", "Magmar", "Pinsir", "Tauros", "Magikarp", "Gyarados", "Lapras", "Ditto", "Eevee", "Vaporeon", "Jolteon", "Flareon", "Porygon", "Omanyte", "Omastar", "Kabuto", "Kabutops", "Aerodactyl", "Snorlax", "Articuno", "Zapdos", "Moltres", "Dratini", "Dragonair", "Dragonite", "Mewtwo", "Mew"]},
        {title: "gen2", list: ["Chikorita"]}
    ]
  }}
}
</script>
