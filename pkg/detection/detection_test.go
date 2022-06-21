package detection

import (
	"reflect"
	"testing"
)

func TestDetector_GetAllTLD(t *testing.T) {
	cases := []struct {
		payload  string
		expected []string
	}{
		{"Die ist ein Textmessage mit .de und .com tld aber auch falschen tld wie zum beispiel dev", []string{"com", "de"}},
		{"Viel Text aber keine tld . oder vielleicht doch", []string{}},
		{"Und auch 2 stufige tld müssen erkannt werden wie .co.uk oder auch .co.at", []string{"co.uk", "co.at"}},
		{"Mehrfach vorkommende tld .com müssen einzigartig zurückgegeben werden .com", []string{"com"}},
	}

	for _, c := range cases {
		detector := new(Detector)
		tld, err := detector.GetAllTld(c.payload)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if !(len(tld) == 0 && len(c.expected) == 0) && !reflect.DeepEqual(tld, c.expected) {
			t.Log("tld should be equal: ", tld, c.expected)
		}
	}
}

func TestDetector_WordlistContainsTld(t *testing.T) {
	cases := []struct {
		tld      string
		wordlist map[string]bool
		expected bool
	}{
		{"com", map[string]bool{"com": true, "de": true}, true},
		{"co.uk", map[string]bool{"com": true, "co.uk": true}, true},
		{"test", map[string]bool{"com": true, "de": true}, false},
	}

	for _, c := range cases {
		detector := new(Detector)
		result, err := detector.WordlistContainsTld(c.tld, c.wordlist)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if result != c.expected {
			t.Log("tld should be equal: ", result, c.expected)
		}
	}
}

func TestDetector_GetFirstLevelDomain(t *testing.T) {
	cases := []struct {
		payload  string
		tld      string
		expected []string
	}{
		{"Textmessage mit domain.de und test.com tld aber auch falschen tld wie zum beispiel .com", "de", []string{"domain.de"}},
		{"Textmessage mit domain.com und test.com tld aber auch falschen tld wie zum beispiel .com", "com", []string{"domain.com", "test.com"}},
		{"Viel Text aber keine tld . oder vielleicht doch domain.de", "", []string{}},
		{"Und auch 2 stufige .co.uk müssen erkannt werden wie website.co.uk oder auch .co.at", "co.uk", []string{"website.co.uk"}},
	}

	for _, c := range cases {
		detector := new(Detector)
		domain, err := detector.GetFirstLevelDomain(c.payload, c.tld)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if !(len(domain) == 0 && len(c.expected) == 0) && !reflect.DeepEqual(domain, c.expected) {
			t.Log("domain should be equal: ", domain, c.expected)
		}
	}
}

func TestDetector_GetSubdomain(t *testing.T) {
	cases := []struct {
		payload  string
		domain   string
		expected []string
	}{
		{"Textmessage mit sub.domain.de und .domain.de tld aber auch falschen tld wie zum beispiel domain.de", "domain.de", []string{"sub.domain.de"}},
		{"Textmessage mit .test.com und first.second.test.com tld aber auch falschen tld wie zum beispiel .com", "test.com", []string{"first.second.test.com"}},
		{"Textmessage mit .test.com und first.second.test.com tld aber auch zwei tld wie zum beispiel second.first.test.com", "test.com", []string{"first.second.test.com", "second.first.test.com"}},
		{"Viel Text aber keine tld . oder vielleicht doch second.test.com", "", []string{}},
		{"Und auch 2 stufige sub.website.co.uk müssen erkannt werden wie website.co.uk oder auch .co.at", "website.co.uk", []string{"sub.website.co.uk"}},
	}

	for _, c := range cases {
		detector := new(Detector)
		domain, err := detector.GetSubdomain(c.payload, c.domain)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if !(len(domain) == 0 && len(c.expected) == 0) && !reflect.DeepEqual(domain, c.expected) {
			t.Log("domain should be equal: ", domain, c.expected)
		}
	}
}

func TestDetector_GetPath(t *testing.T) {
	cases := []struct {
		payload  string
		domain   string
		expected []string
	}{
		{"Textmessage mit sub.domain.de?param=1 und .domain.de tld aber auch falschen tld wie zum beispiel sub.domain.de", "domain.de", []string{"sub.domain.de?param=1"}},
		{"Textmessage mit .test.com und first.second.test.com?param=1&param=2 tld aber auch falschen tld wie zum beispiel .com", "test.com", []string{"first.second.test.com?param=1&param=2"}},
		{"Textmessage mit .test.com und first.second.test.com/text tld aber auch zwei tld wie zum beispiel second.first.test.com/auch", "test.com", []string{"first.second.test.com/text", "second.first.test.com/auch"}},
		{"Viel Text aber keine tld . oder vielleicht doch second.test.com", "", []string{}},
		{"Und auch 2 stufige sub.website.co.uk/add/follower müssen erkannt werden wie website.co.uk oder auch .co.at", "sub.website.co.uk", []string{"sub.website.co.uk/add/follower"}},
	}

	for _, c := range cases {
		detector := new(Detector)
		domain, err := detector.GetSubdomain(c.payload, c.domain)
		if err != nil {
			t.Log("should not throw err: ", err)
		}
		if !(len(domain) == 0 && len(c.expected) == 0) && !reflect.DeepEqual(domain, c.expected) {
			t.Log("domain should be equal: ", domain, c.expected)
		}
	}
}
