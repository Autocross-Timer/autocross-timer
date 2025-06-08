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
                <span class="run-number">{{ index+1 }}</span>
                <span class="run-name">{{ run.driverName }}</span>
                <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                    <span class="right">
                        <span v-if="run.cones" class="run-cones">+{{ run.cones }}</span>
                        <span class="run-paxtime">{{ run.paxTime }}</span>
                    </span>
                </template>
                <template v-else-if="run.getsRerun==1">
                    <span class="right">
                        <span class="run-rerun">RERUN</span>
                    </span>
                </template>
                <template v-else>
                    <span class="right">
                        <span class="run-dnf">DNF</span>
                    </span>
                </template> 
                <br/>
                <span class="run-class">{{ run.carClass }}</span>
                <span class="run-car-number">#{{ run.carNumber }}</span>
                <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                    <span class="right">
                        <span class="run-rawtime">({{ run.rawTime }})</span>
                    </span>
                </template>
            </p>
        </div>
        <div v-else>
            <p v-for="(run, index) in sort(runs, sortType)" :key="sortType">
                <RouterLink :to="'/event/' + run.eventId + '/car/' + run.carNumber" class="link-container">
                    <span v-if="sortType" class="run-number">{{ index+1 }}</span>
                    <span v-else class="run-number">{{ run.runNumber }}</span>
                    <span class="run-name">{{ run.driverName }}</span>
                    <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                        <span class="right">
                            <span v-if="run.cones" class="run-cones">+{{ run.cones }}</span>
                            <span v-if="sortType==='raw'" class="run-rawtime">{{ run.rawTime }}</span>
                            <span v-else class="run-paxtime">{{ run.paxTime }}</span>
                        </span>
                    </template>
                    <template v-else-if="run.getsRerun==1">
                        <span class="right">
                            <span class="run-rerun">RERUN</span>
                        </span>
                    </template>
                    <template v-else>
                        <span class="right">
                            <span class="run-dnf">DNF</span>
                        </span>
                    </template> 
                    <br/>
                    <span class="run-class">{{ run.carClass }}</span>
                    <span class="run-car-number">#{{ run.carNumber }}</span>
                    <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                        <span class="right">
                            <span v-if="sortType==='raw'" class="run-paxtime">({{ run.paxTime }})</span>
                            <span v-else class="run-rawtime">({{ run.rawTime }})</span>
                        </span>
                    </template>
                </RouterLink>
            </p>
        </div>
    </div>
</template>

<style scoped>
    .run-cones, .run-rerun, .run-dnf {
        color: red;
        padding-right: 10px;
    }

    .run-class {
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