package day11

import (
	"fmt"
	"math/big"
	"sort"
)

func bigger(ints []int) []*big.Int {
	bigs := make([]*big.Int, len(ints))
	for i, num := range ints {
		bigs[i] = big.NewInt(int64(num))
	}
	return bigs
}

func bigg(x int) *big.Int {
	return big.NewInt(int64(x))
}
func setupMonkeies() []*Monkey {
	Troop := make([]*Monkey, 8)
	counter := 0
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{56, 56, 92, 65, 71, 61, 79}),
		operation: "*",
		value:     bigg(7),
		test:      bigg(3),
		friend:    3,
		rival:     7,
	}
	counter++

	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{61, 85}),
		operation: "+",
		value:     bigg(5),
		test:      bigg(11),
		friend:    6,
		rival:     4,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{54, 96, 82, 78, 69}),
		operation: "^",
		value:     bigg(0),
		test:      bigg(7),
		friend:    0,
		rival:     7,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{57, 59, 65, 95}),
		operation: "+",
		value:     bigg(4),
		test:      bigg(2),
		friend:    5,
		rival:     1,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{62, 67, 80}),
		operation: "*",
		value:     bigg(17),
		test:      bigg(19),
		friend:    2,
		rival:     6,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{91}),
		operation: "+",
		value:     big.NewInt(int64(7)),
		test:      big.NewInt(int64(5)),
		friend:    1,
		rival:     4,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{79, 83, 64, 52, 77, 56, 63, 92}),
		operation: "+",
		value:     big.NewInt(int64(6)),
		test:      big.NewInt(int64(17)),
		friend:    2,
		rival:     0,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{50, 97, 76, 96, 80, 56}),
		operation: "+",
		value:     big.NewInt(int64(3)),
		test:      big.NewInt(int64(13)),
		friend:    3,
		rival:     5,
	}
	return Troop
}
func setupTestMonkeies() []*Monkey {
	Troop := make([]*Monkey, 4)
	counter := 0
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{79, 98}),
		operation: "*",
		value:     big.NewInt(int64(19)),
		test:      big.NewInt(int64(23)),
		friend:    2,
		rival:     3,
	}
	counter++

	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{54, 65, 75, 74}),
		operation: "+",
		value:     big.NewInt(int64(6)),
		test:      big.NewInt(int64(19)),
		friend:    2,
		rival:     0,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{79, 60, 97}),
		operation: "^",
		value:     big.NewInt(int64(0)),
		test:      big.NewInt(int64(13)),
		friend:    1,
		rival:     3,
	}
	counter++
	Troop[counter] = &Monkey{
		index:     counter,
		items:     bigger([]int{74}),
		operation: "+",
		value:     big.NewInt(int64(3)),
		test:      big.NewInt(int64(17)),
		friend:    0,
		rival:     1,
	}
	return Troop
}

func Run() {
	num := 11
	//decided to skip input parsing this time... might come back to it... TODO
	var total *big.Int
	total = big.NewInt(int64(1))

	Troop = setupMonkeies()
	// Troop = setupTestMonkeies()

	mn := big.NewInt(1)
	for _, m := range Troop {
		t := &big.Int{}
		t.GCD(nil, nil, mn, m.test)
		mn.Mul(mn, m.test)
		mn.Div(mn, t)
	}
	MagicNumber = mn

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(Troop); j++ {
			Troop[j].TakeTurn()
		}
	}
	mb := make([][]*big.Int, len(Troop))
	for i, m := range Troop {
		mb[i] = bigger([]int{m.index, m.count})
	}
	sort.Slice(mb, func(i, j int) bool {
		return (*mb[i][1]).Cmp(mb[j][1]) > 0
	})

	total = total.Mul(mb[0][1], mb[1][1])
	fmt.Println(fmt.Sprintf("\nDay %d: %d\n", num, total))
}

var Troop []*Monkey
var MagicNumber *big.Int

type Monkey struct {
	index     int
	items     []*big.Int
	operation string
	value     *big.Int
	test      *big.Int
	friend    int
	rival     int
	count     int
}

func CreateMonkey() *Monkey {
	return &Monkey{}
}

func (m *Monkey) TakeTurn() {
	for _, item := range m.items {
		new := m.Inspect(item)
		new.Mod(new, MagicNumber)
		t := &big.Int{}
		t.Mod(new, m.test)
		if t.Sign() == 0 {
			Troop[m.friend].Catch(new)
		} else {
			Troop[m.rival].Catch(new)
		}
	}
	m.items = make([]*big.Int, 0)
}

func (m *Monkey) Catch(x *big.Int) {
	m.items = append(m.items, x)
}

func (m *Monkey) Inspect(item *big.Int) *big.Int {
	m.count++
	switch m.operation {
	case "/":
		return item.Div(item, m.value)
	case "*":
		return item.Mul(item, m.value)
	case "+":
		return item.Add(item, m.value)
	case "-":
		return item.Sub(item, m.value)
	case "^":
		return item.Mul(item, item)
	default:
		fmt.Println("How did you get here?")
		return nil
	}
}
