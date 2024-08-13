<template>
  <div class="review">
    <div v-if="review">
      <h2 class="title">{{ review.Title }}</h2>
      <p>Score: {{ review.Score }}</p>
      <p>{{ review.Opnion }}</p>
    </div>
    <div v-else>
      <p>Nenhuma revisão encontrada</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    review: {
      type: Object,
      required: true
    }
  }
}
</script>

<style scoped>
.review {
  background-color: #feeecd;
  padding: 15px;
  text-align: center;
  border-radius: 20px;
  margin-bottom: 10px;
  width: 100px;
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
