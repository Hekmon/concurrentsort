# ConcurrentSort / Quicksort

Compute the ideal slice size limit for concurrency on IntSlice type.

The quicksort algo will create 2 sub slices after partitionning the current one. The algo will then check if one of theses two slices can be launched in another goroutine by asking the concurrent manager, which itself use a mutex. As the tree progress to the last leaves, the sub slices will be shorter and shorter. Therefore, the calls to the mutex will be more and more frequents. This is not an issue with only one worker, but the more workers there is, the more calls to this mutex will happen in the last level of the tree, significantly slower the process by making all the goroutines waiting for this mutex most of their time.

To mitigate this issue, a limit is used to prevent concurrency and so calls to the manager and it's mutex. For example, if the limit is 12, a goroutine will not ask the manager if a slot is free in case the sub slice is less than 12 and launch the quicksort partitionning within the same goroutine instead. This will prevent launching new goroutines for very small slices which could be sorted for less time than the cost of launching a goroutine but also (and most importantly) prevent a call to the manager and so prevent the mutex lock (and/or wait for the lock).

This limit is dependent on the number of workers and theses benchmarks try to approach the right value for different setups.

* [Run 1](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#run-1)
    * [Setup](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#setup)
    * [Log](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#log)
    * [Chart](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#chart)
* [Run 2](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#run-2)
    * [Setup](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#setup-1)
    * [Log](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#log-1)
    * [Chart](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#chart-1)
* [Run 3](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#run-3)
    * [Setup](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#setup-2)
    * [Log](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#log-2)
    * [Chart](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#chart-2)
* [Run 4](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#run-4)
    * [Setup](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#setup-3)
    * [Log](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#log-3)
    * [Chart](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#chart-3)
* [Conclusion](https://github.com/Hekmon/concurrentsort/tree/master/quicksort.bench#conclusion)

## Run 1

### Setup

* Start limit       : 0
* Increase limit    : 10
* Stop limit        : 100
* Slice size        : 16777216 (2^24)
* Nb runs per limit : 128 (2^7)
* Nb workers        : runtime.NumCPU() -> 8

### Log

```
* Determining the best limit for concurrency on a slice of size 16777216 with 8 workers

Benchmarking with limit value at 0
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 8.426385338s with a limit of 0

Benchmarking with limit value at 10
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.689428945s with a limit of 10

Benchmarking with limit value at 20
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.520202066s with a limit of 20

Benchmarking with limit value at 30
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.530961026s with a limit of 30

Benchmarking with limit value at 40
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.519737656s with a limit of 40

Benchmarking with limit value at 50
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.571648364s with a limit of 50

Benchmarking with limit value at 60
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.535444325s with a limit of 60

Benchmarking with limit value at 70
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.532187449s with a limit of 70

Benchmarking with limit value at 80
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.568538891s with a limit of 80

Benchmarking with limit value at 90
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.600706091s with a limit of 90

Benchmarking with limit value at 100
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.647073148s with a limit of 100

Summary :
0	8.426385338s
10	2.689428945s
20	2.520202066s
30	2.530961026s
40	2.519737656s *
50	2.571648364s
60	2.535444325s
70	2.532187449s
80	2.568538891s
90	2.600706091s
100	2.647073148s

* Best limit is 40 with 2.519737656s
```

### Chart

[![Benchmark 1 chart](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_1_chart.png)](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_1_chart.png)

## Run 2

### Setup

* Start limit       : 0
* Increase limit    : 2
* Stop limit        : 20
* Slice size        : 16777216 (2^24)
* Nb runs per limit : 128 (2^7)
* Nb workers        : runtime.NumCPU() -> 8

### Log

```
* Determining the best limit for concurrency on a slice of size 16777216 with 8 workers

Benchmarking with limit value at 0
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 8.471243603s with a limit of 0

Benchmarking with limit value at 2
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.926005148s with a limit of 2

Benchmarking with limit value at 4
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 3.560316773s with a limit of 4

Benchmarking with limit value at 6
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 3.172370422s with a limit of 6

Benchmarking with limit value at 8
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.982243337s with a limit of 8

Benchmarking with limit value at 10
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.895457477s with a limit of 10

Benchmarking with limit value at 12
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.863083786s with a limit of 12

Benchmarking with limit value at 14
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.887618324s with a limit of 14

Benchmarking with limit value at 16
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.862066133s with a limit of 16

Benchmarking with limit value at 18
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.8832129s with a limit of 18

Benchmarking with limit value at 20
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 2.910157836s with a limit of 20

Summary :
0       8.471243603s
2       4.926005148s
4       3.560316773s
6       3.172370422s
8       2.982243337s
10      2.895457477s
12      2.863083786s
14      2.887618324s
16      2.862066133s *
18      2.8832129s
20      2.910157836s

* Best limit is 16 with 2.862066133s
```

### Chart

[![Benchmark 2 chart](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_2_chart.png)](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_2_chart.png)


## Run 3

### Setup

* Start limit       : 0
* Increase limit    : 10
* Stop limit        : 100
* Slice size        : 16777216 (2^24)
* Nb runs per limit : 128 (2^7)
* Nb workers        : 4

### Log
```
* Determining the best limit for concurrency on a slice of size 16777216 with 4 workers

Benchmarking with limit value at 0
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 6.790331836s with a limit of 0

Benchmarking with limit value at 10
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 3.96346049s with a limit of 10

Benchmarking with limit value at 20
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 3.907300657s with a limit of 20

Benchmarking with limit value at 30
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 3.991401562s with a limit of 30

Benchmarking with limit value at 40
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.058877141s with a limit of 40

Benchmarking with limit value at 50
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.069578393s with a limit of 50

Benchmarking with limit value at 60
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.127535815s with a limit of 60

Benchmarking with limit value at 70
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.136819795s with a limit of 70

Benchmarking with limit value at 80
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.219506177s with a limit of 80

Benchmarking with limit value at 90
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.259649689s with a limit of 90

Benchmarking with limit value at 100
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.339977659s with a limit of 100

Summary :
0	6.790331836s
10	3.96346049s
20	3.907300657s *
30	3.991401562s
40	4.058877141s
50	4.069578393s
60	4.127535815s
70	4.136819795s
80	4.219506177s
90	4.259649689s
100	4.339977659s

* Best limit is 20 with 3.907300657s
```

### Chart

[![Benchmark 3 chart](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_3_chart.png)](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_3_chart.png)

## Run 4

### Setup

* Start limit       : 0
* Increase limit    : 2
* Stop limit        : 20
* Slice size        : 16777216 (2^24)
* Nb runs per limit : 128 (2^7)
* Nb workers        : 4

### Log
```
* Determining the best limit for concurrency on a slice of size 16777216 with 4 workers

Benchmarking with limit value at 0
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 6.821426289s with a limit of 0

Benchmarking with limit value at 2
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 5.016777754s with a limit of 2

Benchmarking with limit value at 4
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.446274447s with a limit of 4

Benchmarking with limit value at 6
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.300276791s with a limit of 6

Benchmarking with limit value at 8
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.358062882s with a limit of 8

Benchmarking with limit value at 10
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.312607206s with a limit of 10

Benchmarking with limit value at 12
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.31651437s with a limit of 12

Benchmarking with limit value at 14
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.356993478s with a limit of 14

Benchmarking with limit value at 16
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.415888833s with a limit of 16

Benchmarking with limit value at 18
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.431664694s with a limit of 18

Benchmarking with limit value at 20
 1IS 2IS 3IS 4IS 5IS 6IS 7IS 8IS 9IS 10IS 11IS 12IS 13IS 14IS 15IS 16IS 17IS 18IS 19IS 20IS 21IS 22IS 23IS 24IS 25IS 26IS 27IS 28IS 29IS 30IS 31IS 32IS 33IS 34IS 35IS 36IS 37IS 38IS 39IS 40IS 41IS 42IS 43IS 44IS 45IS 46IS 47IS 48IS 49IS 50IS 51IS 52IS 53IS 54IS 55IS 56IS 57IS 58IS 59IS 60IS 61IS 62IS 63IS 64IS 65IS 66IS 67IS 68IS 69IS 70IS 71IS 72IS 73IS 74IS 75IS 76IS 77IS 78IS 79IS 80IS 81IS 82IS 83IS 84IS 85IS 86IS 87IS 88IS 89IS 90IS 91IS 92IS 93IS 94IS 95IS 96IS 97IS 98IS 99IS 100IS 101IS 102IS 103IS 104IS 105IS 106IS 107IS 108IS 109IS 110IS 111IS 112IS 113IS 114IS 115IS 116IS 117IS 118IS 119IS 120IS 121IS 122IS 123IS 124IS 125IS 126IS 127IS 128IS
Average run is 4.514325263s with a limit of 20

Summary :
0	6.821426289s
2	5.016777754s
4	4.446274447s
6	4.300276791s *
8	4.358062882s
10	4.312607206s
12	4.31651437s
14	4.356993478s
16	4.415888833s
18	4.431664694s
20	4.514325263s

* Best limit is 6 with 4.300276791s
```

### Chart

[![Benchmark 4 chart](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_4_chart.png)](https://github.com/Hekmon/concurrentsort/raw/master/quicksort.bench/bench_4_chart.png)

## Conclusion

It's seems a good value would be `nbWorkers * 1.5`. But only if `nbWorkers <= nbCPU`, indeed if `nbWorkers > nbCPU` not all workers can request the mutex in the same time and therefor increasing the limitation is pointless.
