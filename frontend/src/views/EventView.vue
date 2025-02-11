<script setup>
    import GetRun from '../components/FetchRuns.vue'
    import { useRoute } from 'vue-router';
    import NavLinks from '../components/NavLinks.vue';

    const route = useRoute();
    const eventId = route.params.eventId;
</script>

<template>
    <NavLinks />
    <GetRun :eventId="eventId"/>
</template>