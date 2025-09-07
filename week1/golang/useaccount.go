package account

import (
	account "_/C_/Users/au649790/OneDrive_-_Aarhus_universitet/Desktop/DISSY/week1/golang"
	"fmt"
)

func main() {
	l := account.MakeLedger()
	t := account.Transaction{To: "Jesper", From: "Ivan", Amount: 1000000}
	l.Transaction(&t)
	fmt.Println(l)
}
