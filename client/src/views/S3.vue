<template>
  <div>
    <v-card class="ma-5">
      <v-breadcrumbs :items="s3dir.breadcrumbs" divider="/">
        <template v-slot:item="props">
          <v-breadcrumbs-item
            :class="[props.item.disabled && 'disabled']"
          >
            <router-link :to="props.item.to" v-if="!props.item.disabled">{{ props.item.text }}</router-link>
            <span v-if="props.item.disabled">{{ props.item.text }}</span>
          </v-breadcrumbs-item>
        </template>
      </v-breadcrumbs>
    </v-card>
    <v-alert class="ma-5" type="warning" v-if="s3dir.error">
      {{s3dir.error}}
    </v-alert>
    <v-card class="ma-5">

      <v-layout
      >
        <v-flex
          xs12
          sm12
          md12
        >
          <v-list>
            <v-list-item v-if="this.s3dir.currentDir !== ''" @click="upDirectory()">
              <v-list-item-action><v-icon>chevron_left</v-icon></v-list-item-action>
              <v-list-item-title>..</v-list-item-title>
            </v-list-item>
          </v-list>
          <v-list v-for="file in this.s3dir.files" :key="file.name">
            <v-list-item v-if="file.type === 'file'" @click="downloadFile(file)">
              <v-list-item-action><v-icon>cloud_download</v-icon></v-list-item-action>
              <v-list-item-title>{{ file.name }}</v-list-item-title>
            </v-list-item>
            <v-list-item v-if="file.type === 'directory'" @click="moveTo(`${file.fullpath}`)">
              <v-list-item-action><v-icon>folder</v-icon></v-list-item-action>
              <v-list-item-title>{{ file.name }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-flex>
      </v-layout>
    </v-card>
  </div>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue, Watch} from 'vue-property-decorator';

import {S3dirStore, S3Item} from '../store/modules/s3dir';
import {ProfileStore} from '../store/modules/profile';

Component.registerHooks([
  'beforeRouteEnter',
  'beforeRouteLeave',
  'beforeRouteUpdate',
]);

@Component
export default class S3Dir extends Vue {
  public profile = ProfileStore.CreateProxy( this.$store, ProfileStore );
  public s3dir = S3dirStore.CreateProxy( this.$store, S3dirStore );

  public renderComponent: boolean = true;

  @Watch('$route')
  public onRouteChange(to: any, from: any): void {
    this.refleshPath();
  }

  public downloadFile(item: S3Item): void {
    this.s3dir.download(item);
  }

  public moveTo(path: string): void {
    const profile = this.s3dir.s3profile || {profileid: ''};
    path = path.replace(/^\/|\/$/gi, '');
    this.$router.replace(`/s3/${profile.profileid || ''}/${path}`);
  }

  public upDirectory(): void {
    const currentDir = this.s3dir.currentDir || '';
    const idx = currentDir.lastIndexOf('/');

    const path = idx >= 0 ? currentDir.substring(0, idx) : '';
    this.moveTo(path);
  }

  public mounted() {
    this.refleshPath();
  }

  private refleshPath() {
    (async () => {
      if (this.$route.params.profileid) {
        const profileid = this.$route.params.profileid;
        const profile = await this.profile.findById(profileid);
        if (profile) {
          this.s3dir.updateProfile({profile});
        }
      }

      this.s3dir.setCurrentDir({
        path: this.$route.params.path || '',
      });
    })();
  }

}
</script>
