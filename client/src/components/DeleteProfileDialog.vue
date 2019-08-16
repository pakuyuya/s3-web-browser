<template>
    <v-dialog v-model="dialog" width="500">
        <v-card>
            <v-toolbar dark color="primary">
                <v-btn icon dark @click="close">
                    <v-icon>close</v-icon>
                </v-btn>
                <v-toolbar-title>Confirm</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                S3接続プロファイル {{selprofile.profilename}} を削除しますか？
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="deleteProfile">Delete</v-btn>
                <v-btn @click="close">Cancel</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue, Watch} from 'vue-property-decorator';

import {ProfileStore, S3Profile} from '../store/modules/profile';


@Component
export default class DeleteProfileDialog extends Vue {
  public profile = ProfileStore.CreateProxy( this.$store, ProfileStore );
  public selprofile: S3Profile = {};

  public dialog: boolean = false;

  public openWith(profile: S3Profile) {
    this.dialog = true;
    this.selprofile = profile;
  }

  public close() {
    this.dialog = false;
  }

  public deleteProfile() {
    this.profile.delete(this.selprofile)
        .then(() => {
            this.close();
        })
        .catch((error: any) => {
            console.error(error);
            const msg = (error && error.response && error.response.data && error.response.data.message) || '';
            this.$emit('show-error', msg);
        });
  }

}
</script>
