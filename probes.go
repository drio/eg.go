package eg

import (
  "bufio"
  "fmt"
  "github.com/drio/drio.go/common/files"
  "strings"
)

type probeError struct {
	Where string
	What string
}

func (e probeError) Error() string {
	return fmt.Sprintf("Error in %s: %s ", e.Where, e.What)
}

// 1   100006955   rs4908018   TTTGTCTAAAACAAC CTTTCACTAGGCTCA C   A
// Slice of ids [ id1, id2, .... ]
// Probes is a map from [SEQ N SEQ]     -> slice position
//                      [RC(SEQ N SEQ)] -> slice position
// init(filename)  -> go over file and set ids and seq
// matches(string) -> tell me if the string matches any probe (return the id if so)
type Probes struct {
  ids      []string
  seq      map[string]int
  p_size   int
  n_loaded int
}

func (p *Probes) Init() {
  p.ids = make([]string, 0)
  p.seq = make(map[string]int)
}

func (p *Probes) NumLoaded() int {
  return p.n_loaded
}

func (p *Probes) Psize() int {
  return p.p_size
}

// CheckHit checks if subread match a probe if so, it return its id
func (p *Probes) CheckHit(subread string) (string, bool) {
  n_subread := []byte(subread)
  n_subread[p.p_size/2] = 'N'
  id, found := p.seq[string(n_subread)]
  return p.ids[id], found
}

// Load loads the probes in r to probes
func (p *Probes) Load(r *bufio.Reader) error {
  p.Init()
  p.n_loaded = 0
  expected_n_columns_in_line := 7
  for l := range files.IterLines(r) {
    ss := strings.Split(l, "\t")
    if len(ss) != expected_n_columns_in_line {
      return probeError {
				"Probes.Load()",
				fmt.Sprintf("Invalid # of fields in input probes; line: %d", p.n_loaded),
			}
    }
    id, five, three := ss[2], ss[3], ss[4]
    p.add(id, five, three)
  }
	return nil
}

// add adds a new probe to p given the id of the probe and
// the five and three prime sequence
func (p *Probes) add(id, five, three string) error {
  if p.IsEmpty() { // First time we load a probe, set the size
    p.p_size = len(five) + len(three) + 1
  } else {
    if p.p_size != len(five)+len(three)+1 {
      return probeError {
				"Probes.add()",
				"Different probe lengths.",
			}
    }
  }
  //TODO: reverse complement !!
  p.ids = append(p.ids, id)
  p.seq[five+"N"+three] = p.n_loaded
  p.n_loaded++
	return nil
}

func (p *Probes) IsEmpty() bool {
  return p.n_loaded == 0
}
