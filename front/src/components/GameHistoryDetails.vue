<template>
  <div class="outline-margin">
    <gameTable :isVotable="isVotable" :gameTitle="gameTitle" :fields="fields" :items="items" />
  </div>
</template>


<script>
import GameTable from './game/GameTable.vue'

export default {
  name: 'gameHistoryDetails',
  components: {
    GameTable
  },
  data() {
      return {
        fields: [
          {
            key : 'stock',
            label : '종목명',
            class: 'align-middle'
          },
          {
            key : 'price',
            label : '시초가',
            class: 'align-middle'
          },
          {
            key : 'comparison',
            sortable: true,
            label : '대비(차액)',
            class: 'align-middle text-right',
          },
          {
            key : 'vote',
            label : '투표수',
            class: 'align-middle text-center'
          }
        ],
        items: [
          {
              stock: '삼성전자',
              price: '48000',
              comparison: {
                  rate : '7',
                  difference : '2000'
              },
              vote: {
                  count : '3',
                  items : [  { userId : 'user123', reason : '그냥' }, { userId : 'user456', reason : '그냥요' }, { userId : 'user789', reason : '그냥이요' }]
              }
          },
          {
              stock: '삼성SDS',
              price: '198000',
              comparison:  {
                rate : '10',
                difference : '1980'
              },
              vote: {
                count : '1',
                items : [  { userId : 'user123', reason : '내맘' }]
              }
          },
          {
              stock: '삼성SDI',
              price: '330000',
              comparison:  {
                rate : '-3',
                difference : '8000'
              },
              vote: {
                count : '2',
                items : [  { userId : 'user123', reason : 'ㅇㅇ' }, { userId : 'user123', reason : 'ㅇㅇㅇ' }]
              }
          },
          {   stock: '삼성전기',
              price: '110000',
              comparison:  {
                rate : '-5',
                difference : '3000'
              },
              vote: {
                count : '1',
                items : [  { userId : 'user123', reason : 'ㅇㅇ' }, { userId : 'user123', reason : 'ㅇㅇㅇ' }]
              }
          }
        ],
        gameTitle : '2020.02.13',
        isVotable : false
      }
  }
}
</script>
