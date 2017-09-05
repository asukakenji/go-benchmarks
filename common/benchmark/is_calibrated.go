package benchmark

import "reflect"

var _isCalibrated = map[reflect.Kind]map[reflect.Kind]map[uintptr]struct{}{}

func isCalibrated(type0, type1 reflect.Kind, funcPointer uintptr) bool {
	_isCalibrated0, ok := _isCalibrated[type0]
	if !ok {
		return false
	}
	_isCalibrated1, ok := _isCalibrated0[type1]
	if !ok {
		return false
	}
	_, ok = _isCalibrated1[funcPointer]
	return ok
}

func setCalibrated(type0, type1 reflect.Kind, funcPointer uintptr) {
	_isCalibrated0, ok := _isCalibrated[type0]
	if !ok {
		_isCalibrated[type0] = map[reflect.Kind]map[uintptr]struct{}{}
		_isCalibrated0 = _isCalibrated[type0]
	}
	_isCalibrated1, ok := _isCalibrated0[type1]
	if !ok {
		_isCalibrated0[type1] = map[uintptr]struct{}{}
		_isCalibrated1 = _isCalibrated0[type1]
	}
	_isCalibrated1[funcPointer] = struct{}{}
}
