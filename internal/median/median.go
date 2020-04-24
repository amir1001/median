package median

import (
	"errors"
	"math"
)

type medianCalculator struct {
	short               []float64
	long                []float64
	parity              int
	start               int
	end                 int
	shortPartitionIndex int
	longPartitionIndex  int
	shortLeftMax        float64
	shortRightMin       float64
	longLeftMax         float64
	longRightMin        float64
}

// GetMedian returns the median of 2 sorted arrays and an error
func GetMedian(a []float64, b []float64) (float64, error) {
	var medianCalc medianCalculator
	return medianCalc.getMedian(a, b)
}

func (m *medianCalculator) getMedian(a []float64, b []float64) (float64, error) {
	if err := m.initialise(a, b); err != nil {
		return math.Inf(-1), err
	}
	for {
		done, err := m.pursue()
		switch {
		case err != nil:
			return math.Inf(+1), err
		case done == true:
			return m.calculate(), nil
		}
	}
}

func (m *medianCalculator) initialise(a []float64, b []float64) error {
	if err := m.validate(a, b); err != nil {
		return err
	}
	m.short, m.long = m.align(a, b)
	m.parity = (len(m.short) + len(m.long)) % 2
	m.start, m.end = 0, len(m.short)
	return nil
}

func (m *medianCalculator) validate(a []float64, b []float64) error {
	if len(a)+len(b) <= 0 {
		return errors.New("invalid input error")
	}
	return nil
}

func (m *medianCalculator) align(a []float64, b []float64) (short []float64, long []float64) {
	if len(a) > len(b) {
		return b, a
	}
	return a, b
}

func (m *medianCalculator) pursue() (done bool, err error) {
	// Using binary search in order to find the median
	// by partitioning each input slice into left and right
	// partitions and taking advantage of their internal order.
	m.initPursuit()
	switch {
	case m.shortLeftMax > m.shortRightMin || m.longLeftMax > m.longRightMin:
		done, err = false, errors.New("unordered input")
	case m.shortLeftMax <= m.longRightMin && m.longLeftMax <= m.shortRightMin:
		done, err = true, nil
	case m.shortLeftMax > m.longRightMin:
		// this means the solution cannot be located in the short-right partition
		m.end = m.shortPartitionIndex
		done, err = false, nil
	case m.longLeftMax > m.shortRightMin:
		// this means the solution cannot be located in the short-left partition
		if m.start == m.shortPartitionIndex {
			m.start++
		} else {
			m.start = m.shortPartitionIndex
		}
		done, err = false, nil
	}
	return done, err
}

func (m *medianCalculator) initPursuit() {
	// invariant: len(shortLeft) + len(longLeft) = len(shortRight) + len(longRight), hence:
	// shortPartitionIndex + longPartitionIndex = len(short)-shortPartitionIndex+len(long)-longPartitionIndex
	m.shortPartitionIndex = (m.start + m.end) / 2
	m.longPartitionIndex = int((len(m.short)+len(m.long)-2*m.shortPartitionIndex)/2 + m.parity)
	m.shortLeftMax = m.getMaxLeftElement(m.short, m.shortPartitionIndex)
	m.shortRightMin = m.getMinRightElement(m.short, m.shortPartitionIndex)
	m.longLeftMax = m.getMaxLeftElement(m.long, m.longPartitionIndex)
	m.longRightMin = m.getMinRightElement(m.long, m.longPartitionIndex)
}

func (m *medianCalculator) calculate() float64 {
	leftMax := math.Max(m.shortLeftMax, m.longLeftMax)
	if m.parity == 1 {
		return leftMax
	}
	rightMin := math.Min(m.shortRightMin, m.longRightMin)
	return (leftMax + rightMin) / 2
}

func (m *medianCalculator) getMaxLeftElement(sortedInput []float64, partitionIndex int) float64 {
	if partitionIndex <= 0 {
		return math.Inf(-1)
	}
	return sortedInput[partitionIndex-1]
}

func (m *medianCalculator) getMinRightElement(sortedInput []float64, partitionIndex int) float64 {
	if partitionIndex >= len(sortedInput) {
		return math.Inf(+1)
	}
	return sortedInput[partitionIndex]
}
