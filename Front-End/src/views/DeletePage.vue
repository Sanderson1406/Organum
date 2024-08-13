<template>
  <div class="container">
    <div class="card">
      <form @submit.prevent="deleteBook">
        <p>Tem certeza que deseja excluir o livro?</p>
        <h3>"{{ deletet.title }}"</h3>
        <p>A exclusão de um livro também excluirá todas as suas dependências.</p>
        <button type="submit" class="delete">Excluir</button>
      </form>
      <button @click="goBack" class="back-button">Voltar</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      deletet: {
        title: this.$route.params.title
      }
    }
  },
  methods: {
    deleteBook() {
      const token = localStorage.getItem('token')
      const config = {
        headers: {
          Authorization: `Bearer ${token}`
        },
        data: { title: this.deletet.title }
      }
      axios
        .delete('http://localhost:8084/lectio/book', config)
        .then((response) => {
          console.log(response)
          this.goBack()
          alert('Livro deletado com sucesso.')
        })
        .catch((error) => {
          console.log('Erro ao fazer solicitação', error)
          alert('Erro ao deletar o livro.')
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
  font-size: 20px;
  margin: 5%;
  text-align: center;
  font-weight: 300;
}

h3 {
  font-size: 25px;
  font-weight: 500;
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
  width: 100%;
}

.delete {
  background-color: #ff0000;
}

.back-button {
  background-color: #515151;
}
</style>
