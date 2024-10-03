<script setup lang="ts">
const { id } = useRoute().params;
function convertCurrency(value: number) {
  const formatter = new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
    minimumFractionDigits: 0,
  });

  return formatter.format(value);
}

const { data } = await useFetch(`http://localhost:8080/item/${id}`);
</script>

<template>
  <h1 class="item-name">{{ data.name }}</h1>
  <p class="item-description">{{ data.description }}</p>
  <p>
    Seller:
    <strong>{{ data.user.name }}</strong>
  </p>
  <p>
    Initial price:
    <strong>{{convertCurrency(data.price)}}</strong>
  </p>

  <figure class="item-image">
    <img :src="data.image_url" alt="" />
  </figure>
</template>

<style lang="scss" scoped>
h1.item-name {
  @apply font-bold text-2xl;
}

p.item-description {
  @apply text-lg;
}

figure.item-image {
  @apply my-6;

  &,
  img {
    @apply max-h-80 rounded-2xl;
  }
}
</style>
