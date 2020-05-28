<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <QuerySelector />
    <TableView />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import QuerySelector from './components/QuerySelector.vue';
import TableView from './components/TableView.vue';

export default Vue.extend({
  name: 'App',
  components: {
    QuerySelector,
    TableView,
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
