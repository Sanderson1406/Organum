<template>
  <div class="all">
    <h2 class="title">{{ book.title }}</h2>
    <p>Autor(a): {{ book.author }}</p>
    <p>Páginas: {{ book.pages }}</p>
    <div class="progress-container">
      <div class="progress-bar" :style="{ width: progressWidth }"></div>
    </div>
    <div class="statistics">
      <p>{{ pgl }} páginas lidas de {{ book.pages }}</p>
    </div>
    <div>
      <button type="button" class="other" @click="update">Atualizar</button>
      <button type="button" class="other" @click="addCitations">Adicionar Citações</button>
    </div>
    <button type="button" class="delete" @click="deleteBook">Excluir</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props: {
    book: {
      type: Object,
      required: true,
      default: () => ({
        title: '',
        author: '',
        pages: 0
      })
    }
  },
  data() {
    return {
      pgl: 0,
      progressWidth: '0%'
    }
  },
  mounted() {
    this.getProgress()
  },
  methods: {
    update() {
      this.$router.push({ name: 'atualizar', params: { title: this.book.title } })
    },
    addCitations() {
      this.$router.push({ name: 'citacao', params: { title: this.book.title } })
    },
    getProgress() {
      axios
        .get(`http://localhost:8084/lectio/book/progress?title=${this.book.title}`)
        .then((response) => {
          this.pgl = response.data
          this.progressWidth = `${(this.pgl / this.book.pages) * 100}%`
        })
        .catch((error) => {
          console.error('Erro ao obter o progresso do livro:', error)
        })
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
  height: auto;
}

.statistcs p {
  font-size: 15px;
  margin: 1%;
  text-align: left;
  font-weight: 300;
  text-align: center; /* Esta regra estava aplicada erroneamente ao seletor errado */
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

.other {
  background-color: #2e3d50;
  border-radius: 20px;
  margin: 3% auto;
}

.delete {
  background-color: #ff0000;
  border-radius: 20px;
  margin: 2% auto;
  padding: auto;
  width: 40%;
  display: block;
}

.progress-container {
  width: 100%;
  height: 25px;
  background-color: #f0f0f0;
  border-radius: 20px;
  overflow: hidden;
  margin-top: 15px;
  margin-bottom: 15px;
}

.progress-bar {
  height: 100%;
  background-color: #c4762b;
  border-radius: 20px;
  transition: width 0.3s ease;
}
</style>
