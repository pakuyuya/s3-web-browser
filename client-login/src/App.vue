<template>
  <v-app>
    <v-content>
      <Login/>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import Login from './components/Login.vue';

export default Vue.extend({
  name: 'App',
  components: {
    Login,
  },
  data: () => ({
    //
  }),
});
</script>
