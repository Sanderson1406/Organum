<template>
  <div>
    <h1 class="lectio-title">Lectio</h1>
    <div class="form-container">
      <h2>Cadastro de Livros</h2>
      <form class="book-form" @submit.prevent="addBook">
        <label for="title">Título</label>
        <input type="text" id="title" v-model="newBook.title" required />
        <label for="author">Autor(a)</label>
        <input
          type="text"
          id="author"
          v-model="newBook.author"
          @keydown="validateAuthorKeydown($event)"
          required
        />
        <div class="select-container">
          <select id="tag" v-model="newBook.tag" required>
            <option value="">Selecione o status</option>
            <option value="e">Lido</option>
            <option value="l">Lendo</option>
            <option value="v">Vou Ler</option>
          </select>
        </div>
        <label for="pages">Páginas Totais</label>
        <input type="number" id="pages" v-model="newBook.pages" required />
        <button type="submit">Adicionar Livro</button>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      newBook: {
        title: '',
        author: '',
        tag: '',
        pages: 0
      }
    }
  },
  methods: {
    addBook() {
      const token = localStorage.getItem('token') // Obtém o token do localStorage
      const config = {
        headers: {
          Authorization: `Bearer ${token}` // Adiciona o token no cabeçalho da requisição
        }
      }

      axios
        .post('http://localhost:8084/lectio/book', this.newBook, config) // Passa config para incluir o token
        .then((response) => {
          console.log(response)
          this.goback()
          alert('Livro Cadastrado')
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error)
          alert('Erro ao adicionar livro!')
        })
    },
    validateAuthorKeydown(event) {
      const regex = /^[a-zA-Z\sçÇ]*$/
      if (!regex.test(event.key) && !this.isFunctionalKey(event)) {
        event.preventDefault()
      }
    },
    isFunctionalKey(event) {
      return event.key === 'Tab' || event.key === 'Backspace' || event.key === ' '
    },
    goback() {
      this.$router.push('/lectio')
    }
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&family=Vibes&display=swap');

.lectio-title {
  font-family: 'Vibes', cursive;
  font-size: 60px;
  color: white;
  font-weight: 500;
  text-align: center;
  margin-top: 10px;
  margin-bottom: -20px;
}

.all {
  margin-top: 3%;
}

.select-container {
  display: flex;
  justify-content: center;
  margin: 20px;
}

select {
  background-color: #515151;
  color: white;
  border-radius: 50px;
  align-content: center;
  font-size: 18px;
  width: 200px;
  height: 40px;
  text-align: center;
}

.form-container {
  max-width: 350px;
  margin: 45px auto;
  padding: 20px 40px;
  border-radius: 20px;
  background-color: #feeecd;
}

.book-form label,
p {
  display: block;
  font-size: 16px;
  color: #333;
  margin-bottom: 10px;
  margin-top: 5px;
  text-align: center;
}

.book-form {
  margin-bottom: 12px;
}

.form-container h2 {
  text-align: center;
  font-size: 35px;
  color: black;
  margin-bottom: 20px;
  font-style: normal;
}

input[type='text'],
input[type='number'] {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 20px;
  width: 100%;
  font-size: 16px;
}

button {
  font-family: 'Roboto Condensed', sans-serif;
  width: 100%;
  padding: 10px;
  background-color: #515151;
  color: white;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 18px;
  margin-top: 25px;
}

button:hover {
  background-color: #333333;
}
</style>
