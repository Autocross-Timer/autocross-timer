# Testing scripts

Supporting script to simulate an autocross sessions with fake data.

## Example (sim.py)

```
➜ python3 -m venv ./venv
➜ source ./venv/bin/activate
➜ python3 -m pip install -r requirements.txt
➜ python3 sim.py -h
usage: sim.py [-h] -u URL -c CARS -r RUNS

Description of your program

optional arguments:
  -h, --help            show this help message and exit
  -u URL, --url URL     URL for API
  -c CARS, --cars CARS  Number of cars in event
  -r RUNS, --runs RUNS  Number of runs for each car
➜ python3 sim.py -c 6 -r 2 -u http://localhost:8000
2025-07-13 11:15:08,910 INFO: ##############################################
2025-07-13 11:15:08,911 INFO: # test_race_1752419708.9109545
2025-07-13 11:15:08,911 INFO: # Event 6
2025-07-13 11:15:08,911 INFO: # test_eventId: pjhfibpvoawm
2025-07-13 11:15:08,911 INFO: ##############################################
2025-07-13 11:15:08,919 INFO: Run 1 for car 1
2025-07-13 11:15:08,928 INFO: Run 1 for car 5
2025-07-13 11:15:08,935 INFO: Run 1 for car 2
2025-07-13 11:15:08,943 INFO: Run 1 for car 4
2025-07-13 11:15:08,951 INFO: Run 1 for car 6
2025-07-13 11:15:08,959 INFO: Run 1 for car 3
2025-07-13 11:15:08,965 INFO: Run 2 for car 4
2025-07-13 11:15:08,977 INFO: Run 2 for car 3
2025-07-13 11:15:08,991 INFO: Run 2 for car 2
2025-07-13 11:15:09,003 INFO: Run 2 for car 6
2025-07-13 11:15:09,020 INFO: Run 2 for car 5
2025-07-13 11:15:09,030 INFO: Run 2 for car 1
Delete Event? [Y/n]y
2025-07-13 11:15:11,222 INFO: Event pjhfibpvoawm deleted
2025-07-13 11:15:11,222 INFO: Done
```
