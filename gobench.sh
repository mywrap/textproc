for func in TextToSyls TextToNGrams GenRandomWord; do
    go test -v -run=nomatch -bench="Benchmark${func}" -cpuprofile=cpu_${func}.out
done
