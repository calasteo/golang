package main

//basic package
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//initiate type of struct
type locker struct {
	no             int
	identityType   string
	identityNumber string
}

//initiate global variable
//i is total locker
var i = 0

//set array of locker
var lockers []locker
var colorRed = "\033[31m"
var colorReset = "\033[0m"
var colorGreen = "\033[32m"

func main() {
	//call main menu
	displayMain(false)
}

func displayMain(clear bool) {

	if clear {
		// clear console for better reading
		clearConsole()
	}
	//set variable to get input data
	var input int
	separator()
	fmt.Println("Locker Simple Application")
	separator()
	if i > 0 {
		fmt.Println("1. Init Locker [X]")
	} else {
		fmt.Println("1. Init Locker")
	}

	fmt.Println("2. Input Data")
	fmt.Println("3. Remove Data")
	fmt.Println("4. Find Data by ID")
	fmt.Println("5. Search Data by Type")
	fmt.Println("6. Check Status Locker")
	fmt.Println("7. Exit Application")
	separator()
	fmt.Print("Your Input (1-7): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())
	if input > 0 && input <= 7 {
		switch input {
		case 1:
			initLocker()
			break
		case 2:
			inputData()
			break
		case 3:
			removeData()
			break
		case 4:
			findDataById()
			break
		case 5:
			searchByType()
			break
		case 6:
			checkStatus()
			break
		case 7:
			exitApplication()
			break
		}
	} else {
		fmt.Println(colorRed, "Wrong Input", colorReset)
		displayMain(false)
	}

}

//to set up initial total locker
func initLocker() {
	if i > 0 {
		fmt.Println()
		fmt.Println(colorRed, "Locker already set", colorReset)
		fmt.Println()
		returnBack()
	} else {
		var canGo = false
		var question = "Please Input number of locker : "
		fmt.Println("Init Locker")
		separator()
		//if input is invalid then loop
		for !canGo {
			fmt.Print(question)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			i, _ = strconv.Atoi(scanner.Text())
			if i <= 0 {
				canGo = false
				question = "Wrong input, please input number of locker : "
			} else {
				canGo = true
				fmt.Println()
				fmt.Println(colorGreen, "Total Locker : "+strconv.Itoa(i), colorReset)
				fmt.Println()
				returnBack()
			}
		}
	}
}
func inputData() {
	//check if locker already initialized
	checkInit()
	//check if locker are full
	if len(lockers) >= i {
		fmt.Println()
		fmt.Println(colorRed, "Locker full", colorReset)
		fmt.Println()
		returnBack()
	} else {
		fmt.Println("Input Data Locker")
		separator()
		var iType, iNumber string
		fmt.Print("Input Identity Type : ")
		_, _ = fmt.Scan(&iType)
		fmt.Print("Input Identity Number : ")
		_, _ = fmt.Scan(&iNumber)
		//check if data already exist in locker
		found, index, lowest := checkExists(iType, iNumber)
		if found {
			fmt.Println(colorRed, "Data exists on locker "+strconv.Itoa(index), colorReset)
			returnBack()
		} else {
			//input new data
			dataAppend := locker{
				no:             lowest,
				identityNumber: iNumber,
				identityType:   iType,
			}
			lockers = append(lockers, dataAppend)
			fmt.Println(colorGreen,
				"Identity "+dataAppend.identityType+" with ID "+dataAppend.identityNumber+" inserted for Locker "+strconv.Itoa(lowest),
				colorReset)
			returnBack()
		}
	}

}
func findDataById() {
	checkInit()
	var iNumber string
	fmt.Print("Input Identity Number : ")
	_, _ = fmt.Scan(&iNumber)
	var arr []string
	//compare the input data from array , if match append to new array
	for a := range lockers {
		if strings.ToLower(lockers[a].identityNumber) == strings.ToLower(iNumber) {
			arr = append(arr, strconv.Itoa(lockers[a].no))
		}
	}
	if len(arr) > 0 {
		//data found
		fmt.Println()
		fmt.Println(colorGreen, "Data Identify Number "+iNumber+" found in locker "+strings.Join(arr, ", "), colorReset)
		fmt.Println()
		returnBack()
	} else {
		//data not found
		fmt.Println()
		fmt.Println(colorRed, "Data not found", colorReset)
		fmt.Println()
		returnBack()
	}
}
func searchByType() {
	checkInit()
	var iType string
	fmt.Print("Input Identity Type : ")
	_, _ = fmt.Scan(&iType)
	var arr []string
	//compare the input data identity type from array , if match append to new array
	for a := range lockers {
		if strings.ToLower(lockers[a].identityType) == strings.ToLower(iType) {
			arr = append(arr, lockers[a].identityNumber)
		}
	}
	if len(arr) > 0 {
		//data found
		fmt.Println()
		fmt.Println(colorGreen, "Data Identify Type "+iType+" found : "+strings.Join(arr, ", "), colorReset)
		fmt.Println()
		returnBack()
	} else {
		//data not found
		fmt.Println()
		fmt.Println(colorRed, "Data not found", colorReset)
		fmt.Println()
		returnBack()
	}
}

func removeData() {
	checkInit()
	if len(lockers) == 0 {
		fmt.Println(colorRed, "Empty Data", colorReset)
	} else {
		var no int
		fmt.Print("Input Locker Number : ")
		_, _ = fmt.Scan(&no)
		var arr []locker
		//check if data not match with the locker number, then set to new array
		for a := 0; a < len(lockers); a++ {
			if lockers[a].no != no {
				arr = append(arr, lockers[a])
			}
		}
		if len(arr) == len(lockers) {
			fmt.Println(colorRed, "Data Not Found", colorReset)
		} else {
			//replace array of locker with new array
			lockers = arr
			fmt.Println(colorGreen, "Data Locker number "+strconv.Itoa(no)+" has been deleted", colorReset)
		}
	}

	returnBack()
}
func checkStatus() {
	checkInit()
	separator()
	fmt.Println()
	fmt.Println("Data Locker")
	separator()
	//sorting data from locker number
	sorting()
	if len(lockers) == 0 {
		fmt.Println("                    No Data                       ")
		fmt.Println("--------------------------------------------------")
	} else {
		for a := 0; a < len(lockers); a++ {
			fmt.Println("No Loker        :" + strconv.Itoa(lockers[a].no))
			fmt.Println("Identity Type   :" + lockers[a].identityType)
			fmt.Println("Identity Number :" + lockers[a].identityNumber)
			fmt.Println("--------------------------------------------------")
		}
	}

	returnBack()
}

func exitApplication() {
	fmt.Println(colorRed, "Exiting .. Thank you")
	os.Exit(0)
}

func clearConsole() {
	for a := 0; a < 30; a++ {
		fmt.Println()
	}
}

func separator() {
	fmt.Println("==================================================")
}

//check if data locker already exists
func checkExists(iType string, iNumber string) (bool, int, int) {
	found := false
	var foundNumber, lowestNumber int
	lowestNumber = 0
	//sorting data from locker number
	sorting()
	setup := false
	for c := 0; c < len(lockers); c++ {
		//if identity type and identity number, then set if data found, and set the locker number
		if strings.ToLower(lockers[c].identityNumber) == strings.ToLower(iNumber) && strings.ToLower(lockers[c].identityType) == strings.ToLower(iType) {
			found = true
			foundNumber = lockers[c].no
		}
		//because array index starting with 0,
		//then if lowestNumber have not been set up,
		//and the current index of array +1  not the locker number of current array
		//then lowest number must be the current index of array +1
		if lowestNumber == 0 && lockers[c].no != (c+1) {
			lowestNumber = c + 1
			setup = true
		}
	}
	//if empty lockers
	if len(lockers) == 0 {
		lowestNumber = 1
	} else if !setup {
		//if lowest number has not been setup from looping above
		lowestNumber = len(lockers) + 1
	}
	return found, foundNumber, lowestNumber
}

func sorting() {
	//sorting by locker number
	sort.Slice(lockers, func(a, b int) bool {
		return lockers[a].no < lockers[b].no
	})
}

//display return menu after an action
func returnBack() {
	var input int
	var canGo = false
	var question = "Your Input (1-2): "
	fmt.Println("1. Back to Menu")
	fmt.Println("2. Exit Application")
	for !canGo {
		fmt.Print(question)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input, _ = strconv.Atoi(scanner.Text())
		if input > 0 && input <= 2 {
			canGo = true
			switch input {
			case 1:
				displayMain(false)
				break
			case 2:
				exitApplication()
				break
			}
		} else {
			question = "Wrong input, please input (1-2): "
		}
	}
}

//check whether locker has been initialized or not
func checkInit() {
	if i <= 0 {
		fmt.Println()
		fmt.Println(colorRed, "Locker not ready initialized", colorReset)
		fmt.Println()
		returnBack()
	}
}
