/* -*- Indent-tabs-mode:nil; coding: utf-8 -*-
 * Copyleft (C) 2015
 * "Tag bao" known as "wackonline" <bb.qnyd@gmail.com>
 * Goway is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License and GNU
 * Lesser General Public License published by the Free Software
 * Foundation, either version 3 of the License, or (at your option)
 * any later version.
 * Goway is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License and GNU Lesser General Public License
 * for more details.
 * You should have received a copy of the GNU General Public License
 * and GNU Lesser General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */
package goway

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	//E_ALL     = 1
	E_ERROR   = 1
	E_WARNING = 2
	E_STRICT  = 4
	E_NOTICE  = 8
)

type Logger interface {
	Setloglevel(string)
	IsLogger(int) bool
	StartLogger() Handler
	Error(string, ...interface{})
	Warning(string, ...interface{})
	Strict(string, ...interface{})
	Notice(string, ...interface{})
	Use([]Tlogs)
	Print()
}
type Tlogs map[int]string
type Logs struct {
	logger *log.Logger
	lvs    int
	logs   []Tlogs
}

func (lg *Logs) Error(str string, a ...interface{}) {
	lg.addInfo(E_ERROR, str, a...)
}

func (lg *Logs) Warning(str string, a ...interface{}) {
	lg.addInfo(E_WARNING, str, a...)
}

func (lg *Logs) Strict(str string, a ...interface{}) {
	lg.addInfo(E_STRICT, str, a...)
}

func (lg *Logs) Notice(str string, a ...interface{}) {
	lg.addInfo(E_NOTICE, str, a...)
}

func (lg *Logs) Use(lgs []Tlogs) {
	for _, v := range lgs {
		lg.logs = append(lg.logs, v)
	}
}

func (lg *Logs) addInfo(ty int, str string, a ...interface{}) {
	log := make(map[int]string)
	if len(a) > 0 {
		log[ty] = fmt.Sprintf(str, a...)
	} else {
		log[ty] = str
	}

	lg.logs = append(lg.logs, log)
}

// Determine whether excluded set log level
// Example:
//   all := A|B|C|D|E
//   all & C not eq 0 (C in ALL)
//   all1 := A|B|C
//   all1 & D eq 0 (D not in ALL)
func (lg *Logs) IsLogger(v int) bool {
	perm := lg.lvs & v
	if perm == 0 {
		return false
	}
	return true
}

//Excluded error message is set to false
func (lg *Logs) Setloglevel(lv string) {
	str := strings.Split(lv, "|")
	if len(str) >= 1 {
		for _, v := range str {
			if v == "E_NOTICE" {
				lg.lvs = lg.lvs | E_NOTICE
			} else if v == "E_ERROR" {
				lg.lvs = lg.lvs | E_ERROR
			} else if v == "E_WARNING" {
				lg.lvs = lg.lvs | E_WARNING
			} else if v == "E_STRICT" {
				lg.lvs = lg.lvs | E_STRICT
			} else if v == "E_ALL" {
				lg.lvs = E_ERROR | E_WARNING | E_STRICT | E_NOTICE
			} else {
				lg.lvs = 0
			}
		}
	}
	//lg.all = false
}

func NewLogger() Logger {
	logs := &Logs{logger: log.New(os.Stdout, "[*GOWAY*] ", 0)}
	logs.lvs = 0 //E_ERROR | E_WARNING | E_STRICT | E_NOTICE
	return logs
}

func (lg *Logs) setLogInfo(k int, v string) {
	switch k {
	case E_ERROR:
		lg.logger.SetPrefix("[*Goway*][Error] ")
		lg.logger.Printf(v)
	case E_WARNING:
		lg.logger.SetPrefix("[*Goway*][Warning] ")
		lg.logger.Printf(v)
	case E_STRICT:
		lg.logger.SetPrefix("[*Goway*][Strict] ")
		lg.logger.Printf(v)
	case E_NOTICE:
		lg.logger.SetPrefix("[*Goway*][Notice] ")
		lg.logger.Printf(v)

	}
}

func (lg *Logs) Print() {
	for _, log := range lg.logs {
		for k, v := range log {
			if lg.IsLogger(E_ERROR) && lg.IsLogger(E_WARNING) && lg.IsLogger(E_STRICT) && lg.IsLogger(E_NOTICE) {
				lg.setLogInfo(k, v)
			} else if lg.IsLogger(E_ERROR) && k == E_ERROR {
				lg.setLogInfo(k, v)
			} else if lg.IsLogger(E_NOTICE) && k == E_NOTICE {
				lg.setLogInfo(k, v)
			} else if lg.IsLogger(E_STRICT) && k == E_STRICT {
				lg.setLogInfo(k, v)
			} else if lg.IsLogger(E_WARNING) && k == E_WARNING {
				lg.setLogInfo(k, v)
			}
		}
	}
	lg.logs = []Tlogs{}
}

func (lg *Logs) StartLogger() Handler {
	return func(res http.ResponseWriter, req *http.Request, c Context, lgs Logger) {
		start := time.Now()
		addr := req.Header.Get("X-Real-IP")
		if addr == "" {
			addr = req.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = req.RemoteAddr
			}
		}
		info := fmt.Sprintf("Started %s %s for %s", req.Method, req.URL.Path, addr)
		lgs.Notice(info)

		rw := res.(ResponseWriter)
		c.Next()
		info = fmt.Sprintf("Completed %v %s, Content-Length: %v bytes in %v\n", rw.Status(), http.StatusText(rw.Status()), rw.Size(), time.Since(start))
		lgs.Notice(info)
		lgs.Print()
	}
}
