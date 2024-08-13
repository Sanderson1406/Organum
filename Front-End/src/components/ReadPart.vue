<template>
  <div class="card-container">
    <div class="all">
      <h2 class="title">{{ book.title }}</h2>
      <p>Autor(a): {{ book.author }}</p>
      <div class="button-container">
        <button type="button" class="other" @click="update">Atualizar Status</button>
        <button type="button" class="other" @click="addReview">Escrever Review</button>
      </div>
      <button type="button" class="delete" @click="deleteBook">Excluir</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    book: {
      type: Object,
      required: true
    }
  },
  methods: {
    update() {
      this.$router.push({ name: 'status', params: { title: this.book.title } })
    },
    addReview() {
      this.$router.push({ name: 'review', params: { title: this.book.title } })
    },
    deleteBook() {
      console.log('aqui chegou')
      this.$router.push({ name: 'apagar', params: { title: this.book.title } })
    }
  }
}
</script>

<style scoped>
.card-container {
  width: 250px;
  font-family: 'Roboto Condensed', sans-serif;
}

.all {
  background-color: #feeecd;
  padding: 20px;
  border-radius: 20px;
  margin-bottom: 10px;
  width: 100%;
  box-sizing: border-box;
}

.title {
  font-size: 27px;
  font-weight: 400;
  margin: 1%;
}

p {
  font-size: 18px;
  margin-top: 3%;
  text-align: left;
  font-weight: 300;
  color: black;
}

.button-container {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
}

.other,
.delete {
  background-color: #2e3d50;
  border-radius: 20px;
  margin: 3% auto;
  width: 45%;
}

.delete {
  background-color: #ff0000;
  border-radius: 20px;
  padding: auto;
  display: block;
}
</style>
