package helper

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

func SendError(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	io.WriteString(w, errMsg)
}

func SliceToString(slice []int) string {
	var strSlice []string
	for _, num := range slice {
		strSlice = append(strSlice, strconv.Itoa(num))
	}

	resultStr := strings.Join(strSlice, "")

	return resultStr
}
