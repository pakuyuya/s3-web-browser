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
                    @keydown="keydown"
                  ></v-text-field>

                  <v-text-field
                    id="password"
                    label="Password"
                    name="password"
                    prepend-icon="lock"
                    type="password"
                    v-model="password"
                    @keydown="keydown"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" type="submit" @click="login">Login</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';

@Component
export default class Login extends Vue {
  public loginid: string = '';
  public password: string = '';

  public login(): void {
    // TODO:
  }
  public keydown(event: KeyboardEvent): void {
    if (event.key === 'Enter') {
      this.login();
    }
  }

  get redirectTo() {
    return this.$store.state.redirectTo;
  }

  @Watch('redirectTo')
  public onChangeLogined(redirectTo: string) {
    if (redirectTo !== '') {
      window.location.replace(redirectTo);
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
