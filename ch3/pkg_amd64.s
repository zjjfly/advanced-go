#include "textflag.h"
//用汇编声明一个int
GLOBL ·Id(SB),$8

DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

//用汇编声明一个string
//NOPTR是必须的,因为汇编声明的只是内存块,而go的垃圾回收器需要知道对象是否包含指针
//而NOPTR就是表明NameData不包含指针
GLOBL ·NameData(SB),NOPTR,$8
DATA ·NameData(SB)/8,$"gopher"
//Name是一个StringHeader
GLOBL ·Name(SB),$16
DATA ·Name+0(SB)/8,$·NameData(SB)
DATA ·Name+8(SB)/8,$6

//还有一种解决方案是在go代码中声明一个Name1Data
GLOBL ·Name1Data(SB),NOPTR,$8
DATA ·Name1Data(SB)/8,$"123456"

GLOBL ·Name1(SB),$16
DATA ·Name1+0(SB)/8,$·Name1Data(SB)
DATA ·Name1+8(SB)/8,$6

//最后还有一种方案,把实际的字符串和StringHeader定义在一起
GLOBL ·Name2(SB),$24
DATA ·Name2+0(SB)/8,$·Name2+16(SB)
DATA ·Name2+8(SB)/8,$6
DATA ·Name2+16(SB)/8,$"gopher"

//使用汇编定义函数
TEXT ·asmPrint(SB), $16-0
    MOVQ ·str+0(SB), AX
    MOVQ AX, 0(SP)
    MOVQ ·str+8(SB), BX
    MOVQ BX, 8(SP)
    CALL ·print(SB)
    RET
