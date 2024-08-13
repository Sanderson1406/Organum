<template>
  <div class="all">
    <h2 class="title">{{ book.title }}</h2>
    <p>Autor(a): {{ book.author }}</p>
    <div class="button-container">
      <button type="submit" class="other" @click="updateStatus">Atualizar Status</button>
      <button type="submit" class="delete" @click="deleteBook">Excluir</button>
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
    updateStatus() {
      this.$router.push({ name: 'status', params: { title: this.book.title } })
    },
    deleteBook() {
      this.$router.push({ name: 'apagar', params: { title: this.book.title } })
      console.log('aqui chegou')
    }
  }
}
</script>

<style scoped>
.all {
  background-color: #feeecd;
  padding: 20px;
  border-radius: 20px;
  font-family: 'Roboto Condensed', sans-serif;
  width: 250px;
  height: auto;
  flex-direction: column;
  justify-content: space-between;
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

.statistcs {
  font-size: 15px;
  margin: 1%;
  text-align: left;
  font-weight: 300;
  text-align: center;
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
  width: 45%;
  margin: 0;
}

.delete {
  background-color: #ff0000;
}
</style>
