package controllers

import (
	"github.com/astaxie/beegae"
	"runtime"
	"strconv"
)

type vec struct {
	x int
	y int
}

type result struct {
	count int
	flag  bool
	x     int
	y     int
}

type BoardController struct {
	beegae.Controller
}

func getBoard(str string) [][]int {
	board := make([][]int, 8)
	for i := range board {
		board[i] = make([]int, 8)
	}
	var index = 0
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			s := str[index : index+1]
			if s == "1" {
				board[y-1][x-1] = 1
			} else if s == "2" {
				board[y-1][x-1] = 2
			} else {
				board[y-1][x-1] = 0
			}
			index++
		}
	}
	return board
}

func toString(board [][]int) string {
	str := ""
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			s := board[y-1][x-1]
			str += strconv.Itoa(s)
		}
	}
	return str
}

func getKoma(board [][]int, x int, y int) int {
	if x <= 0 {
		return -1
	}
	if y <= 0 {
		return -1
	}
	if y > len(board) {
		return -1
	}
	if x > len(board[0]) {
		return -1
	}
	return board[y-1][x-1]
}

func search(board [][]int, x int, y int, v vec, my int) bool {
	x += v.x
	y += v.y
	if getKoma(board, x, y) == 0 {
		return false
	}
	if getKoma(board, x, y) == my {
		return false
	}
	if getKoma(board, x, y) == -1 {
		return false
	}
	for i := 0; i < 8; i++ {
		x += v.x
		y += v.y
		if getKoma(board, x, y) == my {
			return true
		}
		if getKoma(board, x, y) == 0 {
			return false
		}
		if getKoma(board, x, y) == -1 {
			return false
		}
	}
	return false
}

func canPut(board [][]int, x int, y int, my int) bool {
	if getKoma(board, x, y) != 0 {
		return false
	}
	flag := false
	directions := []vec{vec{0, -1}, vec{1, -1}, vec{1, 0}, vec{1, 1},
		vec{0, 1}, vec{-1, 1}, vec{-1, 0}, vec{-1, -1}}
	for d := range directions {
		flag = flag || search(board, x, y, directions[d], my)
	}
	return flag
}

func reverse(board [][]int, x int, y int, v vec, my int) [][]int {
	if search(board, x, y, v, my) {
		for i := 0; i < 8; i++ {
			x += v.x
			y += v.y
			if getKoma(board, x, y) == my {
				break
			}
			board[y-1][x-1] = my
		}
	}
	return board
}

func put(board [][]int, x int, y int, my int) [][]int {
	board[y-1][x-1] = my
	directions := []vec{vec{0, -1}, vec{1, -1}, vec{1, 0}, vec{1, 1},
		vec{0, 1}, vec{-1, 1}, vec{-1, 0}, vec{-1, -1}}
	for d := range directions {
		board = reverse(board, x, y, directions[d], my)
	}
	return board
}

func clone(org [][]int) [][]int {
	board := make([][]int, 8)
	for i := range board {
		board[i] = make([]int, 8)
	}
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			board[y-1][x-1] = org[y-1][x-1]
		}
	}
	return board
}

func count(board [][]int, my int) int {
	sum := 0
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			if board[y-1][x-1] == my {
				sum++
			}
		}
	}
	return sum
}

func assumePutSub(ch chan result, org [][]int, x int, y int, my int, n int) {
	board := clone(org)
	put(board, x, y, my)
	c := 0
	f := false
	if n == 0 || runtime.NumGoroutine() > 1000 {
		c = count(board, my)
	} else {
		v := vec{-1, -1}
		v, _, c, f = cal(board, 3-my, n-1)
		if v.x == -1 {
			c, f = assumePut(board, my, n-1)
		}
	}
	ch <- result{c, f, x, y}
}

func assumePut(org [][]int, my int, n int) (int, bool) {
	ch := make(chan result)
	total := 0
	priority := []vec{vec{1, 1}, vec{1, 8}, vec{8, 1}, vec{8, 8}}
	for p := range priority {
		vp := priority[p]
		if canPut(org, vp.x, vp.y, my) {
			total++
			go assumePutSub(ch, org, vp.x, vp.y, my, n)
		}
	}
	if total > 0 {
		cot := 0
		for i := 0; i < total; i++ {
			r := <-ch
			c := r.count
			if c > cot {
				cot = c
			}
		}
		return cot, true
	}
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			if !canPut(org, x, y, my) {
				continue
			}
			total++
			go assumePutSub(ch, org, x, y, my, n)
		}
	}
	if total == 0 {
		v, _, c, f := cal(org, 3-my, n-1)
		if v.x == -1 {
			c = count(org, my)
			f = false
		}
		return c, f
	}
	cot := -1
	flag := false
	for i := 0; i < total; i++ {
		r := <-ch
		c, f := r.count, r.flag
		flag = flag || f
		if c > cot {
			cot = c
		}
	}
	return cot, flag
}

func exist(board [][]int, my int) bool {
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			if canPut(board, x, y, my) {
				return true
			}
		}
	}
	return false
}

func cal(board [][]int, my int, n int) (vec, string, int, bool) {
	log := ""
	priority := []vec{vec{1, 1}, vec{1, 8}, vec{8, 1}, vec{8, 8}}
	for p := range priority {
		vp := priority[p]
		if canPut(board, vp.x, vp.y, my) {
			return vec{vp.x, vp.y}, log, 0, false
		}
	}
	ch := make(chan result)
	total := 0
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			if !canPut(board, x, y, my) {
				continue
			}
			total++
			go func(ch chan result, board [][]int, x int, y int, my int, n int) {
				vb := clone(board)
				put(vb, x, y, my)
				c, f := assumePut(vb, 3-my, n)
				ch <- result{c, f, x, y}
			}(ch, board, x, y, my, n)
		}
	}

	fx := -1
	fy := -1
	fcount := 65
	fflag := true
	for i := 0; i < total; i++ {
		r := <-ch
		count, flag, x, y := r.count, r.flag, r.x, r.y
		if fflag {
			if flag == false {
				fx = x
				fy = y
				fcount = count
				fflag = false
				continue
			}
		}
		if (fflag == false) && (flag == true) {
			continue
		}
		if count < fcount {
			fx = x
			fy = y
			fcount = count
		}
	}
	// log = strconv.Itoa(fcount)
	return vec{fx, fy}, log, fcount, fflag
}

func (this *BoardController) Get() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	var board [][]int
	log := ""
	v := this.GetSession("board")
	p := vec{-1, -1}
	if v == nil || this.GetString("reset") == "true" {
		board = make([][]int, 8)
		for i := range board {
			board[i] = make([]int, 8)
		}
		board[3][3] = 1
		board[3][4] = 2
		board[4][3] = 2
		board[4][4] = 1
		this.SetSession("board", toString(board)+"1")
	} else {
		board = getBoard(v.(string))
		param := this.GetString("p")
		if len(param) == 3 {
			x, _ := strconv.Atoi(param[1:2])
			y, _ := strconv.Atoi(param[0:1])
			my, _ := strconv.Atoi(param[2:3])
			if canPut(board, x, y, my) {
				put(board, x, y, my)
				for {
					// c := count(board, my) + count(board, 3-my)
					p, _, _, _ = cal(board, 3-my, 1)
					// this.Data["log"] = log
					// if flag {
					// 	this.Data["log"] = "Sumi!"
					// }
					if p.x == -1 {
						if !exist(board, my) {
							log += " End!"
						} else {
							log += " Pass!"
						}
						break
					}
					put(board, p.x, p.y, 3-my)
					if !exist(board, my) {
						if !exist(board, 3-my) {
							log += " End!"
							break
						} else {
							log += " You Pass!"
							continue
						}
					}
					break
				}
			} else {
				log += "Can't put"
			}
		}
		this.SetSession("board", toString(board)+"1")
	}
	log += " " + strconv.Itoa(count(board, 1)) + ":" + strconv.Itoa(count(board, 2))

	this.Data["log"] = log
	for y := 1; y <= 8; y++ {
		for x := 1; x <= 8; x++ {
			key := "B" + strconv.Itoa(y) + strconv.Itoa(x)
			if p.x == x && p.y == y {
				this.Data[key] = "◎"
			} else if board[y-1][x-1] == 1 {
				this.Data[key] = "●"
			} else if board[y-1][x-1] == 2 {
				this.Data[key] = "○"
			} else {
				this.Data[key] = "　"
			}
		}
	}
	this.TplName = "board.tpl"
}
