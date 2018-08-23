package priority_test

import (
	"log/syslog"
	"testing"

	p "github.com/cirocosta/syslog-send/priority"
)

func assert(t *testing.T, val bool, message string) {
	if val {
		return
	}

	t.Fatal(message)
}

func TestNewPriority(t *testing.T) {
	testCases := []struct {
		desc       string
		facility   string
		severity   string
		shouldFail bool
		expected   syslog.Priority
	}{
		{
			desc:       "no facility provided",
			facility:   "",
			severity:   "",
			shouldFail: true,
		},
		{
			desc:       "no severity provided",
			facility:   "local0",
			severity:   "",
			shouldFail: true,
		},
		{
			desc:       "invalid facility",
			facility:   "something-invalid",
			severity:   "emerg",
			shouldFail: true,
		},
		{
			desc:       "invalid severity",
			facility:   "local0",
			severity:   "something-invalid",
			shouldFail: true,
		},
		{
			desc:       "valid",
			facility:   "local0",
			severity:   "emerg",
			shouldFail: false,
			expected:   syslog.LOG_LOCAL0 | syslog.LOG_EMERG,
		},
	}

	var (
		err      error
		priority syslog.Priority
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			priority, err = p.NewPriority(tc.facility, tc.severity)
			if tc.shouldFail {
				assert(t, err != nil, "must fail")
				return
			}

			assert(t, err == nil, "must not fail")
			assert(t, priority == tc.expected, "priority must match")
		})
	}
}
