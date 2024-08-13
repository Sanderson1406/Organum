<template>
  <div class="quote">
    <h3 class="title">{{ quote.Title }}</h3>
    <div v-if="quote.Quotes">
      <div v-for="(quoteItem, index) in quote.Quotes" :key="index" class="quote-card">
        <p>{{ quoteItem }}</p>
      </div>
    </div>
    <div v-else>Nenhuma citação salva</div>
  </div>
</template>

<script>
export default {
  props: {
    quote: {
      type: Object,
      required: true,
      default: () => ({
        Title: '',
        Quotes: ['']
      })
    }
  }
}
</script>

<style scoped>
.quote {
  background-color: #feeecd;
  padding: 15px;
  text-align: center;
  border-radius: 20px;
  margin-bottom: 10px;
  max-width: 100%;
  box-sizing: border-box;
}

.title {
  font-size: 27px;
  font-weight: 400;
  margin: 4%;
}

p {
  font-size: 18px;
  margin-top: 3%;
  margin-bottom: 3%;
  text-align: center;
  font-weight: 300;
  color: black;
}
</style>
