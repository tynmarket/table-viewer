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
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';
import QueryItem from './QueryItem.vue';

interface QueryString {
  column: string;
  operator: string;
  val: string;
}

type Query = { value: QueryString };

interface DataType {
  table: string;
  newQuery: QueryString;
  queries: Query[];
}
interface MethodType {
  search: () => void;
  addQuery: () => void;
  deleteQuery: (index: number) => void;
  whereQuery: () => string;
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
    };
  },
  methods: {
    search() {
      console.log(`table: ${this.table}`);
      console.log(`query: ${this.whereQuery()}`);
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

    whereQuery: function(): string {
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
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>
