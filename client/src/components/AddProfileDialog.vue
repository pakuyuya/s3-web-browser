<template>
    <v-dialog v-model="dialog" width="500">
        <v-form ref="form" v-model="valid" lazy-validation>
            <v-card>
                <v-toolbar dark color="primary">
                    <v-btn icon dark @click="close">
                        <v-icon>close</v-icon>
                    </v-btn>
                    <v-toolbar-title>>Add S3 Contents</v-toolbar-title>
                </v-toolbar>
                <v-card-text>
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs12 sm6 md4>
                                <v-text-field label="Profile name" v-model="profilename" :rules="rulesProfilename"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4>
                                <v-text-field label="Bucket" v-model="bucket" :rules="rulesBucket"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4>
                                <v-text-field label="Region" v-model="region" :rules="rulesRegion"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4>
                                <v-radio-group v-model="conn.type" :rules="rulesConnType">
                                    <v-radio
                                        v-for="profileType in listProfileType"
                                        :key="profileType.value"
                                        :label="profileType.caption"
                                        :value="profileType.value"
                                    ></v-radio>
                                </v-radio-group>
                            </v-flex>
                            <v-flex xs12 sm6 md4 v-if="conn.type === 'accesskey'">
                                <v-text-field label="accesskey" v-model="conn.accesskey" :rules="rulesConnAccesskey"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4 v-if="conn.type === 'accesskey'">
                                <v-text-field label="secretkey" v-model="conn.secretkey" :rules="rulesConnSecretkey"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4 v-if="conn.type === 'credentialfile'" :rules="rulesConnProfile">
                                <v-text-field label="profile name in cledentials file" v-model="conn.profile"></v-text-field>
                            </v-flex>
                        </v-layout>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary" @click="submitForm">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-form>
    </v-dialog>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue, Watch} from 'vue-property-decorator';

import {ProfileStore} from '../store/modules/profile';


@Component
export default class AddProfileDialog extends Vue {
  public profile = ProfileStore.CreateProxy( this.$store, ProfileStore );

  public dialog: boolean = false;
  public valid: boolean = true;
  public profilename: string = '';
  public bucket: string = '';
  public region: string = '';
  public conn: {
    type: string,
    profile: string,
    accesskey: string,
    secretkey: string,
  } = {
      type: 'acccesskey',
      profile: '',
      accesskey: '',
      secretkey: '',
  };

  public rulesProfilename = [
      (v: string) => !!v || 'Profilename is required',
      (v: string) => (v && v.length <= 32) || 'Profilename is must be less than 32 characters',
  ];

  public rulesBucket = [
      (v: string) => !!v || 'Bucket is required',
      (v: string) => (v && v.length <= 128) || 'Region is must be less than 128 characters',
  ];

  public rulesRegion = [
      (v: string) => !!v || 'Region is required',
      (v: string) => (v && v.length <= 32) || 'Region is must be less than 32 characters',
  ];

  public rulesConnType = [
      (v: string) => !!v || 'Type is required',
  ];

  public rulesConnAccesskey = [
      (v: string) => !!v || 'Accesskey is required',
      (v: string) => (v && v.length <= 128) || 'Accesskey is must be less than 128 characters',
  ];

  public rulesConnSecretkey = [
      (v: string) => !!v || 'Secretkey is required',
      (v: string) => (v && v.length <= 128) || 'Secretkey is must be less than 128 characters',
  ];

  public rulesConnProfile = [
      (v: string) => !!v || 'Profile is required',
      (v: string) => (v && v.length <= 128) || 'Profile is must be less than 128 characters',
  ];

  public listProfileType: Array<{value: string, caption: string}> = [
      {value: 'accesskey', caption: 'Using Access key'},
      {value: 'credentialfile', caption: 'Using Credential file'},
    ];

  public initopen() {
    this.dialog = true;

    this.profilename = '';
    this.bucket = '';
    this.region = '';
    this.conn = {
        type: 'acccesskey',
        profile: '',
        accesskey: '',
        secretkey: '',
    };
  }

  public close() {
    this.dialog = false;
  }

  public submitForm() {
    if (!(this.$refs.form as any).validate()) {
      return;
    }

    this.saveProfile();
  }

  public saveProfile() {
    const jsonmodel: any = { type: this.conn.type };
    switch (this.conn.type) {
    case 'accesskey':
        jsonmodel.accesskey = this.conn.accesskey;
        jsonmodel.secretkey = this.conn.secretkey;
        break;
    case 'credentialfile':
        jsonmodel.profile = this.conn.profile;
        break;
    }

    const model = {
        profilename: this.profilename,
        bucket: this.bucket,
        region: this.region,
        connjson: JSON.stringify(jsonmodel),
    };
    this.profile.insert(model)
        .then(() => {
            this.close();
        });
  }

}
</script>
