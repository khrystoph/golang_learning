= GoSort project = 

== Intro == 

The purpose of this little project is to play around more with various features of golang. What better way to learn about features than by putting them in action, right?! 
A long time ago, when I was in my Data Structures and Algorithms class in college, we wrote a project for ABET accreditation that compared various algorithms and their sorting speeds. The one of focus was heapsort and comparing against slower algorithms to compare times. I got fancy with it and tried to do threading in java and made a misstep in my assumptions because I forgot to account for the main thread. 

The program executed fast, but the results from the algorithms was skewed by the main thread jumping in randomly and causing context switches on the thread that was executing..which would occasionally cause mergesort and heapsort to be slower than the slow algorithms. I'm hoping to correct that here as well as understand if/how threading can impact the performance of known good algorithms (at least, mergesort and heapsort as they are recursive and are prime candidates for the threading model).

Finally, each of these algorithms will be run in their own thread (to leverage the power of threading in golang and get experience with channels) some of which will run concurrently with different affinities (assuming it makes sense to set affinity) and will block on things that need to fan out (like threaded mergesort). Also, sorts that do end up using threads will not have other algorithms running at the same time to maximize the gains of threading in terms of execution time by maximizing available threads. 

== Algorithms and Explanations == 

Bubblesort - classic starting sorting algorithm that is not efficient, but it is stable. This is the first sorting algorithm discussed in classes due to its simplicity in implementation as well as ability to follow the logic.  

Quicksort - as its name implies it's quick...it partitions the data using a pivot and sorts, which is one of the many "divide and conquer" sorting methods. It also does this in-place, as opposed to creating new arrays. This may also benefit from golang's threading as you can use the partitions created as individual threads. 

Heapsort - this uses the heap datastructure concept and takes a divide and conquer method by using the power of binary tree to take the largest element from the heap and place it in a sorted array.  

Mergesort - this is probably the most easily recognizable and definitely one of the most stable sorting algorithms. It's worst and average sort times are both almost the same. This is also a "divide and conquer" sorting method and leverages recursion to split the array into smaller faster pieces to sort and re-insert into the main array. 

Threaded Mergesort - This takes the mergesort and splits it out into threads via the recursion call. When we recurse into the next call of itself, we spawn this call as a thread. We then leverage channels to ensure thread safety, but we allow the go scheduler to optimize the threads to minimize context switching. 

I will re-evaluate adding additional algorithms down the road (I was shooting for 6) once I figure out which one(s) I want to add. I may add the hybrid heapsort/mergesort method described in the heapsort portion of the Performance Expectations section. Threaded Quicksort would also be another potential candidate for adding into this experiment as it uses divide and conquer to sort the array. 

== Performance Expectations == 

Bubblesort - This algorithm was one of the original and slowest algorithms. results should always be slower than all of the other algorithms as expected O is O(n^2). 
Heapsort - This algorithm expects a roughly similar completion time of O(nlog n). This mirrors a few other algorithms in the execution time, but it does not get the performance gains that can be achieved with threading, unlike a few of the other methods that exist. It seems that a combination of mergesort and heapsort can achieve parallelization on the order of O((nlog n)/m) where m is the number of threads that the system can handle. The advantage of heapsort is that it uses n+1 memory (n+constant) for auxilliary memory...which means that it works well on memory constrained systems. It's worst performance is also O(nlog n), but it's not considered a stable sort. 
Quicksort - The expectation of this sort is that on average (over a large enough sample size), we should see that Quicksort has an execution time of O(nlog n), just like heapsort and the divide and conquer methods. Optimizations can speed this up, but make it unstable. Also, the speed is dependent upon the chosen pivot point. If the pivot is at the start or the end, it will end up with the worst time, which is O(n^2). 
Mergesort - This algorithm is considered both stable and carries roughly the same speed as the other algorithms like Quicksort or Heapsort. The execution time is also expected to be (on worst and average case) as O(nlog n). This method does carry a relatively high footprint for memory, however, as it needs 2n space for memory. 
Threaded Mergesort - This algorithm is expected to get variable performance, but better than nlog n. I expect that with a 4-core processor, we should see about a 4x gain in speedover mergesort as we now have split the work into at least four threads that can run at the same time without interacting with one another, until it completes. Likely it will be a bit slower than 4x due to context switching and memory access patterns, but there will be tests run on several sizes of processor (including 20+ cores). 

== Hypothesis == 

We're going to see bubble sort take a long time...much longer than any of the other sorts. Mergesort, I believe, will come out on top overall, but I think we might collect a few datapoints where Quicksort will be faster than mergesort or even Threaded Mergesort. Heapsort is going to be significantly faster than Bubblesort, but is going to be behind Quicksort and Mergesort (and all variants). 

== Methodology == 
The non-threaded sorting algorithms will be split up into threads that run concurrently. I still need to learn more about the go scheduler as well as if/how I can assure affinity, but program counters has been pointed out to me as a method for testing the performance. I'd like to do testing in two ways: 
1. total time for thread to complete 
2. actual time spent computing 

I plan on doing this because actual execution time is great, but it isn't a representation of what goes on in reality in terms of things coming in and interrupting a task. While I don't expect there to be much difference as these tests will generally be run on isolated hardware that doesn't do anything else, but the OS and hardware can come in and interrupt things, which is why I like the realistic approach of measuring total completion time as well as understand real computing time...you can use the two values to understand how often the execution of the algorithm was interrupted.

Once the non-threaded algorithms complete (I'll be using channels to achieve this interlock), I'll start each threaded algorithm in series. Once all the threads from the first algorithm complete, it will move to the next one, and so on until all threaded algorithms have run. At the end, I'll display the results from each of the algorithms that run for analysis/comparison. 

So, another major point is that I'm going to keep the array size the same, but this will be the only input that you supply to run the program as the program will generate random numbers to insert into the array on each run. I'll continue running tests with larger size arrays until there is a distinguishable difference between results (if the array sizes are too small, you won't notice the compute time difference as all the sorting algorithms are fast with small arrays. It's not until your arrays are very large (like 10m+ elements) that you will see a noticeable difference between the execution times. I'll also run several different size arrays after dialing in a set where it's small. There will be 20 runs at each size. Three total sizes. Each size will be 10x the smallest, so the tests might get extraordinarily long (especially as bubblesort will just sort of fall over at a point).

Additionally, these tests will be run in three different places. Once on my HPC server at home (I'll power on the Xeon E5v4 w/128GB ram for this. It has 20 physical cores and 40 threads...which should provide us with plenty of concurrency to get interesting data from our threaded algorithms), the second test will be on a c5.4xlarge (16vCPUs, 32GB ram) running in amazon's AWS cloud, the third test will be run from my personal laptop at home (4 physical Cores, 8 threads, 16GB ram). I might also throw in a m5.24xlarge and/or an r5.24xlarge to see if the added RAM+CPU makes a significant difference in a few of our algorithms. The added RAM may reduce the number of disk access times in the situation where we are using very, very large array sizes. 

== Data ==
Not completed yet.

== Conclusions ==
Not completed yet.
