for func in RemoveRedundantSpace GenRandomWord TextToWords TextToNGrams; do
    go test -v -run=nomatch -bench="Benchmark${func}" -cpuprofile=cpu_${func}.out
    # break # uncomment if you just want to run the first bench func
done
