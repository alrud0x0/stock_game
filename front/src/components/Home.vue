<template>
<div>
  <div class="outline-margin">
    <userPoint />
  </div>
  <div class="outline-margin">
    <gameTable />
  </div>
</div>
</template>


<script>
import GameTable from './game/GameTable.vue'
import UserPoint from './user/UserPoint.vue'

export default {
  name: 'home',
  components: {
    GameTable,
    UserPoint
  },
  data () {
    return {
      isVotingAvaliable : true
    }
  }
}
</script>
