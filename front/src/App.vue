<template>
  <div id="app" class="container">
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
  authUser: string;
  authPassword: string;
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
      authUser: '',
      authPassword: '',
    };
  },
  mounted() {
    const paramStr = location.href.split('?');

    if (paramStr[1]) {
      const params = new URLSearchParams(paramStr[1]);
      const user = params.get('user');
      const password = params.get('password');

      if (user) {
        this.authUser = user;
      }
      if (password) {
        this.authPassword = password;
      }
    }
  },
  methods: {
    async updateQuery(query: QueryParam) {
      console.log('updateQuery');
      console.log(query);
      try {
        const auth = {
          auth: { username: this.authUser, password: this.authPassword },
        };
        const response = await axios.post<ResponseData>(
          '/api/select',
          query,
          auth
        );
        console.log('header');
        console.log(response.data.data.header);
        console.log('records');
        console.log(response.data.data.records);
        this.header = response.data.data.header;
        this.records = response.data.data.records;
      } catch (error) {
        return null;
      }
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
