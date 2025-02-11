<template>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">There are no runs for this event.</div>
    </div>
    <div v-if="runs" class="content">
        <DisplayRun :runs="runs" :sortType="sortType"/>
    </div>
</template>

<script setup>
    import DisplayRun from './DisplayRuns.vue';
    import { ref, watch } from 'vue';
    import { useInterval } from './useInterval.js';
    import { getAllRuns, getClass, getLeaderboardRuns } from './apiHandler.js';

    const props = defineProps(['sortType', 'filterClass', 'leaderboard', 'eventId']);

    const runs = ref([]);
    
    function updateRuns() {
        if(props.leaderboard){
            getLeaderboardRuns(props.eventId)
            .then(data => {
                runs.value = data
            });
        }
        else if(!props.filterClass) {
            getAllRuns(props.eventId)
            .then(data => {
                runs.value = data
            });
        }
        else{
            getClass(props.eventId, props.filterClass)
            .then(data => {
                runs.value = data
            });
        };
    };
    
    useInterval(updateRuns);    
</script>