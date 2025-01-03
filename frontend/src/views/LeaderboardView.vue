<script setup>
  import { ref, watch } from 'vue';
  import GetRun from '../components/FetchRuns.vue'
  import { useRoute } from 'vue-router';

  const sortType = ref('pax');
  const route = useRoute();
  watch(() => route.params.sortClass, updateActiveButton(), {immediate: true});

  function updateActiveButton() {
    if(route.params.sortClass === 'raw'){
      sortType.value = 'raw';
    }
    else{
      sortType.value = 'pax';
    }
  }
</script>

<template>
    <RouterLink :to="'/leaderboard/pax'" class="link-container">PAX</RouterLink>
    <RouterLink :to="'/leaderboard/raw'" class="link-container">Raw</RouterLink>
    <GetRun :sortType="sortType" />
</template>

<style scoped>
  .button {
    background:none;
    border:none;
    margin:0;
    padding:0px 5px;
    cursor: pointer;
    color: inherit;
  }

  .link-container {
    display: block;
    text-decoration: none; /* Remove default underline */
    color: inherit; /* Inherit text color */
  }

  .router-link-active {
    text-decoration: underline;
  }
</style>