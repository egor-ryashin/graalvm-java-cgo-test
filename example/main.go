package main

// #cgo CFLAGS: -I${SRCDIR}/../target
// #cgo LDFLAGS: -L${SRCDIR}/../target -lcgotest
//
// #include <stdlib.h>
// #include <libcgotest.h>
//
//
//  static char* pass_str(graal_isolatethread_t* thread) {
//    return java_cgo_str(thread, malloc);
//  }
//
import "C"
import (
	"fmt"
	"unsafe"
)

type javaCgo struct {
	isolate *C.graal_isolate_t
}

type JavaCgo interface {
	Str(s string) (string, error)
}

func New() (JavaCgo, error) {
	var isolate *C.graal_isolate_t
	var thread *C.graal_isolatethread_t

	param := &C.graal_create_isolate_params_t{
		reserved_address_space_size: 1024 * 1024 * 500,
	}

	if C.graal_create_isolate(param, &isolate, &thread) != 0 {
		return nil, fmt.Errorf("failed to initialize")
	}

	return &javaCgo{
		isolate: isolate,
	}, nil
}

func (j *javaCgo) attachThread() (*C.graal_isolatethread_t, error) {
	thread := C.graal_get_current_thread(j.isolate)
	if thread != nil {
		return thread, nil
	}

	var newThread *C.graal_isolatethread_t
	if C.graal_attach_thread(j.isolate, &newThread) != 0 {
		return nil, fmt.Errorf("failed to attach thread")
	}

	return newThread, nil
}

func (j *javaCgo) Str(s string) (string, error) {
	thread, err := j.attachThread()
	if err != nil {
		return "", err
	}

  cstr := C.pass_str(thread)
  s = C.GoString(cstr)
  C.free(unsafe.Pointer(cstr))
  return s, nil
}

func main() {
	javaCgo, err := New()
	if err != nil {
		println(err)
		return
	}

	fmt.Println(javaCgo.Str("hey"))
}
