package eg

import (
	"log"
	"bufio"
	"github.com/drio/drio.go/common/files"
	"strings"
	"fmt"
	"strconv"
)

// 1   100006955   rs4908018   TTTGTCTAAAACAAC CTTTCACTAGGCTCA C   A
// Slice of ids [ id1, id2, .... ]
// Probes is a map from [SEQ N SEQ]     -> slice position
//                      [RC(SEQ N SEQ)] -> slice position
// init(filename)  -> go over file and set ids and seq
// matches(string) -> tell me if the string matches any probe (return the id if so)
type Probes struct {
	ids []string
	seq map[string]int
}

func (p *Probes) Init(r *bufio.Reader) {
	log.Printf("Loading probes.")
	p.ids = make([]string, 100000)
	p.seq = make(map[string]int)
	n_probes := 0
	expected_n_columns_in_line := 7
	for l := range(files.IterLines(r)) {
		ss := strings.Split(l, "\t")
		if len(ss) != expected_n_columns_in_line {
			panic(fmt.Sprintf("Invalid probe line: %s", l))
		}
		id, five, three := ss[2], ss[3], ss[4]
		p.ids = append(p.ids, id)
		p.seq[five + "N" + three] = n_probes
		n_probes++
	}
	log.Printf("Number of probes loaded: " + strconv.Itoa(n_probes))
}

