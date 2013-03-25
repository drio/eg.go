package eg

import (
  "fmt"
  "github.com/drio/drio.go/bio/fasta"
  "io"
)

type Hit struct {
  ProbeId string
  Base    byte
}

// Screen looks for subreads in read to see if they match
// any of the probes, if so, it saves them. We return the
// list of hits when done.
func Screen(probes *Probes, read string) []Hit {
  hits := []Hit{}
  probe_len := probes.Psize()
  middle := probe_len / 2
  for i := 0; i <= len(read)-probe_len; i++ {
    subread := []byte(read[i : i+probe_len])
    nt := byte(subread[middle])
    subread[middle] = 'N'
    //fmt.Printf("subread: %s\n", string(subread))
    if probe_id, matches := probes.CheckHit(string(subread)); matches {
      hits = append(hits, Hit{probe_id, nt})
    }
  }
  return hits
}

// Compute iterates over fq and screens each read against the probes. All the
// hits found are dump to outputF
func Compute(fqr fasta.FqReader, probes Probes, outputF io.Writer) int64 {
	n_reads := int64(0)
  for r, done := fqr.Iter(); !done; r, done = fqr.Iter() {
		n_reads++
    for _, hit := range Screen(&probes, r.Seq) {
      fmt.Fprintf(outputF, "%s\t%c\n", hit.ProbeId, hit.Base)
    }
  }
	return n_reads
}
