<template>
  <div class="sprite">
    <img :src="'https://img.pokemondb.net/sprites/sword-shield/icon/' + sprite + '.png'" 
         :alt="name" 
         :title="name" 
         v-bind:class="{ clickable: clickable }"
         v-on:click="$emit('clicked')"
    >
  </div>
</template>

<script>
export default {
  name: 'Sprite',
  props: {
    name: String,
    clickable: Boolean
  },
  computed: {
    sprite: function() {
      return this.name.toLowerCase().replace(" ", "-").replace(".", "").replace(":", "").replace("galar", "galarian").replace("alola", "alolan").replace("’", "").replace("♀", "-f").replace("♂", "-m");
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.sprite {
  display: inline-block;
}
.clickable {
  border: solid transparent 2px;
  cursor: pointer;
}
.clickable:hover {
  border: solid black 2px;
}
</style>
