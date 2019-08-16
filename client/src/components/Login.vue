<template>
  <v-app id="inspire">
    <v-content>
      <v-container
        fluid
        fill-height
      >
        <v-layout
          align-center
          justify-center
        >
          <v-flex
            xs12
            sm8
            md4
          >
            <v-card class="elevation-12">
              <v-toolbar
                color="primary"
                dark
                flat
              >
                <v-toolbar-title>ログイン</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form @submit="login" ref="loginForm">
                  <v-text-field
                    label="Login"
                    name="loginid"
                    prepend-icon="person"
                    type="text"
                    v-model="loginid"
                    autocomplete="off"
                    @keydown="keydown"
                  ></v-text-field>

                  <v-text-field
                    id="password"
                    label="Password"
                    name="password"
                    prepend-icon="lock"
                    type="password"
                    v-model="password"
                    autocomplete="off"
                    @keydown="keydown"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions class="pb-5 mr-2">
                <v-spacer></v-spacer>
                <v-btn color="primary" type="submit" @click="login">Login</v-btn>
              </v-card-actions>
              <v-card-text v-if="warnMsg">
                <v-alert type="warning">{{warnMsg}}</v-alert>
              </v-card-text>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import { UserStore } from '../store/modules/user';

@Component
export default class Login extends Vue {
  public loginid: string = '';
  public password: string = '';
  public userStore: UserStore = UserStore.CreateProxy(this.$store, UserStore );

  public warnMsg: string = '';

  public login(): void {
    this.warnMsg = '';
    (async () => {
      this.userStore.login({
        loginid: this.loginid,
        password: this.password,
      }).then((redirectTo) => {
        window.location.href = redirectTo;
      }).catch((error) => {
        if (error.response && error.response.status === 401) {
          this.warnMsg = 'IDまたはパスワードが違います。';
        }
        console.error(error);
      });
    })();
  }

  public keydown(event: KeyboardEvent): void {
    if (event.key === 'Enter') {
      this.login();
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
