<template>
  <main>
    <slot />
  </main>
</template>

<script setup lang="ts">
useHead({
  title: "Borsodoy Auction House",
});
</script>

<style lang="scss">
main {
  @apply flex flex-col items-center w-full bg-neutral-200 pt-10;
}
</style>
