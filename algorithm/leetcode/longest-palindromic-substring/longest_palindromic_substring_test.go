package longest_palindromic_substring

import "testing"

/*
 * @author taohu
 * @date 2020/4/13
 * @time 下午8:29
 * @desc please describe function
**/

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "babad"},
			want: "aba",
		},
		{
			name: "case1",
			args: args{s: "cbbd"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
