package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ReadInput(inputFile string) ([4]Bucket, error) {
	var buckets [4]Bucket

	file, err := os.Open(inputFile)
	if err != nil {
		return buckets, err
	}

	defer file.Close()

	var teams []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			teams = append(teams, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return buckets, err
	}

	if len(teams) != 36 {
		return buckets, fmt.Errorf("expected 36 teams, got %d", len(teams))
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	r.Shuffle(len(teams), func(i, j int) {
		teams[i], teams[j] = teams[j], teams[i]
	})

	for i := range 36 {
		buckets[i/9][i%9] = teams[i]
	}

	return buckets, nil
}

func WriteOutput(buckets [4]Bucket, matches []TeamPair, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	bucketNames := []string{"A", "B", "C", "D"}

	for i, bucket := range buckets {
		fmt.Fprintf(file, "Bucket %s:\n", bucketNames[i])
		for _, team := range bucket {
			fmt.Fprintln(file, team)
		}
		fmt.Fprintln(file) // add an empty line between buckets
	}

	fmt.Fprintln(file, "Matches:")
	for _, match := range matches {
		fmt.Fprintf(file, "%s\n", match)
	}

	return nil
}

func GenerateBucketPairs(b Bucket) []TeamPair {
	var pairs []TeamPair

	for i := range 3 {
		group := b[i*3 : (i+1)*3]
		pairs = append(pairs,
			TeamPair{group[0], group[1]},
			TeamPair{group[2], group[0]},
			TeamPair{group[1], group[2]},
		)
	}

	return pairs
}

func GenerateCrossBucketsPairs(a, b Bucket) []TeamPair {
	var pairs []TeamPair

	for i := range a {
		pairs = append(pairs, TeamPair{a[i], b[i]})
		pairs = append(pairs, TeamPair{b[(i+1)%len(b)], a[i]})
	}

	return pairs
}
