<template>
  <div class="field is-grouped">
    <p class="control">
      <label class="label">WHERE</label>
    </p>
    <p class="control">
      <input v-model="column" type="txt" class="input" />
    </p>
    <p class="control">
      <span class="select">
        <select v-model="operator">
          <option value="=">=</option>
          <option value="!=">≠</option>
          <option value=">">&gt;</option>
          <option value="<">&lt;</option>
          <option value=">=">≧</option>
          <option value="<=">≦</option>
          <option value="IN">IN</option>
          <option value="LIKE">LIKE</option>
          <option value="BETWEEN">BETWEEN</option>
          <option value="IS NULL">IS NULL</option>
          <option value="IS NOT NULL">IS NOT NULL</option>
        </select>
      </span>
    </p>
    <p class="control">
      <input v-model="val" type="txt" class="input" />
    </p>
    <p class="control">
      <button v-if="newItem" v-on:click="addQuery" class="button">＋</button>
      <button v-else v-on:click="deleteQuery(index)" class="button">-</button>
    </p>
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
