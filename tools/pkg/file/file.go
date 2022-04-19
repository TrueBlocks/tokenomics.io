package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TrueBlocks/tokenomics.io/tools/pkg/types"
)

func LineCounts(folder, chain, addr string) (types.Counts, error) {
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}
	base := "./" + folder + "exports/" + chain
	if !FolderExists(base) {
		fmt.Println("SOMSOMTEOMTOME")
		return types.Counts{}, fmt.Errorf("data folder (%s) not found", base)
	}
	counts := types.Counts{}
	counts.Appearances, _ = lineCount(folder+"exports/"+chain+"/apps/"+addr+".csv", true)
	counts.Neighbors, _ = lineCount(folder+"exports/"+chain+"/neighbors/"+addr+".csv", true)
	counts.Logs, _ = lineCount(folder+"exports/"+chain+"/logs/"+addr+".csv", true)
	counts.Txs, _ = lineCount(folder+"exports/"+chain+"/txs/"+addr+".csv", true)
	counts.Statements, _ = lineCount(folder+"exports/"+chain+"/statements/"+addr+".csv", true)
	return counts, nil
}

func lineCount(fileName string, ignoreHeader bool) (int, error) {
	// fmt.Println("-------------------------------------------")
	// fmt.Println(fileName, FileExists(fileName))
	// fmt.Println("-------------------------------------------")
	r, _ := os.Open(fileName)
	defer r.Close()

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	if ignoreHeader && count > 0 {
		count--
	}
	return count, nil
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func AsciiFileToLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	file.Close()
	return ret
}
