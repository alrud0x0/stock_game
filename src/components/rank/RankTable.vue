<template>
  <div id="rankTable">
    <div>
      <b-table striped hover :items="items" :fields="fields"></b-table>
    </div>
    </div>
</template>


<script>
export default {
  name: 'rankTable',
  data() {
      return {
        fields: [
          {
            key : 'rating',
            sortable: true
          },
          {
            key : 'member',
            sortable: true
          },
          {
            key : 'honor point',
            sortable: true
          },
          {
            key : 'winer point',
            sortable: true
          }
        ],
        items: [
          { rating: 1, member: 'user123', 'honor point': '100', 'winer point': '200' },
          { rating: 2, member: 'user456', 'honor point': '250', 'winer point': '600' },
          { rating: 3, member: 'user789', 'honor point': '400', 'winer point': '250' },
          { rating: 4, member: 'user101', 'honor point': '500', 'winer point': '0' },
        ],
        rankTitle : ''
      }
  }
}
</script>
