package rabbit

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func recovery(writer http.ResponseWriter) {
	if err := recover(); err != nil {
		message := fmt.Sprintf("%s", err)
		var pcs [32]uintptr
		n := runtime.Callers(3, pcs[:])
		var str strings.Builder
		str.WriteString(message + "\nTraceback:")
		for _, pc := range pcs[:n] {
			fn := runtime.FuncForPC(pc)
			file, line := fn.FileLine(pc)
			str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
		}
		log.Printf("%s\n\n", str.String())

		http.Error(writer, message, http.StatusInternalServerError)
	}
}
