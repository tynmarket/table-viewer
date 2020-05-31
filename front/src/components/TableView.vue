<template>
  <table>
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
</template>

<script lang="ts">
import Vue from 'vue';
import { ThisTypedComponentOptionsWithRecordProps } from 'vue/types/options';

interface DataType {
  allRecords: (string | number)[][];
}
interface MethodType {}
interface ComputedType {}
interface PropType {}

export default Vue.extend({
  name: 'TableView',
  data() {
    return {
      allRecords: [
        ['id', 'code', 'name'],
        [1, '4368', '扶桑化学'],
        [2, '9985', 'ソフトバンク'],
      ],
    };
  },
  computed: {
    header() {
      return this.allRecords[0];
    },
    records() {
      return this.allRecords.slice(1);
    },
  },
} as ThisTypedComponentOptionsWithRecordProps<Vue, DataType, MethodType, ComputedType, PropType>);
</script>

<style lang="scss"></style>
