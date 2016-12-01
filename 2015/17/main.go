package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Container struct {
	ID       int
	Capacity int
}

type ContainerSet struct {
	Containers []Container
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	containers := ContainerSet{}
	i := 0
	for s.Scan() {
		cap, _ := strconv.Atoi(s.Text())
		containers = containers.With(Container{
			ID:       i,
			Capacity: cap,
		})
		i++
	}

	liters, _ := strconv.Atoi(os.Args[1])
	combinations := make(map[string]ContainerSet)
	CombineContainers(containers, ContainerSet{}, liters, combinations)

	minContainers := 0
	minContainerCombinations := 0
	for _, combination := range combinations {
		if minContainers == 0 || minContainers > len(combination.Containers) {
			minContainerCombinations = 1
			minContainers = len(combination.Containers)
		} else if minContainers == len(combination.Containers) {
			minContainerCombinations++
		}
	}

	log.Printf("There are %d possible ways to combine these containers.", len(combinations))
	log.Printf("This requires at least %d containers. There are %d combinations with only %d containers.", minContainers, minContainerCombinations, minContainers)
}

func CombineContainers(containers ContainerSet, prefix ContainerSet, liters int, result map[string]ContainerSet) {
	for _, container := range containers.Containers {
		newContainerSet := prefix.With(container)
		if newContainerSet.Capacity() == liters {
			result[newContainerSet.String()] = newContainerSet
		} else if newContainerSet.Capacity() > liters {
			continue
		} else {
			go CombineContainers(containers.Without(container), newContainerSet, liters, result)
		}
	}
}

func (c ContainerSet) With(container Container) ContainerSet {
	return ContainerSet{
		Containers: append(c.Containers, container),
	}
}

func (c ContainerSet) Without(container Container) ContainerSet {
	new := ContainerSet{}
	for _, origContainer := range c.Containers {
		if container != origContainer {
			new = new.With(origContainer)
		}
	}
	return new
}

func (c ContainerSet) Capacity() int {
	result := 0
	for _, container := range c.Containers {
		result += container.Capacity
	}
	return result
}

func (c ContainerSet) String() string {
	containerStrings := []string{}
	for _, container := range c.Containers {
		containerStrings = append(containerStrings, fmt.Sprintf("%s", container))
	}
	sort.Strings(containerStrings)
	return fmt.Sprintf("Containers: %s, Capacity: %d", strings.Join(containerStrings, ", "), c.Capacity())
}

func (c Container) String() string {
	return fmt.Sprintf("%d: %d", c.ID, c.Capacity)
}
