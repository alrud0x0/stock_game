<template lang="html">
<b-navbar toggleable="md" type="light" sticky variant="faded" >
  <b-navbar-brand tag="h1" class="mb-0">Stock Game</b-navbar-brand>

  <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

  <b-collapse id="nav-collapse" is-nav>
    <b-navbar-nav class="ml-auto">
      <b-nav-item :to="{path : '/home'}" :active="$route.name == 'home'">HOME</b-nav-item>
      <b-nav-item :to="{path : '/game/upcoming'}" :active="$route.name == 'upcomingGame'">Current Game</b-nav-item>
      <b-nav-item :to="{path : '/game/current'}">Proceeding Game</b-nav-item>
      <b-nav-item :to="{path : '/rank'}" :active="$route.name == 'rank'">Ranking</b-nav-item>
      <b-nav-item :to="{path : '/info'}">Info</b-nav-item>
    </b-navbar-nav>
  </b-collapse>
</b-navbar>
</template>

<script>
export default {
  name: 'topTab'
}


</script>
