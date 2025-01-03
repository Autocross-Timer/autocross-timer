<script setup>
    import DisplayRun from '../components/DisplayRuns.vue';
    import { getCarRuns } from '../components/apiHandler.js';
    import { ref, watch } from 'vue';
    import { useInterval } from '../components/useInterval.js';
    import { useRoute } from 'vue-router';

    const runs = ref(null);
    const route = useRoute();

    watch(() => route.params.carId, updateRuns(), {immediate: true});

    function updateRuns() {
        getCarRuns(route.params.carId)
        .then(data => {
            runs.value = data
            });
    };

    useInterval(updateRuns);

    
</script>

<template>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">loading</div>
    </div>
    <div v-if="runs" class="content">
        <DisplayRun :runs="runs" :displayType="'individual-car'"/>
    </div>
</template>