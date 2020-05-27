<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <HelloWorld msg="Welcome to Your Vue.js + TypeScript App" />
    <table>
      <thead>
        <tr>
          <th v-for="(header, i) in items[0]" :key="i">
            {{ header }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, i) in data" :key="i">
          <td v-for="(item, j) in row" :key="j">
            {{ item }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import HelloWorld from "./components/HelloWorld.vue";

export default Vue.extend({
  name: "App",
  components: {
    HelloWorld
  },
  data() {
    return {
      headers: ["ID", "Code", "Name"],
      items: [
        ["ID", "Code", "Name"],
        [1, "4368", "扶桑化学"],
        [2, "9985", "ソフトバンク"]
      ]
    };
  },
  computed: {
    data() {
      // @ts-ignore
      return this.items.slice(1);
    }
  }
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
