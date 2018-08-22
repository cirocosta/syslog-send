// Package priority exposes functionality for converting human-provided
// syslog priorities into the integer used by `log/syslog`.
package priority

import (
	"fmt"
	"log/syslog"
	"strings"
)

// NewPriority does the job of converting a tuple of facility and severity
// into a syslog priority.
//
// Both facility and severity must valid non-empty syslog entities.
//
// See the `log/syslog` godoc for more information.
func NewPriority(facility, severity string) (prio syslog.Priority, err error) {
	if facility == "" || severity == "" {
		err = fmt.Errorf("facility and severity must be non-empty")
		return
	}

	facility = strings.ToLower(facility)
	severity = strings.ToLower(severity)

	switch facility {
	case "local0":
		prio = syslog.LOG_LOCAL0
	case "local1":
		prio = syslog.LOG_LOCAL1
	case "local2":
		prio = syslog.LOG_LOCAL2
	case "local3":
		prio = syslog.LOG_LOCAL3
	case "local4":
		prio = syslog.LOG_LOCAL4
	case "local5":
		prio = syslog.LOG_LOCAL5
	case "local6":
		prio = syslog.LOG_LOCAL6
	case "local7":
		prio = syslog.LOG_LOCAL7
	default:
		err = fmt.Errorf("unknown facility %s", facility)
		return
	}

	switch severity {
	case "emerg":
		prio = prio | syslog.LOG_EMERG
	case "alert":
		prio = prio | syslog.LOG_ALERT
	case "crit":
		prio = prio | syslog.LOG_CRIT
	case "err":
		prio = prio | syslog.LOG_ERR
	case "warning":
		prio = prio | syslog.LOG_WARNING
	case "notice":
		prio = prio | syslog.LOG_NOTICE
	case "info":
		prio = prio | syslog.LOG_INFO
	case "debug":
		prio = prio | syslog.LOG_DEBUG
	default:
		err = fmt.Errorf("unknown severity %s", severity)
		return
	}

	return
}
