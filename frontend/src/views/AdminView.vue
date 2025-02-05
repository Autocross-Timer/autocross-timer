<script setup>
import { ref } from 'vue'
import { createEvent, createRun, deleteEvent, deleteRun } from '../components/apiHandler.js';

const clubName = ref('');
const eventLocation = ref('');
const eventDate = ref('');
const eventNumber = ref('');
const eventId = ref('');

function makeEvent() {

    const event = {
        clubName: clubName.value,
        eventLocation: eventLocation.value,
        eventDate: eventDate.value,
        eventNumber: eventNumber.value,
        eventId: eventId.value
    };

    console.log('Making Event: ', event);
    createEvent(event);

};

const delEventEventId = ref('');

function delEvent() {
    console.log('Deleting Event: ', delEventEventId.value);
    deleteEvent(delEventEventId.value);
};

const runEventId = ref('');
const runNumber = ref('');
const carNumber = ref('');
const rawTime = ref('');
const paxTime = ref('');
const carClass = ref('');
const driverName = ref('');
const cones = ref('');

function makeRun(){
    const run = { 
        eventId: runEventId.value,
        runNumber: runNumber.value,
        carNumber: carNumber.value,
        rawTime: rawTime.value,
        paxTime: paxTime.value,
        carClass: carClass.value,
        driverName: driverName.value,
        cones: cones.value,
    };

    console.log('Making Run: ', run);
    createRun(run, run.eventId);
};

const delRunEventId = ref('');
const delRunNumber = ref('');

function delRun(){
    console.log('Deleting Run: ', delRunEventId.value, delRunNumber.value);
    deleteRun(delRunEventId.value, delRunNumber.value);
};

</script>

<template>
    <div class="make-event">
        <h2>Make Event</h2><br>
        <form @submit.prevent="makeEvent">
            <div class="form-group">
                <label for="clubName">Club Name: </label>
                <input type="text" id="clubName" v-model="clubName" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <label for="eventLocation">Event Location: </label>
                <input type="text" id="eventLocation" v-model="eventLocation" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="eventDate">Date: </label>
                <input type="date" id="eventDate" v-model="eventDate" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <label for="eventNumber">Event Number: </label>
                <input type="text" id="eventNumber" v-model="eventNumber" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <label for="eventId">Event ID: </label>
                <input type="text" id="eventId" v-model="eventId" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary">Make Event</button>
            </div>
        </form>
    </div>
    <div class="delete-event">
        <h2>Delete Event</h2><br>
        <form @submit.prevent="delEvent">
            <div class="form-group">
                <label for="eventId">Event ID: </label>
                <input type="text" id="eventId" v-model="delEventEventId" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary">Delete Event</button>
            </div>
        </form>
    </div>
    <div class="make-run">
        <h2>Make Run</h2><br>
        <form @submit.prevent="makeRun">
            <div class="form-group">
                <label for="eventId">Event ID: </label>
                <input type="text" id="eventId" v-model="runEventId" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <label for="runNumber">Run Number: </label>
                <input type="text" id="runNumber" v-model="runNumber" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="carNumber">Car Number: </label>
                <input type="text" id="carNumber" v-model="carNumber" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="rawTime">Raw Time: </label>
                <input type="text" id="rawTime" v-model="rawTime" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="paxTime">Pax Time: </label>
                <input type="text" id="paxTime" v-model="paxTime" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="carClass">Car Class: </label>
                <input type="text" id="carClass" v-model="carClass" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="driverName">Name: </label>
                <input type="text" id="driverName" v-model="driverName" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <label for="cones">Cones Hit: </label>
                <input type="text" id="cones" v-model="cones" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary">Make Run</button>
            </div>
        </form>
    </div>
    <div class="delete-run">
        <h2>Delete Run</h2><br>
        <form @submit.prevent="delRun">
            <div class="form-group">
                <label for="eventId">Event ID: </label>
                <input type="text" id="eventId" v-model="delRunEventId" class="form-control" required placeholder="">
            </div>
            <div class="form-group">
                <label for="runNumber">Run Number: </label>
                <input type="text" id="runNumber" v-model="delRunNumber" class="form-control" required placeholder=""></input>
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary">Delete Run</button>
            </div>
        </form>
    </div>
</template>