package webapp

import "testing"

func Test_generateLogsFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generateLogsFile()
		})
	}
}

func Test_logStdout(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logStdout()
		})
	}
}

func Test_logFile(t *testing.T) {
	type args struct {
		logtype string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logFile(tt.args.logtype)
		})
	}
}
