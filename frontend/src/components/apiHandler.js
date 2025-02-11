const API_URL = import.meta.env.VITE_API_URL;

export function getAllRuns(eventId) {
    return fetch(`${API_URL}/runs/${eventId}/`)
        .then(response => response.json())
};

export function getLeaderboardRuns(eventId) {
    return fetch(`${API_URL}/runs/leaderboard/${eventId}/`)
        .then(response => response.json())
}

export function getRun(eventId, run_number) {
    return fetch(`${API_URL}/runs/${eventId}/${run_number}/`)
        .then(response => response.json())
};

export function getCarRuns(eventId, car_number) {
    return fetch(`${API_URL}/car/${eventId}/${car_number}`)
        .then(response => response.json())
}

export function getClass(eventId, class_name) {
    return fetch(`${API_URL}/class/${eventId}/${class_name}`)
        .then(response => response.json())
}

export function getClasses(eventId) {
    return fetch(`${API_URL}/classes/${eventId}`)
        .then(response => response.json())
}

export function getEvents() {
    return fetch(`${API_URL}/events/`)
        .then(response => response.json())
}

export function getEvent(eventId) {
    return fetch(`${API_URL}/events/${eventId}/`)
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