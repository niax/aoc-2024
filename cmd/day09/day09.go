package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type file struct {
	Id       int
	Size     int
	BlockIds []int
}

type freeSpace struct {
	BlockId int
	Size    int
}

func (f *freeSpace) Merge(other freeSpace) bool {
	nextBlockId := f.BlockId + f.Size
	if nextBlockId == other.BlockId {
		f.Size += other.Size
		return true
	}
	return false
}

type filesystem struct {
	Files     []file
	Blocks    []*file
	FreeSpace []freeSpace
}

func (fs *filesystem) InsertFreespace(newFreeSpace freeSpace) {
	idx, exists := slices.BinarySearchFunc(fs.FreeSpace, newFreeSpace.BlockId, func(f freeSpace, target int) int {
		if f.BlockId > target {
			return 1
		} else if f.BlockId < target {
			return -1
		}
		return 0
	})
	if exists {
		fs.FreeSpace[idx].Size += newFreeSpace.Size
	} else {
		previousIdx := idx - 1
		if newFreeSpace.Merge(fs.FreeSpace[previousIdx]) {
			fs.FreeSpace[previousIdx] = newFreeSpace
		} else {
			// Insert
			fs.FreeSpace = slices.Insert(fs.FreeSpace, previousIdx, newFreeSpace)
		}
	}
}

func (fs *filesystem) AddFile(size int) {
	fs.Files = append(fs.Files, file{
		Id:       len(fs.Files),
		BlockIds: make([]int, size),
		Size:     size,
	})
	filePtr := &fs.Files[len(fs.Files)-1]
	for i := 0; i < size; i++ {
		filePtr.BlockIds[i] = len(fs.Blocks)
		fs.Blocks = append(fs.Blocks, filePtr)
	}
}

func (fs *filesystem) AddFreeSpace(size int) {
	fs.FreeSpace = append(fs.FreeSpace, freeSpace{
		BlockId: len(fs.Blocks),
		Size:    size,
	})
	for i := 0; i < size; i++ {
		fs.Blocks = append(fs.Blocks, nil)
	}
}

func (fs *filesystem) Checksum() int {
	checksum := 0
	for i := range fs.Blocks {
		if fs.Blocks[i] == nil {
			continue
		}
		checksum += (i * fs.Blocks[i].Id)
	}
	return checksum
}

func (fs *filesystem) Print() {
	for i := range fs.Blocks {
		if fs.Blocks[i] == nil {
			fmt.Print("nil ")
		} else {
			fmt.Printf("%d ", fs.Blocks[i].Id)
		}
	}
	fmt.Print("\n")
}

func (fs *filesystem) Compact() {
	firstIdx := 0
	lastIdx := len(fs.Blocks) - 1
	for lastIdx > firstIdx {
		if fs.Blocks[firstIdx] != nil {
			firstIdx++
		} else if fs.Blocks[lastIdx] == nil {
			lastIdx--
		} else {
			fs.Blocks[firstIdx] = fs.Blocks[lastIdx]
			fs.Blocks[lastIdx] = nil
		}
	}
}

func (fs *filesystem) CompactWholeFiles() {
	for i := len(fs.Files) - 1; i >= 0; i-- {
		filePtr := &fs.Files[i]

		currentBlockIdx := -1
		for j := range fs.FreeSpace {
			if fs.FreeSpace[j].BlockId > filePtr.BlockIds[0] {
				break
			}
			if fs.FreeSpace[j].Size >= filePtr.Size {
				currentBlockIdx = fs.FreeSpace[j].BlockId
				fs.FreeSpace[j].BlockId += filePtr.Size
				fs.FreeSpace[j].Size -= filePtr.Size
				break
			}
		}

		if currentBlockIdx < 0 {
			continue
		}

		oldBlockStart := filePtr.BlockIds[0]
		newBlockId := currentBlockIdx
		for j := 0; j < filePtr.Size; j++ {
			if fs.Blocks[newBlockId] != nil {
				panic("INVALID!")
			}
			fs.Blocks[newBlockId] = filePtr
			oldBlockId := filePtr.BlockIds[j]
			filePtr.BlockIds[j] = newBlockId
			fs.Blocks[oldBlockId] = nil
			newBlockId++
		}

		fs.InsertFreespace(freeSpace{
			BlockId: oldBlockStart,
			Size:    filePtr.Size,
		})
	}
}

func main() {
	inputFd, err := os.Open("inputs/09")
	if err != nil {
		panic(err)
	}
	defer inputFd.Close()

	input := bufio.NewScanner(inputFd)

	p1Fs := filesystem{
		Files:     make([]file, 0, 1024*1024),
		Blocks:    make([]*file, 0, 1024*1024),
		FreeSpace: make([]freeSpace, 0, 1024),
	}
	p2Fs := filesystem{
		Files:     make([]file, 0, 1024*1024),
		Blocks:    make([]*file, 0, 1024*1024),
		FreeSpace: make([]freeSpace, 0, 1024),
	}

	for input.Scan() {
		line := input.Text()

		freeSpace := false
		for i := range line {
			blockCount, _ := strconv.Atoi(line[i : i+1])

			if freeSpace {
				p1Fs.AddFreeSpace(blockCount)
				p2Fs.AddFreeSpace(blockCount)
			} else {
				p1Fs.AddFile(blockCount)
				p2Fs.AddFile(blockCount)
			}
			freeSpace = !freeSpace
		}
	}

	p1Fs.Compact()
	p2Fs.CompactWholeFiles()
	fmt.Printf("%d\n", p1Fs.Checksum())
	fmt.Printf("%d\n", p2Fs.Checksum())
}
