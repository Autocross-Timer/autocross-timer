<script setup>
    import DisplayRun from '../components/DisplayRuns.vue';
    import { getCarRuns } from '../components/apiHandler.js';
    import { ref, watch } from 'vue';
    import { useInterval } from '../components/useInterval.js';
    import { useRoute } from 'vue-router';
    import NavLinks from '../components/NavLinks.vue';
    import { addCar, removeCar, checkCar } from '../components/manageWatchlist.js';
    import DisplayDriver from '../components/DisplayDriver.vue';

    const runs = ref(null);
    const route = useRoute();

    var carId = route.params.carId;
    var eventId = route.params.eventId;
    var isWatched = ref(checkCar(eventId, carId));

    function updateRuns() {
        getCarRuns(route.params.eventId, route.params.carId)
        .then(data => {
            runs.value = data
            });
    };

    useInterval(updateRuns);

    function followCar() {
        addCar(eventId, carId);
        isWatched.value = checkCar(eventId, carId);
    };

    function unfollowCar() {
        removeCar(eventId, carId);
        isWatched.value = checkCar(eventId, carId);
    };
    
</script>

<template>
    <NavLinks />
    <button v-if="!isWatched" @click="followCar()">Follow car {{ carId }}</button>
    <button v-if="isWatched" @click="unfollowCar()">Unfollow car {{ carId }}</button>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">loading</div>
    </div>
    <div v-if="runs" class="content">
        <DisplayDriver :runs="runs"/>
        <DisplayRun :runs="runs" :displayType="'individual-car'"/>
    </div>
</template>