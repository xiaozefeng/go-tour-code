package errhandling

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type ErrWriter struct {
	io.Writer
	err error
}

func (e *ErrWriter) Write(p []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	var n int
	n, e.err = e.Writer.Write(p)
	return n, nil
}

// solution 2
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &ErrWriter{Writer: w}
	_, _ = fmt.Fprintf(ew, "HTTP/1.1 %d %d\r\n", st.Code, st.Reason)

	for _, h := range headers {
		_, _ = fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}
	_, _ = fmt.Fprint(ew, "\r\n")
	_, _ = io.Copy(ew, body)
	return ew.err
}

//solution 1
//func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
//	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
//	if err != nil {
//		return err
//	}
//	for _, h := range headers {
//		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
//		if err != nil {
//			return err
//		}
//	}
//	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
//		return err
//	}
//	_, err = io.Copy(w, body)
//	return err
//}
