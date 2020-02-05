<template>
<div>
  <div class="outline-margin">
    <div align="middle" class="common-margin">
      <h3>{{ nickname }} 님의 점수</h3>
    </div>
    <userPoint />
  </div>
  <div class="outline-margin">
    <div align="middle" class="common-margin">
      <h3>전체 랭킹</h3>
    </div>
    <rankTable />
  </div>
</div>
</template>


<script>
import UserPoint from './user/UserPoint.vue'
import RankTable from './rank/RankTable.vue'

export default {
  name: 'rank',
  components : {
    UserPoint,
    RankTable
  },
  data() {
    return {
      nickname: 'user123'
    }
  }
}
</script>
