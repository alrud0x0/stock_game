<template>
  <div id="gameTable">
    <div align="middle" class="common-margin">
      <h3>{{ gameTitle }}</h3>
    </div>
    <div>
      <b-table striped hover :items="items" :fields="fields"></b-table>
    </div>
    </div>
</template>


<script>
export default {
  name: 'gameTable',
  data() {
      return {
        fields: [
          {
            key : 'Stock',
            sortable: true
          },
          {
            key : 'Last Price',
            sortable: true
          },
          {
            key : 'Rate',
            sortable: true
          },
          {
            key : 'Vote',
            sortable: true
          }
        ],
        items: [
          { Stock: 40, 'Last Price': '2000', Rate: '10%', Vote: 'Macdonald' },
          { Stock: 20, 'Last Price': '56000', Rate: '30%', Vote: 'Macdonald' },
          { Stock: 10, 'Last Price': '300', Rate: '-10%', Vote: 'Macdonald' },
          { Stock: 41, 'Last Price': '250', Rate: '-50%', Vote: 'Macdonald' }
        ],
        gameTitle : '20.02.28 1차'
      }
  }
}
</script>
