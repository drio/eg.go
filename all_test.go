package eg

import (
  "bytes"
  "testing"
)

func testRC(t *testing.T) {
  f := func(r, rc, res string) {
    t.Errorf("RC(%s) should be [%s] not [%s]", r, rc, res)
  }

  r, rc := "AAANCCC", "GGGNTTT"
  if res := ReverseComplement(&r); bytes.Compare([]byte(r), []byte(*res)) != 0 {
    f(r, rc, *res)
  }

  r, rc = "GATTTGGGGTTCAAAGCAGTATCGATCAAATAGTAAATCCATTTGTTCAACTCACAGTTT",
    "AAACTGTGAGTTGAACAAATGGATTTACTATTTGATCGATACTGCTTTGAACCCCAAATC"
  if res := ReverseComplement(&r); bytes.Compare([]byte(r), []byte(*res)) != 0 {
    f(r, rc, *res)
  }
}

func setProbes() *Probes {
  probes := Probes{}
  probes.Init()
  return &probes
}

func TestProbeEmpty(t *testing.T) {
  probes := setProbes()
  if !probes.IsEmpty() {
    t.Errorf("The probes structure should be empty.")
  }
  probes.add("id1", "ACT", "CCC")
  if probes.IsEmpty() {
    t.Errorf("The probes structure should NOT be empty.")
  }
}

func testHit(probes *Probes, id, read string, t *testing.T, should_hit bool) {
  if rid, found := probes.CheckHit(read); found == !should_hit {
    t.Errorf("%s should hit", read)
  } else {
    if bytes.Compare([]byte(id), []byte(rid)) != 0 {
      t.Errorf("the id should be [%s] not [%s] ", id, rid)
    }
  }
}

func TestDifferentProbeLengths(t *testing.T) {
  probes := setProbes()
  if err := probes.add("id1", "AAA", "CCC"); err != nil {
    t.Errorf("Getting error when adding first probe ")
  }
  if err := probes.add("id2", "AAAT", "GCCC"); err == nil {
    t.Errorf("Should have gotten an error when inserting probe of different size")
  }
}

func TestBasicCheckHit(t *testing.T) {
  should_hit := true
  probes := setProbes()

  id := "id1"
  probes.add(id, "AAA", "CCC")
  testHit(probes, id, "AAATCCC", t, should_hit)
  testHit(probes, id, "GGGNTTT", t, should_hit) // Reverse Complement
  testHit(probes, id, "ACATCCC", t, !should_hit)
  testHit(probes, id, "AAATCCT", t, !should_hit)
  testHit(probes, id, "CAATCCC", t, !should_hit)

}

// Test eg.Screen()
func TestScreening(t *testing.T) {
  probes := setProbes()

  probes.add("id1", "CT", "GA")
  read := "ACTAGAAT"
  //   CTXGA
  hits := Screen(probes, read)
  if len(hits) != 1 {
    t.Errorf("read [%s] should have 1 hit, has %d", read, len(hits))
  }

  probes.add("id2", "AG", "AT")
  read = "ACTAGAAT"
  //   CTXGA
  //     AGXAT
  hits = Screen(probes, read)
  if len(hits) != 2 {
    t.Errorf("Probes: %d read [%s] should have 2 hits, has %d", probes, read, len(hits))
  }

}
