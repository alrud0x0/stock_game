<template>
  <div id="gameTable">
    <div align="middle" class="common-margin">
      <h3>{{ gameTitle }}</h3>
    </div>
    <div>
      <b-table striped hover
        :items="items"
        :fields="fields">

        <template v-slot:cell(comparison)="row">
          <span :style="{color : [row.value.rate < 0 ? 'blue' : 'red']}">
            {{ row.value.rate }}% <br/>
            {{ row.value.difference }}
          </span>
        </template>


      </b-table>
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
            key : 'stock',
            sortable: true,
            label : '종목명',
            class: 'align-middle'
          },
          {
            key : 'price',
            sortable: true,
            label : '현재값',
            class: 'align-middle'
          },
          {
            key : 'comparison',
            sortable: true,
            label : '대비(차액)',
            class: 'align-middle'
          },
          {
            key : 'vote',
            sortable: true,
            label : '투표수',
            class: 'align-middle'
          }
        ],
        items: [
          { stock: '삼성전자', price: '60000', comparison: { rate : '10', difference : '3000'}, vote: '3' },
          { stock: '삼성SDS', price: '198000', comparison:  { rate : '10', difference : '198000'} , vote: '7' },
          { stock: '삼성SDI', price: '330000', comparison:  { rate : '-3', difference : '8000'} , vote: '5' },
          { stock: '삼성전기', price: '110000', comparison:  { rate : '-5', difference : '3000'}, vote: '1' }
        ],
        gameTitle : '20.02.28 1차'
      }
  }
}
</script>
