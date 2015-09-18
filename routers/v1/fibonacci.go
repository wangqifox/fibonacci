package v1

import (
	"github.com/Unknown/com"
	"github.com/wangqifox/fibonacci/middleware"
	"github.com/wangqifox/fibonacci/models"
)

func Fibonacci(ctx *middleware.Context) {
	number := ctx.QueryInt("number")
	result := models.fibonacci(number)
	if result < 0 {
		ctx.JSON(422, map[string]string{
			"error": "Unprocessable number: " + com.ToStr(number),
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"result": result,
	})
}

// func getFibonacci(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	number, err := strconv.Atoi(r.FormValue("number"))
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	fmt.Println(number)
// 	result := fibonacci(number)
// 	if result < 0 {
// 		http.Error(w, "结果溢出", 500)
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%d", result))
// }
