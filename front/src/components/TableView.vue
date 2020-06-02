<template>
  <div class="table-container is-size-7">
    <table class="table is-striped">
      <thead>
        <tr>
          <th v-for="(column, i) in header" :key="i">
            {{ column }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(record, i) in records" :key="i">
          <td v-for="(value, j) in record" :key="j">
            {{ value }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';

interface DataType {}
interface MethodType {}
interface ComputedType {}
interface PropType {
  header: string[];
  records: string[][];
}

export default Vue.extend({
  name: 'TableView',
  props: {
    header: {
      type: Array,
      default: [],
    },
    records: {
      type: Array,
      default: [],
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss">
.table {
  margin-top: 30px;
  margin-right: auto;
  margin-left: auto;

  td {
    white-space: nowrap;
  }
}
</style>
