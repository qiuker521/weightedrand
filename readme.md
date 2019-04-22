# Weighted Random

Thanks to [weightedrand](https://github.com/mroth/weightedrand)

### Origin

Use it to make decisions like "choose what upstream of a proxy by weight".

为了实现一个类似nginx选择upstream一样，带有权重的随机选择器。

### Perface

We have to methods of a pick:

1. totally random pick by weight
2. totally consistency hash pick by weight

Rand pick always picks randomly, follows the weight.

Hash pick always picks the same choice when conditions are the same. 
You should make the conditions randomly yourself.

有两种方案

1. 纯随机选择
2. 一致性hash选择

纯随机保证结果完全随机，但是符合权重。

一致性hash保证相同条件下，相同哈希选择的结果最终相同。
请自己保证自己填入的数据是随机的。


### Usage

```
package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/go-ego/murmur"
	"github.com/qiuker521/weightedrand"
)

func main() {
	randpick()
	hashpick()
}
func randpick() {
	c := weightedrand.Chooser{}
	var choices = []weightedrand.Choice{}
	for i := 1; i <= 5; i++ {
		//every choice has a random weight
        //纯随机条件
		choices = append(choices, weightedrand.Choice{strconv.Itoa(i), rand.Float64()})
	}
	c.NewChooser(choices...)
	//This is always a random pick
    //纯随机返回随机结果
	log.Println(c.Pick())
}

func hashpick() {
	c := weightedrand.Chooser{}

	var choices = []weightedrand.Choice{}
	for i := 1; i <= 5; i++ {
		//every choice has the same weight to test hash pick
        //我们假设所有选项都有同样的权重（保证条件相同）
		choices = append(choices, weightedrand.Choice{strconv.Itoa(i), 1})
	}
	c.NewChooser(choices...)

	//This is a hash pick which will always return choice "4".
    //条件相同的时候，这个字符串的随机选择结果永远为4.
	log.Println(c.PickByHash(float64(murmur.Sum32("This is a string, such as a user id"))))
}

```