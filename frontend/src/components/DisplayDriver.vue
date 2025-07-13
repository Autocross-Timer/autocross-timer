<script setup>
    import { ref } from 'vue';
    const props = defineProps(['runs']);
    const runs = props.runs;

    const driversName = ref(runs[0].driverName);
    const carNumber = ref(runs[0].carNumber);
    const carClass = ref(runs[0].carClass);
</script>

<template>
    <p>Driver: {{ driversName }}</p>
    <p>Car Number: {{ carNumber }}</p>
    <p>Car Class: {{ carClass }}</p>
</template>
