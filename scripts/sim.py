import logging
import argparse
import requests
from enum import Enum
import random
import string
from random_username.generate import generate_username
from datetime import datetime
import time

LOG = logging.getLogger(__name__)
logging.basicConfig(format='%(asctime)s %(levelname)s: %(message)s', level=logging.INFO)

BASE_URL = "USER_PROVIDED"

class CarClass(Enum):
    SS    = ("SS", 0.836)
    SSP   = ("SSP", 0.857)
    XP    = ("XP", 0.890)
    AM    = ("AM", 1.000)
    AS    = ("AS", 0.826)
    ASP   = ("ASP", 0.851)
    BP    = ("BP", 0.874)
    BM    = ("BM", 0.978)
    BS    = ("BS", 0.819)
    BSP   = ("BSP", 0.855)
    CP    = ("CP", 0.863)
    CM    = ("CM", 0.899)
    CS    = ("CS", 0.813)
    CSP   = ("CSP", 0.859)
    DP    = ("DP", 0.865)
    DM    = ("DM", 0.906)
    DS    = ("DS", 0.811)
    DSP   = ("DSP", 0.847)
    EP    = ("EP", 0.858)
    EM    = ("EM", 0.916)
    ES    = ("ES", 0.790)
    ESP   = ("ESP", 0.840)
    FP    = ("FP", 0.877)
    FM    = ("FM", 0.917)
    FS    = ("FS", 0.817)
    FSP   = ("FSP", 0.831)
    FSAE  = ("FSAE", 0.980)
    GS    = ("GS", 0.794)
    HCR   = ("HCR", 0.814)
    HS    = ("HS", 0.784)
    CAM_T = ("CAM-T", 0.821)
    KM    = ("KM", 0.937)
    HCS   = ("HCS", 0.789)
    CAM_C = ("CAM-C", 0.825)
    SMF   = ("SMF", 0.850)
    CAM_S = ("CAM-S", 0.844)
    SM    = ("SM", 0.868)
    SSC   = ("SSC", 0.809)
    SSM   = ("SSM", 0.878)
    XA    = ("XA", 0.846)
    SST   = ("SST", 0.837)
    XB    = ("XB", 0.848)
    CSM   = ("CSM", 0.800)
    AST   = ("AST(STR)", 0.834)
    XU    = ("XU", 0.869)
    CSX   = ("CSX", 0.812)
    BST   = ("BST(STU)", 0.833)
    CST   = ("CST", 0.830)
    EVX   = ("EVX", 0.839)
    DST   = ("DST(STX)", 0.818)
    EST   = ("EST(STS)", 0.815)
    GST   = ("GST(STH)", 0.810)

    def __new__(cls, class_name, multi):
        entry = object.__new__(cls)
        entry._value_ = class_name
        entry.class_name = str(class_name)
        entry.multi = multi
        return entry

    def __str__(self):
        return self.class_name

class CarEntry:
    def __init__(self, event_id, car_num):
        self.event_id = event_id
        self.car_num = car_num
        self.total_runs = 0
        self.car_class = random.choice(list(CarClass))
        self.driver_name = generate_username()

    def add_run(self, run_num):
        self.total_runs +=1
        base_time = random.randint(30000, 50000) / 1000.0
        pax_time = round(base_time * self.car_class.multi, 3)
        cones = random.randint(0, 10) // random.randint(1, 15)

        runData = {
            "eventId":str(self.event_id),
            "runNumber":str(run_num),
            "carNumber":str(self.car_num),
            "rawTime":str(base_time),
            "paxTime":str(pax_time),
            "carClass":str(self.car_class),
            "driverName":self.driver_name,
            "cones":str(cones),
        }
        req = requests.post(f"{BASE_URL}/api/runs/{self.event_id}", json=runData)

        if req.status_code == 200:
            return
        LOG.error(req.text)

class Event:
    def __init__(self, num_cars, num_runs):
        self.event_id = ''.join(random.choice(string.ascii_lowercase) for i in range(12))
        self.num_runs = num_runs
        self.cars = []
        for car_num in range(1, num_cars + 1):
            self.cars.append(CarEntry(event_id = self.event_id, car_num=car_num))

    def do_runs(self):
        run = 0
        for run_num in range(1, self.num_runs + 1):
            random.shuffle(self.cars)
            for car in self.cars:
                LOG.info(f"Run {run_num} for car {car.car_num}")
                run += 1
                car.add_run(run_num=run)

    def add_event(self):
        club = f"test_race_{time.time()}"
        event_num = random.randint(1, 20)
        locaton = f"test_eventId: {self.event_id}"
        runData = {
            "clubName":club,
            "eventLocation":locaton,
            "eventDate": datetime.today().strftime('%Y-%m-%d'),
            "eventNumber":str(event_num),
            "eventId":str(self.event_id),
        }
        LOG.info(f"##############################################")
        LOG.info(f"# {club}")
        LOG.info(f"# Event {event_num}")
        LOG.info(f"# {locaton}")
        LOG.info(f"##############################################")

        req = requests.post(f"{BASE_URL}/api/events", json=runData)
        if req.status_code == 200:
            return
        raise Exception(f"Could not create event: {req.text}")

    def delete_event(self):
        req = requests.get(f"{BASE_URL}/api/event/delete/{self.event_id}")
        if req.status_code == 200:
            LOG.info(f"Event {self.event_id} deleted")
            return
        raise Exception(f"Could not delete event: {req.text}")

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Description of your program')
    parser.add_argument('-u', '--url', help='URL for API', required=True)
    parser.add_argument('-c', '--cars', type=int, help='Number of cars in event', required=True)
    parser.add_argument('-r', '--runs', type=int, help='Number of runs for each car', required=True)
    args = vars(parser.parse_args())
    BASE_URL = args['url']
    event = Event(num_cars=int(args['cars']), num_runs=int(args['runs']))
    event.add_event()
    event.do_runs()
    prompt = ""
    while prompt not in ['y', 'n']:
        prompt = input("Delete Event? [Y/n]").lower()
        if prompt == '':
            prompt = 'y'
    if prompt == 'y':
        event.delete_event()
    LOG.info("Done")
