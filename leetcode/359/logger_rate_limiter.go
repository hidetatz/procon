package main

type Logger struct {
	history map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{history: map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	last, ok := this.history[message]
	if !ok || timestamp-last >= 10 {
		this.history[message] = timestamp
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
