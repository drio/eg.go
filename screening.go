package eg

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
    subread := []byte(read[i : probe_len+1])
    nt := byte(subread[middle])
    subread[middle] = 'N'
    if probe_id, matches := probes.CheckHit(string(subread)); matches {
      hits = append(hits, Hit{probe_id, nt})
    }
  }
  return hits
}
