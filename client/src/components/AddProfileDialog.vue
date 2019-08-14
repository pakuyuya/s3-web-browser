<template>
    <v-dialog v-model="dialog" width="500">
        <v-form ref="form" v-model="valid" lazy-validation>
            <v-card>
                <v-toolbar dark color="primary">
                    <v-btn icon dark @click="close">
                        <v-icon>close</v-icon>
                    </v-btn>
                    <v-toolbar-title>Add S3 Contents..</v-toolbar-title>
                </v-toolbar>
                <v-card-text>
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs12 sm12 md12>
                                <v-text-field label="プロファイル名" v-model="profilename" :rules="rulesProfilename"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm12 md12>
                                <v-text-field label="S3 Bucket" v-model="bucket" :rules="rulesBucket"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm12 md12>
                                <v-text-field label="S3 Region" v-model="conn.region" :rules="rulesRegion"></v-text-field>
                            </v-flex>
                            <v-flex>
                                <v-radio-group v-model="conn.type" :rules="rulesConnType">
                                    <v-radio xs12 sm6 md4
                                        v-for="profileType in listProfileType"
                                        :key="profileType.value"
                                        :label="profileType.caption"
                                        :value="profileType.value"
                                    ></v-radio>
                                </v-radio-group>
                            </v-flex>
                            <v-flex xs12 sm12 md12>
                                <v-text-field label="accesskey" v-model="conn.accesskey" :rules="rulesConnAccesskey" :disabled="conn.type !== 'accesskey'"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm12 md12>
                                <v-text-field label="secretkey" v-model="conn.secretkey" :rules="rulesConnSecretkey" :disabled="conn.type !== 'accesskey'"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm12 md12 :rules="rulesConnProfile">
                                <v-text-field label="profile name in cledentials file" v-model="conn.profile" :disabled="conn.type !== 'credentialfile'"></v-text-field>
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
  public conn: {
    type: string,
    profile: string,
    region: string,
    accesskey: string,
    secretkey: string,
  } = {
      type: 'acccesskey',
      profile: '',
      region: '',
      accesskey: '',
      secretkey: '',
  };

  public rulesProfilename = [
      (v: string) => !!v || '必須です',
      (v: string) => (v && v.length <= 32) || '32文字以下で入力してください',
  ];

  public rulesBucket = [
      (v: string) => !!v || '必須です',
      (v: string) => (v && v.length <= 128) || '128文字以下で入力してください',
  ];

  public rulesRegion = [
      (v: string) => !!v || '必須です',
      (v: string) => (v && v.length <= 32) || '32文字以下で入力してください',
  ];

  public rulesConnType = [
      (v: string) => !!v || '必須です',
  ];

  public rulesConnAccesskey = [
      (v: string) => !!v || '必須です',
      (v: string) => (v && v.length <= 128) || '128文字以下で入力してください',
  ];

  public rulesConnSecretkey = [
      (v: string) => !!v || '必須です',
      (v: string) => (v && v.length <= 128) || '128文字以下で入力してください',
  ];

  public rulesConnProfile = [
      (v: string) => !!v || 'Profile is required',
      (v: string) => (v && v.length <= 128) || 'Profile is must be less than 128 characters',
  ];

  public listProfileType: Array<{value: string, caption: string}> = [
      {value: 'accesskey', caption: 'アクセスキー認証'},
      {value: 'credentialfile', caption: 'Credentials ファイル'},
    ];

  public initopen() {
    this.dialog = true;

    this.profilename = '';
    this.bucket = '';
    this.conn = {
        type: 'acccesskey',
        profile: '',
        region: '',
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
    const jsonmodel: any = {
        type: this.conn.type,
        region: this.conn.region,
    };
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
        connjson: JSON.stringify(jsonmodel),
    };
    this.profile.insert(model)
        .then(() => {
            this.close();
        })
        .catch((error) => {
            console.log(error);
            const msg = (error && error.response && error.response.data && error.response.data.message) || '';
            this.$emit('show-error', msg);
        });
  }

}
</script>
