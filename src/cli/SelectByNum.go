package cli

import (
	"fmt"
	"strconv"
)

func SelectByNum(selectionObj string, maxNum int, listFunc func(), itemlist []string) (int, error) {
	var tmp string

	fmt.Printf("Select %s by entering corresponding number\n\n", selectionObj)
	if listFunc != nil {
		listFunc()
	} else if len(itemlist) > 0 {
		for i, item := range itemlist {
			fmt.Printf("%d. %s\n", i, item)
		}
	}

	fmt.Printf("\n: ")
	_, err := fmt.Scanln(&tmp)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, err
	}

	if num >= maxNum || num < 0 {
		return 0, fmt.Errorf("invalid assembly number")
	}
	return num, nil
}
