package generics

import "sync"

func split[A any](slice []A, chunkSize int) (r [][]A) {
	length := len(slice)

	pos := chunkSize
	for pos < length {
		r = append(r, slice[pos-chunkSize:pos])
		pos += chunkSize
	}

	r = append(r, slice[pos-chunkSize:length])
	return
}

type itemStruct[A any] struct {
	order int
	value []A
}

func mapChannel[A, B any](fx func(A) B, iStruct itemStruct[A], channel chan itemStruct[B], waitGroup *sync.WaitGroup) {
	var r []B

	for _, item := range iStruct.value {
		r = append(r, fx(item))
	}

	channel <- itemStruct[B]{
		order: iStruct.order,
		value: r,
	}

	waitGroup.Done()
}

func collectAsync[A any, B any](channel <-chan itemStruct[B], resultChannel chan<- []B, storageSize int) {
	var result []B
	storage := make([][]B, storageSize)

	for rItemStruct := range channel {
		order := rItemStruct.order

		storage[order] = rItemStruct.value
	}

	for _, chain := range storage {
		result = append(result, chain...)
	}

	resultChannel <- result
}

func Pmap[A, B any](fx func(A) B, slice []A, chunkSize int) []B {
	var waitGroup sync.WaitGroup
	channel := make(chan itemStruct[B], chunkSize)
	resultChannel := make(chan []B)

	chains := split(slice, chunkSize)
	for i, chain := range chains {
		waitGroup.Add(1)

		item := itemStruct[A]{
			order: i,
			value: chain,
		}

		go mapChannel(fx, item, channel, &waitGroup)
	}

	go collectAsync[A, B](channel, resultChannel, chunkSize)
	waitGroup.Wait()
	close(channel)

	return <-resultChannel
}
