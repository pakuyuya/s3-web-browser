<template>
    <v-dialog v-model="dialog" width="500">
        <v-card>
            <v-toolbar dark color="error">
                <v-btn icon dark @click="close">
                    <v-icon>close</v-icon>
                </v-btn>
                <v-toolbar-title>Error</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                {{message}}
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
// import Vue from 'vue';
import {Component, Prop, Vue, Watch} from 'vue-property-decorator';

@Component
export default class ErrorDialog extends Vue {
  public message: string = '';
  public dialog: boolean = false;

  public openWith(message?: string) {
    this.dialog = true;
    this.message = message || 'システムエラーが発生しました。再試行していただくか、復旧しない場合管理者にご連絡ください。'
  }

  public close() {
    this.dialog = false;
  }
}
</script>
