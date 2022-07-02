package logging

import (
	"context"
	"reflect"
	"testing"
)

type mock struct {
	NullLogger

	fields map[string]interface{}
}

func newMock() *mock {
	return &mock{fields: make(map[string]interface{})}
}

func (l *mock) WithFields(fields map[string]interface{}) Logger {
	l2 := newMock()

	for k, v := range l.fields {
		l2.fields[k] = v
	}

	for k, v := range fields {
		l2.fields[k] = v
	}

	return l2
}

func (l *mock) WithField(k string, v interface{}) Logger {
	return l.WithFields(map[string]interface{}{k: v})
}

// testWithLogger is a helper with reset the origin default logger
func testWithLogger(t *testing.T, logger Logger) {
	origin := l
	t.Cleanup(func() {
		l = origin
	})

	SetDefaultLogger(logger)
}

func TestSetDefaultLoggerToNil(t *testing.T) {
	var want Logger = NullLogger{}
	testWithLogger(t, nil)

	if l != want {
		t.Fatal("SetDefaultLogger(nil) should set the default Logger to NullLogger")
	}
}

func TestSetDefaultLogger(t *testing.T) {
	var want Logger = newMock()
	testWithLogger(t, want)

	if l != want {
		t.Fatal("SetDefaultLogger(nil) should set the default Logger to NullLogger")
	}
}

func TestFromContext(t *testing.T) {
	t.Run("context doesn't contain Logger", func(t *testing.T) {
		want := newMock()
		testWithLogger(t, want)

		if got := FromContext(context.Background()); got != want {
			t.Error("got should be the same instance with want")
		}
	})

	t.Run("context contains a Logger", func(t *testing.T) {
		want := newMock()
		ctx := IntoContext(context.Background(), want)

		if got := FromContext(ctx); got != want {
			t.Error("got should be the same instance with want")
		}
	})
}

func TestWithFields(t *testing.T) {
	testWithLogger(t, newMock())

	ctx, _ := WithFields(context.Background(), map[string]interface{}{
		"foo1": "bar1",
	})

	ctx, got := WithField(ctx, "foo2", "bar2")

	if !reflect.DeepEqual(got.(*mock).fields, map[string]interface{}{
		"foo1": "bar1",
		"foo2": "bar2",
	}) {
		t.Error("the returned Logger should contains injected fields")
	}

	if !reflect.DeepEqual(FromContext(ctx).(*mock).fields, map[string]interface{}{
		"foo1": "bar1",
		"foo2": "bar2",
	}) {
		t.Error("the returned context should contains injected fields")
	}
}

func TestCopy(t *testing.T) {
	testWithLogger(t, newMock())

	src, _ := WithField(context.Background(), "foo1", "bar1")
	dst := context.Background()

	want := FromContext(src)
	got := FromContext(Copy(dst, src))

	if got != want {
		t.Error("got should be the same instance with want")
	}
}

func TestNullLogger_ShouldNotPanic(t *testing.T) {
	var l Logger = NullLogger{}
	l = l.WithField("foo", "bar").WithFields(map[string]interface{}{
		"baz": "foo",
	})
	l.Info("Info")
	l.Infof("Infof")
	l.Warn("Warn")
	l.Warnf("Warnf")
	l.Error("Error")
	l.Errorf("Errorf")
	l.Fatal("Fatal")
	l.Fatalf("Fatalf")
}
