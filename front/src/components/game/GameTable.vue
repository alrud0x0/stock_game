<template>
  <div id="gameTable">
    <div align="middle" class="common-margin">
      <h3>{{ gameTitle }}</h3>
    </div>
    <div align="middle" class="common-margin">
      <userVoteTable />
    </div>
    <div>
      <b-table noCollapse hover
        :items="items"
        :fields="fields"
        @row-clicked="item=>$set(item, '_showDetails', !item._showDetails)">

        <template v-slot:cell(price)="row">
            {{ row.value | currencyFormat}}
        </template>

        <template v-slot:cell(comparison)="row">
          <span :style="{color : [row.value.rate < 0 ? 'blue' : 'red']}">
            {{ row.value.rate }}% <br/>
            {{ row.value.difference | currencyFormat}}
          </span>
        </template>

        <template v-slot:cell(vote)="row">
          {{ row.value.count }}
          <!--<b-button size="sm" @click="row.toggleDetails" class="mb-2" variant="outline-info">
            <b-icon icon="chevron-down" v-if="row.detailsShowing"></b-icon>
            <b-icon icon="chevron-up" v-if="!row.detailsShowing"></b-icon>
          </b-button>-->
        </template>

        <!-- voting info -->
        <template v-slot:row-details="row">
          <voteInfo :items="row.item.vote.items" />
        </template>

      </b-table>
    </div>
    </div>
</template>


<script>
import VoteInfo from './VoteInfo.vue'
import UserVoteTable from './UserVoteTable.vue'

export default {
  name: 'gameTable',
  components: {
    VoteInfo,
    UserVoteTable
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
            label : '현재값',
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
              price: '60000',
              comparison: {
                  rate : '10',
                  difference : '3000'
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
        gameTitle : '20.02.28 1차'
      }
  },
  filters : {
    currencyFormat: function(value) {
      var num = new Number(value);
      return num.toFixed(0).replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    }
  }
}
</script>
