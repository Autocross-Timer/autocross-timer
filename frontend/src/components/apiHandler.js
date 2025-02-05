const API_URL = import.meta.env.VITE_API_URL;
const event_id = 1;

export function getAllRuns() {
    return fetch(`${API_URL}/runs/${event_id}/`)
        .then(response => response.json())
};

export function getLeaderboardRuns() {
    return fetch(`${API_URL}/runs/leaderboard/${event_id}/`)
        .then(response => response.json())
}

export function getRun(run_number) {
    return fetch(`${API_URL}/runs/${event_id}/${run_number}/`)
        .then(response => response.json())
};

export function getCarRuns(car_number) {
    return fetch(`${API_URL}/car/${event_id}/${car_number}`)
        .then(response => response.json())
}

export function getClass(class_name) {
    return fetch(`${API_URL}/class/${event_id}/${class_name}`)
        .then(response => response.json())
}

export function getEvents() {
    return fetch(`${API_URL}/events/`)
        .then(response => response.json())
}

export function getEvent(event_id) {
    return fetch(`${API_URL}/events/${event_id}/`)
        .then(response => response.json())
}

export function createEvent(event) {
    return fetch(`${API_URL}/events`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(event),
        mode: 'no-cors'
    })
}

export function createRun(run, eventId){
    return fetch(`${API_URL}/runs/${eventId}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(run),
        mode: 'no-cors'
    })
}

export function deleteEvent(eventId){
    return fetch(`${API_URL}/event/delete/${eventId}`)
        .then(response => response.json())
}

export function deleteRun(eventId, runNumber){
    return fetch(`${API_URL}/run/delete/${eventId}/${runNumber}`)
        .then(response => response.json())
}