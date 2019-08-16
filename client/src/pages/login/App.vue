<template>
  <v-app>
    <v-content>
      <Login/>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'vue-property-decorator';
import Login from '../../components/Login.vue';

@Component({
  components: {
    Login,
  },
})
export default class App extends Vue {
}
</script>
