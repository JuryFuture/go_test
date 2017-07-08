package module

type Item map[string]interface{}

// 用于判断条目是否有效
func (item Item) Valid() bool {
	return item != nil
}
