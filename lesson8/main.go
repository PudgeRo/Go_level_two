package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	ErrorOpenFile  = fmt.Errorf("cannot open file")
	ErrorCloseFile = fmt.Errorf("cannot close file")
)

type File struct {
	fileName string
	filePath string
}

type Result struct {
	File
	crc32 [32]byte
}

func main() {
	var (
		wg      = sync.WaitGroup{}
		input   = make(chan File)
		results = make(chan *Result)
		path    = flag.String("path", ".", "path to files")
	)

	flag.Parse()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(input, results)
		}()
	}

	go func() {
		if err := filePathWalkDir(*path, input); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	sameHash := make(map[[32]byte][]File)
	for result := range results {
		sameHash[result.crc32] = append(sameHash[result.crc32], struct {
			fileName string
			filePath string
		}{fileName: result.fileName, filePath: result.filePath})
	}

	for crc, files := range sameHash {
		if len(files) > 1 {
			fmt.Printf("%v duplicates with hash %v\n", len(files), hex.EncodeToString(crc[:]))
			for i, f := range files {
				fmt.Printf("%v) The name of file: %v; its path: %v\n", i+1, f.fileName, f.filePath)
			}
			fmt.Println("")
		}
	}
}

func worker(inputFiles chan File, results chan *Result) {
	for file := range inputFiles {
		var (
			h   = crc32.NewIEEE()
			sum = [32]byte{}
		)
		f, err := os.Open(file.filePath)
		if err != nil {
			fmt.Printf("%w: %v\n", ErrorOpenFile, f)
		}
		if _, err := io.Copy(h, f); err != nil {
			fmt.Println(err)
			if err := f.Close(); err != nil {
				fmt.Printf("%w: %v\n", ErrorCloseFile, f)
				continue
			}
			continue
		}
		if err := f.Close(); err != nil {
			fmt.Printf("%w: %v\n", ErrorCloseFile, f)
			continue
		}
		copy(sum[:], h.Sum(nil))
		results <- &Result{
			File:  File{fileName: file.fileName, filePath: file.filePath},
			crc32: sum,
		}
	}
}

func filePathWalkDir(path string, inputFiles chan File) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			inputFiles <- File{info.Name(), path}
		}
		return nil
	})
	close(inputFiles)
	return err
}
