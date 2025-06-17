package util

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type overwrite struct {
	base
	str string
}

func (o overwrite) describe() string {
	return fmt.Sprintf("overwrite with num=%v, str=%v", o.num, o.str)
}

func TryCallMethodDirectly() {
	co := container{
		base: base{num: 7},
		str:  "hello",
	}

	// ── 1. 直接呼叫 ─────────────────────
	fmt.Println(co.describe())      // → container num=7 str=hello
	fmt.Println(co.base.describe()) // → base num=7

	// ── 2. Method Value ───────────────
	mv := co.describe // 型別是 func() string，接收者已經綁定到 co
	fmt.Println(mv()) // → container num=7 str=hello

	// ── 3. Method Expression ──────────
	me1 := (container).describe // func(container) string
	fmt.Println(me1(co))        // → container num=7 str=hello

	me2 := (base).describe    // func(base) string
	fmt.Println(me2(co.base)) // → base num=7
}

func TryEmbedding() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	ov := overwrite{
		base: base{
			num: 2,
		},
		str: "overwrite name",
	}

	var dov describer = ov
	fmt.Println("overwrite describer:", dov.describe())
}
