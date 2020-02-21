<template>
  <div id="userPoint">
    <div align="middle" class="common-margin">
      <h3>{{ nickname }} 님의 점수</h3>
    </div>
    <div class="point-div">
    <div>
      <b-img :src="require('../../assets/img/red_point_48.png')"></b-img>
      <span class="point-rate font-color-red">{{ honorPoint }}</span>
    </div>
    <div>
      <b-img :src="require('../../assets/img/yellow_point_48.png')"></b-img>
      <span class="point-rate font-color-yellow">{{ winPoint }}</span>
    </div>
    </div>
  </div>
</template>


<script>
export default {
  name: 'userPoint',
  data() {
      return {
        honorPoint : '290',
        winPoint : '300',
        nickname: 'user123'
      }
  }
}
</script>
