<template>
  <div class="container">
    <h1>Organum</h1>
    <div class="form-container">
      <h2>Cadastro</h2>
      <form @submit.prevent="createuser">
        <div class="form-group">
          <label for="email">E-mail:</label>
          <input type="email" id="email" v-model="newuser.email" required />

          <label for="nome">Nome:</label>
          <input type="text" id="nome" v-model="newuser.name" required />

          <label for="sobrenome">Sobrenome:</label>
          <input type="text" id="sobrenome" v-model="newuser.sobrenome" required />

          <label for="senha">Senha:</label>
          <input type="password" id="senha" v-model="newuser.password" required />

          <label for="confirmarsenha">Confirmar Senha:</label>
          <input type="password" id="confirmarsenha" v-model="confirmarsenha" required />

          <button type="submit">Cadastrar</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      newuser: {
        email: '',
        name: '',
        sobrenome: '',
        password: ''
      },
      confirmarsenha: ''
    }
  },
  methods: {
    createuser() {
      if (this.newuser.password !== this.confirmarsenha) {
        alert('As senhas não coincidem!')
        return
      }

      axios
        .post('http://localhost:8084/createuser', this.newuser)
        .then((response) => {
          console.log(response)
          this.choose()
          alert('Usuário cadastrado com sucesso')
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error)
          alert('Erro ao fazer a solicitação')
        })
    },
    choose() {
      this.$router.push('/')
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&family=Vibes&display=swap');

body {
  font-family: 'Roboto Condensed', sans-serif;
}

.container {
  margin-top: 2%;
}

.form-container {
  max-width: 350px;
  margin: 45px auto;
  padding: 20px 40px;
  border-radius: 20px;
  background-color: #feeecd;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.form-container h2 {
  text-align: center;
  font-size: 35px;
  color: black;
  margin-bottom: 20px;
  font-style: normal;
}
.form-group {
  margin-bottom: 12px;
}

.form-group label {
  display: block;
  font-size: 16px;
  color: #333;
  margin-bottom: 10px;
  margin-top: 5px;
  text-align: center;
}

.form-group input {
  border-radius: 20px;
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  font-size: 16px;
}

button {
  font-family: 'Roboto Condensed', sans-serif;
  width: 100%;
  padding: 10px;
  background-color: #515151;
  color: white;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-size: 16px;
  margin-top: 25px;
}

button:hover {
  background-color: #333333;
}

h1 {
  font-family: 'Vibes', cursive;
  font-size: 48px;
  color: white;
  font-weight: 500;
  text-align: center;
  margin-top: 10px;
  margin-bottom: -20px;
}
</style>
