<template>
  <div id="voteInfo">
    <b-table noCollapse headVariant="dark"
      :items="items"
      :fields="fields">
    </b-table>
  </div>
</template>


<script>
export default {
  name: 'voteInfo',
  props: {
    items: {
      type: Array,
      default: () => ({})
    }
  },
  data() {
    return {
      fields : [ {key : 'userId', label: '사용자'} , {key : 'reason', label: '선택사유'} ]
    }
  }
}
</script>
