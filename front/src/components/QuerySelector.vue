<template>
  <div>
    <div class="field is-grouped">
      <p class="control">
        <label class="label">テーブル名</label>
      </p>
      <p class="control">
        <input
          v-model="table"
          v-on:keyup.enter="search"
          v-focus
          type="txt"
          class="input"
        />
      </p>
      <p class="control">
        <button v-on:click="search" class="button">検索</button>
      </p>
    </div>
    <div v-for="(query, i) in queries" :key="query.key" class="field">
      <QueryItem
        v-model="query.value"
        v-bind:index="i"
        v-bind:newItem="i == 0"
        v-on:add-query="addQuery"
        v-on:delete-query="deleteQuery"
      ></QueryItem>
    </div>
    <div class="field is-grouped">
      <p class="control">
        <label class="label">ORDER BY id</label>
      </p>
      <p class="control">
        <span class="select">
          <select v-model="order">
            <option value="ASC">ASC</option>
            <option value="DESC">DESC</option>
          </select>
        </span>
      </p>
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

type Query = { value: QueryString; key: string };

interface QueryParam {
  table: string;
  where: QueryString[];
  order: string;
}

interface DataType {
  table: string;
  queries: Query[];
  order: string;
}
interface MethodType {
  search: () => void;
  addQuery: () => void;
  deleteQuery: (index: number) => void;
  whereQueryString: () => string;
  buildQuery: () => QueryParam;
  whereQuery: () => QueryString[];
  validQueries: () => Query[];
  orderQuery: () => string;
}
interface ComputedType {}
interface PropType {}

const generateKey = (): string => {
  return Math.random()
    .toString(32)
    .substring(2);
};

const newQuery = (): Query => {
  return {
    value: {
      column: '',
      operator: '=',
      val: '',
    },
    key: generateKey(),
  };
};

export default Vue.extend({
  name: 'QuerySelector',
  components: {
    QueryItem,
  },
  props: {},
  data() {
    return {
      table: null,
      queries: [newQuery()],
      order: 'ASC',
    };
  },
  methods: {
    search() {
      console.log(`table: ${this.table}`);
      console.log(`query: ${this.whereQueryString()}`);
      console.log(`order: ${this.orderQuery()}`);
      const query = this.buildQuery();
      this.$emit('update-query', query);
    },

    addQuery(query: QueryString) {
      const queries = [
        newQuery(),
        ...this.queries.slice(1),
        {
          value: query,
          key: generateKey(),
        },
      ];
      this.queries = queries;
    },

    deleteQuery(index: number) {
      const queries = this.queries.filter(
        (query: Query, i: number): boolean => {
          return i != index;
        }
      );
      queries.forEach((value: Query) => {
        const query = value.value;
        console.log(`${query.column} ${query.operator} ${query.val}`);
      });
      this.queries = queries;
    },

    buildQuery(): QueryParam {
      return {
        table: this.table,
        where: this.whereQuery(),
        order: this.order,
      };
    },

    whereQuery(): QueryString[] {
      return this.validQueries().map(
        (query: Query): QueryString => query.value
      );
    },

    validQueries(): Query[] {
      return this.queries.filter((query: Query): boolean => {
        return query.value.column != '';
      });
    },

    whereQueryString(): string {
      const queries = this.validQueries();
      if (queries.length > 0) {
        return queries
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
