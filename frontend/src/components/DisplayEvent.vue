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
    #run-number {
        padding: 0px 5px 0px 5px;
        align-items: center;
        border: solid red 2px;
    }

    #run-name {
        padding-left: 10px;
    }

    .right {
        float: right;
    }

    #run-cones {
        color: red;
        padding-right: 10px;
    }

    #run-class {
        padding-right: 10px;
    }

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