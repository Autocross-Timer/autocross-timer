<script setup>
    import { RouterLink, useRoute } from 'vue-router';

    const route = useRoute();
    const eventId = route.params.eventId;    
</script>

<template>
    <header>
        <div class="top-links">
            <nav>
                <RouterLink :to="'/event/' + eventId">Runs</RouterLink>
                <RouterLink :to="'/event/' + eventId + '/leaderboard/pax'">Leaderboard</RouterLink>
                <RouterLink :to="'/event/' + eventId + '/watchlist'">Watch List</RouterLink>
            </nav>
        </div>
    </header>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 20px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}
</style>