package speech

// #cgo CFLAGS: -I/usr/include/speech-dispatcher
// #cgo LDFLAGS: -lspeechd
// #include <speech-dispatcher/libspeechd.h>
// #include <stdlib.h> // Required for C.free
/*
int say_c(char *message, char *lang)
{
    SPDConnection *conn;

    // Open Speech Dispatcher connection in SINGLE MODE mode.
    conn = spd_open("go-say-wrapper", "conn", NULL, SPD_MODE_SINGLE);

    // Set Language
    if (lang != NULL) {
        spd_set_language(conn, lang);
    }

    // Say message.
    // Note: SPD_MESSAGE is typically used for general text.
    // We cast to (char*) although it's passed as a const char* in the original C++
    // for compatibility with the C library function signature if it's older.
    spd_say(conn, SPD_MESSAGE, message);

    // Close connection
    spd_close(conn);

    return 0;
}
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

// Say speaks the given message using Speech Dispatcher.
// The language is optional and defaults to the system's setting if an empty
// string is passed. Example languages: "en", "es", "fr".
func Say(message, lang string) error {
	// CGo calls must not be called from Go threads that are created by C code,
	// and they can block. Locking the OS thread is a good practice.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Convert Go strings to C strings (char*).
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage)) // Free the C-allocated memory for the message

	var cLang *C.char
	if lang != "" {
		cLang = C.CString(lang)
		defer C.free(unsafe.Pointer(cLang)) // Free the C-allocated memory for the language
	} else {
		// Pass NULL if no language is specified, letting the C function handle the default.
		cLang = nil
	}

	// Call the C function. The C code is embedded above in the multi-line C comment.
	result, err := C.say_c(cMessage, cLang)

	if err != nil {
		return fmt.Errorf("error calling speechd 'say_c': %w", err)
	}
	if result != 0 {
		return fmt.Errorf("speechd 'say_c' returned non-zero exit code: %d", result)
	}

	return nil
}
