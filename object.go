package objc

// #include <stdlib.h>
// #include <objc/runtime.h>
import "C"
import "unsafe"

type Id C.id

func Object_copy(obj Id, size uint) Id {
	return Id(C.object_copy(obj, C.size_t(size)))
}

func Object_dispose(obj Id) Id {
	return Id(C.object_dispose(obj))
}

func Object_setInstanceVariable(obj Id, name string, value unsafe.Pointer) Ivar {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	return Ivar(C.object_setInstanceVariable(obj, cname, value))
}

func Object_getInstanceVariable(obj Id, name string) (ivar Ivar, outValue unsafe.Pointer) {
	cname := C.CString(name)
	defer free(unsafe.Pointer(cname))

	ivar = Ivar(C.object_getInstanceVariable(obj, cname, &outValue))
	return
}

func Object_getIndexedIvars(obj Id) unsafe.Pointer {
	return C.object_getIndexedIvars(obj)
}

func Object_getIvar(object Id, ivar Ivar) unsafe.Pointer {
	return unsafe.Pointer(C.object_getIvar(object, ivar))
}

func Object_setIvar(object Id, ivar Ivar, value unsafe.Pointer) {
	C.object_setIvar(object, ivar, C.id(value))
}

func Object_getClassName(obj Id) string {
	return C.GoString(C.object_getClassName(obj))
}

func Object_getClass(object Id) Class {
	return Class(C.object_getClass(object))
}

func Object_setClass(object Id, cls Class) Class {
	return Class(C.object_setClass(object, cls))
}
