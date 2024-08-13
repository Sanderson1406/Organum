<template>
  <div>
    <div class="head">
      <div class="logo-container">
         <h1>Organum</h1>
      </div>
      <div class="logout-button">
        <Logout></Logout>
      </div>
    </div>
    <div class="button-container">
      <button type="button" @click="registerBook" class="register-button">
        Cadastrar novo Livro
      </button>
    </div>

    <!-- Seção para Livros em Leitura -->
    <div class="head-other">
      <h1>Livros em leitura</h1>
    </div>
    <div v-if="booksLLoading">
      <Load class="load"></Load>
    </div>
    <div v-else-if="booksL.length" class="books-container" id="reading-books">
      <ReadingPart
        v-for="(book, index) in booksL"
        :key="index"
        :book="book"
        class="book-item"
      ></ReadingPart>
    </div>
    <div v-else>
      <Nothing class="nt"></Nothing>
    </div>

    <!-- Seção para Livros que Você Pretende Ler -->
    <div class="head-other">
      <h1>Livros que você pretende ler</h1>
    </div>
    <div v-if="booksVLoading">
      <Load class="load"></Load>
    </div>
    <div v-else-if="booksV.length" class="books-container" id="to-read-books">
      <WillPart
        v-for="(book, index) in booksV"
        :key="index"
        :book="book"
        class="book-item"
      ></WillPart>
    </div>
    <div v-else>
      <Nothing class="nt"></Nothing>
    </div>

    <!-- Seção para Livros que Você Leu -->
    <div class="head-other">
      <h1>Livros que você pretende ler</h1>
    </div>
    <div v-if="booksELoading">
      <Load class="load"></Load>
    </div>
    <div v-else-if="booksE.length" class="books-container" id="to-read-books">
      <ReadPart
        v-for="(book, index) in booksE"
        :key="index"
        :book="book"
        class="book-item"
      ></ReadPart>
    </div>
    <div v-else>
      <Nothing class="nt"></Nothing>
    </div>

    <!-- Seção para Resenhas -->
    <div>
      <div class="head-other">
        <h1>Resenhas</h1>
      </div>
      <div v-if="reviewsLoading">
        <Load></Load>
      </div>
      <div v-else-if="reviews.length" class="reviews-container">
        <ReviewPart
          v-for="(review, index) in reviews"
          :key="index"
          :review="review"
          class="review-item"
        ></ReviewPart>
      </div>
      <div v-else>
        <Nothing class="nt"></Nothing>
      </div>
    </div>

    <!-- Seção para Citações Salvas -->
    <div class="head-other">
      <h1>Citações Salvas</h1>
    </div>
    <div v-if="quotesLoading">
      <Load></Load>
    </div>
    <div v-else-if="quotes.length" class="quote-container">
      <QuotePart
        v-for="(quote, index) in quotes"
        :key="index"
        :quote="quote"
        class="quote-item"
      ></QuotePart>
    </div>
    <div v-else>
      <Nothing class="nt"></Nothing>
    </div>
  </div>
</template>

<script>
import ReadingPart from '../components/ReadingPart.vue'
import WillPart from '../components/WillPart.vue'
import ReadPart from '../components/ReadPart.vue'
import QuotePart from '../components/QuotePart.vue'
import ReviewPart from '../components/ReviewPart.vue'
import Nothing from '@/components/NothingPart.vue'
import Load from '@/components/LoadPart.vue'
import Logout from '@/components/LogoutPart.vue'
import axios from 'axios'

export default {
  components: {
    ReadingPart,
    WillPart,
    ReadPart,
    QuotePart,
    ReviewPart,
    Nothing,
    Load,
    Logout
  },
  data() {
    return {
      booksL: [],
      booksV: [],
      booksE: [],
      quotes: [],
      reviews: [],
      booksLLoading: true,
      booksVLoading: true,
      booksELoading: true,
      quotesLoading: true,
      reviewsLoading: true
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    registerBook() {
      this.$router.push('/registerbook')
    },
    async getData() {
      try {
        await Promise.all([
          this.getBooksL(),
          this.getBooksV(),
          this.getBooksE(),
          this.getQuotes(),
          this.getReviews()
        ])
        console.log('Todas as interações no servidor foram concluídas!')
      } catch (error) {
        this.handleError(error)
      }
    },
    async getBooksL() {
      try {
        const response = await axios.get('http://localhost:8084/lectio/book/l?tag=l', { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
        this.booksL = response.data || []
      } catch (error) {
        this.handleError(error)
      } finally {
        this.booksLLoading = false
      }
      console.log('Livros em leitura:', this.booksL)
    },
    async getBooksV() {
      try {
        const response = await axios.get('http://localhost:8084/lectio/book/l?tag=v', { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
        this.booksV = response.data || []
      } catch (error) {
        this.handleError(error)
      } finally {
        this.booksVLoading = false
      }
      console.log('Livros que você pretende ler:', this.booksV)
    },
    async getBooksE() {
      try {
        const response = await axios.get('http://localhost:8084/lectio/book/l?tag=e', { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
        this.booksE = response.data || []
      } catch (error) {
        this.handleError(error)
      } finally {
        this.booksELoading = false
      }
      console.log('Livros que você leu:', this.booksE)
    },
    async getQuotes() {
      try {
        const response = await axios.get('http://localhost:8084/lectio/quote', { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
        this.quotes = response.data || []
      } catch (error) {
        this.handleError(error)
      } finally {
        this.quotesLoading = false
      }
      console.log('Citações Salvas:', this.quotes)
    },
    async getReviews() {
      try {
        const response = await axios.get('http://localhost:8084/lectio/review', { headers: { Authorization: `Bearer ${localStorage.getItem('token')}` } })
        this.reviews = response.data || []
      } catch (error) {
        this.handleError(error)
      } finally {
        this.reviewsLoading = false
      }
      console.log('Resenhas:', this.reviews)
    },
    handleError(error) {
      console.error('Erro ao fazer a solicitação:', error)
      alert('Erro ao obter os dados do servidor!')
    }
  }
}
</script>

<style scoped>
.head {
  background-color: #2e3d50;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 20px;
  width: 100%;
  height: 70px;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
}

.logo-container {
  flex: 1; 
  display: flex;
  align-items: center;
}

.head h1 {
  margin: 0;
  font-size: 40px;
  text-align: center;
  color: #fff;
}

.logout-button {
  margin-left: 20px;
}

.register-button {
  background-color: #feeecd;
  color: black;
  padding: 10px;
  border: none;
  border-radius: 20px;
  cursor: pointer;
  margin-bottom: 20px;
  max-width: 350px;
  height: 50px;
  font-size: 25px;
}

.button-container {
  text-align: center;
  margin-top: 70px;
}

.head-other {
  background-color: #2e3d50;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px 0;
  width: 100%;
  height: 55px;
  margin-top: 20px;
  position: sticky;
  top: 65px; 
  left: 0;
  z-index: 1000;
}

.head-other h1 {
  font-family: 'Roboto Condensed', sans-serif;
  font-weight: 200;
  font-size: 25px;
  margin: 0;
  color: #fff;
  text-align: center;
}

.books-container,
.reviews-container,
.quote-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  margin-top: 20px;
  margin-bottom: 20px;
}

.review-item,
.book-item,
.quote-item {
  margin: 10px;
  flex: 1 1 calc(25% - 20px);
  max-width: calc(25% - 20px);
  box-sizing: border-box;
}

@media (max-width: 1024px) {
  .review-item,
  .book-item,
  .quote-item {
    flex: 1 1 calc(33.333% - 20px);
    max-width: calc(33.333% - 20px);
  }
}

@media (max-width: 768px) {
  .review-item,
  .book-item,
  .quote-item {
    flex: 1 1 calc(50% - 20px);
    max-width: calc(50% - 20px);
  }
}

@media (max-width: 480px) {
  .review-item,
  .book-item,
  .quote-item {
    flex: 1 1 calc(100% - 20px);
    max-width: calc(100% - 20px);
  }
}
</style>
