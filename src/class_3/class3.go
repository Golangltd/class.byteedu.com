package main

import (
	"fmt"
)

/*
第三节  初始化游戏牌型、结构定义
http://www.byteedu.com/thread-1095-1-1.html
(出处: www.ByteEdu.Com)
*/

// 棋盘的定义
/*
   ----------------------------------------------------------------------------------
   |                                                                                          |
   |        [0,0]01        [1,0]02        [2,0]03        [3,0]04                |
   |                                                                                          |
   |                                                                                          |
   |        [0,1]05        [1,1]06        [2,1]07        [3,1]08                |
   |                                                                                          |
   |                                                                                          |
   |        [0,2]09        [1,2]10        [2,2]11        [3,2]12                |
   |                                                                                          |
   |                                                                                          |
   |        [0,3]13        [1,3]14        [2,3]15        [3,3]16                |
   |                                                                                          |
   ----------------------------------------------------------------------------------
*/

// 斗兽棋的棋子类型
const (
	DSQINIT_QZ = iota // DSQ_QZ == 0
	Elephant          // elephant == 1   大象
	Lion              // lion == 2       狮子
	Tiger             // tiger == 3      老虎
	Leopard           // leopard == 4    豹子
	Wolf              // wolf == 5       狼
	Dog               // dog == 6        狗
	Cat               // cat == 7        猫
	Mouse             // mouse == 8      老鼠
)

// 棋子的行动的方向
const (
	FANGXIANGINIT = iota // FANGXIANGINIT == 0
	UP                   // UP            == 1
	DOWN                 // DOWN          == 2
	LEFT                 // LEFT          == 3
	RIGHT                // RIGHT         == 4
)

// 棋子的攻击方式
const (
	ITYPEINIY    = iota // ITYPEINIY == 0
	MOVE                // MOVE == 1         正常移动
	DISAPPEAR           // DISAPPEAR == 2    自残
	ALLDISAPPEAR        // ALLDISAPPEAR == 3 同归于尽
	BEAT                // BEAT == 4         击败对方
	TEAMMATE            // TEAMMATE == 5     队友
	MOVESUCC            // MOVESUCC == 6     移动成功
	MOVEFAIL            // MOVEFAIL == 7     移动失败
	DATAERROR           // DATAERROR == 8    数据错误    玩家的棋子已经被吃掉不存在了
	DATANOEXIT          // DATANOEXIT == 9   数据不存在  棋子的数据大于 16或者小于0
)

var DSQ_qi = []int{ // 1-8 A ;9-16 B ; 17 未翻牌; 18 已翻牌
	Proto2.Elephant, Proto2.Lion, Proto2.Tiger, Proto2.Leopard, Proto2.Wolf, Proto2.Dog, Proto2.Cat, Proto2.Mouse,
	Proto2.Mouse + Proto2.Elephant, Proto2.Mouse + Proto2.Lion, Proto2.Mouse + Proto2.Tiger, Proto2.Mouse + Proto2.Leopard,
	Proto2.Mouse + Proto2.Wolf, Proto2.Mouse + Proto2.Dog, Proto2.Mouse + Proto2.Cat, 2 * Proto2.Mouse}

func init() {
	InitDSQ(DSQ_qi)
	return
}

func main() {
	return
}

// 初始化牌型
func InitDSQ(data1 []int) [4][4]int {

	data, erdata, j, k := data1, [4][4]int{}, 0, 0

	for i := 0; i < Proto2.Mouse*2; i++ {
		icount := util.RandInterval_LollipopGo(0, int32(len(data))-1)
		//fmt.Println("随机数：", icount)
		if len(data) == 1 {
			erdata[3][3] = data[0]
		} else {
			//------------------------------------------------------------------
			if int(icount) < len(data) {
				erdata[j][k] = data[icount]
				k++
				if k%4 == 0 {
					j++
					k = 0
				}
				data = append(data[:icount], data[icount+1:]...)
			} else {
				erdata[j][k] = data[icount]
				k++
				if k%4 == 0 {
					j++
					k = 0
				}
				data = data[:icount-1]
			}
			//------------------------------------------------------------------
		}
		//fmt.Println("生成的数据", erdata)
	}

	return erdata
}
