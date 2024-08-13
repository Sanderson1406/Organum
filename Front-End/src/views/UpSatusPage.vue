<template>
  <div class="container">
    <div class="card">
      <form @submit.prevent="updateStatus">
        <label for="statusInput">Novo Status do livro:</label>
        <select id="tag" v-model="update.newtag">
          <option value="">Escolha a nova opção</option>
          <option value="e">Lido</option>
          <option value="v">Vou ler</option>
          <option value="l">Lendo</option>
        </select>
        <button type="submit">Atualizar</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      update: {
        title: this.$route.params.title,
        newtag: ''
      }
    }
  },
  methods: {
    updateStatus() {
      if (!this.update.newtag) {
        alert('Por favor, selecione o novo status do livro.')
        return
      }
      axios
        .post('http://localhost:8084/lectio/book/updattag', this.update)
        .then((response) => {
          console.log(response)
          this.goback()
          alert('Status do livro atualizado com sucesso')
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error)
          alert('Erro ao atualizar status do livro!')
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
  justify-content: center;
  align-items: center;
  min-height: 100vh;
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
