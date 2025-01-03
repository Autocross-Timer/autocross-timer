<template>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">loading</div>
    </div>
    <div v-if="runs" class="content">
        <DisplayRun :runs="runs" :sortType="sortType"/>
    </div>
</template>

<script setup>
    import DisplayRun from '../components/DisplayRuns.vue';
    import { ref } from 'vue';
    import { useInterval } from './useInterval.js';
    import { getAllRuns, getClass } from '../components/apiHandler.js';

    const props = defineProps(['sortType', 'filterClass']);

    const runs = ref([]);

    updateRuns(); 
    
    function updateRuns() {
        if(!props.filterClass) {
            getAllRuns()
            .then(data => {
                runs.value = data
            });
        }
        else{
            getClass(props.filterClass)
            .then(data => {
                runs.value = data
            });
        }
    };
    
    useInterval(updateRuns);    
</script>