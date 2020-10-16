package cron

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestWithLocation(t *testing.T) {
	c := New(WithLocation(time.UTC))
	if c.location != time.UTC {
		t.Errorf("expected UTC, got %v", c.location)
	}
}

func TestWithParser(t *testing.T) {
	var parser = NewParser(Dow)
	c := New(WithParser(parser))
	if c.parser != parser {
		t.Error("expected provided parser")
	}
}

func TestWithVerboseLogger(t *testing.T) {
	var buf syncWriter
	var logger = log.New(&buf, "", log.LstdFlags)
	c := New(WithLogger(VerbosePrintfLogger(logger)))
	if c.logger.(printfLogger).logger != logger {
		t.Error("expected provided logger")
	}

	c.AddFunc(EntryID("0"), "@every 1s", func() {})
	c.Start()
	time.Sleep(OneSecond)
	c.Stop()
	out := buf.String()
	if !strings.Contains(out, "schedule,") ||
		!strings.Contains(out, "run,") {
		t.Error("expected to see some actions, got:", out)
	}
}

func TestWithLocker(t *testing.T) {
	var locker = LockerMock{}
	c := New(WithLocker(locker))
	if c.locker != locker {
		t.Error("expected provided locker")
	}
}

type LockMock struct {
}

func (l *LockMock) TTL() (time.Duration, error) {
	return 0, nil
}

func (l *LockMock) Release() error {
	return nil
}

type LockerMock struct {
}

func (l LockerMock) Obtain(key string, ttl time.Duration) (Lock, error) {
	return &LockMock{}, nil
}
