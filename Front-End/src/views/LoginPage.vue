<template>
  <div class="login-page">
    <div class="all">
      <div class="welcome">
        <h3>Bem vindo(a) ao</h3>
        <h1>Organum</h1>
      </div>
      <div class="login-container">
        <h2>Login</h2>
        <form @submit.prevent="login">
          <div class="form-group">
            <label for="email">Email:</label>
            <input type="email" id="email" v-model="getuser.email" required aria-label="Email" />
          </div>
          <div class="form-group">
            <label for="password">Senha:</label>
            <input type="password" id="password" v-model="getuser.password" required aria-label="Senha" />
          </div>
          <button type="submit">Entrar</button>
        </form>
      </div>
      <div class="other">
        <router-link to="/registrar" class="botao">Não tenho cadastro</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      getuser: {
        email: '',
        password: ''
      }
    };
  },
  methods: {
    async login() {
      try {
        const response = await axios.post('http://localhost:8084/login', this.getuser);
        const token = response.data.token;

        if (token) {
          localStorage.setItem('token', token);
          this.$router.push('/choose'); 
        } else {
          alert('Erro ao obter o token');
        }
      } catch (error) {
        console.error('Erro ao fazer a solicitação:', error);
        alert('Senha ou email incorretos');
      }
    }
  }
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&family=Vibes&display=swap');

body {
  font-family: 'Roboto Condensed', sans-serif;
  background-color: #515151;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.all {
  text-align: center;
  max-width: 350px;
  margin: 0 auto;
  margin-top: 5%;
}

.welcome {
  margin-bottom: 30px;
}

.welcome h3 {
  font-size: 30px;
  color: white;
  font-weight: 300;
}

.welcome h1 {
  font-family: 'Vibes', cursive;
  font-size: 48px;
  color: white;
  font-weight: 500;
}

.login-container {
  padding: 20px 40px;
  border-radius: 20px;
  background-color: #feeecd;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

.login-container h2 {
  font-size: 35px;
  color: black;
  margin-bottom: 20px;
  font-weight: 400;
}

.other {
  margin-top: 20px;
}

.botao {
  background: #feeecd;
  border: none;
  border-radius: 20px;
  color: black;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
  outline: none;
  transition: opacity 0.3s;
}
</style>
