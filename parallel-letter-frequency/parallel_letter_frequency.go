package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// recieves a word from the word Channel and builds a frequency map of letters.
// returns the frequency map to the aggregator Channel
func LetterFrequency(wordChan <-chan string, aggChan chan<- FreqMap) {
	for {
		word := <-wordChan
		letterMap := make(FreqMap)
		for _, letter := range word {
			letterMap[letter] += 1
		}
		aggChan <- letterMap
	}
}

// Returns the aggregated letter map from the words being processed
// takes the word count being aggregated
// recieves frequency maps from a channel for processed words
// returns an aggregated frequency map of all the words
func AggregateFrequency(wordCount int, aggChan <-chan FreqMap, resultChan chan<- FreqMap) {
	resultFreqMap := make(FreqMap)
	for i := 0; i < wordCount; i++ {
		letterMap := <-aggChan
		for letter, count := range letterMap {
			resultFreqMap[letter] += count
		}
	}

	resultChan <- resultFreqMap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	// channel takes the word
	var wordChan chan string = make(chan string)

	// frequency is written too and read from
	var freqChan chan FreqMap = make(chan FreqMap)
	var aggChan chan FreqMap = make(chan FreqMap)

	// start the Aggregator
	go AggregateFrequency(len(l), freqChan, aggChan)

	// start the letter counters (5)
	for i := 0; i < 5; i++ {
		go LetterFrequency(wordChan, freqChan)
	}

	// feed the words into the channel
	for _, word := range l {
		wordChan <- word
	}

	// wait for the aggregated results
	freq := <-aggChan

	return freq
}
