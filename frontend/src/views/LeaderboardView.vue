<script setup>
  import { ref, watch } from 'vue';
  import GetRun from '../components/FetchRuns.vue'
  import { useRoute } from 'vue-router';
  import NavLinks from '../components/NavLinks.vue';

  const sortType = ref('pax');
  const route = useRoute();
  const eventId = route.params.eventId;
  watch(() => route.params.sortClass, (newSortClass) => {
    updateActiveButton(newSortClass);
  });

  function updateActiveButton(newSortClass) {
    if(newSortClass === 'raw'){
      sortType.value = 'raw';
    }
    else{
      sortType.value = 'pax';
    }
  }
</script>

<template>
    <NavLinks />
    <RouterLink :to="'/event/' + eventId + '/leaderboard/pax'" class="link-container">PAX</RouterLink>
    <RouterLink :to="'/event/' + eventId + '/leaderboard/raw'" class="link-container">Raw</RouterLink>
    <GetRun :sortType="sortType" :leaderboard="true" :eventId="eventId"/>
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
    text-decoration: none; /* Remove default underline */
    color: inherit; /* Inherit text color */
  }

  .router-link-active {
    text-decoration: underline;
  }
</style>