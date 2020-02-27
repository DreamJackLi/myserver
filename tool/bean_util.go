package tool

import (
	"reflect"
)

/**
 * 对象复制方法
 *
 * @param[in] source 复制源
 * @param[in] target 生成对象指针
 */
func CopyProperties(source interface{}, target interface{}) {
	// 目标类型数组
	targetTypes := reflect.TypeOf(target)
	if targetTypes.Elem().Kind().String() == reflect.Array.String() ||
		targetTypes.Elem().Kind().String() == reflect.Slice.String() {
		copyArrProperties(source, target)
		return
	}
	// 目标值数组
	targetValues := reflect.ValueOf(target)
	// 源类型数组
	sourceTypes := reflect.TypeOf(source)
	// 源值数组
	sourceValues := reflect.ValueOf(source)
	// 遍历目标值数组
	for i := 0; i < targetTypes.Elem().NumField(); i++ {
		// 获取目标字段元素
		typeField := targetTypes.Elem().Field(i)
		// 获取目标字段元素名称
		targetFiledName := typeField.Name
		// 根据目标字段元素名称，获取源数据字段元素
		sourceFieldByFiledName, isExist := sourceTypes.Elem().FieldByName(targetFiledName)
		if isExist {
			// 能获取到信息，拿到对应数据的下表
			index := sourceFieldByFiledName.Index[0]

			sourceValue := sourceValues.Elem().Field(index)
			if sourceValue.Kind().String() == targetValues.Elem().Field(i).Kind().String() {
				targetValues.Elem().Field(i).Set(sourceValue)
			}
		}
	}
}

/**
 * 数组对象复制方法
 *
 * @param[in] source 复制源
 * @param[in] target 生成对象指针
 */
func copyArrProperties(source interface{}, target interface{}) {

	arrSourceValue := reflect.ValueOf(source).Elem()
	arrTargetValue := reflect.ValueOf(target).Elem()
	//targetValue:=reflect.ValueOf(target).Elem()
	for i := 0; i < arrSourceValue.Len(); i++ {
		//targetValue.Index(i).Set(value.Index(i))
		// 目标值数组
		targetValues := arrTargetValue.Index(i)
		// 目标类型数组
		targetTypes := reflect.TypeOf(targetValues.Interface())
		if reflect.TypeOf(targetTypes).Kind() == reflect.Ptr {
			targetTypes = targetTypes.Elem()
		}
		// 源值数组
		sourceValues := arrSourceValue.Index(i)
		// 源类型数组
		sourceTypes := reflect.TypeOf(sourceValues.Interface())
		if reflect.TypeOf(sourceTypes).Kind() == reflect.Ptr {
			sourceTypes = sourceTypes.Elem()
		}
		temp := reflect.New(targetTypes)
		targetValues.Set(temp)

		// 遍历目标值数组
		for i := 0; i < targetTypes.NumField(); i++ {
			// 获取目标字段元素
			typeField := targetTypes.Field(i)
			// 获取目标字段元素名称
			targetFiledName := typeField.Name
			// 根据目标字段元素名称，获取源数据字段元素
			sourceFieldByFiledName, isExist := sourceTypes.FieldByName(targetFiledName)
			if isExist {
				// 能获取到信息，拿到对应数据的下表
				index := sourceFieldByFiledName.Index[0]
				sourceValue := sourceValues.Elem().Field(index)
				//fmt.Println("targetFiledName: ", sourceValue)
				//fmt.Println("targetValues.Elem().Field(index): ", targetValues.Elem().Field(index).Kind())
				//fmt.Println("sourceValue: ", sourceValue.Elem().Field(index).Kind())
				targetValues.Elem().Field(i).Set(sourceValue)
			}
		}
	}
}
