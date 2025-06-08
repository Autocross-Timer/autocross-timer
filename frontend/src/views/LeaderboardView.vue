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
  const sortClass = ref(false);
  const leaderboard = ref(false);

  if(route.params.sortClass !== 'raw' && route.params.sortClass !== 'pax'){
    sortClass.value = route.params.sortClass;
  }
  else {
    leaderboard.value = true;
  };

  watch(() => route.params.sortClass, (newSortClass) => {
    updateActiveButton(newSortClass);
  });

  function updateActiveButton(newSortClass) {
    if(newSortClass === 'raw'){
      sortType.value = 'raw';
      sortClass.value = false;
      leaderboard.value = true;
    }
    else if(newSortClass === 'pax'){
      sortType.value = 'pax';
      sortClass.value = false;
      leaderboard.value = true;
    }
    else{
      sortType.value = 'pax';
      sortClass.value = newSortClass;
      leaderboard.value = false;
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
    <GetRun :key="sortClass" :sortType="sortType" :leaderboard="leaderboard" :filterClass="sortClass" :eventId="eventId"/>
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