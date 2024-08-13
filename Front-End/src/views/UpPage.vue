<template>
  <div class="container">
    <div class="card">
      <form @submit.prevent="updatePages">
        <label for="numberInput">Páginas Lidas:</label>
        <input
          type="number"
          v-model="progress.pgl"
          id="numberInput"
          placeholder="Digite um número"
        />
        <button type="submit">Atualizar</button>
      </form>
      <p>
        *Páginas lidas se refere à quantidade de páginas lidas em uma sessão de leitura, não o total
      </p>
    </div>
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
      },
      progress: {
        title: this.$route.params.title,
        pgl: null
      }
    }
  },
  created() {
    // Configura um interceptor para adicionar o token de autenticação a todas as requisições
    axios.interceptors.request.use(config => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });
  },
  methods: {
    updatePages() {
      if (!this.progress.pgl) {
        alert('Por favor, preencha o número de páginas lidas.');
        return;
      }
      axios
        .put('http://localhost:8084/lectio/book/progress', this.progress)
        .then((response) => {
          console.log(response);
          this.goBack();
          alert('Progresso atualizado');
        })
        .catch((error) => {
          console.error('Erro ao fazer solicitação', error);
          alert('Erro ao atualizar');
        });
    },
    updateStatus() {
      if (!this.update.newtag) {
        alert('Por favor, selecione o novo status do livro.');
        return;
      }
      axios
        .post('http://localhost:8084/lectio/book/updattag', this.update)
        .then((response) => {
          console.log(response);
          this.goBack();
          alert('Status do livro atualizado com sucesso');
        })
        .catch((error) => {
          console.error('Erro ao fazer a solicitação:', error);
          alert('Erro ao atualizar status do livro!');
        });
    },
    goBack() {
      this.$router.push('/lectio');
    }
  }
}
</script>
S

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
