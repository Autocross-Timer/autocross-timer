<script setup>
    const props = defineProps(['event']);
</script>

<template>
    <div v-if="!event">
        <div class="spinner-border spinner-border-sm">loading</div>
    </div>
    <div class="event-container" v-if="event">
        <RouterLink :to="'/event/' + event.eventId" class="link-container">
            <h2>{{ event.clubName }}</h2>
            <p>Event {{ event.eventNumber }}</p>
            <p>{{ event.eventLocation }}</p>
        </RouterLink>
    </div>
</template>

<style scoped>
    .link-container {
        display: block;
        text-decoration: none; /* Remove default underline */
        color: inherit; /* Inherit text color */
        border-style: solid;
        border-radius: 10px;
        border-width: 2px;
    }

    .event-container {
        padding-bottom: 10px;
    }
</style>