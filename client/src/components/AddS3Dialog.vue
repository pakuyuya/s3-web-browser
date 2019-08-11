<template>
    <v-dialog v-model="dialog" width="500">
        <v-card>
            <v-toolbar dark color="primary">
                <v-btn icon dark @click="dialog = false">
                    <v-icon>close</v-icon>
                </v-btn>
                <v-toolbar-title>>Add S3 Contents</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                <v-container grid-list-md>
                    <v-layout wrap>
                        <v-flex xs12 sm6 md4>
                            <v-text-field label="Profile name" v-model="profilename" required></v-text-field>
                        </v-flex>
                        <v-flex xs12 sm6 md4>
                            <v-text-field label="Region" v-model="region" required></v-text-field>
                        </v-flex>
                        <v-flex xs12 sm6 md4>
                            <v-radio-group v-model="connjson.type">
                                <v-radio
                                    v-for="profileType in listProfileType"
                                    :key="profileType.value"
                                    :label="profileType.caption"
                                    :value="profileType.value"
                                ></v-radio>
                            </v-radio-group>
                        </v-flex>
                        <v-flex xs12 sm6 md4 v-if="connjson.type === 'credentialfile'">
                            <v-text-field label="profile name in cledentials file" v-model="connjson.profile" required></v-text-field>
                        </v-flex>
                        <v-flex xs12 sm6 md4 v-if="connjson.type === 'accesskey'">
                            <v-text-field label="accesskey" v-model="connjson.accesskey" required></v-text-field>
                        </v-flex>
                        <v-flex xs12 sm6 md4 v-if="connjson.type === 'accesskey'">
                            <v-text-field label="secretkey" v-model="connjson.secretkey" required></v-text-field>
                        </v-flex>
                    </v-layout>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" @click="saveProfile"></v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue, Watch} from 'vue-property-decorator';

import {S3dirStore} from '../store/modules/s3dir';
import {ProfileStore} from '../store/modules/profile';


@Component
export default class S3Dir extends Vue {
  public profile = ProfileStore.CreateProxy( this.$store, ProfileStore );
  public s3dir = S3dirStore.CreateProxy( this.$store, S3dirStore );

  public renderComponent: boolean = true;

  public listProfileType: {value:string, caption:string}[] = [
      {value: 'credentialfile', caption: 'Using Credential file'},
      {value: 'accesskey', caption: 'Using Access key'},
    ];

  public saveProfile() {
    // TODO:
  }
}
</script>
