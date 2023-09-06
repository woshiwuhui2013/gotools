package httptest

import (
	"fmt"
	"net/http"
)

func HttpHandle() error {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello respond"))
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
