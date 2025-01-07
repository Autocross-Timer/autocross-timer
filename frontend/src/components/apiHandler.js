const API_URL = `https://live.jaredkelly.ca/api`;

export function getAllRuns() {
    return fetch(`${API_URL}/runs/`)
        .then(response => response.json())
};

export function getRun(id) {
    return fetch(`${API_URL}/runs/${id}/`)
        .then(response => response.json())
};

export function getCarRuns(id) {
    return fetch(`${API_URL}/car/${id}`)
        .then(response => response.json())
}

export function getClass(name) {
    return fetch(`${API_URL}/class/${name}`)
        .then(response => response.json())
}