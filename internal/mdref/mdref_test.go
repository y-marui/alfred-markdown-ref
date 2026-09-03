package mdref

import "testing"

const sampleInput = `This is sample[B] of this workflow[1].
Some links[A] nomatter character[5].
You don't have to specify ref links[D].
It skips the corresponding number[C].
This[AAAA] will be ignored.
![alt text also need sufficient length](https://image.url)

[5]: some url1
[1]: some url2
[A]: some url3
[B]: some url4
[C]: some url5
`

func TestConvert_DefaultStart(t *testing.T) {
	want := `This is sample[1] of this workflow[2].
Some links[3] nomatter character[4].
You don't have to specify ref links[5].
It skips the corresponding number[6].
This[AAAA] will be ignored.
![alt text also need sufficient length](https://image.url)

[1]: some url4
[2]: some url2
[3]: some url3
[4]: some url1
[5]:
[6]: some url5`

	got, err := Convert(sampleInput, 1)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if got != want {
		t.Errorf("Convert() =\n%q\nwant\n%q", got, want)
	}
}

func TestConvert_CustomStart(t *testing.T) {
	want := `This is sample[3] of this workflow[4].
Some links[5] nomatter character[6].
You don't have to specify ref links[7].
It skips the corresponding number[8].
This[AAAA] will be ignored.
![alt text also need sufficient length](https://image.url)

[3]: some url4
[4]: some url2
[5]: some url3
[6]: some url1
[7]:
[8]: some url5`

	got, err := Convert(sampleInput, 3)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if got != want {
		t.Errorf("Convert() =\n%q\nwant\n%q", got, want)
	}
}

func TestConvert_InvalidStart(t *testing.T) {
	for _, start := range []int{0, -1} {
		if _, err := Convert("text", start); err == nil {
			t.Errorf("Convert(text, %d) error = nil, want error", start)
		}
	}
}

func TestConvert_NoReferences(t *testing.T) {
	got, err := Convert("plain text with no refs", 1)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if want := "plain text with no refs\n"; got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvert_UndefinedReferenceGetsBlankEntry(t *testing.T) {
	input := "a[X] b[Y] c[Z]\n\n[X]: url-x\n[Z]: url-z\n"
	want := "a[1] b[2] c[3]\n\n[1]: url-x\n[2]:\n[3]: url-z"

	got, err := Convert(input, 1)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

// A regression case for the numeric-collision risk: an undefined reference
// whose original label already looks like a plausible renumbered target
// (here "[2]") must not be left as-is, or it would be visually
// indistinguishable from — and could literally collide with — a
// genuinely-renumbered reference.
func TestConvert_UndefinedNumericLabelStillRenumbered(t *testing.T) {
	input := "a[X] b[2] c[Y]\n\n[X]: url-x\n[Y]: url-y\n"
	want := "a[1] b[2] c[3]\n\n[1]: url-x\n[2]:\n[3]: url-y"

	got, err := Convert(input, 1)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvert_RepeatedKeyUsesFirstOccurrenceOrder(t *testing.T) {
	input := "a[X] b[Y] c[X]\n\n[X]: url-x\n[Y]: url-y\n"
	want := "a[1] b[2] c[1]\n\n[1]: url-x\n[2]: url-y"

	got, err := Convert(input, 1)
	if err != nil {
		t.Fatalf("Convert() error = %v", err)
	}
	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}
