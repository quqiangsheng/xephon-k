
10s, 10c, 1407

````
Total: 1407
0.015683903 	 
0.0253293515 	 
0.03133023744444444 	 .
0.036211688423076926 	 ...
0.040174712999999994 	 ......
0.04557385726315788 	 ........................
0.04997236222962964 	 ...................
0.0533814260448718 	 ......................
0.057669672732793525 	 ...................................
0.06191504331007749 	 ..................
0.06644933333936652 	 ...............................
0.07135827299999999 	 ..............
0.07485128701754387 	 ........
0.0785662916 	 ....
0.08195361674193548 	 ....
0.08741449133333332 	 ....
0.09515455137500001 	 .
0.10306686749999999 	 
0.1087344955 	 
0.124013587 	 
INFO[0013] total time 10.118765 s pkg=k.bench.reporter 
INFO[0013] total request 1407 pkg=k.bench.reporter 
INFO[0013] fastest 0.015145 s pkg=k.bench.reporter 
INFO[0013] slowest 0.124014 s pkg=k.bench.reporter 
INFO[0013] average 0.059257 s pkg=k.bench.reporter 
INFO[0013] total request size 253399293 pkg=k.bench.reporter 
INFO[0013] toatl response size 4221 pkg=k.bench.reporter 
INFO[0013] 200: 1407 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
  reporter: basic
  limitBy: time
  points: 100000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 10000
  numSeries: 10
targets:
  influxdb:
    host: localhost
    port: 8086
    url: write?db=xb
    timeout: 30
  xephonk:
    host: localhost
    port: 2333
    url: write
    timeout: 30
  kairosdb:
    host: localhost
    port: 8080
    url: api/v1/datapoints
    timeout: 30
INFO[0013] bench finished pkg=k.cmd.bench 
````