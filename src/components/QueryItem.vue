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
    <button>＋</button>
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
interface MethodType {}
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
    this.$emit('input', {
      column: this.column,
      operator: this.operator,
      val: this.val,
    });
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>
