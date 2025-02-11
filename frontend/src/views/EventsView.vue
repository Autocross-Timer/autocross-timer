<script setup>
    import { getEvents } from '../components/apiHandler.js';
    import { ref } from 'vue';
    import DisplayEvent from '../components/DisplayEvent.vue';

    const finalTodaysEvents = ref([]);
    const finalEvents = ref([]);
    const date = new Date().toJSON().slice(0, 10);
    var todaysEvents = [];
    var events = [];
    getEvents().then(data => {
        for (let i = 0; i < data.length; i++) {
            if (data[i].eventDate == date) {
                todaysEvents.push(data[i]);
            }
            else if (data[i].eventDate < date) {
                events.push(data[i]);
            }
        };
        finalTodaysEvents.value = todaysEvents;
        finalEvents.value = events;
    });
</script>

<template>
    <div v-if="finalTodaysEvents">
        <h1>Todays Events</h1>
        <div v-for="event in finalTodaysEvents" :key="event.eventId">
            <DisplayEvent :event="event" />
        </div>
    </div>
    <div v-if="!finalEvents">
        <h1>Todays Events</h1>
        <p>No events today</p>
    </div>
    <div v-if="finalEvents">
        <h1>Past Events</h1>
        <div v-for="event in finalEvents" :key="event.eventId">
            <DisplayEvent :event="event" />
        </div>
    </div>
</template>