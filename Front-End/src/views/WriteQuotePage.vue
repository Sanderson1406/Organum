<template>
  <div class="container">
    <div class="card">
      <form @submit.prevent="writequote">
        <label for="quoteInput">Escreva sua Citação:</label>
        <input type="text" v-model="addquote.quote" id="quoteInput" />
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
      addquote: {
        title: this.$route.params.title,
        quote: ''
      }
    }
  },
  methods: {
    writequote() {
      axios
        .post('http://localhost:8084/lectio/quote', this.addquote)
        .then((response) => {
          console.log(response)
          this.goback()
          alert('Citação Salva')
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error)
          alert('Erro ao adicionar citação!')
        })
    },
    goback() {
      this.$router.push('/lectio')
    }
  }
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  gap: 20px;
  padding: 20px;
  box-sizing: border-box;
}

.card {
  background-color: #feeecd;
  border-radius: 20px;
  padding: 20px;
  min-width: 300px;
  max-width: 30%;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-sizing: border-box;
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
  box-sizing: border-box;
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
  max-width: 100%;
  margin-top: 2%;
  margin-bottom: 2%;
}
</style>
