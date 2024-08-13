<template>
  <div class="container">
    <div class="card">
      <form @submit.prevent="writeReview">
        <label for="reviewInput">Escreva seu Review:</label>
        <input type="text" v-model="addreview.opnion" id="reviewInput" />
        <label for="scoreInput">Avaliação (0 a 5):</label>
        <input type="number" step="0.1" min="0" max="5" v-model="addreview.score" id="scoreInput" />
        <button type="submit">Salvar</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      addreview: {
        title: this.$route.params.title,
        opnion: '',
        score: ''
      }
    }
  },
  methods: {
    writeReview() {
      if (!this.addreview.opnion || !this.addreview.score) {
        alert('Por favor, preencha todos os campos!')
        return
      }
      axios
        .post('http://localhost:8084/lectio/review', this.addreview)
        .then((response) => {
          console.log(response)
          this.goBack()
          alert('Review salvo com sucesso')
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error)
          alert('Erro ao adicionar review!')
        })
    },
    goBack() {
      this.$router.push('/lectio')
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 20px;
}

.card {
  background-color: #feeecd;
  border-radius: 20px;
  padding: 20px;
  max-width: 30%;
  width: 100%;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 20px;
  margin-bottom: 10px;
}

button {
  background: #515151;
  border: none;
  border-radius: 20px;
  color: white;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  outline: none;
  transition: opacity 0.3s;
  width: 100%;
}

button:hover {
  opacity: 0.8;
}
</style>

<style scoped>
.container {
  display: flex;
  flex-direction: row;
  gap: 20px;
  max-width: 800px;
  margin: 0 auto;
  flex-wrap: wrap;
}

.card {
  flex: 1;
  padding: 20px;
  border-radius: 20px;
  background-color: #feeecd;
  min-width: 300px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

p {
  color: black;
  font-size: 15px;
  margin: 2%;
  text-align: center;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  text-align: center;
  font-size: 20px;
}

input,
select {
  width: 100%;
  padding: 15px;
  border: 1px solid #ccc;
  border-radius: 20px;
  margin-top: 2%;
  margin-bottom: 2%;
  text-align: center;
}

button {
  background: #515151;
  border: none;
  border-radius: 20px;
  color: white;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  outline: none;
  transition: opacity 0.3s;
  margin-top: 2%;
  margin-bottom: 2%;
}
</style>
