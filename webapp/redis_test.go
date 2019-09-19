package webapp

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func Test_initializeCache(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initializeCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisIndex(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisIndex(tt.args.w, tt.args.r)
		})
	}
}

func Test_redisUserShow(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisUserShow(tt.args.w, tt.args.r)
		})
	}
}

func Test_redisEditUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisEditUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_redisInsertUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisInsertUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_redisUpdateUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisUpdateUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_redisDeleteUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			redisDeleteUser(tt.args.w, tt.args.r)
		})
	}
}

func Test_getRedisKey(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRedisKey(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("getRedisKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_covertString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := covertString(tt.args.key); got != tt.want {
				t.Errorf("covertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
