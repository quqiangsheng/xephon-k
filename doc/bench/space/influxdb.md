# InfluxDB

1,000,000 6.1M WAL, data folder is empty

````
⇒  xkb --limit points --target influxdb
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: points
  points: 1000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 100
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
Do you want to proceed? [Y/N]y
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0029] generator stopped after 1000000 points pkg=k.bench 
INFO[0029] close data channel pkg=k.bench 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0058959335000000005 	 
0.009159969666666668 	 
0.011229679 	 
0.013282145 	 
0.015270097005449566 	 ..............
0.016677509277120688 	 ...........................................................
0.018215746064766215 	 ...............................................................................
0.019876091 	 
0.02079301199631902 	 ................................
0.023975321691131493 	 ......
0.02617547595833334 	 .
0.027930318166666666 	 
0.030007455722222225 	 
0.0320365560506329 	 .
0.033529532788235276 	 .
0.035604549571428584 	 
0.037761762250000004 	 
0.03995495111764706 	 
0.04229246 	 
0.055023617 	 
INFO[0030] run time 18.733643 s pkg=k.bench.reporter 
INFO[0030] total request 10000 pkg=k.bench.reporter 
INFO[0030] fastest 5892809 pkg=k.bench.reporter 
INFO[0030] slowest 55023617 pkg=k.bench.reporter 
INFO[0030] total request size 47000000 pkg=k.bench.reporter 
INFO[0030] toatl response size 0 pkg=k.bench.reporter 
INFO[0030] 204: 10000 pkg=k.bench.reporter 
INFO[0030] bench finished pkg=k.cmd.bench 
````

10,000,000  1.8 M 


````
root@329e05187d49:/var/lib/influxdb/data/xb/autogen/2# du -sh *
1.5M	000000004-000000003.tsm
332K	000000005-000000001.tsm
````

```` 
⇒  xkb --limit points --target influxdb
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: points
  points: 10000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 1000
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
Do you want to proceed? [Y/N]y
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0023] generator stopped after 10000000 points pkg=k.bench 
INFO[0023] close data channel pkg=k.bench 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.006332456714285712 	 
0.011334941534883722 	 
0.0152038685 	 
0.01617709979157388 	 ..................................................................................................
0.019085289336168654 	 ......................................................
0.021819765044131394 	 .....................
0.025140076313765188 	 .........
0.028374251141078836 	 ....
0.031259074407035146 	 ...
0.03442091589719627 	 ..
0.03699015820895522 	 .
0.0399347663125 	 
0.04243660850000001 	 
0.045809186999999994 	 
0.050121871307692314 	 
0.05558600208333333 	 
0.061915671000000005 	 
0.06706112425 	 
0.076331226 	 
0.0790071195 	 
INFO[0024] run time 19.655630 s pkg=k.bench.reporter 
INFO[0024] total request 10000 pkg=k.bench.reporter 
INFO[0024] fastest 4459598 pkg=k.bench.reporter 
INFO[0024] slowest 79678882 pkg=k.bench.reporter 
INFO[0024] total request size 470000000 pkg=k.bench.reporter 
INFO[0024] toatl response size 0 pkg=k.bench.reporter 
INFO[0024] 204: 10000 pkg=k.bench.reporter 
INFO[0024] bench finished pkg=k.cmd.bench 
````

100,000,000 19M

- tar -czf a.tar.gz *.tsm

````
root@816048fc9f7e:/var/lib/influxdb/data/xb/autogen/2# du -sh *
15M	000000032-000000005.tsm
2.2M	000000036-000000003.tsm
1.1M	000000038-000000002.tsm
488K	000000039-000000001.tsm
2.2M	a.tar.gz
````

```` 
⇒  xkb --limit points --target influxdb
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: points
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
Do you want to proceed? [Y/N]y
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0072] generator stopped after 100000000 points pkg=k.bench 
INFO[0072] close data channel pkg=k.bench 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.02547461117073171 	 
0.031097282741052626 	 .........
0.03665486298137381 	 .................
0.042203592756335215 	 ....................
0.04817877791572606 	 ..........................
0.05572437511756023 	 ........................................
0.06319053637218329 	 .........................
0.06934404148698481 	 ..................
0.07529120360185178 	 ...............
0.08208258025285177 	 ..........
0.08906499071720102 	 ......
0.09612240634000006 	 ....
0.10460719207407412 	 ..
0.11306831372000001 	 .
0.12126799583333335 	 
0.12983624591666668 	 
0.140761655 	 
0.15038328033333334 	 
0.17062228799999998 	 
0.179243064 	 
INFO[0073] run time 63.458617 s pkg=k.bench.reporter 
INFO[0073] total request 10000 pkg=k.bench.reporter 
INFO[0073] fastest 20303930 pkg=k.bench.reporter 
INFO[0073] slowest 179243064 pkg=k.bench.reporter 
INFO[0073] total request size 4700000000 pkg=k.bench.reporter 
INFO[0073] toatl response size 0 pkg=k.bench.reporter 
INFO[0073] 204: 10000 pkg=k.bench.reporter 
INFO[0073] bench finished pkg=k.cmd.bench 
````