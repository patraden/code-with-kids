package main

import (
	"log"
)

func main() {
	buckets, err := ReadInput("example/teams.txt")
	if err != nil {
		log.Fatalf("input file error: %v", err)
	}

	a, b, c, d := buckets[0], buckets[1], buckets[2], buckets[3]
	bucketPairs := [][2]Bucket{
		{a, b},
		{a, c},
		{a, d},
		{b, c},
		{b, d},
		{c, d},
	}

	matchesFinal := []TeamPair{}
	matchesFinal = append(matchesFinal, GenerateBucketPairs(a)...)
	matchesFinal = append(matchesFinal, GenerateBucketPairs(b)...)
	matchesFinal = append(matchesFinal, GenerateBucketPairs(c)...)
	matchesFinal = append(matchesFinal, GenerateBucketPairs(d)...)

	for _, pair := range bucketPairs {
		matches := GenerateCrossBucketsPairs(pair[0], pair[1])
		matchesFinal = append(matchesFinal, matches...)
	}

	WriteOutput(buckets, matchesFinal, "example/results.txt")
}
