<template>
  <v-app>
    <v-navigation-drawer absolute>
      <v-toolbar flat>
        <v-list>
          <v-list-tile>
            <v-list-tile-title class="title">
              S3 Web Browser
            </v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-toolbar>
    </v-navigation-drawer>

    <v-divider></v-divider>

    <v-list dense class="pt-0">
      <v-list-tile v-for="profile in profiles" :key="profile.id" @click="selectProfile">
        <v-list-tile-action>
          <v-icon>cloud_circle</v-icon>
        </v-list-tile-action>
        <v-list-tile-component>
          <v-list-tile-title>{{ profile.name }}</v-list-tile-title>
        </v-list-tile-component>
      </v-list-tile>
    </v-list>

    <v-container>
      <v-layout justify-center v-if="selectedProfile === undefined" >
        接続先を選択してください。
      </v-layout>
      <v-list v-if="selectedProfile !== undefined" class="text-xs-center">
        <s3list />
      </v-list>
    </v-container>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import S3list from './components/S3list.vue';

export default Vue.extend({
  name: 'App',
  components: {
    S3list,
  },
  data: () => ({
    profiles : [],
    selectedProfile: undefined,
  }),
  methods: {
    selectProfile () {

    },
  },
});
</script>
