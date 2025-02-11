<template>
    <NavLinks />
    <!--<div><pre v-if="watchedUsers" lang="json">{{ watchedUsers }}</pre></div>
    <input type="text" id="carNumber" v-model="carNumber" placeholder="Car Number" />
    <button @click="addNewCar()">Add Car</button>
    <button @click="removeOldCar()">Remove Car</button>
    <button @click="checkNewCar()">Check Car</button>
    <p v-if="checkValue">{{ checkValue }}</p>-->
    <div v-if="!filteredRuns" class="text-center">
        <div class="spinner-border spinner-border-sm">loading</div>
    </div>
    <div v-if="filteredRuns" class="content">
        <DisplayRun :runs="filteredRuns"/>
    </div>
</template>

<script setup>
    import NavLinks from '../components/NavLinks.vue';
    import DisplayRun from '../components/DisplayRuns.vue';
    import { getWatchedCars, addCar, removeCar, checkCar } from '../components/manageWatchlist.js';
    import { useInterval } from '../components/useInterval.js';
    import { ref } from 'vue';
    import { useRoute } from 'vue-router';
    import { getAllRuns } from '../components/apiHandler.js';

    const route = useRoute();
    var eventId = route.params.eventId;
    const carNumber = ref('');
    const watchedUsers = ref(getWatchedCars(eventId));
    const checkValue = ref(null);
    const filteredRuns = ref(null);

    //updateRuns();

    function addNewCar() {
        addCar(eventId, carNumber.value);
        watchedUsers.value = getWatchedCars(eventId);
    };

    function removeOldCar() {
        removeCar(eventId, carNumber.value);
        watchedUsers.value = getWatchedCars(eventId);
    };

    function checkNewCar() {
        checkValue.value = checkCar(eventId, carNumber.value);
    };

    function updateRuns() {
        getAllRuns(eventId)
        .then(data => {
            filteredRuns.value = data.filter(run => watchedUsers.value.includes(run.carNumber));
            });
    };

    useInterval(updateRuns);
</script>