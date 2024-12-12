package day9

func parseInput(input string) []int {
	lengths := make([]int, 0)
	for _, c := range input {
		lengths = append(lengths, int(c-'0'))
	}
	return lengths
}

func calculateChecksum(disk []int) int {
	checksum := 0
	for pos, fileID := range disk {
		if fileID != -1 {
			checksum += pos * fileID
		}
	}
	return checksum
}

func PartOne(input string) (int, error) {
	lengths := parseInput(input)

	disk := make([]int, 0)
	fileID := 0
	
	for i := 0; i < len(lengths); i += 2 {
		for j := 0; j < lengths[i]; j++ {
			disk = append(disk, fileID)
		}
		fileID++
		
		if i+1 < len(lengths) {
			disk = append(disk, make([]int, lengths[i+1])...)
		}
	}

	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			for j := 0; j < i; j++ {
				if disk[j] == -1 {
					disk[j] = disk[i]
					disk[i] = -1
					break
				}
			}
		}
	}

	return calculateChecksum(disk), nil
}

func PartTwo(input string) (int, error) {
	lengths := parseInput(input)

	disk := make([]int, 0)
	fileID := 0
	fileStarts := make(map[int]int)
	fileLengths := make(map[int]int)
	
	for i := 0; i < len(lengths); i += 2 {
		fileStarts[fileID] = len(disk)
		fileLengths[fileID] = lengths[i]
		
		for j := 0; j < lengths[i]; j++ {
			disk = append(disk, fileID)
		}
		fileID++
		
		if i+1 < len(lengths) {
			disk = append(disk, make([]int, lengths[i+1])...)
		}
	}

	for id := fileID - 1; id >= 0; id-- {
		fileLen := fileLengths[id]
		fileStart := fileStarts[id]
		
		bestPos := -1
		for i := 0; i < fileStart; i++ {
			if disk[i] == -1 && i+fileLen <= len(disk) {
				allEmpty := true
				for j := 0; j < fileLen; j++ {
					if disk[i+j] != -1 {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					bestPos = i
					break
				}
			}
		}

		if bestPos != -1 {
			for i := 0; i < fileLen; i++ {
				disk[fileStart+i] = -1
				disk[bestPos+i] = id
			}
		}
	}

	return calculateChecksum(disk), nil
}
