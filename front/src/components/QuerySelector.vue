<template>
  <div>
    <div>
      <label>
        テーブル名
        <input v-model="table" type="txt" />
      </label>
      <button v-on:click="search">検索</button>
    </div>
    <QueryItem
      v-model="newQuery"
      v-bind:newItem="true"
      v-on:add-query="addQuery"
    ></QueryItem>
    <div v-for="(query, i) in queries" :key="query.value.val">
      <QueryItem
        v-model="query.value"
        v-bind:index="i"
        v-on:delete-query="deleteQuery"
      ></QueryItem>
    </div>
    <div>
      <label>
        ORDER BY id
        <select v-model="order">
          <option value="ASC">ASC</option>
          <option value="DESC">DESC</option>
        </select>
      </label>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import axios from 'axios';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';
import QueryItem from './QueryItem.vue';

interface QueryString {
  column: string;
  operator: string;
  val: string;
}

type Query = { value: QueryString };

interface QueryParam {
  table: string;
  where: QueryString[];
  order: string;
}

interface ResponseData {
  data: {
    header: string[];
    record: (string | number | boolean | null)[][];
  };
}

interface DataType {
  table: string;
  newQuery: QueryString;
  queries: Query[];
  order: string;
}
interface MethodType {
  search: () => void;
  addQuery: () => void;
  deleteQuery: (index: number) => void;
  whereQueryString: () => string;
  sendQuery: () => Promise<ResponseData>;
  buildQuery: () => QueryParam;
  whereQuery: () => QueryString[];
  orderQuery: () => string;
}
interface ComputedType {}
interface PropType {}

export default Vue.extend({
  name: 'QuerySelector',
  components: {
    QueryItem,
  },
  props: {},
  data() {
    return {
      table: null,
      newQuery: {
        column: '',
        operator: '=',
        val: '',
      },
      queries: [],
      order: 'ASC',
    };
  },
  methods: {
    search() {
      console.log(`table: ${this.table}`);
      console.log(`query: ${this.whereQueryString()}`);
      console.log(`order: ${this.orderQuery()}`);
      this.sendQuery();
    },

    addQuery(query: QueryString) {
      const queries = [
        ...this.queries,
        {
          value: {
            column: query.column,
            operator: query.operator,
            val: query.val,
          },
        },
      ];
      this.queries = queries;
      // TODO 親側からnewQueryをクリアできない？
    },

    deleteQuery(index: number) {
      const len = this.queries.length + 1;
      const queries = [
        ...this.queries.slice(0, index),
        ...this.queries.slice(index + 1, len),
      ];
      queries.forEach((value: Query) => {
        const query = value.value;
        console.log(`${query.column} ${query.operator} ${query.val}`);
      });
      this.queries = queries;
    },

    async sendQuery(): Promise<ResponseData | null> {
      const body = this.buildQuery();

      try {
        const response = await axios.post('/v1/select', body);
        return response.data;
      } catch (error) {
        return null;
      }
    },

    buildQuery(): QueryParam {
      return {
        table: this.table,
        where: this.whereQuery(),
        order: this.order,
      };
    },

    whereQuery(): QueryString[] {
      return this.queries.map((query: Query): QueryString => query.value);
    },

    whereQueryString(): string {
      if (this.queries.length > 0) {
        return this.queries
          .map((query: Query): string => {
            const value = query.value;
            return `${value.column} ${value.operator} ${value.val}`;
          })
          .join(' AND ');
      } else {
        return '';
      }
    },

    orderQuery(): string {
      return `ORDER BY id ${this.order}`;
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>