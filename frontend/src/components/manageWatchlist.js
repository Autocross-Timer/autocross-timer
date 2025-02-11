export function addCar(eventId, carNumber){
    //console.log("Adding car " + carNumber + " from event " + eventId);
    var retString = localStorage.getItem(eventId);
    var retArray = JSON.parse(retString);
    if(retArray.includes(carNumber)){
        return;
    }
    if (retArray == null){
        retArray = [];
    }
    retArray.push(carNumber);
    localStorage.setItem(eventId, JSON.stringify(retArray));
};

export function removeCar(eventId, carNumber){
    //console.log("Removing car " + carNumber + " from event " + eventId);
    var retString = localStorage.getItem(eventId);
    var retArray = JSON.parse(retString);
    if (retArray == null){
        retArray = [];
    }
    var index = retArray.indexOf(carNumber);
    if (index > -1){
        retArray.splice(index, 1);
    }
    localStorage.setItem(eventId, JSON.stringify(retArray));
};

export function checkCar(eventId, carNumber){
    var retString = localStorage.getItem(eventId);
    var retArray = JSON.parse(retString);
    if (retArray == null){
        retArray = [];
    }
    return retArray.includes(carNumber);
};

export function getWatchedCars(eventId){
    var retString = localStorage.getItem(eventId);
    var retArray = JSON.parse(retString);
    if (retArray == null){
        retArray = [];
    }
    return retArray;
};