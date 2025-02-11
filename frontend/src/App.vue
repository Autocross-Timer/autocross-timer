<script setup>
import { RouterView } from 'vue-router';
</script>

<template>
  <header>
    <div class="wrapper">
      <RouterLink to="/" class="no-decoration"><h1 class="title green">Live Timing</h1></RouterLink>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.title {
  margin: 0;
  padding: 0;
}

h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: absolute;
  top: -10px;
  left: 10px;
}

.wrapper {
  padding-bottom: 10px;
}

.no-decoration {
  all: unset;
  cursor: pointer;
}

</style>
