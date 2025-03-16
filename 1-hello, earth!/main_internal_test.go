package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello, Earth!",
		},
		"Spanish": {
			lang: "es",
			want: "¡Hola, Tierra!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour, Terre!",
		},
		"Greek": {
			lang: "gr",
			want: "Γειά σου, Γη!",
		},
		"Hebrew": {
			lang: "he",
			want: "שלום, ארץ!",
		},
		"Italian": {
			lang: "it",
			want: "Ciao, Terra!",
		},
		"Akkadian": {
			lang: "akk",
			want: "𒀭𒊏𒁲𒀭𒊏𒁲",
		},
		"Unsupported": {
			lang: "",
			want: "Language \"\" not supported",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("greet(%q) = %q; want %q", tc.lang, got, tc.want)
			}
		})
	}
}
