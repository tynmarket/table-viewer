<template>
  <div>
    <label>
      クエリ
      <input v-model="column" type="txt" />
      <select v-model="operator">
        <option value="=">=</option>
        <option value="!=">!=</option>
      </select>
      <input v-model="val" type="txt" />
    </label>
    <button v-on:click="addQuery">＋</button>
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
}
interface ComputedType {}
interface PropType {
  value: QueryString;
}

export default Vue.extend({
  name: 'QueryItem',
  props: {
    value: {
      type: Object,
      required: true,
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
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>
