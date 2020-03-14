<template>
<div>
  <div align="middle" class="common-margin">
    <h3>지난 게임 목록</h3>
  </div>

  <b-list-group>
      <b-list-group-item class="d-flex justify-content-between align-items-center"
          v-for="(value, index) in games"
          :key="value"
          :href="'history/' + index">{{ value }}
        <b-badge v-if="index == 0" variant="primary" pill>NEW</b-badge>
      </b-list-group-item>
  </b-list-group>
</div>
</template>


<script>
export default {
  name: 'gameHistory',
  components: {

  },
  data () {
    return {
      games : ['2020.03.13', '2020.03.06', '2020.02.28', '2020.02.21', '2020.02.14', '2020.02.07']
    }
  }
}
</script>
