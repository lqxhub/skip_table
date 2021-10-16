package skiptable

type SkipListLevel struct {
	//指向下一个结点
	forward *SkipListNode

	/*
	 * 到下一个node的距离;
	 * 思考,为啥是记录到下一个node, 而不是记录上一个node到这的距离
	 */
	span int64
}

type SkipListNode struct {
	//指向上一个结点
	backward *SkipListNode
	//索引用的层
	level []SkipListLevel
	//存储的值
	value ISkipListNode
	//排名用的分数
	score float64
}

func NewSkipListNode(level int, score float64, value ISkipListNode) *SkipListNode {
	return &SkipListNode{
		backward: nil,
		level:    make([]SkipListLevel, level),
		value:    value,
		score:    score,
	}
}

//第i层的下一个元素
func (node *SkipListNode) Next(i int) *SkipListNode {
	return node.level[i].forward
}

//设置第i层的下一个元素
func (node *SkipListNode) SetNext(i int, next *SkipListNode) {
	node.level[i].forward = next
}

//第i层的span值
func (node *SkipListNode) Span(i int) int64 {
	return node.level[i].span
}

//设置第i层的span值
func (node *SkipListNode) SetSpan(i int, span int64) {
	node.level[i].span = span
}

//上一个元素    想一下,为啥指向上一个的元素不需要i呢???
func (node *SkipListNode) Pre() *SkipListNode {
	return node.backward
}
