package module

type Data interface{
	// 用于判断数据数据是否有效
	Valid() bool
}
