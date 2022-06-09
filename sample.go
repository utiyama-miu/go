package main

import (
	"fmt"
	"math/rand"
	"time"
)

const CARD_NUM int = 5
const HAND_NUM int = 52

var card [CARD_NUM]string

type Pile struct {
	next  [HAND_NUM]int
	index int
}

func (p *Pile) getNext() int {
	p.index++
	return p.next[p.index]
}

func (p *Pile) init_card() {
	for i := 0; i < CARD_NUM; i++ {
		p.next[i] = i
	}
	p.index = -1
}

func (p *Pile) shuffle() {
	for i := CARD_NUM - 1; i > 0; i-- {
		j := rand.Intn(i)
		tmp := p.next[i]
		p.next[i] = p.next[j]
		p.next[j] = tmp
	}
}

type Game struct {
	name  string
	Pile  Pile
	hands [CARD_NUM]int
}
type Hands struct {
	hand [HAND_NUM]int
}

func (h *Hands) setup(p *Pile) {
	for i := 0; i < len(h.hand); i++ {
		h.hand[i] = p.getNext()
	}
}
func (h *Hands) sort() {
	for i := 0; i < HAND_NUM-1; i++ {
		for j := HAND_NUM - 1; j > i; j-- {
			if h.hand[j] < h.hand[j-1] {
				tmp := h.hand[j]
				h.hand[j] = h.hand[j-1]
				h.hand[j-1] = tmp
			}
		}
	}
}
func (h *Hands) show() {
	fmt.Println(card[h.hand[0]] + card[h.hand[1]] + card[h.hand[2]] + card[h.hand[3]] + card[h.hand[4]])

}
func (h *Hands) check() string {
	return "ブタ"
}

func (g Game) exec() {
	rand.Seed(time.Now().UnixMicro())
	g.Pile.init_card()
	g.Pile.shuffle()

	var hand [HAND_NUM]int = [HAND_NUM]int{g.Pile.getNext(), g.Pile.getNext(), g.Pile.getNext(), g.Pile.getNext(), g.Pile.getNext()}
	fmt.Println(">>", g.name, "START<<")
	//g.hands.setup(&g.Pile)
	fmt.Println(" 1 2 3 4 5 ")
	var c int = 0
	var cha int = 0
	for {
		fmt.Print("何枚交換しますか(0-", HAND_NUM, ")")
		fmt.Scan(&c)
		if 0 <= c && c <= HAND_NUM {
			cha = c
			break
		}
	}
	var ex [5]int
	for i := 0; i < cha; i++ {
		for {
			var result bool = true
			fmt.Print("どれを交換しますか(1-5)")
			fmt.Scan(&c)
			ex[i] = c - 1
			if c < 1 || 5 < c {
				result = false //resultを使って処理を分ける
			}
			//前回までの選択被ったら
			for j := 0; j < i; j++ {
				if ex[j] == ex[i] {
					result = false
				}
			}
			if result {
				break
			}
		}
	}
	for i := 0; i < cha; i++ {
		hand[ex[i]] = g.Pile.getNext()
	}
	//g.hands.show()
	fmt.Println(card[hand[0]] + card[hand[1]] + card[hand[2]] + card[hand[3]] + card[hand[4]])
	fmt.Println(" 1 2 3 4 5 ")
	fmt.Print("です。") //g.hands.check(),
	for i := 0; i < HAND_NUM-1; i++ {
		for j := HAND_NUM - 1; j > i; j-- {
			if hand[j] < hand[j-1] {
				tmp := hand[j]
				hand[j] = hand[j-1]
				hand[j-1] = tmp
			}
		}
	}

	fmt.Println(card[hand[0]] + card[hand[1]] + card[hand[2]] + card[hand[3]] + card[hand[4]])
	fmt.Println("1  2  3  4  5")
	fmt.Println(">>", g.name, "OVER<<")
}
func init_card() {
	var sult [4]string = [4]string{"♠", "♦", "♥", "☘"}
	var num [13]string = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "j", "Q", "K"}
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(sult); j++ {
			card[i*4+j] = "[" + sult[j] + "][" + num[i] + "]"
			fmt.Print(card[i*4+j])
		}
	}
}
func main() {
	for {
		fmt.Printf("開始:1 終了:0")
		var cmd int = 0
		fmt.Scan(&cmd)
		if cmd == 0 {
			break
		}
		// Game{"GoPoker", Pile{}}.exec()
		init_card()
	}
	fmt.Printf("終了")
}
