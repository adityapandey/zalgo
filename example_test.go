package zalgo_test

import (
	"math/rand"
	"os"

	"github.com/adityapandey/zalgo"
)

func ExampleWriter() {
	z := zalgo.NewWriter(os.Stdout)
	z.Opt.Up = true
	z.Opt.Mid = true
	z.Opt.Down = true
	mzga := []byte("MAKE ZALGO GREAT AGAIN!!!")
	rand.Seed(666)
	z.Write(mzga)
	// Output: M̩͓̯͎̟̳͌̑ͨ̉Á̠̮̹̙̾ͪͫ̚͘Ḳ͈ͥ͒̂͊͗́E̦̩̬͓̒ͤ̌̈́ͥ ̛̹̝̠̬ͭ͛̄ͅZ̧̞̓ͤͪ͐ͣÁ̫̪̣̤̣̗̆ͩLͦ͒̈̌͐̆̚G̭̯̞̻̹̭̟Ȯ̹̯̺̜̈ͬ͜ͅͅ ̡̯͔̜̩̝G͍͙ͅR̛͕͈̹̝̼ͫ̎ͬͫĚ̵̖͊̍A̟̦̙̯̪̥̿͘Ṫ̛̄̄̿ ̮̘A͛G̳͕͕̲̠͝A̸̮̖̙͙̞͕̬ͣ̚̚Ī̦͔̻ͫ̅̃̚N͕͔̦͓̻̳͉̍ͩͦ̍̄!̻ͮͧ̅̅̉͊̇!̨͚̩͊͆!̠ͣͬ͐ͪ̌̇̋
}
