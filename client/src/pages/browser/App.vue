<template>
  <v-app>
    <v-navigation-drawer app clipped fixed v-model="leftDrawer" :mini-variant="leftMiniVariant">
      <v-list dense>
        <v-list-item>
          <v-btn tile outlined color="success" @click="$refs.addProfileDialog.initopen()">
            <v-icon left>add</v-icon> Add S3 profile
          </v-btn>
        </v-list-item>
        <v-list-item v-for="profile in $store.state.profile.list" :key="profile.profileid" @click.stop="selectProfile(profile)">
          <v-list-item-action>
            <v-icon>cloud_circle</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-action>{{ profile.profilename }}</v-list-item-action>
          </v-list-item-content>
          <v-list-item-action @click.stop="$refs.deleteProfileDialog.openWith(profile)">
            <v-icon>close</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar dark app clipped-left clipped-right absolute color="#2277e0">
      <v-btn icon @click.stop="leftDrawer = true">
        <v-icon>menu</v-icon>
      </v-btn>
      <v-toolbar-title>
        S3 Web Browser
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-title>{{ $store.state.user.name }}</v-toolbar-title>
      <v-btn icon @click.stop="rightDrawer = !rightDrawer">
        <v-icon>account_circle</v-icon>
      </v-btn>
    </v-app-bar>

    <v-content><router-view @show-error="onShowError"/></v-content>

    <v-navigation-drawer app fixed right v-model="rightDrawer">
      <v-list>
        <v-list-item @click.stop="rightDrawer = !rightDrawer">
          <v-spacer></v-spacer>
          <v-list-item-action><v-icon>chevron_right</v-icon></v-list-item-action>
        </v-list-item>
        <v-list-item @click.stop="logout">
          <v-list-item-action><v-icon>dashboard</v-icon></v-list-item-action>
          <v-list-item-content><v-list-item-title>ログアウト</v-list-item-title></v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <AddProfileDialog @show-error="onShowError" ref="addProfileDialog"></AddProfileDialog>
    <DeleteProfileDialog @show-error="onShowError" ref="deleteProfileDialog"></DeleteProfileDialog>
    <ErrorDialog ref="errorDialog"></ErrorDialog>
  </v-app>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue} from 'vue-property-decorator';
import AddProfileDialog from '../../components/AddProfileDialog.vue';
import DeleteProfileDialog from '../../components/DeleteProfileDialog.vue';
import ErrorDialog from '../../components/ErrorDialog.vue';

import {ProfileStore, S3Profile} from '../../store/modules/profile';
import {S3dirStore} from '../../store/modules/s3dir';
import {UserStore} from '../../store/modules/user';

@Component({
  components: {
    AddProfileDialog,
    DeleteProfileDialog,
    ErrorDialog,
  },
})
export default class App extends Vue {
  public profile = ProfileStore.CreateProxy(this.$store, ProfileStore);
  public s3dir = S3dirStore.CreateProxy(this.$store, S3dirStore);
  public user = UserStore.CreateProxy(this.$store, UserStore);

  public leftDrawer: boolean = true;

  public leftMiniVariant: boolean = false;

  public rightDrawer: boolean = false;

  public mounted() {
    this.profile.reload();
    this.user.loadLoginInfo();
  }

  public selectProfile(profile: S3Profile) {
    this.s3dir.updateProfile({profile});
    this.$router.push(`/s3/${profile.profileid}/`);
  }

  public onShowError(message: string) {
    console.error(this.$refs.errorDialog);
    (this.$refs.errorDialog as ErrorDialog).openWith(message);
  }

  public logout() {
    window.location.href = '/logout';
  }
}
</script>
