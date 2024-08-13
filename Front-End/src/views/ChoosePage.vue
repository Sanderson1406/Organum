<template>
  <div class="all">
    <div class="lectio">
      <h1>Lectio</h1>
      <p>
        Com o Lectio, você pode: Adicionar livros à sua lista pessoal. Categorizar os livros com
        tags como ‘Lidos’, ‘Quero Ler’ ou ‘Estou Lendo’. Atualizar o status de leitura. Escrever
        resenhas. Registrar o número de páginas lidas diariamente. Acompanhar suas estatísticas de
        leitura, como tempo gasto com os livros. Salvar citações favoritas para inspiração e
        compartilhamento.
      </p>
      <div class="but">
        <button type="submit" @click="lectio">Organizar</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.all {
  display: flex;
  height: 100vh;
  width: 100vw;
}

.lectio {
  background-color: #2e3d50;
  width: 100%;
  padding: 10px;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.lectio h1 {
  font-size: 80px;
}

.lectio p {
  color: #fff;
  margin-left: 15%;
  margin-right: 15%;
  margin-top: 5%;
  margin-bottom: 5%;
  text-align: center;
  font-size: 25px;
  font-weight: 200;
}

.but {
  padding: 10px;
  justify-content: center;
  display: flex;
  align-items: center;
  width: 60%;
}

button {
  padding: 2%;
  max-width: 60%;
  width: 30%;
  margin: auto;
  font-size: 25px;
}
</style>

<script>
export default {
  methods: {
    lectio() {
      this.$router.push('/lectio')
    }
  }
}
</script>
