<script setup>
    import { RouterLink } from 'vue-router';
    
    const props = defineProps(['runs', 'displayType', 'sortType']);

    function sort(runs, sortType){
        // sort by created time
        if (sortType === undefined) {
            const newRuns = [].slice.call(runs).sort((a, b) => {
                const runCreateA = Number(a.created);
                const runCreateB = Number(b.created);
                if (runCreateA < runCreateB) {
                    return 1;
                }
                if (runCreateA > runCreateB) {
                    return -1;
                }
                return 0;
            });
            return newRuns
        }

        // sort by run time, only show fastest per-car
        const newRuns = [].slice.call(runs).sort((a, b) => {
            const timeA = Number(sortType == 'raw' ? a.rawTime : a.paxTime);
            const timeB = Number(sortType == 'raw' ? b.rawTime : b.paxTime);

            // Add 2 seconds per cone
            const timePenA = Number(a.cones * 2);
            const timePenB = Number(b.cones * 2);

            if (timeA + timePenA > timeB + timePenB) {
                return 1;
            }
            if (timeA + timePenA < timeB + timePenB) {
                return -1;
            }
            return 0;
        });

        // filter to only fastest run
        const cars_seen = []
        const filterRuns = newRuns.filter((run) => {
            if (cars_seen.includes(run.carNumber)) {
                return false
            }
            cars_seen.push(run.carNumber)
            return true
        });
        return filterRuns
    }

    function individualmeta(runs){
        // Find best run
        var best_time = 999999999
        runs.forEach((run) => {
            const runTime = Number(run.rawTime) + run.cones * 2;
            if (runTime < best_time && run.isDnf!==1 && run.getsRerun!==1) {
                if (best_time !== 999999999) {
                    run.delta = Math.round((best_time - runTime) * 1000) / 1000
                }
                best_time = runTime
            }
        });
        runs.forEach((run) => {
            const runTime = Number(run.rawTime) + run.cones * 2;
            if (runTime == best_time) {
                run.best = 1
            }
        });
        return runs
    }
</script>

<template>
    <div v-if="!runs" class="text-center">
        <div class="spinner-border spinner-border-sm">loading runs</div>
    </div>
    <div v-if="runs" class="content">
        <div v-if="displayType==='individual-car'">
            <p class="run-container" v-for="(run, index) in individualmeta(runs)" :key="runs.runNumber">
                <span class="run-number">Run {{ index+1 }}</span>
                <span class="run-name">{{ run.driverName }}</span>
                <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                    <span class="right">
                        <span v-if="run.cones >= 2" class="run-cones">+{{ run.cones }} Cones</span>
                        <span v-else-if="run.cones == 1" class="run-cones">+{{ run.cones }} Cone</span>
                        <span class="run-rawtime">Raw: {{ run.rawTime }}</span>
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
                        <span class="run_alt_time">
                            <span class="run-paxtime">(PAX: {{ run.paxTime }})</span>
                        </span>
                    </span>
                </template>
                <br/>
                <span v-if="run.best==1" class="best-run">🔥 BEST TIME 🔥</span>
                <template v-if="run.isDnf!==1 && run.getsRerun!==1 && run.delta > 0">
                    <span class="right">
                        <span class="run-delta"> NEW BEST (-{{ run.delta }})</span>
                    </span>
                </template>
                <br/>
            </p>
        </div>
        <div v-else>
            <p v-for="(run, index) in sort(runs, sortType)" :key="sortType">
                <RouterLink :to="'/event/' + run.eventId + '/car/' + run.carNumber" class="link-container">
                    <span v-if="sortType" class="run-number">{{ index+1 }}</span>
                    <span v-else class="run-number">Run {{ run.runNumber }}</span>
                    <span class="run-name">{{ run.driverName }}</span>
                    <span v-if="sortType && index == 0">🥇</span>
                    <span v-else-if="sortType && index == 1">🥈</span>
                    <span v-else-if="sortType && index == 2">🥉</span>
                    <template v-if="run.isDnf!==1 && run.getsRerun!==1">
                        <span class="right">
                            <span v-if="run.cones >= 2" class="run-cones">+{{ run.cones }} Cones</span>
                            <span v-else-if="run.cones == 1" class="run-cones">+{{ run.cones }} Cone</span>
                            <span v-if="sortType==='raw'" class="run-rawtime">Raw {{ run.rawTime }}</span>
                            <span v-else class="run-paxtime">PAX {{ run.paxTime }}</span>
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
                            <span class="run_alt_time">
                                <span v-if="sortType==='raw'" class="run-paxtime">(PAX: {{ run.paxTime }})</span>
                                <span v-else class="run-rawtime">(Raw: {{ run.rawTime }})</span>
                            </span>
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

    .run_alt_time {
        color: grey;
    }

    .run-delta {
        color: green;
    }
    .best-run {
        color: green;
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