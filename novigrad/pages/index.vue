<script setup lang="ts">
const { data } = await useFetch("http://localhost:8080/items");

function convertCurrency(value: number) {
  const formatter = new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "BRL",
    minimumFractionDigits: 0,
  });

  return formatter.format(value);
}
</script>

<template>
  <h1 class="title">Borsodoy Auction House</h1>

  <p class="description">
    Welcome to our Online Auction House, a digital space where buyers and
    sellers come together to trade exclusive and valuable items. Our system is
    designed to
    <i>provide a safe, transparent, and exciting auction experience,</i>
    whether you are a collector looking for rare pieces or a seller interested
    in reaching a global audience.
  </p>

  <h3>Active Auctions</h3>

  <ul class="items">
    <li v-for="item in data">
      <NuxtLink :to="`/items/${item.id}`">
        <img :src="item.image_url" />

        <div class="item-data">
          <p class="item-name">{{ item.name }}</p>

          <span class="item-description">
            {{ item.description }}
          </span>

          <span class="item-price">
            Initial value:
            <strong>{{ convertCurrency(item.price) }}</strong>
          </span>
        </div>
      </NuxtLink>
    </li>
  </ul>
</template>

<style lang="scss">
h1.title {
  @apply text-2xl font-bold;
}

p.description {
  @apply text-center w-6/12 my-4;
}

h3 {
  @apply text-xl font-bold;
}

ul.items {
  @apply grid grid-cols-4 gap-8 my-8;

  li {
    @apply flex flex-col items-center w-64 h-80 rounded-xl p-4 transition-all duration-700;

    border: 1px solid #1b263b;
    background-color: #e9ecef;

    &:hover {
      @apply rounded-3xl cursor-pointer;

      background-color: #1b263b;
      border-color: #e9ecef;
      color: #f8f9fa;
    }

    img {
      @apply rounded-xl mb-2 max-h-32 max-w-full;
    }

    div.item-data {
      @apply flex flex-col;

      p.item-name {
        @apply font-semibold;
      }
    }
  }
}
</style>
