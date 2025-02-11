<script setup>
    import { RouterLink } from 'vue-router';
    
    const props = defineProps(['runs', 'displayType', 'sortType']);

    function sort(runs, sortType){
        if(sortType === "raw"){
            const newRuns = [].slice.call(runs).sort((a, b) => a.rawTime - b.rawTime);
            return newRuns;
        }
        else if(sortType === "pax"){
            const newRuns = [].slice.call(runs).sort((a, b) => a.paxTime - b.paxTime);
            return newRuns;
        }
        else {
            return runs;
        }
    };
</script>

<template>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">loading runs</div>
    </div>
    <div v-if="runs" class="content">
        <div v-if="displayType==='individual-car'">
            <p class="run-container" v-for="(run, index) in runs" :key="runs.runNumber">
                <span id="run-number">{{ index+1 }}</span>
                <span id="run-name">{{ run.driverName }}</span>
                <span class="right">
                    <span v-if="run.cones" id="run-cones">+{{ run.cones }}</span>
                    <span id="run-paxtime">{{ run.paxTime }}</span>
                </span>
                <br/>
                <span id="run-class">{{ run.carClass }}</span>
                <span id="run-car-number">#{{ run.carNumber }}</span>
                <span class="right">
                    <span id="run-rawtime">({{ run.rawTime }})</span>
                </span>
            </p>
        </div>
        <div v-else>
            <p v-for="(run, index) in sort(runs, sortType)" :key="sortType">
                <RouterLink :to="'/event/' + run.eventId + '/car/' + run.carNumber" class="link-container">
                    <span v-if="sortType" id="run-number">{{ index+1 }}</span>
                    <span v-else id="run-number">{{ run.runNumber }}</span>
                    <span id="run-name">{{ run.driverName }}</span>
                    <span class="right">
                        <span v-if="run.cones" id="run-cones">+{{ run.cones }}</span>
                        <span v-if="sortType==='raw'" id="run-rawtime">{{ run.rawTime }}</span>
                        <span v-else id="run-paxtime">{{ run.paxTime }}</span>
                    </span>
                    <br/>
                    <span id="run-class">{{ run.carClass }}</span>
                    <span id="run-car-number">#{{ run.carNumber }}</span>
                    <span class="right">
                        <span v-if="sortType==='raw'" id="run-paxtime">({{ run.paxTime }})</span>
                        <span v-else id="run-rawtime">({{ run.rawTime }})</span>
                    </span>
                </RouterLink>
            </p>
        </div>
    </div>
</template>

<style scoped>
    #run-number {
        padding: 0px 5px 0px 5px;
        align-items: center;
        border: solid red 2px;
        color: white;
    }

    #run-name {
        padding-left: 10px;
        color: white;
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

    .link-container, .run-container {
        display: block;
        text-decoration: none; /* Remove default underline */
        color: inherit; /* Inherit text color */
        border-style: solid;
        border-width: 0 0 1px 0;
        padding-top: 5px;
    }
</style>