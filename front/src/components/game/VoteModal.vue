<template>
  <div>
    <b-button block v-b-modal.vote-modal variant="info">게임 참여하기 <b-icon icon="arrow-bar-right"></b-icon></b-button>

    <b-modal
      id="vote-modal"
      ref="modal"
      title="한주간 가장 많이 상승할 것 같은 종목에 투표하세요."
      centered
      @show="resetModal"
      @hidden="resetModal"
      @ok="handleOk"
    >
      <b-form @submit.stop.prevent="handleSubmit">
        <b-form-group
          id="input-group-1"
          label="종목:"
          label-for="stock-name"
          description="3개까지 투표가능"
        >
            <b-form-input
              id="stock-name"
              v-model="stock"
              required
              placeholder="주식명을 입력하세요."
            ></b-form-input>

        </b-form-group>

        <b-form-group
          id="input-group-1"
          label="이유:"
          label-for="vote-reason"
          description="option 입니다."
        >

          <b-form-input
            id="vote-reason"
            v-model="reason"
            placeholder="이 종목을 고른 이유?"
          ></b-form-input>

        </b-form-group>

      </b-form>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: 'voteModal',
  data() {
     return {
       voteInfo: {
          stock: '',
          reason: ''
        },
        stock: '',
        reason: '',
        stockList: [
          { id: '123', name: '삼성전자', price: '66500' },
          { id: '456', name: '삼성SDS', price: '198000' },
          { id: '789', name: 'SK하이닉스', price: '88000' },
          { id: '111', name: 'SK네트웍스', price: '5500' }
        ]
     }
   },
   methods: {
     resetModal() {
             this.stock = ''
             this.reason = ''
      },
      handleOk(bvModalEvt) {
        // Prevent modal from closing
        bvModalEvt.preventDefault()
        // Trigger submit handler
        this.handleSubmit()
      },
      handleSubmit() {
        alert(this.stock + ", " + this.reason)
        this.$nextTick(() => {
          this.$bvModal.hide('vote-modal')
        })
      }
   }
}
</script>
