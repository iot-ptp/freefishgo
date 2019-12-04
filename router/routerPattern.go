package router

type Pattern struct {
	PatternString  string
	ControllerName string
	ActionName     string
	//正则匹配出来的变量地址映射变量映射
	PatternMap map[string]int
}
type freeFishUrl struct {
	controllerName   string
	controllerAction string
	OtherKeyMap      map[string]interface{}
}

// 获取控制器名称
func (f *freeFishUrl) GetControllerName() string {
	if v, ok := f.OtherKeyMap["Controller"]; ok {
		return v.(string)
	} else {
		return f.controllerName
	}
}

// 获取动作名称
func (f *freeFishUrl) GetControllerAction() string {
	if v, ok := f.OtherKeyMap["Action"]; ok {
		return v.(string)
	} else {
		return f.controllerName
	}
}
