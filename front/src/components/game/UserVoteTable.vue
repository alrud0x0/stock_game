<template>
  <div id="userVoteTable">
    <b-button v-b-toggle.collapse-1 variant="info">내 투표 정보 보기</b-button>
    <b-collapse id="collapse-1" class="mt-2">
      <b-card>
        <p class="card-text">내가 투표한 주식 랭킹</p>
        <b-table noCollapse
          :items="items"
          :fields="fields">

          <template v-slot:cell(comparison)="row">
            <span :style="{color : [row.value.rate < 0 ? 'blue' : 'red']}">
              {{ row.value.rate }}% <br/>
              {{ row.value.difference}}
            </span>
          </template>

        </b-table>
      </b-card>
    </b-collapse>
  </div>
</template>

<script>
export default {
  name: 'userVoteTable',
  data() {
    return {
      fields : [
        {
          key : 'rank',
          label: '순위',
          class: 'align-middle'
        },
        {
          key : 'stock',
          label : '종목명',
          class: 'align-middle'
        },
        {
          key : 'comparison',
          label : '대비(차액)',
          class: 'align-middle text-right'
        }
      ],
      items : [
        {
            rank : 3,
            stock: '삼성전기',
            comparison:  {
              rate : '5',
              difference : '13000'
            }
        },
        {
            rank : 7,
            stock: '삼성전자',
            comparison:  {
              rate : '2',
              difference : '3000'
            }
        }
      ]
    }
  }
}
</script>
