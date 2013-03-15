package eg

import (
  "bytes"
  "testing"
)

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

func TestBasicCheckHit(t *testing.T) {
  should_hit := true
  probes := setProbes()
  id := "id1"
  probes.add("id1", "AAA", "CCC")
  testHit(probes, id, "AAATCCC", t, should_hit)
  testHit(probes, id, "ACATCCC", t, !should_hit)
  testHit(probes, id, "AAATCCT", t, !should_hit)
  testHit(probes, id, "CAATCCC", t, !should_hit)
}
