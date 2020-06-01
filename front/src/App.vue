<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <QuerySelector v-on:update-query="updateQuery" />
    <TableView v-bind:header="header" v-bind:records="records" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios from 'axios';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';
import QuerySelector from './components/QuerySelector.vue';
import TableView from './components/TableView.vue';

interface QueryString {
  column: string;
  operator: string;
  val: string;
}

interface QueryParam {
  table: string;
  where: QueryString[];
  order: string;
}

interface ResponseData {
  data: {
    header: string[];
    records: string[][];
  };
}

interface DataType {
  header: string[];
  records: string[][];
}
interface MethodType {
  updateQuery: (query: QueryParam) => void;
}
interface ComputedType {}
interface PropType {}

export default Vue.extend({
  name: 'App',
  components: {
    QuerySelector,
    TableView,
  },
  data() {
    return {
      header: [],
      records: [],
    };
  },
  methods: {
    async updateQuery(query: QueryParam) {
      console.log('updateQuery');
      console.log(query);
      try {
        const response = await axios.post<ResponseData>('/select', query);
        console.log('records');
        console.log(response.data.data);
        this.records = response.data.data.records;
      } catch (error) {
        return null;
      }
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
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
