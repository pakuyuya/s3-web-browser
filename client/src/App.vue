<template>
  <v-app>
    <v-navigation-drawer app clipped fixed v-model="leftDrawer" :mini-variant="leftMiniVariant">
      <v-list dense>
        <v-list-item v-for="profile in $store.state.profile.list" :key="profile.id" @click="selectProfile(profile.id)">
          <v-list-item-action>
            <v-icon>cloud_circle</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-action>{{ profile.name }}</v-list-item-action>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app clipped-left clipped-right absolute>
      <v-toolbar-title>
        S3 Web Browser
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-title>{{ $store.state.user.name }}</v-toolbar-title>
      <v-btn icon @click.stop="rightDrawer = !rightDrawer">
        <v-icon>account_circle</v-icon>
      </v-btn>
    </v-app-bar>

    
    <v-content><router-view/></v-content>

    <v-navigation-drawer app fixed right v-model="rightDrawer">
      <v-list>
        <v-list-item to="/logout">
          <v-list-item-action><v-icon>dashboard</v-icon></v-list-item-action>
          <v-list-item-content><v-list-item-title>ログアウト</v-list-item-title></v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </v-app>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue} from 'vue-property-decorator';
import S3list from './components/S3list.vue';


@Component
export default class App extends Vue {
  public leftDrawer: boolean = true;

  public leftMiniVariant: boolean = false;

  public rightDrawer: boolean = false;

  public selectProfile(id: string) {
    // TODO:
  }
}
</script>
