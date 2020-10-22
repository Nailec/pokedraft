<template>
  <div class="pool">
    <ul class="nav nav-tabs nav-justified">
      <li class="nav-item">
        <a class="nav-link" @click.prevent="setActive('all')" :class="{ active: isActive('all') }" href="#all">All</a>
      </li>
      <li v-for="section in pool" v-bind:key="section.title" class="nav-item">
        <a class="nav-link" @click.prevent="setActive(section.title)" :class="{ active: isActive(section.title) }" href="#section.title">{{ section.title }}</a>
      </li>
    </ul>
    <div class="tab-content">
      <div class="tab-pane fade" :class="{ 'active show': isActive('all') }" id="all">
        <Sprite v-for="pokemon in full" v-bind:key="pokemon" :name="pokemon" v-bind:clickable="true" @clicked="$emit('do-pick', pokemon)"/>
      </div>
      <div v-for="section in pool" v-bind:key="section.title" class="tab-pane fade" :class="{ 'active show': isActive(section.title) }" id="section.title">
        <Sprite v-for="pokemon in section.list" v-bind:key="pokemon" :name="pokemon" v-bind:clickable="true" @clicked="$emit('do-pick', pokemon)"/>
      </div>
      <div>
      </div>
    </div>
  </div>
</template>

<script>
import Sprite from './Sprite.vue'

export default {
  name: 'Pool',
  components: {Sprite},
  props: {
    pool: Array
  },
  data: function () {
    return {activeTab: 'all'}
  },
  computed: {
    full: function() {
        var res = [];
        for (var section of this.pool) {
            res = res.concat(section.list);
        }

        return res.filter((item, pos) => res.indexOf(item) === pos);
    }
  },
  methods: {
    isActive(tab) {
      return this.activeTab === tab
    },
    setActive(tab) {
      this.activeTab = tab
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.tab-pane {
  text-align: center;
}
</style>
