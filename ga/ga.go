package ga

import (
	"math/rand"
	"sync"
	"time"
)

var (
	BottomLimit     int
	UpperLimit      int
	ChromosomeRange int
)

type Chromosome struct {
	body    []int
	fitness float32
}

type DatasetNode struct {
	Chromosome
	answer bool
	rank   int
	mtx    sync.RWMutex
}

func (chr *Chromosome) countFitness(dataset []DatasetNode) {
	finalAnswer := true
	chr.fitness = 0
	var additionalPoints float32
	var tp, fn, fp, tn int
	//The problem with this approach is that additional_points can be calculated as summ with rank before and after operation datasetChromosome.rank = 0
	//But this should not make critical errors
	for _, datasetChromosome := range dataset {
		for genePos, geneVal := range datasetChromosome.body {
			if geneVal <= chr.body[genePos]+ChromosomeRange && geneVal >= chr.body[genePos]-ChromosomeRange {
				datasetChromosome.mtx.RLock()
				additionalPoints += float32(datasetChromosome.rank) / float32(len(dataset)*len(datasetChromosome.body))
				datasetChromosome.mtx.RUnlock()
			} else {
				finalAnswer = false
			}
		}
		if finalAnswer == datasetChromosome.answer {
			datasetChromosome.mtx.Lock()
			datasetChromosome.rank = 0
			datasetChromosome.mtx.Unlock()
		}
		switch {
		case finalAnswer == datasetChromosome.answer && finalAnswer == true:
			tp++
		case finalAnswer == datasetChromosome.answer && finalAnswer == false:
			tn++
		case finalAnswer != datasetChromosome.answer && finalAnswer == false:
			fn++
		case finalAnswer != datasetChromosome.answer && finalAnswer == true:
			fp++
		}
	}
	if tp+fn != 0 {
		chr.fitness = float32(tp) + additionalPoints
	}
}

func GeneratePopulation(chromosomeSize, populationSize int) []Chromosome {
	generatedPopulation := make([]Chromosome, populationSize)
	//Fill population with random chromosomes
	for i := 0; i < populationSize; i++ {
		for j := 0; j < chromosomeSize; j++ {
			generatedPopulation[i].body = append(generatedPopulation[i].body, rand.Intn(UpperLimit-BottomLimit)+BottomLimit)
		}
	}
	return generatedPopulation
}

func fitnessCalculator(ch <-chan *Chromosome, dataset []DatasetNode, wg *sync.WaitGroup) {
	defer wg.Done()
	for chr := range ch {
		chr.countFitness(dataset)
	}
}

func selection(population []Chromosome) []Chromosome {
	parentCandidates := make([]Chromosome, 0, len(population)/2)
	for len(population) > 1 {
		//Tournament selection
		first := rand.Intn(len(population))
		second := rand.Intn(len(population))
		for second == first {
			if second < len(population)-1 {
				second++
			} else if second > 0 {
				second--
			} else {
				break
			}
		}
		parentCandidates = append(parentCandidates, population[first], population[second])
		population[first], population[len(population)-1] = population[len(population)-1], population[first]
		population[second], population[len(population)-1] = population[len(population)-1], population[second]
		population = population[:len(population)-2]
	}
	return parentCandidates
}

func crossover(population []Chromosome) []Chromosome {

}

func GaEducate(dataset []DatasetNode, epochNum int) []Chromosome {
	rand.Seed(time.Now().UnixNano())
	currentPopulation := GeneratePopulation(len(dataset[0].body), 100)
	for epoch := 0; epoch < epochNum; epoch++ {
		//count fitness for each Chromosome
		chrChannel := make(chan *Chromosome)
		wg := new(sync.WaitGroup)
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go fitnessCalculator(chrChannel, dataset, wg)
		}
		for _, curChromosome := range currentPopulation {
			chrChannel <- &curChromosome
		}
		close(chrChannel)
		wg.Wait()
		//selection phase
		selectedPopulation := selection(currentPopulation)
		//make offspring
		children := crossover(selectedPopulation)
	}
	return nil
}
