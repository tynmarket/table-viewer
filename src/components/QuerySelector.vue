<template>
  <div>
    <div>
      <label>
        テーブル名
        <input v-model="table" type="txt" />
      </label>
      <button v-on:click="search">検索</button>
    </div>
    <div v-for="(query, i) in queries" :key="i">
      <QueryItem v-model="query.value"></QueryItem>
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
  queries: Query[];
}
interface MethodType {
  search: () => void;
  addQuery: () => void;
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
      queries: [
        {
          value: {
            column: 'a',
            operator: '!=',
            val: 'c',
          },
        },
      ],
    };
  },
  methods: {
    search() {
      console.log(`table: ${this.table}`);
      console.log(`query: ${this.whereQuery()}`);
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
