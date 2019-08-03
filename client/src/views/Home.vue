<template>
  <v-list-item>
    <v-list-item-content>
      <v-list-item-title>
        接続先を選択してください。
      </v-list-item-title>
    </v-list-item-content>
  </v-list-item>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue} from 'vue-property-decorator';


@Component
export default class Home extends Vue {
  public leftDrawer: boolean = true;
}
</script>
