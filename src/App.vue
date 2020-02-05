<template lang="html">
  <div id="app">
    <div class="container-scroller">
      <topTab/>
    </div>
    <b-container fluid>
      <router-view></router-view>
    </b-container>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


import TopTab from './components/common/TopTab.vue'


export default {
  name: 'app',
  components: {
    TopTab,
  }
}
</script>

<style>
  @import "assets/css/main.css";
</style>
