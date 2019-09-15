package soap

import (
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func logAndUnmarshal(req *http.Request) (receivedEnvelope, error) {
	rawReq, _ := ioutil.ReadAll(req.Body)
	log.
		WithFields(log.Fields{
			"bytes": len(rawReq),
			"from":  req.RemoteAddr,
		}).
		Trace("received request\n", string(rawReq))
	fmt.Println(string(rawReq))

	return Unmarshal(rawReq)
}
