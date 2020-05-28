<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <TableView />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import TableView from './components/TableView.vue';

export default Vue.extend({
  name: 'App',
  components: {
    TableView,
  },
  data() {
    return {
      headers: ['ID', 'Code', 'Name'],
      items: [
        ['ID', 'Code', 'Name'],
        [1, '4368', '扶桑化学'],
        [2, '9985', 'ソフトバンク'],
      ],
    };
  },
  computed: {
    data() {
      // @ts-ignore
      return this.items.slice(1);
    },
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
