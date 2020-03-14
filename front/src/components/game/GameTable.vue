<template>
  <div id="gameTable">
    <div align="middle" >
      <h3>{{ gameTitle }}</h3>
    </div>
    <div v-if="isVotable" align="middle" class="common-margin">
      <voteModal />
    </div>
    <div v-else align="middle" class="common-margin">
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
import VoteModal from './VoteModal.vue'
import VoteInfo from './VoteInfo.vue'
import UserVoteTable from './UserVoteTable.vue'

export default {
  name: 'gameTable',
  components: {
    VoteModal,
    VoteInfo,
    UserVoteTable
  },
  props: {
    isVotable: {
      type: Boolean,
      default: false
    },
    gameTitle: {
      type: String,
      default: "",
    },
    items: {
      type: Array,
      default: () => ({})
    },
    fields: {
      type: Array,
      default: () => ({})
    }
  },
  data() {
      return {

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
