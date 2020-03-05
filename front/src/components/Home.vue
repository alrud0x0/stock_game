<template>
<div>
  <userPoint />
  <div v-if="isVotingAvaliable">
    <b-button block variant="info" :to="{path : '/game/upcoming'}">게임 참여하기 <b-icon icon="arrow-bar-right"></b-icon></b-button>
  </div>
  <gameTable />
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
