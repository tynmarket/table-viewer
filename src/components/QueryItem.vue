<template>
  <div>
    <label>
      クエリ
      <input v-model="column" type="txt" />
      <select v-model="operator">
        <!-- TODO 演算子追加 -->
        <option value="=">=</option>
        <option value="!=">!=</option>
      </select>
      <input v-model="val" type="txt" />
    </label>
    <button v-if="newItem" v-on:click="addQuery">＋</button>
    <button v-else v-on:click="deleteQuery(index)">-</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';

interface QueryString {
  column: string;
  operator: string;
  val: string;
}

type DataType = QueryString;
interface MethodType {
  queryString: () => QueryString;
  addQuery: () => void;
  deleteQuery: (index: number) => void;
}
interface ComputedType {}
interface PropType {
  value: QueryString;
  newItem: boolean;
  index: number;
}

export default Vue.extend({
  name: 'QueryItem',
  props: {
    value: {
      type: Object,
      required: true,
    },
    newItem: {
      type: Boolean,
    },
    index: {
      type: Number,
    },
  },
  data() {
    return {
      column: '',
      operator: '',
      val: '',
    };
  },
  mounted() {
    this.column = this.value.column;
    this.operator = this.value.operator;
    this.val = this.value.val;
  },
  updated() {
    this.$emit('input', this.queryString());
  },
  methods: {
    queryString(): QueryString {
      return {
        column: this.column,
        operator: this.operator,
        val: this.val,
      };
    },
    addQuery() {
      const query = this.queryString();
      this.column = '';
      this.operator = '=';
      this.val = '';
      this.$emit('add-query', query);
    },
    deleteQuery(index: number) {
      this.$emit('delete-query', index);
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>
