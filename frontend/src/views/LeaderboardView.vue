<script setup>
  import { ref, watch } from 'vue';
  import GetRun from '../components/FetchRuns.vue'
  import { useRoute } from 'vue-router';
  import NavLinks from '../components/NavLinks.vue';
  import { getClasses } from '../components/apiHandler.js';

  const sortType = ref('pax');
  const route = useRoute();
  const eventId = route.params.eventId;
  const classes = ref(null)
  const sortClass = ref(route.params.sortClass);
  watch(() => route.params.sortClass, (newSortClass) => {
    sortClass.value = newSortClass;
    updateActiveButton(newSortClass);
  });

  function updateActiveButton(newSortClass) {
    if(newSortClass === 'raw'){
      sortType.value = 'raw';
      sortClass.value = false;
    }
    else if(newSortClass === 'pax'){
      sortType.value = 'pax';
      sortClass.value = false;
    }
    else{
      sortType.value = 'pax';
    }
  }

  getClasses(eventId)
    .then(data => {
      classes.value = data;
  });
</script>

<template>
    <NavLinks />
    <RouterLink :to="'/event/' + eventId + '/leaderboard/pax'" class="link-container">PAX</RouterLink>
    <RouterLink :to="'/event/' + eventId + '/leaderboard/raw'" class="link-container">Raw</RouterLink>
    <RouterLink v-for="carClass in classes" :to="'/event/' + eventId + '/leaderboard/' + carClass" class="link-container">{{ carClass }}</RouterLink>
    <GetRun :key="sortClass" :sortType="sortType" :leaderboard="false" :filterClass="sortClass" :eventId="eventId"/>
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