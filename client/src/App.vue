<template>
  <v-app>
    <v-navigation-drawer app clipped fixed v-model="leftDrawer" :mini-variant="leftMiniVariant">
      <v-list dense>
        <v-list-item>
          <v-btn tile outlined color="success" @click="$refs.addProfileDialog.initopen()">
            <v-icon left>add</v-icon> Add S3 profile
          </v-btn>
        </v-list-item>
        <v-list-item v-for="profile in $store.state.profile.list" :key="profile.profileid" @click="selectProfile(profile)">
          <v-list-item-action>
            <v-icon>cloud_circle</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-action>{{ profile.profilename }}</v-list-item-action>
          </v-list-item-content>
          <v-list-item-action>
            <v-icon>close</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar dark app clipped-left clipped-right absolute color="#2277e0">
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
    <AddProfileDialog ref="addProfileDialog"></AddProfileDialog>
  </v-app>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue} from 'vue-property-decorator';
import AddProfileDialog from './components/AddProfileDialog.vue';

import {ProfileStore, S3Profile} from './store/modules/profile';
import {S3dirStore} from './store/modules/s3dir';

@Component({
  components: {
    AddProfileDialog,
  },
})
export default class App extends Vue {
  public profile = ProfileStore.CreateProxy( this.$store, ProfileStore );
  public s3dir = S3dirStore.CreateProxy( this.$store, S3dirStore );

  public leftDrawer: boolean = true;

  public leftMiniVariant: boolean = false;

  public rightDrawer: boolean = false;

  public mounted() {
    this.profile.reload();
  }

  public selectProfile(profile: S3Profile) {
    this.s3dir.updateProfile({profile});
    this.$router.push(`/s3/${profile.profileid}/`);
  }
}
</script>
