package main

import (
	// "fmt"
	// "fmt"
	// "fmt"
	// "fmt"
	// "fmt"
	// "fmt"

	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolutionAdd(param1 int, param2 int) int {
	return param1 + param2
}

func SolutionCenturyFromYear(year int) int {
	century := year / 100
	if year%100 != 0 {
		century++
	}
	return century
}

func SolutionCheckPalindrome(inputString string) bool {
	stringReverse := make([]byte, len(inputString))
	for index, value := range inputString {
		stringReverse[len(stringReverse)-1-index] = byte(value)
	}
	if strings.Compare(inputString, string(stringReverse)) == 0 {
		return true
	}
	return false
}

func SolutionAdjProduct(inputArray []int) int {
	max := inputArray[0] * inputArray[1]
	for i := 1; i < len(inputArray)-1; i++ {
		multiply := inputArray[i] * inputArray[i+1]
		if multiply > max {
			max = multiply
		}
	}
	return max
}

func SolutionShapeArea(n int) int {
	area := n*n + (n-1)*(n-1)
	return area
}

func SolutionMakeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	statuesNeeded := 0

	for i := 0; i < len(statues)-1; i++ {
		if statues[i+1]-statues[i] != 1 {
			statusAdd := statues[i+1] - statues[i] - 1
			statuesNeeded += statusAdd
		}

	}
	return statuesNeeded
}

func SolutionAlmostIncreasingSequence(sequence []int) bool {
	countElementNotRestrict := 0
	sizeSequence := len(sequence)
	for i := 0; i < sizeSequence-1; i++ {
		if sequence[i+1] <= sequence[i] {
			countElementNotRestrict++

			skipNeighbor := i+2 < sizeSequence && sequence[i+2] <= sequence[i]
			skipBack := i-1 >= 0 && sequence[i+1] <= sequence[i-1]

			if (skipBack && skipNeighbor) || countElementNotRestrict >= 2 {
				return false
			}
		}
	}

	return true
}

func SolutionMatrixElementSum(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			isTopHaunted := (i > 0 && matrix[i-1][j] == 0) || (matrix[0][j] == 0)
			if matrix[i][j] != 0 && !isTopHaunted {
				sum += matrix[i][j]
			}
		}
	}

	return sum
}

func SolutionInputArray(inputArray []string) []string {
	maxLength := 0
	result := []string{}
	for _, value := range inputArray {
		if maxLength < len(value) {
			maxLength = len(value)
			result = []string{}
		}
		if len(value) == maxLength {
			result = append(result, value)
		}
	}

	return result
}

func SolutionCommonCharacterCount(s1 string, s2 string) int {
	mapFrequentS1 := map[string]int{}
	mapFrequentS2 := map[string]int{}
	countCommonChar := 0
	for _, value := range s1 {
		mapFrequentS1[string(value)] = mapFrequentS1[string(value)] + 1
	}

	for _, value := range s2 {
		mapFrequentS2[string(value)] = mapFrequentS2[string(value)] + 1
	}

	for key, freq := range mapFrequentS1 {
		if mapFrequentS2[key] != 0 {
			min := mapFrequentS2[key]

			if min > freq {
				min = freq
			}
			countCommonChar += min
		}
	}

	return countCommonChar
}

func SolutionIsLucky(n int) bool {
	sumFirstHalf := 0

	stringNumberTicket := strconv.Itoa(n)
	halfIndex := len(stringNumberTicket) / 2
	for i := 0; i < len(strconv.Itoa(n)); i++ {
		number, _ := strconv.Atoi(string(stringNumberTicket[i]))
		if i < halfIndex {
			sumFirstHalf += number
		} else {
			sumFirstHalf -= number
		}
	}

	if sumFirstHalf == 0 {
		return true
	}
	return false
}

func SolutionSortByHeight(a []int) []int {

	arrSort := make([]int, 0)
	for _, i := range a {
		if i != -1 {
			arrSort = append(arrSort, i)
		}
	}
	sort.Ints(arrSort)
	x := 0
	for n, i := range a {
		if i != -1 {
			a[n] = arrSort[x]
			x++
		}
	}
	return a
}

func SolutionReverseInParentheses(inputString string) string {
	stack, tmp := make([]rune, 0, len(inputString)), make([]rune, 0, len(inputString))
	for _, c := range inputString {
		if c == ')' {
			i := len(stack) - 1
			tmp = tmp[:0]
			for ; stack[i] != '('; i-- {
				tmp = append(tmp, stack[i])
			}
			stack = append(stack[:i], tmp...)
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func SolutionAlternatingSums(a []int) []int {
	result := make([]int, 2)
	isIntoTeam1 := true
	for _, value := range a {
		if isIntoTeam1 {
			result[0] += value
			isIntoTeam1 = false
		} else {
			result[1] += value
			isIntoTeam1 = true
		}
	}

	return result
}

func SolutionAddBorder(picture []string) []string {
	lenBorder := len(picture[0]) + 2
	result := make([]string, len(picture)+2)
	result[0] = strings.Repeat("*", lenBorder)
	result[len(result)-1] = strings.Repeat("*", lenBorder)

	for i := 1; i < len(result)-1; i++ {
		result[i] = "*" + picture[i-1] + "*"
	}

	return result
}

func SolutionAreSimilar(a []int, b []int) bool {
	c := 0
	s := make(map[int]bool)
	for n := range a {
		if a[n] != b[n] {
			c++
			s[a[n]] = true
			s[b[n]] = true
		}

	}

	return (c == 2 || c == 0) && (len(s) == 0 || len(s) == 2)
}

func SolutionArrayChange(inputArray []int) int {
	moves := 0
	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i] >= inputArray[i+1] {
			fmt.Println(i)
			moves += inputArray[i] - inputArray[i+1] + 1
			inputArray[i+1] = inputArray[i] + 1
		}
	}

	return moves
}

func SolutionPalindromeRearranging(inputString string) bool {
	m := make(map[rune]int)
	for _, r := range inputString {
		m[r]++
	}

	oddFrequent := 0
	for _, v := range m {
		if v%2 == 1 {
			oddFrequent++
		}
	}
	return oddFrequent < 2
}

func SolutionAreEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {

	weakestArmYour := yourLeft
	strongestArmYour := yourRight

	if weakestArmYour > yourRight {
		weakestArmYour = yourRight
		strongestArmYour = yourLeft
	}

	weakestArmFriends := friendsLeft
	strongestArmFriends := friendsRight

	if weakestArmFriends > friendsRight {
		weakestArmFriends = friendsRight
		strongestArmFriends = friendsLeft
	}

	if weakestArmFriends == weakestArmYour && strongestArmFriends == strongestArmYour {
		return true
	}

	return false
}

func SolutionArrayMaximalAdjacentDifference(inputArray []int) int {
	max := 0

	for i := 0; i < len(inputArray)-1; i++ {
		absoluteBetweenAdjacent := int(math.Abs(float64(inputArray[i+1] - inputArray[i])))

		if absoluteBetweenAdjacent > max {
			max = absoluteBetweenAdjacent
		}
	}

	return max
}

func SolutionIsIPv4Address(inputString string) bool {
	sliceString := strings.Split(inputString, ".")

	if len(sliceString) != 4 {
		return false
	}

	for _, value := range sliceString {

		if len(value) != 2 {
			if len(value) == 0 {
				return false
			}
		} else {
			if string(value[0]) == "0" {
				return false
			}
		}
		intParse, err := strconv.Atoi(value)

		if err != nil {
			return false
		}
		if intParse > math.MaxUint8 {
			return false
		}
	}
	return true
}

func SolutionAvoidObstacles(inputArray []int) int {
	step := 1
	for {
		isJumpEnough := true
		for i := 0; i < len(inputArray); i++ {
			if inputArray[i]%step == 0 {
				isJumpEnough = false
				fmt.Println("??", step, inputArray[i])
			}
		}
		if isJumpEnough {
			return step
		}
		step++
	}
}

func SolutionBoxBlur(image [][]int) [][]int {
	result := make([][]int, len(image)-2)
	for i := range result {
		result[i] = make([]int, len(image[i])-2)
	}
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			if i > 0 && j > 0 && i < len(image)-1 && j < len(image[i])-1 {
				// [i - 1][j - 1]
				// [i][j - 1]
				// [i+1][j - 1]

				// [i - 1][j]
				// [i + 1][j]

				// [i - 1][j + 1]
				// [i][j + 1]
				// [i+1][j + 1]
				average := math.Floor(float64((image[i][j] + image[i+1][j + 1] + image[i][j + 1] + image[i - 1][j + 1] + image[i - 1][j - 1] + image[i][j - 1] + image[i+1][j - 1] + image[i - 1][j] + image[i + 1][j]) / 9))
				result[i - 1][j - 1] = int(average)
			}
		}
	}

	return result
}

func SolutionMineSweeper(matrix [][]bool) [][]int {
	res := make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ {
        res[i] = make([]int, len(matrix[0]))
    }
    
    for r := 0; r < len(matrix); r++ {
        for c := 0; c < len(matrix[0]); c++ {
            sum := 0

            if r > 0 {
                if matrix[r-1][c] {
                    sum++
                }
                if c > 0  && matrix[r-1][c-1] {
                    sum++
                }
                if c < len(matrix[0])-1 && matrix[r-1][c+1] {
                    sum++
                }
            }
         
            if c > 0 && matrix[r][c-1] {
                sum++
            }
            if c < len(matrix[0])-1 && matrix[r][c+1] {
                sum++
            }
         
            if r < len(matrix)-1 {
                if matrix[r+1][c] {
                    sum++
                }
                if c > 0 && matrix[r+1][c-1] {
                    sum++
                }
                if c < len(matrix[0])-1 && matrix[r+1][c+1] {
                    sum++
                }
            }
            res[r][c] = sum
        }
    }
    return res
}

func ArrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == elemToReplace {
			inputArray[i] =substitutionElem
		}
	}
	
	return inputArray
}

func SolutionEvenDigitsOnly(n int) bool {
	stringDigits := strconv.Itoa(n)

	for i := 0; i < len(stringDigits); i++ {
		value, _ := strconv.Atoi(string(stringDigits[i]))

		if value % 2 != 0 {
			return false
		}
	}

	return true
}

func SolutionVariableName(name string) bool {
	if name[0] >= 48 && name[0] <= 57 {
		return false
	}

	for i := 0; i < len(name); i++ {
		decimalASCII := name[i] 
		
		if decimalASCII < 48 || (decimalASCII > 57 && decimalASCII < 65) || (decimalASCII > 90 && decimalASCII < 97 && decimalASCII != 95) || decimalASCII > 122 {
			return false
		} 
	}
	return true
}

func SolutionAlphabeticShift(inputString string) string {
	Z := 122
	A := 97
	bytesArray :=make([]byte, len(inputString))
	copy(bytesArray, inputString)
	
	for index, value := range inputString {
		if int(value) == Z {
			bytesArray[index] = byte(A)
		} else {
			bytesArray[index] = bytesArray[index] + 1
		}
		fmt.Println(value)
	}

	return string(bytesArray)
}

func SolutionChessBoardCellColor(cell1 string, cell2 string) bool {
	matrixChessBoard := [][]bool{
		{false, true, false, true, false, true, false, true},
		{true, false, true, false, true, false, true, false},
		{false, true, false, true, false, true, false, true},
		{true, false, true, false, true, false, true, false},
		{false, true, false, true, false, true, false, true},
		{true, false, true, false, true, false, true, false},
		{false, true, false, true, false, true, false, true},
		{true, false, true, false, true, false, true, false},
	}
	
	mapHorizontal := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
	}

	cell1Horizontal := mapHorizontal[string(cell1[0])]
	cell1Vertical, _ := strconv.Atoi(string(cell1[1]))

	cell2Horizontal := mapHorizontal[string(cell2[0])]
	cell2Vertical, _ := strconv.Atoi(string(cell2[1]))


	if matrixChessBoard[cell1Horizontal][8 - cell1Vertical] == matrixChessBoard[cell2Horizontal][8 - cell2Vertical] {
		return true
	}
	
	return false
}

func SolutionCircleOfNumber(n int, firstNumber int) int {
	if firstNumber < n / 2 {
        return firstNumber + n / 2
    } else {
        return firstNumber - n / 2
    }
}

func SolutionDepositProfit(deposit int, rate int, threshold int) int {
	years := 0
    balance := float64(deposit)
    for balance < float64(threshold) {
        balance = balance * (1 + float64(rate)/100)
        years++
    }
    return years
}

func SolutionAbsoluteValuesSumMinimization(a []int) int {
	mapMin := map[int][]int{}

	min := 0
	for i := 0; i < len(a); i++ {
		sum := 0.0
		for j := 0; j < len(a); j++ {
			sum += math.Abs(float64(a[i] - a[j])) 
		}
		value, isExist := mapMin[int(sum)] 

		if isExist {
			mapMin[int(sum)] = append(value, a[i])
		} else {
			mapMin[int(sum)] = []int{a[i]}
		}

		if i == 0 {
			min = int(sum)
		} else {
			if min > int(sum) {
				min = int(sum)
			}
		}
	}

	for key, value := range mapMin {
		if key == min {
			minSame := value[0]
			for i := 0; i < len(value); i++ {
				if minSame > value[i] {
					minSame = value[i]
				}
			}
			return minSame
		}
	}

	return -1
}

var perms [][]string

func SolutionStringsRearrangement(inputArray []string) bool {
	perms = make([][]string, 0)
    n := len(inputArray)
    
    permute(inputArray, 0, n-1)
    for _, perm := range perms {
        fail := false
        prev := perm[0]
        for _, str := range perm[1:] {
            if !oneDiff(prev, str) {
                fail = true
            }
            prev = str
        }
        if !fail { 
            return true
        }
    }
    return false
}

func permute(arr []string, l, r int) {
    if l == r {
        cp := make([]string, len(arr))
        copy(cp, arr)
        perms = append(perms, cp)
        return 
    }
    for i := l;i <= r; i++ {
        swap(arr, l, i)
        permute(arr, l+1, r)
        swap(arr, l, i)
    } 
}

func swap(arr []string, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func oneDiff(a, b string) bool {
    dif := 0
    for n := range a {
        if a[n] != b[n] {
            dif++
        }
    }
    return dif == 1
}


func SolutionExtractEachKth(inputArray []int, k int) []int {
	result := []int{}
    for index, v := range inputArray {
        if (index+1) % k != 0 {
            result = append(result, v)
        }       
    }
    return result
}

func SolutionFirstDigit(inputString string) string {
	for i := 0; i < len(inputString); i++ {
		if inputString[i] >= 48 && inputString[i] <= 57 {
			return string(inputString[i])
		}
	}
	return ""
}

func SolutionDifferentSymbolsNaive(s string) int {
	mapFrequent := map[string]bool{}

	for i := 0; i < len(s); i++ {
		mapFrequent[string(s[i])] = true
	}

	return len(mapFrequent)
}

func SolutionArrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := 0
	for i := 0; i < len(inputArray) - k + 1; i++ {
		consecutiveSlice := inputArray[i:i+k]
		sum := 0 

		for j := 0; j < len(consecutiveSlice); j++ {
			sum += consecutiveSlice[j]
		}

		if max < sum {
			max = sum
		}
	}
	
	return max
}

func SolutionGrowingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	height := 0
	days := 0
	for height < desiredHeight {
		height += upSpeed

		if height >= desiredHeight {
			days++
			break
		}

		height -= downSpeed
		days++

		fmt.Println(days, height)
	}

	fmt.Println(days)

	return days
}

func SolutionKnapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if weight1 + weight2 <= maxW {
		return value1 + value2
	}

	if weight1 > maxW && weight2 <= maxW {
		return value2
	}

	if weight2 > maxW && weight1 <= maxW {
		return value1
	}

	if weight1 > maxW && weight2 > maxW {
		return 0
	}

	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}

func SolutionLongestDigitsPrefix(inputString string) string {
	result := ""
	 for _, r := range inputString {
        if r >= 48 && r <= 57 {
            result += string(r)
        } else {
            break;
        }
    }
	return result
}
