// MusicBrainz
package musicbrainz

const (
   API = "http://musicbrainz.org/ws/2/release"
   invert = "\x1b[7m"
   reset = "\x1b[m"
)

var Status = map[string]int{"Official": 0, "Bootleg": 1}
