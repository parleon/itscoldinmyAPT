# itscoldinmyAPT

programs are run by invoking the executable with a single command line argument representing slice size
e.g. ./q1 1000

q1 slice isnt regenerated after first summation since go compiler does not create a constant

q1 slices were tested to see if properly generated with random numbers, however, for ease of readability the slices are not printed



the observations of big-O from the documentation seems to be accurate. When running the code and testing for different inputs the sort stable scales with input at a rate of log(n) more.
for instance, for an x of 100, my computer sorts in ~79 µs and sortsstable for ~165 µs, for 10000 it is ~1.14 ms and ~4.47 ms respectively. the additional factor of log(n) means that sortstable should be 3 times longer, which it is (since n*log(n)*log(n)/n*log(n) = 3)