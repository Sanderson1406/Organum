<template>
    <div>
      <button @click="logout">Sair</button>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    methods: {
      async logout() {
        try {
          const response = await axios.post('http://localhost:8084/logout', null, {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('token')}`, 
            },
          });
          localStorage.removeItem('token');
          this.$router.push('/');
          console.log(response.data); 
        } catch (error) {
          console.error('Erro ao fazer logout:', error)
        }
      }
    }
  }
  </script>
  
  <style scoped>
  button {
    background-color: #ff0000;
    width: 100px;
    font-size: 20px;
    padding: auto;
    margin-bottom: 25px;
  }
</style>
  