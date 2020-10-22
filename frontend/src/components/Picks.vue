<template>
  <div class="picks">
    <table>
      <thead>
        <tr>
          <th v-for="player in playersAndPicks" v-bind:key="player.name">
            {{ player.name }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in grid" v-bind:key="row">
          <td v-for="item in row" v-bind:key="item">
            <Sprite :name="item" v-bind:clickable="false"/>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import Sprite from './Sprite.vue'

export default {
  name: 'Picks',
  components: {
    Sprite
  },
  props: {
    playersAndPicks: Array
  },
  computed: {
    grid: function() {
        var res = [];
        var limit = this.playersAndPicks.length > 0 ? Math.max(this.playersAndPicks[0].picks.length, this.playersAndPicks[this.playersAndPicks.length-1].picks.length) : 0;
        for (var i=0; i<limit; i++) {
            var row = [];
            for (var j in this.playersAndPicks) {
                row.push(this.playersAndPicks[j].picks[i]);
            }
            res.push(row);
        }
        return res;
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.picks {
  text-align: center;
  overflow: auto;
}
table {
  margin: auto;
  border: none;
  border-top: solid black 5px;
  border-bottom: solid black 5px;
  border-collapse: collapse;
  width: 100%;
}
table thead {
  border-bottom: solid black 1px;
}
table th {
  text-align: center;
  font-size: 10px;
  overflow: hidden;
  text-overflow: ellipsis;
  min-width: 60px;
}
table td {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  min-width: 60px;
}
table th:nth-child(even) {
  background: lightgray;
}
table td:nth-child(even) {
  background: lightgray;
}
td div.sprite {
  width: 100%;
}
</style>
