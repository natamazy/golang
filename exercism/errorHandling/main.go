package main

// My code
/*
// func Use(opener ResourceOpener, input string) (err error) {
// 	// Retry opening the resource if a TransientError occurs
// 	var resource Resource

// 	for {
// 		resource, err = opener()
// 		if err == nil {
// 			break
// 		}
// 		if _, ok := err.(TransientError); !ok {
// 			return err
// 		}
// 	}

// 	// Ensure that the resource is closed, even in case of panics or errors
// 	defer func() {
// 		// Recover from panic
// 		if r := recover(); r != nil {
// 			// Check if it's a FrobError
// 			if frobErr, ok := r.(FrobError); ok {
// 				// Call Defrob with the defrobTag from FrobError
// 				resource.Defrob(frobErr.defrobTag)
// 				err = frobErr.inner
// 			} else {
// 				// For other panics, convert to an error
// 				err = r.(error)
// 			}
// 		}

// 		// Always close the resource, regardless of whether there's an error or panic
// 		resource.Close()
// 	}()

// 	// Call the Frob function on the opened resource
// 	resource.Frob(input)

// 	return err
// }
*/

// Better solutin
func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	// Open resource
	for r, err = o(); err != nil; r, err = o() {
		if _, ok := err.(TransientError); ok {
			continue
		}
		return err
	}
	defer r.Close()
	defer func() {
		if e := recover(); e != nil {
			if f, ok := e.(FrobError); ok {
				r.Defrob(f.defrobTag)
			}
			err = e.(error)
		}
	}()
	r.Frob(input)
	return err
}
