<script setup>
    import { ref } from 'vue';
    const props = defineProps(['runs']);
    const runs = props.runs;

    const driversName = ref(runs[0].driverName);
    const carNumber = ref(runs[0].carNumber);
    const carClass = ref(runs[0].carClass);
    const bestRawTime = ref(Math.min(...runs.map(run => run.rawTime)));
    const bestPaxTime = ref(Math.min(...runs.map(run => run.paxTime)));
</script>

<template>
    <p>Driver: {{ driversName }}</p>
    <p>Car Number: {{ carNumber }}</p>
    <p>Car Class: {{ carClass }}</p>
    <p>Best Raw Time: {{ bestRawTime }}</p>
    <p>Best Pax Time: {{ bestPaxTime }}</p>
</template>
